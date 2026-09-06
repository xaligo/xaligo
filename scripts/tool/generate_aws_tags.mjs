// Generate declarative profiles from the bundled AWS catalog. No network I/O.
import fs from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { execFileSync } from 'node:child_process';
import { loadBalancerParameters } from './aws_load_balancer_parameters.mjs';

const root = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../..');
const check = process.argv.includes('--check');
const csv = fs.readFileSync(path.join(root, 'etc/resources/aws/service-catalog.csv'), 'utf8');
// RFC 4180 fields (including quoted commas/newlines and doubled quotes).
const rows = []; let row = [], field = '', quoted = false;
for (let i = 0; i < csv.length; i++) {
  const c = csv[i];
  if (c === '"') { if (quoted && csv[i + 1] === '"') { field += '"'; i++; } else quoted = !quoted; }
  else if (!quoted && (c === ',' || c === '\n')) { row.push(field.replace(/\r$/, '')); field = ''; if (c === '\n') { rows.push(row); row = []; } }
  else field += c;
}
if (field || row.length) { row.push(field); rows.push(row); }
const entries = rows.slice(1).filter(r => /\/(Architecture-Service-Icons|Resource-Icons|Category-Icons|Architecture-Group-Icons)\//.test(r[4] ?? '')).map(([id, category, name, file, source]) => ({ id: Number(id), category, name, file, source }));
const slug = value => value.toLowerCase().replace(/[^a-z0-9]+/g, '-').replace(/^-|-$/g, '');
const short = value => value.replace(/^(Amazon|AWS) /, '').replace(/^Simple Storage Service\b/, 'S3').replace(/^Simple Notification Service\b/, 'SNS').replace(/^Simple Queue Service\b/, 'SQS');
const parameter = (name, description, example, type = 'text', values = []) => ({ name, description, example, type, values });
const enums = (name, description, values) => parameter(name, description, values[0], 'enum', values);
const groupRows = [
  ['aws-cloud', 'AWS Cloud', '#000000', 'solid', 'AWS-Cloud-logo_32.svg', 'global'],
  ['aws-cloud-alt', 'AWS Cloud', '#000000', 'solid', 'AWS-Cloud_32.svg', 'global'],
  ['aws-cloud-dark', 'AWS Cloud (dark icon)', '#000000', 'solid', 'AWS-Cloud-logo_32_Dark.svg', 'global'],
  ['aws-cloud-alt-dark', 'AWS Cloud (dark icon)', '#000000', 'solid', 'AWS-Cloud_32_Dark.svg', 'global'],
  ['region', 'Region', '#00A1C9', 'dashed', 'Region_32.svg', 'region'],
  ['availability-zone', 'Availability Zone', '#00A1C9', 'dashed', '', 'availability-zone'],
  ['security-group', 'Security group', '#CC0000', 'dashed', '', 'vpc'],
  ['auto-scaling-group', 'Auto Scaling group', '#E7601B', 'dashed', 'Auto-Scaling-group_32.svg', 'region'],
  ['vpc', 'VPC', '#8C4FFF', 'solid', 'Virtual-private-cloud-VPC_32.svg', 'vpc'],
  ['private-subnet', 'Private subnet', '#00A1C9', 'solid', 'Private-subnet_32.svg', 'subnet'],
  ['public-subnet', 'Public subnet', '#3F8624', 'solid', 'Public-subnet_32.svg', 'subnet'],
  ['server-contents', 'Server contents', '#7A7C7F', 'solid', 'Server-contents_32.svg', 'host'],
  ['corporate-data-center', 'Corporate data center', '#7A7C7F', 'solid', 'Corporate-data-center_32.svg', 'external'],
  ['ec2-instance-contents', 'EC2 instance contents', '#E7601B', 'solid', 'EC2-instance-contents_32.svg', 'host'],
  ['spot-fleet', 'Spot Fleet', '#E7601B', 'solid', 'Spot-Fleet_32.svg', 'region'],
  ['aws-account', 'AWS account', '#E7008A', 'solid', 'AWS-Account_32.svg', 'account'],
  ['aws-iot-greengrass-deployment', 'AWS IoT Greengrass Deployment', '#3F8624', 'solid', 'AWS-IoT-Greengrass-Deployment_32.svg', 'edge'],
  ['aws-iot-greengrass', 'AWS IoT Greengrass', '#3F8624', 'solid', '', 'edge'],
  ['elastic-beanstalk-container', 'Elastic Beanstalk container', '#E7601B', 'solid', '', 'region'],
  ['aws-step-functions-workflow', 'AWS Step Functions workflow', '#E7008A', 'solid', '', 'region'],
  ['generic-group', 'Generic group', '#AAB7B8', 'dashed', '', 'logical'],
];
const definitions = groupRows.map(([tag, name, stroke, style, icon, scope]) => {
  const entry = entries.find(e => e.file === icon);
  let parameters = [];
  if (['vpc', 'private-subnet', 'public-subnet'].includes(tag)) parameters = [parameter('cidr', 'IPv4/IPv6 CIDR annotation', tag === 'vpc' ? '10.0.0.0/16' : '10.0.1.0/24', 'cidr')];
  if (tag === 'availability-zone') parameters = [parameter('zone', 'Availability Zone name', 'ap-northeast-1a')];
  if (tag === 'region') parameters = [parameter('region-name', 'AWS Region name', 'ap-northeast-1')];
  if (tag === 'aws-account') parameters = [parameter('account-id', 'Account identifier annotation', '123456789012')];
  if (['auto-scaling-group', 'spot-fleet'].includes(tag)) parameters = [parameter('desired-capacity', 'Desired capacity annotation', '2', 'integer')];
  if (tag === 'security-group') parameters = [parameter('ingress', 'Inbound rule summary', 'HTTPS 443'), parameter('egress', 'Outbound rule summary', 'TCP 443')];
  return { tag, name, kind: 'group', category: 'Group', scope, description: 'Logical containment boundary; children remain individually connectable. The header uses the AWS group color and icon.', catalogID: entry?.id ?? 0, catalogIDs: entry ? [entry.id] : [], parameters, group: { stroke, style, icon, width: tag === 'generic-group' ? 1 : 2 } };
});
const occupied = new Set(definitions.map(d => d.tag));
const buckets = new Map();
for (const entry of entries.filter(e => e.category !== 'Group')) {
  const kind = entry.source.includes('/Resource-Icons/') ? 'resource' : entry.source.includes('/Category-Icons/') ? 'category' : 'service';
  const key = kind + ':' + entry.file.replace(/_(16|32|48|64)\.svg$/, '');
  if (!buckets.has(key)) buckets.set(key, []);
  buckets.get(key).push({ ...entry, kind });
}
for (const variants of [...buckets.values()].sort((a, b) => a[0].file.localeCompare(b[0].file, 'en'))) {
  const entry = variants.find(e => /_48\.svg$/.test(e.file)) ?? variants.find(e => /_64\.svg$/.test(e.file)) ?? variants[0];
  let tag = 'aws-' + (entry.kind === 'category' ? 'category-' : '') + slug(short(entry.name));
  if (entry.id === 1579) tag = 'vpc-endpoint';
  if (occupied.has(tag)) tag += '-' + entry.kind;
  if (occupied.has(tag)) throw new Error('Tag collision: ' + tag);
  occupied.add(tag);
  const family = entry.category === 'Resource' ? entry.source.split('/')[5].replace(/^Res_/, '').replaceAll('-', ' ') : entry.category;
  // Scope is a diagram recommendation, not a placement validator. Generic
  // feature/policy/user icons must not be described as deployable services.
  let scope = entry.kind === 'service' ? 'service' : 'logical';
  let description = entry.kind === 'category' ? 'Category marker, not a deployable resource.' : `AWS ${entry.kind} icon. Use a label and explicit annotations to describe its role; scope is selected by the author.`;
  let parameters = [parameter('role', 'Architectural role annotation', entry.kind === 'category' ? 'Service family' : 'Application component')];
  if (/Compute|Containers/.test(family)) parameters = [parameter('workload', 'Workload or process annotation', 'Web application')];
  if (/Storage/.test(family)) parameters = [parameter('data', 'Stored data annotation', 'Application data')];
  if (/Database/.test(family)) parameters = [parameter('data-model', 'Data model annotation', 'Application records')];
  if (/Networking/.test(family)) parameters = [parameter('traffic', 'Traffic/protocol annotation', 'HTTPS')];
  if (/Security/.test(family)) parameters = [parameter('protects', 'Protected resource or policy annotation', 'Application access')];
  if (/App Integration/.test(family)) parameters = [parameter('message', 'Message/event annotation', 'Order events')];
  if (/Analytics|Artificial Intelligence/.test(family)) parameters = [parameter('input', 'Input data annotation', 'Application events'), parameter('output', 'Output data annotation', 'Insights')];
  if (/Management|Developer Tools/.test(family)) parameters = [parameter('target', 'Managed resource or build target', 'Application')];
  if (/Internet of Things/.test(family)) parameters = [parameter('device', 'Device/workload annotation', 'Sensor')];
  if (/^Amazon EC2(?: Instance)?$/.test(entry.name)) { scope = 'subnet'; parameters = [parameter('instance-type', 'EC2 instance type annotation', 't3.micro'), enums('state', 'Diagram instance state', ['running', 'stopped', 'pending', 'terminated'])]; }
  if (entry.name === 'AWS Lambda' || entry.name === 'AWS Lambda Lambda Function') { scope = 'region'; parameters = [parameter('runtime', 'Runtime annotation (no version validation)', 'provided.al2023'), parameter('memory-mb', 'Configured memory annotation', '512', 'integer')]; }
  if (/^Amazon (RDS|Aurora)$/.test(entry.name)) { scope = 'region'; parameters = [parameter('engine', 'Database engine annotation', 'PostgreSQL'), parameter('multi-az', 'Multi-AZ annotation; does not create replicas', 'true', 'boolean')]; }
  if (/^Amazon DynamoDB(?: Table)?$/.test(entry.name)) { scope = 'region'; parameters = [enums('billing-mode', 'Billing mode annotation', ['on-demand', 'provisioned']), parameter('partition-key', 'Partition key annotation', 'id')]; }
  if (/^Amazon Simple Storage Service(?: Bucket)?$/.test(entry.name)) { scope = 'region'; parameters = [parameter('bucket-name', 'Bucket name annotation', 'example-assets'), parameter('versioning', 'Versioning annotation', 'true', 'boolean')]; description = 'Regional storage service/bucket. Access via a VPC endpoint does not place the bucket inside a subnet.'; }
  if (/^Amazon Simple Queue Service(?: Queue)?$/.test(entry.name)) { scope = 'region'; parameters = [enums('queue-type', 'Queue type annotation', ['standard', 'fifo'])]; }
  if (/^Amazon Simple Notification Service(?: Topic)?$/.test(entry.name)) { scope = 'region'; parameters = [enums('topic-type', 'Topic type annotation', ['standard', 'fifo'])]; }
  if (/^Amazon VPC NAT Gateway$/.test(entry.name)) { scope = 'subnet'; parameters = [enums('connectivity', 'NAT connectivity annotation', ['public', 'private'])]; description = 'Normal interior resource, not a VPC-boundary attachment. This sample models a subnet-scoped NAT gateway.'; }
  if (/^Amazon VPC (Elastic Network Interface|Elastic Network Adapter)$/.test(entry.name)) scope = 'subnet';
  if (/^Amazon VPC Customer Gateway$/.test(entry.name)) scope = 'external';
  if (/^Amazon (CloudFront|Route 53)$/.test(entry.name) || entry.name === 'AWS Identity and Access Management') scope = 'global';
  let boundary;
  parameters = loadBalancerParameters(tag) ?? parameters;
  if ([1579, 1581, 1590].includes(entry.id)) {
    scope = 'vpc-boundary'; boundary = { parentTag: 'vpc', defaultSide: entry.id === 1581 ? 'top' : 'right', defaultSize: 48 };
    parameters = entry.id === 1579 ? [enums('endpoint-type', 'Logical endpoint type annotation', ['interface', 'gateway', 'gateway-load-balancer']), parameter('service-name', 'Endpoint service annotation', 'com.amazonaws.ap-northeast-1.s3')] : [parameter('attachment', 'Attachment/route annotation', 'Main VPC')];
    description = 'Icon-only logical VPC boundary attachment. side/anchor/offset move the icon along the border, outside normal flow. Endpoint placement is a diagram convention, not the physical location of endpoint network interfaces.';
  }
  definitions.push({ tag, name: entry.name, kind: entry.kind, category: family, scope, description, catalogID: entry.id, catalogIDs: variants.map(e => e.id).sort((a,b) => a-b), parameters, boundary });
}
definitions.sort((a,b) => a.tag.localeCompare(b.tag, 'en'));
const str = JSON.stringify;
const goStrings = values => `[]string{${values.map(str).join(',')}}`;
let go = '// Code generated by npm run generate:aws-tags; DO NOT EDIT.\npackage aws\n\nvar definitions = []Definition{\n';
for (const d of definitions) {
  go += `{Tag:${str(d.tag)},Name:${str(d.name)},Kind:${str(d.kind)},Category:${str(d.category)},Scope:${str(d.scope)},Description:${str(d.description)},CatalogID:${d.catalogID},CatalogIDs:[]int{${d.catalogIDs.join(',')}},Parameters:[]Parameter{${d.parameters.map(p => `{Name:${str(p.name)},Type:${str(p.type)},Description:${str(p.description)},Example:${str(p.example)},Values:${goStrings(p.values)}}`).join(',')}},`;
  if (d.group) go += `Group:&GroupStyle{Stroke:${str(d.group.stroke)},Style:${str(d.group.style)},Icon:${str(d.group.icon)},Width:${d.group.width}},`;
  if (d.boundary) go += `Boundary:&BoundaryAttachment{Tag:${str(d.tag)},ParentTag:"vpc",CatalogID:${d.catalogID},DefaultSide:${str(d.boundary.defaultSide)},DefaultSize:48},`;
  go += '},\n';
}
go += '}\n';
const generated = new Map([
  ['internal/core/profiles/aws/catalog_generated.go', execFileSync('gofmt', { input: go, encoding: 'utf8' })],
  ['docs/src/examples/samples/aws/catalog.json', JSON.stringify(definitions, null, 2) + '\n'],
]);
for (const [file, content] of generated) {
  const target = path.join(root, file);
  if (check) { if (!fs.existsSync(target) || fs.readFileSync(target, 'utf8') !== content) throw new Error('Stale generated file: ' + file); }
  else { fs.mkdirSync(path.dirname(target), { recursive: true }); fs.writeFileSync(target, content); }
}
console.log(`${definitions.length} AWS tags cover ${entries.length} catalog entries (${definitions.filter(d => d.group).length} groups).`);
