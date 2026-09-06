// Update the entire review library from checked-in, offline AWS research data.
// Normal rendering still preserves author edits. This explicit migration makes
// recoverable backups and records hashes to detect subsequent manual changes.
import fs from 'node:fs';
import path from 'node:path';
import os from 'node:os';
import crypto from 'node:crypto';
import {fileURLToPath} from 'node:url';
import {buildDesigns, orderedFields, diagramOperationRank} from './aws_design_model.mjs';
import {loadBalancerDesigns} from './aws_load_balancer_designs.mjs';
import {nativeLoadBalancerSample, nativeLoadBalancerFlows} from './aws_native_components.mjs';

const root = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../..');
const base = path.join(root, 'docs/src/examples/samples/aws');
const read = file => JSON.parse(fs.readFileSync(path.join(base, file), 'utf8'));
const catalog = read('catalog.json');
const cfn = read('research/cloudformation-models.json');
const api = read('research/api-models.json');
const designs = buildDesigns(catalog, cfn, api);
const check = process.argv.includes('--check');
const update = process.argv.includes('--update');
const selected = process.argv.find(arg => arg.startsWith('--tag='))?.slice(6);
if (selected && !catalog.some(d => d.tag === selected)) throw new Error('Unknown tag: ' + selected);
if (!check && !update) throw new Error('Specify --check or --update. --update backs up edited sources before replacing them.');
const hash = value => crypto.createHash('sha256').update(value).digest('hex');
const xml = value => String(value).replaceAll('&', '&amp;').replaceAll('"', '&quot;').replaceAll('<', '&lt;').replaceAll('>', '&gt;').replaceAll('\n', '&#10;');
const md = value => String(value ?? '').replaceAll('|', '\\|').replaceAll('\n', ' ');
// Break both Latin and CJK text on an explicit, conservative layout budget.
function wrap(value, max = 44) {
  const out = [];
  for (const line of String(value).split('\n')) {
    let row = '', width = 0;
    for (const char of line) {
      const size = /[\u1100-\u115f\u2e80-\ua4cf\uac00-\ud7a3\uf900-\ufaff\ufe10-\ufe6f\uff01-\uff60\uffe0-\uffe6]/u.test(char) ? 2 : 1;
      if (width + size > max) {
        const space = row.lastIndexOf(' ');
        if (space > row.length / 2) { out.push(row.slice(0, space)); row = row.slice(space + 1); width = row.length; }
        else { out.push(row); row = ''; width = 0; }
      }
      row += char; width += size;
    }
    out.push(row);
  }
  return out.join('\n');
}
const rectangle = (id, lines, width = 412, height = 242) => `<rectangle id="${xml(id)}" width="${width}" height="${height}" font-size="16" fill="#f8fafc" stroke="#cbd5e1" title="${xml(wrap(lines.join('\n'), Math.floor((width - 36) / 9)))}" />`;
const typeName = type => type.split('::').at(-1);
const displayName = name => name.replace(/^Amazon/, '').replace(/Configuration/g, 'Config').replace(/Destination/g, 'Dest').replace(/([a-z0-9])([A-Z])/g, '$1 $2');
const displayField = f => `${displayName(f.name)}${f.required ? '*' : ''}: ${/^(String|Boolean|Integer|Long|Double|Timestamp|Json|string|boolean|integer|long|double|timestamp|blob)$/.test(f.type) ? f.type : f.type.startsWith('List<') ? 'list' : f.type.startsWith('Map<') ? 'map' : 'object'}`;
const fieldLines = (fields, limit = 6) => orderedFields(fields).slice(0, limit).map(displayField);
function cardsFor(d, design) {
  if (loadBalancerDesigns[d.tag]) return loadBalancerDesigns[d.tag].cards;
  if (design.primary) {
    const model = cfn.models[design.primary];
    const nested = design.nested.slice(0, 2);
    const operations = design.operations.filter(o => diagramOperationRank(o.name, design.primary) < 2).slice(0, 5);
    return [
      [design.kind === 'resource-schema' ? 'Resource configuration' : 'Service configuration context', typeName(design.primary), ...fieldLines(model.fields, 5)],
      ['Nested configuration', ...nested.flatMap(n => [displayName(n.field) + ' -> object', ...fieldLines(cfn.models[n.model].fields, 2)]), ...(nested.length ? [] : ['No nested object in selected model', 'See scalar fields and API operations'])],
      ['Additional parameters', ...orderedFields(model.fields).slice(5, 11).map(displayField), ...(model.fields.length > 11 ? ['Complete field list in README'] : [])],
      ['Referenced resource types', ...design.related.slice(0, 6).map(typeName), ...(design.related.length ? ['Based on identifier fields', 'References are NOT child ownership'] : ['No peer type selected from ID fields', 'Full service type inventory in README'])],
      ['Operations / lifecycle', ...operations.map(o => o.name), ...(operations.length ? ['Input / output schemas: linked API sheet'] : ['See resource lifecycle / updates', 'CloudFormation properties in README'])],
      ['Icon meaning / interpretation', wrap(d.name, 38), design.kind === 'resource-schema' ? 'Mapped resource model shown above' : 'Service-level schema shown as context', 'Do not infer API fields from icon artwork', 'Solid enclosure = authored membership', 'Type references are not traffic paths'],
    ];
  }
  if (design.apis.length) {
    return design.operations.slice(0, 6).map(o => {
      const operation = api.services[o.service].operations.find(p => p.name === o.name);
      return [o.name, 'API input (not an XAL attribute)', ...fieldLines(operation.input, 5), 'Full input/output schema in API sheet'];
    });
  }
  if (design.guide) return design.guide.concepts.map((concept, i) => [`${i + 1}. ${concept}`, 'Product / workflow concept', 'Use a separate connectable component', 'Not a CloudFormation property', 'See the product-specific guide']);
  if (design.kind === 'category') return design.peers.slice(0, 6).map(tag => ['Service family member', catalog.find(d => d.tag === tag).name, tag, 'Separate service, not a child API field']);
  if (d.group) return [
    ['Containment', d.name, `Scope: ${d.scope}`, 'Children retain individual identities'],
    ['Boundary design', `Border: ${d.group.style}`, `Color: ${d.group.stroke}`, 'Header: icon + measured label'],
    ['Membership / references', 'Enclosure represents diagram membership', 'Connections bind semantic child IDs', 'Not an AWS deployment validator'],
    ['Authoring parameters', ...d.parameters.map(p => p.name + ': ' + p.type), 'id / title / width / height / layout'],
  ];
  return [
    ['Symbol identity', d.name, 'Official architecture-catalog artwork'],
    ['Meaning in this diagram', 'Legend / actor / device / annotation', 'Not a deployable AWS service'],
    ['Relationship', 'Place next to its owning context', 'Add traffic edges only if meaningful'],
    ['Presentation', 'Semantic id + label + size', 'No invented AWS API parameters'],
  ];
}
function sample(d, design) {
  if (loadBalancerDesigns[d.tag]) return nativeLoadBalancerSample(d.tag);
  const lb = loadBalancerDesigns[d.tag];
  const status = design.lifecycle === 'retired' ? 'RETIRED — historical diagrams only / 提供終了' : design.lifecycle === 'renamed-q-developer' ? 'Legacy name — now Amazon Q Developer' : 'Design review / 設計レビュー';
  let component;
  if (d.boundary) component = `<vpc id="boundary-context" title="VPC / 論理境界" width="308" height="180" margin="16"><aws-ec2 id="member" label="Workload" /><${d.tag} id="component" side="${d.boundary.defaultSide}" anchor="0.55" /></vpc>`;
  else if (d.group) component = `<${d.tag} id="component" title="${xml(d.name)}" width="308" height="180" margin="16"><rectangle id="member" width="180" height="64" title="Member / 子要素" /></${d.tag}>`;
  else component = `<${d.tag} id="component" label="${xml(wrap(d.name, 35))}" label-width="340" />`;
  const summary = lb?.summary ?? [`Research: ${design.kind}`, design.primary ?? (design.apis.length ? `API: ${design.apis.join(', ')}` : design.guide ? design.guide.concepts.join(' / ') : `Scope: ${d.scope}`), design.kind === 'service-context' ? 'Configuration context; icon itself may represent a feature.' : 'Functional structure and full reference tables are in README.'];
  const cards = cardsFor(d, design);
  let rows = '', extraHeight = 0;
  for (let i = 0; i < cards.length; i += 3) {
    const group = cards.slice(i, i + 3);
    const height = Math.max(242, ...group.map(card => wrap(card.join('\n'), 41).split('\n').length * 20 + 32));
    extraHeight += height - 242;
    rows += `      <row height="${height}" gap="24">${group.map((card, j) => rectangle(`part-${i + j + 1}`, card, 412, height)).join('\n        ')}</row>\n`;
  }
  return `<!-- AWS functional design library v2. Editable: see README.md and research/designs.json. -->
<!-- Configuration cards describe AWS concepts; their text is not a provisioning command. -->
<xaligo version="2">
<frames>
  <frame id="sample" title="${xml(d.tag)}" width="1500" height="${1160 + extraHeight}">
    <col width="1360" height="${1000 + extraHeight}" margin="32" gap="24">
      <row height="230" gap="32">${component}
        ${rectangle('overview', [lb?.name ?? status, ...summary], 920, 180)}
      </row>
${rows}      ${rectangle('design-legend', ['DESIGN CONTRACT / 図の読み方', '* = required in the source schema, not a required XAL attribute. Values are not deployed to AWS.', 'Configuration / API cards are references, not automatically nested resources or network traffic.', design.lifecycle === 'retired' ? 'AWS lists this product as retired. Keep for historical diagrams; do not treat this as a deployment recommendation.' : 'Availability and region support must be checked before deployment. Research source/version and full parameter tables: README.'], 1284, 150)}
    </col>
  </frame>
</frames>
</xaligo>
`;
}
const fieldTable = fields => `| Field | Type | Required in AWS schema |\n|---|---|---|\n${fields.map(f => `| ${f.documentation ? '[' + md(f.name) + '](' + f.documentation + ')' : '\`' + md(f.name) + '\`'} | \`${md(f.type)}\` | ${f.required ? 'yes' : 'no'} |`).join('\n')}\n`;
function researchNotes(d, design) {
  const lb = loadBalancerDesigns[d.tag];
  let result = `\n<!-- aws-functional-research:start -->\n## 機能調査・構成デザイン（2026-09-06）\n\n分類: \`${design.kind}\`。${design.owner ? 'サービス文脈: [`' + design.owner + '`](../' + design.owner + '/README.md)。' : ''}\n\n`;
  if (design.lifecycle === 'retired') result += `**提供終了済み。** [AWS lifecycle](https://docs.aws.amazon.com/general/latest/gr/full_shutdown_services.html) に基づく歴史的構成図向けの部品です。スキーマやアイコンの残存は現在の提供を意味しません。\n\n`;
  if (design.lifecycle === 'renamed-q-developer') result += '旧 CodeWhisperer のアイコンです。現行名称は Amazon Q Developer。\n\n';
  result += 'サンプルはアイコンと、設定・内包構造・関連リソース・操作を分離したレビューシートです。設定カードは編集可能な `rectangle`、グループは既存の専用タグで実装しています。カードのフィールド名を新しい XAL 属性として受理するわけではありません。専用タグが受理する属性は上の Parameters 表を参照してください。\n\n';
  result += '実線の通信と、設定の参照・同じサービスに属する型一覧を区別します。スキーマの必須項目は AWS 側の仕様であり、図の必須入力ではありません。記載の構成モデル/API は取り込んだ公式資料の範囲であり、全リージョン・全機能の完全性や稼働可否を保証しません。\n\n';
  if (design.kind === 'service-context') result += '**重要:** このアイコンに対応する独立した構成リソースを断定せず、所属サービスの構成モデルを参考表示しています。アイコン名や絵柄から属性・親子関係・通信を推測しません。\n\n';
  if (/^aws-aurora-.*(mariadb|oracle|sql-server|rds)/.test(d.tag)) result += '**カタログ名称の注意:** この既存アイコン名は対応データベースエンジンの保証ではありません。Aurora の MySQL/PostgreSQL 互換性と RDS の他エンジンを混同しないでください。[Aurora guide](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)\n\n';
  if (lb) {
    result += '### コンパクトなネイティブ部品（V2）\n\n現在の XAL/SVG は透かしなしの専用コンポーネントです。枠の左上に公式アイコンとドメインタグ、内部に `aws-listener` のポート/TLS/mTLS バッジを配置します。上の Parameters 表は従来の**アイコン単体形式**の属性です。新しい形式では `id`, `domain`, `view="component"`, `width/height`, 位置・余白・枠スタイルと、1〜32個のリスナーを受理します。\n\nリスナーは `id`, `protocol`, `port` が必須。任意属性は `mtls`, `certificate`, `trust-store`, `target-group`, `backend-tls`, `backend-mtls`, `visible`, `show-title`。サイズは表示内容から自動計算し、TLS/mTLS は各1タグを縦並びにしてONを緑で表示します。既定の「Listener」見出しは `show-title="false"` で非表示にでき、余白も縮みます。TLS はリスナーのプロトコルから決まり、backend属性は転送先の設定情報として保持します（追加バッジなし）。ALB verify にはトラストストアが必要、NLB の mTLS は TCP パススルーで転送先が担当します。設定参照は存在/有効性を検証しません。\n\n[全属性・制約・V1境界](../../../xal/aws-resources.md#native-albnlb-components-v2)。リスナー ID に直接接続できます。ターゲットグループ等はまだ汎用矩形であり、全 AWS サービスがネイティブ実装になったわけではありません。専用設計は `external/engine/src/usc/aws/`、モデルは `external/engine/src/ent/model/aws/` で調整します。\n\n';
    result += '### 専用の構成・接続例\n\n';
    for (const card of lb.cards) result += `#### ${card[0]}\n\n${card.slice(1).map(line => '- ' + line).join('\n')}\n\n`;
    result += '追加の XAL/SVG ペアに通信経路と設定参照を描いています。NLB の TLS 終端と TCP パススルーは別構成であり、同じロードバランサーの同じポートに二つのリスナーを作る例ではありません。ALB のサーバー証明書とクライアント証明書用トラストストアは別概念です。\n\n';
    result += (d.tag.endsWith('-application-load-balancer') ? ['verify'] : ['termination', 'passthrough']).map(mode => `- ${mode}: [XAL](${mode}.xal) / [SVG](${mode}.svg)`).join('\n') + '\n\n';
  }
  if (design.guide) result += `### ソフトウェア・ワークフローの概念\n\n${design.guide.concepts.map(concept => '- ' + concept).join('\n')}\n\nこれらは製品ガイドを基に選んだ図の区画であり、CloudFormation/API の属性名ではありません。\n\n`;
  if (design.primary) {
    result += `### 構成モデル: \`${design.primary}\`\n\n[公式リファレンス](${cfn.models[design.primary].documentation})。全 ${design.fields.length} プロパティを型付きで列挙します（表示カードには主要項目のみ）。\n\n${fieldTable(cfn.models[design.primary].fields)}\n`;
    for (const n of design.nested) result += `#### ${n.field} → \`${n.model}\`\n\n${fieldTable(cfn.models[n.model].fields)}\n`;
    result += `### 関連する構成リソース（${design.resources.length} 型）\n\n同じサービス文脈の型一覧です。すべてがこのアイコンの子リソースという意味ではありません。さらに深い入れ子は [公式スキーマのスナップショット](../research/cloudformation-models.json) に記録しています。\n\n${design.resources.map(name => '- [' + name + '](' + cfn.models[name].documentation + ')').join('\n')}\n\n`;
  }
  if (design.apis.length) result += `### API の操作・パラメータ\n\n${design.apis.map(name => `- [${api.services[name].name}: ${api.services[name].operations.length} 操作の入力・出力一覧](../research/api/${name}.md)（API version ${api.services[name].version}）`).join('\n')}\n\n`;
  result += `### 出典・調査範囲\n\n${[...new Set([...design.sources, ...(lb?.sources ?? [])])].map((url, i) => `- [公式資料 ${i + 1}](${url})`).join('\n')}\n\nCloudFormation 仕様 ${cfn.version}、AWS SDK ${Object.keys(api.services).length} サービスモデルをオフラインで参照。取得日・元データの SHA-256 は [調査マニフェスト](../research/README.md) を参照。API モデル名・フィールド名は仕様から抽出し、説明本文を転載していません。利用可能性は全サービス一律には確認できないため、提供終了の確認がないものも「現在利用可能」と断定していません。\n\n### 次の部品レビュー\n\n- 本アイコンが独立したリソースか、機能・状態・デバイスの記号かを確認する。\n- 詳細カードのうち専用の子タグ・参照属性として実装する範囲を選ぶ。\n- 通信、制御、認証、監視の関係を分け、必要な接続点・配置制約を確認する。\n- 編集後は \`npm run generate:aws-samples -- --render --tag=${d.tag}\`。通常の再描画は XAL/README を上書きしない。\n<!-- aws-functional-research:end -->\n`;
  return result;
}

