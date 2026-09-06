# AWS functional research snapshot

取得日: 2026-09-06。877 タグを、構成リソース・サービス文脈・API・製品ガイド・グループ・記号に分類しています。

| Mapping | Tags |
|---|---|
| resource-schema | 88 |
| group | 9 |
| product-guide | 65 |
| symbol | 90 |
| service-context | 562 |
| api-context | 38 |
| category | 25 |

- [CloudFormation specification](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/cfn-resource-specification.html): version 263.0.0; 11492 resource/property models. [Snapshot](cloudformation-models.json). SHA-256 (uncompressed source): `c95f3e96b333cd5adbfde65b31204669b6c75a8112a0a51fc20722c6a1cf7831`.
- [AWS SDK models](https://github.com/boto/botocore): 431 services. [Snapshot](api-models.json). SHA-256 (downloaded archive): `76284b2144d1d2e8625cf604ac375ce979d646c3aa7005d73e1075a58586019a`. The source branch is mutable; the checked-in snapshot and hash identify this research input.
- [Per-tag mapping / sources / lifecycle](designs.json). `resource-schema` is a model-name mapping, not certification that every feature of the icon has been implemented. `service-context` explicitly means the schema belongs to the service, not necessarily the pictured feature.

## Design and coverage boundaries

The diagrams distinguish configuration, ownership and traffic. All field names/types come from the public model data; manually authored concept cards are explicitly labelled as diagram concepts. This is a review library, not a provisioning engine or an exhaustive service simulator. Region support, quotas, API conditionals and current availability require the linked service guide. Lifecycle notices override stale catalogs and overview pages.

ALB/NLB have separate listener, target group, IP/TLS and trust-store designs with connected examples. The shared ELBv2 CloudFormation schema is not used to infer that NLB supports ALB mTLS.

## Reproduction and edit safety

`npm run generate:aws-designs -- --update` explicitly updates design sources and makes a temporary backup of overwritten files. `--check` checks reproducibility. The original `generate:aws-samples -- --render` only refreshes SVGs and preserves component edits. `sample-hashes.json` records generated content; use it to audit source changes, not to discard them.

Public model imports are offline scripts taking previously downloaded files: `import:aws-cfn-research` and `import:aws-api-research`. They do not access an AWS account or execute SDK code.
