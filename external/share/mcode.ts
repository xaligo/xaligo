export class MCode {
  constructor(
    public readonly code: string,
    public readonly message: string,
  ) {}

  get Code(): string {
    return this.code;
  }

  get Message(): string {
    return this.message;
  }

  PaddedCode(): string {
    if (this.code.length >= maxCodeLength) return this.code;
    return this.code + ' '.repeat(maxCodeLength - this.code.length);
  }
}

let maxCodeLength = 0;

export function NewMCode(code: string, message: string): MCode {
  const mcode = new MCode(code, message);
  RegisterMCodes(mcode);
  return mcode;
}

export function Mcode(mcode: MCode): MCode {
  return mcode;
}

export function RegisterMCodes(...mcodes: MCode[]): void {
  for (const mcode of mcodes) {
    if (mcode.code.length > maxCodeLength) maxCodeLength = mcode.code.length;
  }
}

export function GetMaxCodeLength(): number {
  return maxCodeLength;
}

export const MSYS1 = NewMCode('MSYS1', 'System start');
export const MSYS2 = NewMCode('MSYS2', 'System error');
export const MLOG1 = NewMCode('MLOG1', 'Logger created');
export const MLOG2 = NewMCode('MLOG2', 'Logger output fallback');