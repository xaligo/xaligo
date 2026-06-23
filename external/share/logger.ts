import { MCode } from './mcode';

export interface LoggerConfig {
  component?: string;
  service?: string;
  level?: string;
  structured?: boolean;
  enableCaller?: boolean;
  output?: string;
}

export interface Logger {
  DEBUG(mcode: MCode, optionalMessage: string, fields?: Record<string, unknown>): void;
  INFO(mcode: MCode, optionalMessage: string, fields?: Record<string, unknown>): void;
  WARN(mcode: MCode, optionalMessage: string, fields?: Record<string, unknown>): void;
  ERROR(mcode: MCode, optionalMessage: string, fields?: Record<string, unknown>): void;
  FATAL(mcode: MCode, optionalMessage: string, fields?: Record<string, unknown>): void;
}

export enum LogLevel {
  Debug = 0,
  Info = 1,
  Warn = 2,
  Error = 3,
  Fatal = 4,
}

export interface LogEntry {
  timestamp: string;
  level: string;
  code: string;
  component?: string;
  service?: string;
  message: string;
  fields?: Record<string, unknown>;
  file?: string;
  function?: string;
  line?: number;
  error?: string;
}

type LogWriter = (line: string) => void;

class SharedLogger implements Logger {
  private readonly level: LogLevel;
  private readonly output: LogWriter;

  constructor(private readonly config: LoggerConfig, output?: LogWriter) {
    this.level = parseLogLevel(config.level);
    this.output = output ?? openLogOutput(config.output);
  }

  DEBUG(mcode: MCode, optionalMessage: string, fields?: Record<string, unknown>): void {
    this.log(LogLevel.Debug, mcode, optionalMessage, fields);
  }

  INFO(mcode: MCode, optionalMessage: string, fields?: Record<string, unknown>): void {
    this.log(LogLevel.Info, mcode, optionalMessage, fields);
  }

  WARN(mcode: MCode, optionalMessage: string, fields?: Record<string, unknown>): void {
    this.log(LogLevel.Warn, mcode, optionalMessage, fields);
  }

  ERROR(mcode: MCode, optionalMessage: string, fields?: Record<string, unknown>): void {
    this.log(LogLevel.Error, mcode, optionalMessage, fields);
  }

  FATAL(mcode: MCode, optionalMessage: string, fields?: Record<string, unknown>): void {
    this.log(LogLevel.Fatal, mcode, optionalMessage, fields);
    if (typeof process !== 'undefined') process.exit(1);
    throw new Error(optionalMessageFor(mcode, optionalMessage));
  }

  private log(level: LogLevel, mcode: MCode, message: string, fields?: Record<string, unknown>): void {
    if (level < this.level) return;
    const entry: LogEntry = {
      timestamp: new Date().toISOString(),
      level: logLevelString(level),
      code: mcode.code,
      message: optionalMessageFor(mcode, message),
    };
    if (this.config.component) entry.component = this.config.component;
    if (this.config.service) entry.service = this.config.service;
    if (this.config.enableCaller || level === LogLevel.Debug) setCaller(entry);
    const clonedFields = cloneFields(fields);
    if (clonedFields) entry.fields = clonedFields;
    extractError(entry);
    this.write(entry);
  }

  private write(entry: LogEntry): void {
    if (this.config.structured) {
      this.output(JSON.stringify(entry));
      return;
    }
    let line = `[${entry.timestamp}] [${entry.level}] [${new MCode(entry.code, '').PaddedCode()}] ${entry.message}`;
    if (entry.fields && entry.level === logLevelString(LogLevel.Debug)) line += ` ${JSON.stringify(entry.fields)}`;
    if (entry.error) line += ` error=${JSON.stringify(entry.error)}`;
    this.output(line);
  }
}

export function NewLogger(config: LoggerConfig): Logger {
  return newLogger(config);
}

export function NewEnvLogger(component: string, service: string): Logger {
  const env = processEnv();
  const config: LoggerConfig = {
    component,
    service,
    structured: truthyEnv(env.XALIGO_LOG_STRUCTURED),
    enableCaller: truthyEnv(env.XALIGO_LOG_CALLER),
  };
  if (env.XALIGO_LOG_LEVEL !== undefined) config.level = env.XALIGO_LOG_LEVEL;
  if (env.XALIGO_LOG_OUTPUT !== undefined) config.output = env.XALIGO_LOG_OUTPUT;
  return NewLogger(config);
}

export function newLogger(config: LoggerConfig, output?: LogWriter): Logger {
  return new SharedLogger(config, output);
}

function processEnv(): Record<string, string | undefined> {
  if (typeof process === 'undefined') return {};
  return process.env;
}

function truthyEnv(value: string | undefined): boolean {
  switch ((value ?? '').trim().toLowerCase()) {
    case '1':
    case 'true':
    case 'yes':
    case 'y':
    case 'on':
      return true;
    default:
      return false;
  }
}

function parseLogLevel(level: string | undefined): LogLevel {
  switch ((level ?? '').trim().toUpperCase()) {
    case 'DEBUG':
      return LogLevel.Debug;
    case 'WARN':
    case 'WARNING':
      return LogLevel.Warn;
    case 'ERROR':
      return LogLevel.Error;
    case 'FATAL':
      return LogLevel.Fatal;
    default:
      return LogLevel.Info;
  }
}

function logLevelString(level: LogLevel): string {
  switch (level) {
    case LogLevel.Debug:
      return 'DEBUG';
    case LogLevel.Info:
      return 'INFO';
    case LogLevel.Warn:
      return 'WARN';
    case LogLevel.Error:
      return 'ERROR';
    case LogLevel.Fatal:
      return 'FATAL';
    default:
      return 'UNKNOWN';
  }
}

function openLogOutput(output: string | undefined): LogWriter {
  const target = (output ?? '').trim();
  switch (target) {
    case 'stdout':
      return (line: string) => console.log(line);
    case '':
    case 'stderr':
      return (line: string) => console.error(line);
    default:
      return (line: string) => appendLogFile(target, line);
  }
}

function appendLogFile(path: string, line: string): void {
  if (typeof require !== 'undefined') {
    const fs = require('node:fs') as { appendFileSync(path: string, data: string): void };
    fs.appendFileSync(path, `${line}\n`);
    return;
  }
  console.error(`logger output fallback: ${path}: file output is unavailable`);
  console.log(line);
}

function optionalMessageFor(mcode: MCode, optional: string): string {
  if (optional === '') return mcode.message;
  if (mcode.message === '') return optional;
  return `${mcode.message}: ${optional}`;
}

function cloneFields(fields: Record<string, unknown> | undefined): Record<string, unknown> | undefined {
  if (!fields || Object.keys(fields).length === 0) return undefined;
  return { ...fields };
}

function extractError(entry: LogEntry): void {
  if (!entry.fields || !Object.prototype.hasOwnProperty.call(entry.fields, 'error')) return;
  const value = entry.fields.error;
  if (value instanceof Error) entry.error = value.message;
  else if (typeof value === 'string') entry.error = value;
  else if (value !== undefined && value !== null) entry.error = String(value);
  delete entry.fields.error;
  if (Object.keys(entry.fields).length === 0) delete entry.fields;
}

function setCaller(entry: LogEntry): void {
  const stack = new Error().stack?.split('\n').map((line) => line.trim()) ?? [];
  const caller = stack.find((line) => !line.includes('SharedLogger.') && !line.includes('setCaller'));
  if (!caller) return;
  entry.function = caller;
}