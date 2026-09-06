// Import factual model names/fields from AWS's public CloudFormation spec.
// No account access, deployment, or source/dependency metadata transmission.
import fs from 'node:fs';
import path from 'node:path';
import zlib from 'node:zlib';
import crypto from 'node:crypto';
import { fileURLToPath } from 'node:url';

const root=path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../..');
const file=process.argv[2];
if (!file) throw new Error('Usage: node scripts/tool/import_aws_research.mjs <downloaded CloudFormation spec>');
let bytes=fs.readFileSync(file);
if (bytes[0]===0x1f && bytes[1]===0x8b) bytes=zlib.gunzipSync(bytes);
const input=JSON.parse(bytes);
const typeName=(name,p)=>p.PrimitiveType ?? (p.Type==='List'||p.Type==='Map' ? `${p.Type}<${p.PrimitiveItemType ?? p.ItemType}>` : p.Type);
const models={};
for (const [name, model] of Object.entries({...input.ResourceTypes,...input.PropertyTypes}).sort(([a],[b])=>a.localeCompare(b,'en'))) {
  models[name]={documentation:model.Documentation?.replace(/^http:/,'https:'),fields:Object.entries(model.Properties??{}).sort(([a],[b])=>a.localeCompare(b,'en')).map(([field,p])=>({name:field,type:typeName(name,p),required:p.Required===true,documentation:p.Documentation?.replace(/^http:/,'https:')}))};
}
const snapshot={source:'https://d1uauaxba7bl26.cloudfront.net/latest/gzip/CloudFormationResourceSpecification.json',sourceIndex:'https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/cfn-resource-specification.html',version:input.ResourceSpecificationVersion,retrieved:'2026-09-06',sha256:crypto.createHash('sha256').update(bytes).digest('hex'),models};
const directory=path.join(root,'docs/src/examples/samples/aws/research');
fs.mkdirSync(directory,{recursive:true});
fs.writeFileSync(path.join(directory,'cloudformation-models.json'),JSON.stringify(snapshot)+'\n');
console.log(`Imported ${Object.keys(models).length} models from AWS resource specification ${snapshot.version}.`);