const generated = new Map();
for (const design of designs) {
  if (design.kind === 'unmapped') throw new Error('Missing explicit research mapping: ' + design.tag);
  const d = catalog.find(d => d.tag === design.tag);
  if (selected && selected !== d.tag) continue;
  let previous = fs.readFileSync(path.join(base, d.tag, 'README.md'), 'utf8').split('<!-- aws-functional-research:start -->')[0].trimEnd();
  if (loadBalancerDesigns[d.tag]) {
    previous = previous.replace('## Parameters\n', '## Parameters — icon-only form\n').replace(/- Implementation: [^\n]+/, '- Implementation: V1/V2 icon-only form; V2 native component with domain header and aws-listener children (current preview).');
  }
  const parameters = '| Parameter | Type | Meaning | Example |\n|---|---|---|---|\n' + d.parameters.map(p => `| \`${p.name}\` | ${p.type === 'enum' ? p.values.map(v => '\`' + v + '\`').join(' / ') : p.type} | ${p.description} | \`${p.example}\` |`).join('\n');
  previous = previous.replace(/\| Parameter \| Type \| Meaning \| Example \|\n[\s\S]*?(?=\n\n## Review notes)/, parameters);
  generated.set(`${d.tag}/sample.xal`, sample(d, design));
  if (loadBalancerDesigns[d.tag]) for (const flow of nativeLoadBalancerFlows(d.tag)) generated.set(`${d.tag}/${flow.mode}.xal`, flow.source);
  generated.set(`${d.tag}/README.md`, previous + '\n' + researchNotes(d, design).replace('../../../xal/aws-resources.md#native-albnlb-components-v2', '../../../../xal/aws-resources.md#native-albnlb-components-v2'));
}
if (!selected) {
  generated.set('research/designs.json', JSON.stringify(designs.map(({fields, ...design}) => design), null, 2) + '\n');
  for (const [name, service] of Object.entries(api.services)) generated.set(`research/api/${name}.md`, `# ${service.name}\n\nAPI version: ${service.version}. [Official AWS SDK source](${service.source}).\n\nThese are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.\n\n${service.operations.map(o => `## ${o.name}\n\nInput:\n\n${fieldTable(o.input)}\nOutput:\n\n${fieldTable(o.output)}`).join('\n')}\n`);
  const counts = Object.fromEntries(Object.entries(Object.groupBy(designs, d => d.kind)).map(([k, values]) => [k, values.length]));
  generated.set('research/README.md', `# AWS functional research snapshot\n\n取得日: 2026-09-06。${designs.length} タグを、構成リソース・サービス文脈・API・製品ガイド・グループ・記号に分類しています。\n\n| Mapping | Tags |\n|---|---|\n${Object.entries(counts).map(([name, count]) => `| ${name} | ${count} |`).join('\n')}\n\n- [CloudFormation specification](${cfn.sourceIndex}): version ${cfn.version}; ${Object.keys(cfn.models).length} resource/property models. [Snapshot](cloudformation-models.json). SHA-256 (uncompressed source): \`${cfn.sha256}\`.\n- [AWS SDK models](${api.source}): ${Object.keys(api.services).length} services. [Snapshot](api-models.json). SHA-256 (downloaded archive): \`${api.sha256}\`. The source branch is mutable; the checked-in snapshot and hash identify this research input.\n- [Per-tag mapping / sources / lifecycle](designs.json). \`resource-schema\` is a model-name mapping, not certification that every feature of the icon has been implemented. \`service-context\` explicitly means the schema belongs to the service, not necessarily the pictured feature.\n\n## Design and coverage boundaries\n\nThe diagrams distinguish configuration, ownership and traffic. All field names/types come from the public model data; manually authored concept cards are explicitly labelled as diagram concepts. This is a review library, not a provisioning engine or an exhaustive service simulator. Region support, quotas, API conditionals and current availability require the linked service guide. Lifecycle notices override stale catalogs and overview pages.\n\nALB/NLB have separate listener, target group, IP/TLS and trust-store designs with connected examples. The shared ELBv2 CloudFormation schema is not used to infer that NLB supports ALB mTLS.\n\n## Reproduction and edit safety\n\n\`npm run generate:aws-designs -- --update\` explicitly updates design sources and makes a temporary backup of overwritten files. \`--check\` checks reproducibility. The original \`generate:aws-samples -- --render\` only refreshes SVGs and preserves component edits. \`sample-hashes.json\` records generated content; use it to audit source changes, not to discard them.\n\nPublic model imports are offline scripts taking previously downloaded files: \`import:aws-cfn-research\` and \`import:aws-api-research\`. They do not access an AWS account or execute SDK code.\n`);
}
let backup;
for (const [name, content] of generated) {
  const target = path.join(base, name);
  const previous = fs.existsSync(target) ? fs.readFileSync(target, 'utf8') : undefined;
  if (previous === content) continue;
  if (check) throw new Error('Stale generated design: ' + name);
  if (previous !== undefined && /(\.xal|\/README\.md)$/.test(name) && !name.startsWith('research/')) {
    backup ??= fs.mkdtempSync(path.join(os.tmpdir(), 'xaligo-aws-design-backup-'));
    const copy = path.join(backup, name);
    fs.mkdirSync(path.dirname(copy), {recursive: true}); fs.copyFileSync(target, copy);
  }
  fs.mkdirSync(path.dirname(target), {recursive: true}); fs.writeFileSync(target, content);
}
if (!selected && !check) fs.writeFileSync(path.join(base, 'research/sample-hashes.json'), JSON.stringify(Object.fromEntries([...generated].filter(([name]) => name.endsWith('.xal')).map(([name, content]) => [name, hash(content)])), null, 2) + '\n');
console.log(`${check ? 'Verified' : 'Updated'} ${selected ? 1 : catalog.length} functional AWS designs${backup ? '; previous sources saved in ' + backup : ''}.`);
