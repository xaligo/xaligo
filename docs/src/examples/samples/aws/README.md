# AWS component catalog

877 dedicated tags cover all 1,875 bundled AWS catalog entries, including size variants, category/resource icons, and 21 group tags. Tabler/Yamaha icons are outside this AWS catalog.

Each directory contains an editable `sample.xal`, its rendered `sample.svg`, and parameter/design notes. Numeric `<item id="…">` remains supported. New resource tags use `aws-…`; existing group tags and `vpc-endpoint` retain their names. Same-name service/resource icons are distinct (`-resource` suffix); catalog size variants share one tag.

Start with [VPC endpoint](vpc-endpoint/README.md), [Internet Gateway](aws-vpc-internet-gateway/README.md), [NAT Gateway](aws-vpc-nat-gateway/README.md), [EC2](aws-ec2/README.md), [S3 bucket](aws-s3-bucket/README.md), and [VPC](vpc/README.md).

Functional parameters are diagram annotations, not provisioning commands or a complete AWS API schema. Scope recommendations are deliberately non-enforcing for logical service/feature icons. VPC endpoint border placement is an authoring convention: [interface endpoints create network interfaces in subnets](https://docs.aws.amazon.com/vpc/latest/privatelink/concepts.html), whereas [gateway endpoints are associated with route tables](https://docs.aws.amazon.com/vpc/latest/privatelink/gateway-endpoints.html). [Internet gateways attach to VPCs](https://docs.aws.amazon.com/vpc/latest/userguide/working-with-igw.html); the [subnet-scoped NAT example](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat.html) stays inside a subnet.

Regenerate the registry with `npm run generate:aws-tags`. Bootstrap missing examples with `npm run generate:aws-samples`. Refresh SVGs from current sources with `npm run generate:aws-samples -- --render` (optionally `--tag=vpc-endpoint`). Existing XAL/README files are never overwritten. `catalog.json` is the generated tag/parameter/asset manifest.


## Functional design review

全877タグの機能レビューシート、型付きパラメータ表、関連リソース、API一覧、出典を各READMEへ追加しました。[調査範囲とソース](research/README.md)を参照してください。サービス文脈の設定と、アイコン自身の独立したリソース仕様を明示的に区別しています。

- [ALB: listeners / rules / target groups / mTLS trust store](aws-elastic-load-balancing-application-load-balancer/README.md) · [mTLS接続図](aws-elastic-load-balancing-application-load-balancer/verify.svg)
- [NLB: listeners / addressing / TLS / targets](aws-elastic-load-balancing-network-load-balancer/README.md) · [TLS終端](aws-elastic-load-balancing-network-load-balancer/termination.svg) · [TCPパススルー](aws-elastic-load-balancing-network-load-balancer/passthrough.svg)

通常の描画は編集内容を保持します。全デザインを意図的に更新する場合だけ `npm run generate:aws-designs -- --update` を使用してください（上書き前のソースは一時ディレクトリへ退避）。詳細カードは編集可能な既存XAL部品であり、掲載したAWS APIフィールドをすべて新しいXAL属性として実装したという意味ではありません。

| Tag | Component | Scope | Preview |
|---|---|---|---|
| [`auto-scaling-group`](auto-scaling-group/README.md) | Auto Scaling group | region | [SVG](auto-scaling-group/sample.svg) |
| [`availability-zone`](availability-zone/README.md) | Availability Zone | availability-zone | [SVG](availability-zone/sample.svg) |
| [`aws-account`](aws-account/README.md) | AWS account | account | [SVG](aws-account/sample.svg) |
| [`aws-activate`](aws-activate/README.md) | AWS Activate | service | [SVG](aws-activate/sample.svg) |
| [`aws-alert-48-dark`](aws-alert-48-dark/README.md) | Alert 48 Dark | logical | [SVG](aws-alert-48-dark/sample.svg) |
| [`aws-alert-48-light`](aws-alert-48-light/README.md) | Alert 48 Light | logical | [SVG](aws-alert-48-light/sample.svg) |
| [`aws-alexa-for-business`](aws-alexa-for-business/README.md) | Alexa For Business | service | [SVG](aws-alexa-for-business/sample.svg) |
| [`aws-amplify`](aws-amplify/README.md) | AWS Amplify | service | [SVG](aws-amplify/sample.svg) |
| [`aws-amplify-aws-amplify-studio`](aws-amplify-aws-amplify-studio/README.md) | AWS Amplify AWS Amplify Studio | logical | [SVG](aws-amplify-aws-amplify-studio/sample.svg) |
| [`aws-apache-mxnet-on-aws`](aws-apache-mxnet-on-aws/README.md) | Apache MXNet on AWS | service | [SVG](aws-apache-mxnet-on-aws/sample.svg) |
| [`aws-api-gateway`](aws-api-gateway/README.md) | Amazon API Gateway | service | [SVG](aws-api-gateway/sample.svg) |
| [`aws-api-gateway-endpoint`](aws-api-gateway-endpoint/README.md) | Amazon API Gateway Endpoint | logical | [SVG](aws-api-gateway-endpoint/sample.svg) |
| [`aws-app-mesh`](aws-app-mesh/README.md) | AWS App Mesh | service | [SVG](aws-app-mesh/sample.svg) |
| [`aws-app-mesh-mesh`](aws-app-mesh-mesh/README.md) | AWS App Mesh Mesh | logical | [SVG](aws-app-mesh-mesh/sample.svg) |
| [`aws-app-mesh-virtual-gateway`](aws-app-mesh-virtual-gateway/README.md) | AWS App Mesh Virtual Gateway | logical | [SVG](aws-app-mesh-virtual-gateway/sample.svg) |
| [`aws-app-mesh-virtual-node`](aws-app-mesh-virtual-node/README.md) | AWS App Mesh Virtual Node | logical | [SVG](aws-app-mesh-virtual-node/sample.svg) |
| [`aws-app-mesh-virtual-router`](aws-app-mesh-virtual-router/README.md) | AWS App Mesh Virtual Router | logical | [SVG](aws-app-mesh-virtual-router/sample.svg) |
| [`aws-app-mesh-virtual-service`](aws-app-mesh-virtual-service/README.md) | AWS App Mesh Virtual Service | logical | [SVG](aws-app-mesh-virtual-service/sample.svg) |
| [`aws-app-runner`](aws-app-runner/README.md) | AWS App Runner | service | [SVG](aws-app-runner/sample.svg) |
| [`aws-app-studio`](aws-app-studio/README.md) | AWS App Studio | service | [SVG](aws-app-studio/sample.svg) |
| [`aws-appconfig`](aws-appconfig/README.md) | AWS AppConfig | service | [SVG](aws-appconfig/sample.svg) |
| [`aws-appfabric`](aws-appfabric/README.md) | AWS AppFabric | service | [SVG](aws-appfabric/sample.svg) |
| [`aws-appflow`](aws-appflow/README.md) | Amazon AppFlow | service | [SVG](aws-appflow/sample.svg) |
| [`aws-application-auto-scaling`](aws-application-auto-scaling/README.md) | AWS Application Auto Scaling | service | [SVG](aws-application-auto-scaling/sample.svg) |
| [`aws-application-discovery-service`](aws-application-discovery-service/README.md) | AWS Application Discovery Service | service | [SVG](aws-application-discovery-service/sample.svg) |
| [`aws-application-discovery-service-aws-agentless-collector`](aws-application-discovery-service-aws-agentless-collector/README.md) | AWS Application Discovery Service AWS Agentless Collector | logical | [SVG](aws-application-discovery-service-aws-agentless-collector/sample.svg) |
| [`aws-application-discovery-service-aws-discovery-agent`](aws-application-discovery-service-aws-discovery-agent/README.md) | AWS Application Discovery Service AWS Discovery Agent | logical | [SVG](aws-application-discovery-service-aws-discovery-agent/sample.svg) |
| [`aws-application-discovery-service-migration-evaluator-collector`](aws-application-discovery-service-migration-evaluator-collector/README.md) | AWS Application Discovery Service Migration Evaluator Collector | logical | [SVG](aws-application-discovery-service-migration-evaluator-collector/sample.svg) |
| [`aws-application-migration-service`](aws-application-migration-service/README.md) | AWS Application Migration Service | service | [SVG](aws-application-migration-service/sample.svg) |
| [`aws-application-recovery-controller`](aws-application-recovery-controller/README.md) | Amazon Application Recovery Controller | service | [SVG](aws-application-recovery-controller/sample.svg) |
| [`aws-appstream-2`](aws-appstream-2/README.md) | Amazon AppStream 2 | service | [SVG](aws-appstream-2/sample.svg) |
| [`aws-appsync`](aws-appsync/README.md) | AWS AppSync | service | [SVG](aws-appsync/sample.svg) |
| [`aws-artifact`](aws-artifact/README.md) | AWS Artifact | service | [SVG](aws-artifact/sample.svg) |
| [`aws-athena`](aws-athena/README.md) | Amazon Athena | service | [SVG](aws-athena/sample.svg) |
| [`aws-athena-data-source-connectors`](aws-athena-data-source-connectors/README.md) | Amazon Athena Data Source Connectors | logical | [SVG](aws-athena-data-source-connectors/sample.svg) |
| [`aws-audit-manager`](aws-audit-manager/README.md) | AWS Audit Manager | service | [SVG](aws-audit-manager/sample.svg) |
| [`aws-augmented-ai-a2i`](aws-augmented-ai-a2i/README.md) | Amazon Augmented AI A2I | service | [SVG](aws-augmented-ai-a2i/sample.svg) |
| [`aws-aurora`](aws-aurora/README.md) | Amazon Aurora | region | [SVG](aws-aurora/sample.svg) |
| [`aws-aurora-amazon-aurora-instance-alternate`](aws-aurora-amazon-aurora-instance-alternate/README.md) | Amazon Aurora Amazon Aurora Instance alternate | logical | [SVG](aws-aurora-amazon-aurora-instance-alternate/sample.svg) |
| [`aws-aurora-amazon-rds-instance`](aws-aurora-amazon-rds-instance/README.md) | Amazon Aurora Amazon RDS Instance | logical | [SVG](aws-aurora-amazon-rds-instance/sample.svg) |
| [`aws-aurora-amazon-rds-instance-aternate`](aws-aurora-amazon-rds-instance-aternate/README.md) | Amazon Aurora Amazon RDS Instance Aternate | logical | [SVG](aws-aurora-amazon-rds-instance-aternate/sample.svg) |
| [`aws-aurora-instance`](aws-aurora-instance/README.md) | Amazon Aurora Instance | logical | [SVG](aws-aurora-instance/sample.svg) |
| [`aws-aurora-mariadb-instance`](aws-aurora-mariadb-instance/README.md) | Amazon Aurora MariaDB Instance | logical | [SVG](aws-aurora-mariadb-instance/sample.svg) |
| [`aws-aurora-mariadb-instance-alternate`](aws-aurora-mariadb-instance-alternate/README.md) | Amazon Aurora MariaDB Instance Alternate | logical | [SVG](aws-aurora-mariadb-instance-alternate/sample.svg) |
| [`aws-aurora-mysql-instance`](aws-aurora-mysql-instance/README.md) | Amazon Aurora MySQL Instance | logical | [SVG](aws-aurora-mysql-instance/sample.svg) |
| [`aws-aurora-mysql-instance-alternate`](aws-aurora-mysql-instance-alternate/README.md) | Amazon Aurora MySQL Instance Alternate | logical | [SVG](aws-aurora-mysql-instance-alternate/sample.svg) |
| [`aws-aurora-oracle-instance`](aws-aurora-oracle-instance/README.md) | Amazon Aurora Oracle Instance | logical | [SVG](aws-aurora-oracle-instance/sample.svg) |
| [`aws-aurora-oracle-instance-alternate`](aws-aurora-oracle-instance-alternate/README.md) | Amazon Aurora Oracle Instance Alternate | logical | [SVG](aws-aurora-oracle-instance-alternate/sample.svg) |
| [`aws-aurora-piops-instance`](aws-aurora-piops-instance/README.md) | Amazon Aurora PIOPS Instance | logical | [SVG](aws-aurora-piops-instance/sample.svg) |
| [`aws-aurora-postgresql-instance`](aws-aurora-postgresql-instance/README.md) | Amazon Aurora PostgreSQL Instance | logical | [SVG](aws-aurora-postgresql-instance/sample.svg) |
| [`aws-aurora-postgresql-instance-alternate`](aws-aurora-postgresql-instance-alternate/README.md) | Amazon Aurora PostgreSQL Instance Alternate | logical | [SVG](aws-aurora-postgresql-instance-alternate/sample.svg) |
| [`aws-aurora-sql-server-instance`](aws-aurora-sql-server-instance/README.md) | Amazon Aurora SQL Server Instance | logical | [SVG](aws-aurora-sql-server-instance/sample.svg) |
| [`aws-aurora-sql-server-instance-alternate`](aws-aurora-sql-server-instance-alternate/README.md) | Amazon Aurora SQL Server Instance Alternate | logical | [SVG](aws-aurora-sql-server-instance-alternate/sample.svg) |
| [`aws-aurora-trusted-language-extensions-for-postgresql`](aws-aurora-trusted-language-extensions-for-postgresql/README.md) | Amazon Aurora Trusted Language Extensions for PostgreSQL | logical | [SVG](aws-aurora-trusted-language-extensions-for-postgresql/sample.svg) |
| [`aws-authenticated-user-48-dark`](aws-authenticated-user-48-dark/README.md) | Authenticated User 48 Dark | logical | [SVG](aws-authenticated-user-48-dark/sample.svg) |
| [`aws-authenticated-user-48-light`](aws-authenticated-user-48-light/README.md) | Authenticated User 48 Light | logical | [SVG](aws-authenticated-user-48-light/sample.svg) |
| [`aws-auto-scaling`](aws-auto-scaling/README.md) | AWS Auto Scaling | service | [SVG](aws-auto-scaling/sample.svg) |
| [`aws-b2b-data-interchange`](aws-b2b-data-interchange/README.md) | AWS B2B Data Interchange | service | [SVG](aws-b2b-data-interchange/sample.svg) |
| [`aws-backint-agent`](aws-backint-agent/README.md) | AWS Backint Agent | service | [SVG](aws-backint-agent/sample.svg) |
| [`aws-backup`](aws-backup/README.md) | AWS Backup | service | [SVG](aws-backup/sample.svg) |
| [`aws-backup-audit-manager`](aws-backup-audit-manager/README.md) | AWS Backup Audit Manager | logical | [SVG](aws-backup-audit-manager/sample.svg) |
| [`aws-backup-aws-backup-for-aws-cloudformation`](aws-backup-aws-backup-for-aws-cloudformation/README.md) | AWS Backup AWS Backup for AWS CloudFormation | logical | [SVG](aws-backup-aws-backup-for-aws-cloudformation/sample.svg) |
| [`aws-backup-aws-backup-support-for-amazon-fsx-for-netapp-ontap`](aws-backup-aws-backup-support-for-amazon-fsx-for-netapp-ontap/README.md) | AWS Backup AWS Backup support for Amazon FSx for NetApp ONTAP | logical | [SVG](aws-backup-aws-backup-support-for-amazon-fsx-for-netapp-ontap/sample.svg) |
| [`aws-backup-aws-backup-support-for-amazon-s3`](aws-backup-aws-backup-support-for-amazon-s3/README.md) | AWS Backup AWS Backup support for Amazon S3 | logical | [SVG](aws-backup-aws-backup-support-for-amazon-s3/sample.svg) |
| [`aws-backup-aws-backup-support-for-vmware-workloads`](aws-backup-aws-backup-support-for-vmware-workloads/README.md) | AWS Backup AWS Backup Support for VMware Workloads | logical | [SVG](aws-backup-aws-backup-support-for-vmware-workloads/sample.svg) |
| [`aws-backup-backup-plan`](aws-backup-backup-plan/README.md) | AWS Backup Backup Plan | logical | [SVG](aws-backup-backup-plan/sample.svg) |
| [`aws-backup-backup-restore`](aws-backup-backup-restore/README.md) | AWS Backup Backup Restore | logical | [SVG](aws-backup-backup-restore/sample.svg) |
| [`aws-backup-backup-vault`](aws-backup-backup-vault/README.md) | AWS Backup Backup Vault | logical | [SVG](aws-backup-backup-vault/sample.svg) |
| [`aws-backup-compliance-reporting`](aws-backup-compliance-reporting/README.md) | AWS Backup Compliance Reporting | logical | [SVG](aws-backup-compliance-reporting/sample.svg) |
| [`aws-backup-compute`](aws-backup-compute/README.md) | AWS Backup Compute | logical | [SVG](aws-backup-compute/sample.svg) |
| [`aws-backup-database`](aws-backup-database/README.md) | AWS Backup Database | logical | [SVG](aws-backup-database/sample.svg) |
| [`aws-backup-gateway`](aws-backup-gateway/README.md) | AWS Backup Gateway | logical | [SVG](aws-backup-gateway/sample.svg) |
| [`aws-backup-legal-hold`](aws-backup-legal-hold/README.md) | AWS Backup Legal Hold | logical | [SVG](aws-backup-legal-hold/sample.svg) |
| [`aws-backup-recovery-point-objective`](aws-backup-recovery-point-objective/README.md) | AWS Backup Recovery Point Objective | logical | [SVG](aws-backup-recovery-point-objective/sample.svg) |
| [`aws-backup-recovery-time-objective`](aws-backup-recovery-time-objective/README.md) | AWS Backup Recovery Time Objective | logical | [SVG](aws-backup-recovery-time-objective/sample.svg) |
| [`aws-backup-storage`](aws-backup-storage/README.md) | AWS Backup Storage | logical | [SVG](aws-backup-storage/sample.svg) |
| [`aws-backup-vault-lock`](aws-backup-vault-lock/README.md) | AWS Backup Vault Lock | logical | [SVG](aws-backup-vault-lock/sample.svg) |
| [`aws-backup-virtual-machine`](aws-backup-virtual-machine/README.md) | AWS Backup Virtual Machine | logical | [SVG](aws-backup-virtual-machine/sample.svg) |
| [`aws-backup-virtual-machine-monitor`](aws-backup-virtual-machine-monitor/README.md) | AWS Backup Virtual Machine Monitor | logical | [SVG](aws-backup-virtual-machine-monitor/sample.svg) |
| [`aws-batch`](aws-batch/README.md) | AWS Batch | service | [SVG](aws-batch/sample.svg) |
| [`aws-bedrock`](aws-bedrock/README.md) | Amazon Bedrock | service | [SVG](aws-bedrock/sample.svg) |
| [`aws-billing-conductor`](aws-billing-conductor/README.md) | AWS Billing Conductor | service | [SVG](aws-billing-conductor/sample.svg) |
| [`aws-bottlerocket`](aws-bottlerocket/README.md) | Bottlerocket | service | [SVG](aws-bottlerocket/sample.svg) |
| [`aws-braket`](aws-braket/README.md) | Amazon Braket | service | [SVG](aws-braket/sample.svg) |
| [`aws-braket-chandelier`](aws-braket-chandelier/README.md) | Amazon Braket Chandelier | logical | [SVG](aws-braket-chandelier/sample.svg) |
| [`aws-braket-chip`](aws-braket-chip/README.md) | Amazon Braket Chip | logical | [SVG](aws-braket-chip/sample.svg) |
| [`aws-braket-embedded-simulator`](aws-braket-embedded-simulator/README.md) | Amazon Braket Embedded Simulator | logical | [SVG](aws-braket-embedded-simulator/sample.svg) |
| [`aws-braket-managed-simulator`](aws-braket-managed-simulator/README.md) | Amazon Braket Managed Simulator | logical | [SVG](aws-braket-managed-simulator/sample.svg) |
| [`aws-braket-noise-simulator`](aws-braket-noise-simulator/README.md) | Amazon Braket Noise Simulator | logical | [SVG](aws-braket-noise-simulator/sample.svg) |
| [`aws-braket-qpu`](aws-braket-qpu/README.md) | Amazon Braket QPU | logical | [SVG](aws-braket-qpu/sample.svg) |
| [`aws-braket-simulator`](aws-braket-simulator/README.md) | Amazon Braket Simulator | logical | [SVG](aws-braket-simulator/sample.svg) |
| [`aws-braket-simulator-1`](aws-braket-simulator-1/README.md) | Amazon Braket Simulator 1 | logical | [SVG](aws-braket-simulator-1/sample.svg) |
| [`aws-braket-simulator-2`](aws-braket-simulator-2/README.md) | Amazon Braket Simulator 2 | logical | [SVG](aws-braket-simulator-2/sample.svg) |
| [`aws-braket-simulator-3`](aws-braket-simulator-3/README.md) | Amazon Braket Simulator 3 | logical | [SVG](aws-braket-simulator-3/sample.svg) |
| [`aws-braket-simulator-4`](aws-braket-simulator-4/README.md) | Amazon Braket Simulator 4 | logical | [SVG](aws-braket-simulator-4/sample.svg) |
| [`aws-braket-state-vector`](aws-braket-state-vector/README.md) | Amazon Braket State Vector | logical | [SVG](aws-braket-state-vector/sample.svg) |
| [`aws-braket-tensor-network`](aws-braket-tensor-network/README.md) | Amazon Braket Tensor Network | logical | [SVG](aws-braket-tensor-network/sample.svg) |
| [`aws-budgets`](aws-budgets/README.md) | AWS Budgets | service | [SVG](aws-budgets/sample.svg) |
| [`aws-camera-48-dark`](aws-camera-48-dark/README.md) | Camera 48 Dark | logical | [SVG](aws-camera-48-dark/sample.svg) |
| [`aws-camera-48-light`](aws-camera-48-light/README.md) | Camera 48 Light | logical | [SVG](aws-camera-48-light/sample.svg) |
| [`aws-category-arch-category-analytics`](aws-category-arch-category-analytics/README.md) | Arch Category Analytics | logical | [SVG](aws-category-arch-category-analytics/sample.svg) |
| [`aws-category-arch-category-application-integration`](aws-category-arch-category-application-integration/README.md) | Arch Category Application Integration | logical | [SVG](aws-category-arch-category-application-integration/sample.svg) |
| [`aws-category-arch-category-artificial-intelligence`](aws-category-arch-category-artificial-intelligence/README.md) | Arch Category Artificial Intelligence | logical | [SVG](aws-category-arch-category-artificial-intelligence/sample.svg) |
| [`aws-category-arch-category-blockchain`](aws-category-arch-category-blockchain/README.md) | Arch Category Blockchain | logical | [SVG](aws-category-arch-category-blockchain/sample.svg) |
| [`aws-category-arch-category-business-applications`](aws-category-arch-category-business-applications/README.md) | Arch Category Business Applications | logical | [SVG](aws-category-arch-category-business-applications/sample.svg) |
| [`aws-category-arch-category-cloud-financial-management`](aws-category-arch-category-cloud-financial-management/README.md) | Arch Category Cloud Financial Management | logical | [SVG](aws-category-arch-category-cloud-financial-management/sample.svg) |
| [`aws-category-arch-category-compute`](aws-category-arch-category-compute/README.md) | Arch Category Compute | logical | [SVG](aws-category-arch-category-compute/sample.svg) |
| [`aws-category-arch-category-contact-center`](aws-category-arch-category-contact-center/README.md) | Arch Category Contact Center | logical | [SVG](aws-category-arch-category-contact-center/sample.svg) |
| [`aws-category-arch-category-containers`](aws-category-arch-category-containers/README.md) | Arch Category Containers | logical | [SVG](aws-category-arch-category-containers/sample.svg) |
| [`aws-category-arch-category-customer-enablement`](aws-category-arch-category-customer-enablement/README.md) | Arch Category Customer Enablement | logical | [SVG](aws-category-arch-category-customer-enablement/sample.svg) |
| [`aws-category-arch-category-database`](aws-category-arch-category-database/README.md) | Arch Category Database | logical | [SVG](aws-category-arch-category-database/sample.svg) |
| [`aws-category-arch-category-developer-tools`](aws-category-arch-category-developer-tools/README.md) | Arch Category Developer Tools | logical | [SVG](aws-category-arch-category-developer-tools/sample.svg) |
| [`aws-category-arch-category-end-user-computing`](aws-category-arch-category-end-user-computing/README.md) | Arch Category End User Computing | logical | [SVG](aws-category-arch-category-end-user-computing/sample.svg) |
| [`aws-category-arch-category-front-end-web-mobile`](aws-category-arch-category-front-end-web-mobile/README.md) | Arch Category Front End Web Mobile | logical | [SVG](aws-category-arch-category-front-end-web-mobile/sample.svg) |
| [`aws-category-arch-category-games`](aws-category-arch-category-games/README.md) | Arch Category Games | logical | [SVG](aws-category-arch-category-games/sample.svg) |
| [`aws-category-arch-category-internet-of-things`](aws-category-arch-category-internet-of-things/README.md) | Arch Category Internet of Things | logical | [SVG](aws-category-arch-category-internet-of-things/sample.svg) |
| [`aws-category-arch-category-management-governance`](aws-category-arch-category-management-governance/README.md) | Arch Category Management Governance | logical | [SVG](aws-category-arch-category-management-governance/sample.svg) |
| [`aws-category-arch-category-media-services`](aws-category-arch-category-media-services/README.md) | Arch Category Media Services | logical | [SVG](aws-category-arch-category-media-services/sample.svg) |
| [`aws-category-arch-category-migration-modernization`](aws-category-arch-category-migration-modernization/README.md) | Arch Category Migration Modernization | logical | [SVG](aws-category-arch-category-migration-modernization/sample.svg) |
| [`aws-category-arch-category-networking-content-delivery`](aws-category-arch-category-networking-content-delivery/README.md) | Arch Category Networking Content Delivery | logical | [SVG](aws-category-arch-category-networking-content-delivery/sample.svg) |
| [`aws-category-arch-category-quantum-technologies`](aws-category-arch-category-quantum-technologies/README.md) | Arch Category Quantum Technologies | logical | [SVG](aws-category-arch-category-quantum-technologies/sample.svg) |
| [`aws-category-arch-category-satellite`](aws-category-arch-category-satellite/README.md) | Arch Category Satellite | logical | [SVG](aws-category-arch-category-satellite/sample.svg) |
| [`aws-category-arch-category-security-identity-compliance`](aws-category-arch-category-security-identity-compliance/README.md) | Arch Category Security Identity Compliance | logical | [SVG](aws-category-arch-category-security-identity-compliance/sample.svg) |
| [`aws-category-arch-category-serverless`](aws-category-arch-category-serverless/README.md) | Arch Category Serverless | logical | [SVG](aws-category-arch-category-serverless/sample.svg) |
| [`aws-category-arch-category-storage`](aws-category-arch-category-storage/README.md) | Arch Category Storage | logical | [SVG](aws-category-arch-category-storage/sample.svg) |
| [`aws-certificate-manager`](aws-certificate-manager/README.md) | AWS Certificate Manager | service | [SVG](aws-certificate-manager/sample.svg) |
| [`aws-certificate-manager-certificate-authority`](aws-certificate-manager-certificate-authority/README.md) | AWS Certificate Manager Certificate Authority | logical | [SVG](aws-certificate-manager-certificate-authority/sample.svg) |
| [`aws-chat-48-dark`](aws-chat-48-dark/README.md) | Chat 48 Dark | logical | [SVG](aws-chat-48-dark/sample.svg) |
| [`aws-chat-48-light`](aws-chat-48-light/README.md) | Chat 48 Light | logical | [SVG](aws-chat-48-light/sample.svg) |
| [`aws-chatbot`](aws-chatbot/README.md) | AWS Chatbot | service | [SVG](aws-chatbot/sample.svg) |
| [`aws-chime`](aws-chime/README.md) | Amazon Chime | service | [SVG](aws-chime/sample.svg) |
| [`aws-chime-sdk`](aws-chime-sdk/README.md) | Amazon Chime SDK | service | [SVG](aws-chime-sdk/sample.svg) |
| [`aws-clean-rooms`](aws-clean-rooms/README.md) | AWS Clean Rooms | service | [SVG](aws-clean-rooms/sample.svg) |
| [`aws-client-48-dark`](aws-client-48-dark/README.md) | Client 48 Dark | logical | [SVG](aws-client-48-dark/sample.svg) |
| [`aws-client-48-light`](aws-client-48-light/README.md) | Client 48 Light | logical | [SVG](aws-client-48-light/sample.svg) |
| [`aws-client-vpn`](aws-client-vpn/README.md) | AWS Client VPN | service | [SVG](aws-client-vpn/sample.svg) |
| [`aws-cloud`](aws-cloud/README.md) | AWS Cloud | global | [SVG](aws-cloud/sample.svg) |
| [`aws-cloud-alt`](aws-cloud-alt/README.md) | AWS Cloud | global | [SVG](aws-cloud-alt/sample.svg) |
| [`aws-cloud-alt-dark`](aws-cloud-alt-dark/README.md) | AWS Cloud (dark icon) | global | [SVG](aws-cloud-alt-dark/sample.svg) |
| [`aws-cloud-control-api`](aws-cloud-control-api/README.md) | AWS Cloud Control API | service | [SVG](aws-cloud-control-api/sample.svg) |
| [`aws-cloud-dark`](aws-cloud-dark/README.md) | AWS Cloud (dark icon) | global | [SVG](aws-cloud-dark/sample.svg) |
| [`aws-cloud-development-kit`](aws-cloud-development-kit/README.md) | AWS Cloud Development Kit | service | [SVG](aws-cloud-development-kit/sample.svg) |
| [`aws-cloud-digital-interface`](aws-cloud-digital-interface/README.md) | AWS Cloud Digital Interface | logical | [SVG](aws-cloud-digital-interface/sample.svg) |
| [`aws-cloud-directory`](aws-cloud-directory/README.md) | Amazon Cloud Directory | service | [SVG](aws-cloud-directory/sample.svg) |
| [`aws-cloud-map`](aws-cloud-map/README.md) | AWS Cloud Map | service | [SVG](aws-cloud-map/sample.svg) |
| [`aws-cloud-map-namespace`](aws-cloud-map-namespace/README.md) | AWS Cloud Map Namespace | logical | [SVG](aws-cloud-map-namespace/sample.svg) |
| [`aws-cloud-map-resource`](aws-cloud-map-resource/README.md) | AWS Cloud Map Resource | logical | [SVG](aws-cloud-map-resource/sample.svg) |
| [`aws-cloud-map-service`](aws-cloud-map-service/README.md) | AWS Cloud Map Service | logical | [SVG](aws-cloud-map-service/sample.svg) |
| [`aws-cloud-wan`](aws-cloud-wan/README.md) | AWS Cloud WAN | service | [SVG](aws-cloud-wan/sample.svg) |
| [`aws-cloud-wan-core-network-edge`](aws-cloud-wan-core-network-edge/README.md) | AWS Cloud WAN Core Network Edge | logical | [SVG](aws-cloud-wan-core-network-edge/sample.svg) |
| [`aws-cloud-wan-segment-network`](aws-cloud-wan-segment-network/README.md) | AWS Cloud WAN Segment Network | logical | [SVG](aws-cloud-wan-segment-network/sample.svg) |
| [`aws-cloud-wan-transit-gateway-route-table-attachment`](aws-cloud-wan-transit-gateway-route-table-attachment/README.md) | AWS Cloud WAN Transit Gateway Route Table Attachment | logical | [SVG](aws-cloud-wan-transit-gateway-route-table-attachment/sample.svg) |
| [`aws-cloud9`](aws-cloud9/README.md) | AWS Cloud9 | service | [SVG](aws-cloud9/sample.svg) |
| [`aws-cloud9-cloud9`](aws-cloud9-cloud9/README.md) | AWS Cloud9 Cloud9 | logical | [SVG](aws-cloud9-cloud9/sample.svg) |
| [`aws-cloudformation`](aws-cloudformation/README.md) | AWS CloudFormation | service | [SVG](aws-cloudformation/sample.svg) |
| [`aws-cloudformation-change-set`](aws-cloudformation-change-set/README.md) | AWS CloudFormation Change Set | logical | [SVG](aws-cloudformation-change-set/sample.svg) |
| [`aws-cloudformation-stack`](aws-cloudformation-stack/README.md) | AWS CloudFormation Stack | logical | [SVG](aws-cloudformation-stack/sample.svg) |
| [`aws-cloudformation-template`](aws-cloudformation-template/README.md) | AWS CloudFormation Template | logical | [SVG](aws-cloudformation-template/sample.svg) |
| [`aws-cloudfront`](aws-cloudfront/README.md) | Amazon CloudFront | global | [SVG](aws-cloudfront/sample.svg) |
| [`aws-cloudfront-download-distribution`](aws-cloudfront-download-distribution/README.md) | Amazon CloudFront Download Distribution | logical | [SVG](aws-cloudfront-download-distribution/sample.svg) |
| [`aws-cloudfront-edge-location`](aws-cloudfront-edge-location/README.md) | Amazon CloudFront Edge Location | logical | [SVG](aws-cloudfront-edge-location/sample.svg) |
| [`aws-cloudfront-functions`](aws-cloudfront-functions/README.md) | Amazon CloudFront Functions | logical | [SVG](aws-cloudfront-functions/sample.svg) |
| [`aws-cloudfront-streaming-distribution`](aws-cloudfront-streaming-distribution/README.md) | Amazon CloudFront Streaming Distribution | logical | [SVG](aws-cloudfront-streaming-distribution/sample.svg) |
| [`aws-cloudhsm`](aws-cloudhsm/README.md) | AWS CloudHSM | service | [SVG](aws-cloudhsm/sample.svg) |
| [`aws-cloudsearch`](aws-cloudsearch/README.md) | Amazon CloudSearch | service | [SVG](aws-cloudsearch/sample.svg) |
| [`aws-cloudsearch-search-documents`](aws-cloudsearch-search-documents/README.md) | Amazon CloudSearch Search Documents | logical | [SVG](aws-cloudsearch-search-documents/sample.svg) |
| [`aws-cloudshell`](aws-cloudshell/README.md) | AWS CloudShell | service | [SVG](aws-cloudshell/sample.svg) |
| [`aws-cloudtrail`](aws-cloudtrail/README.md) | AWS CloudTrail | service | [SVG](aws-cloudtrail/sample.svg) |
| [`aws-cloudtrail-cloudtrail-lake`](aws-cloudtrail-cloudtrail-lake/README.md) | AWS CloudTrail CloudTrail Lake | logical | [SVG](aws-cloudtrail-cloudtrail-lake/sample.svg) |
| [`aws-cloudwatch`](aws-cloudwatch/README.md) | Amazon CloudWatch | service | [SVG](aws-cloudwatch/sample.svg) |
| [`aws-cloudwatch-alarm`](aws-cloudwatch-alarm/README.md) | Amazon CloudWatch Alarm | logical | [SVG](aws-cloudwatch-alarm/sample.svg) |
| [`aws-cloudwatch-cross-account-observability`](aws-cloudwatch-cross-account-observability/README.md) | Amazon CloudWatch Cross account Observability | logical | [SVG](aws-cloudwatch-cross-account-observability/sample.svg) |
| [`aws-cloudwatch-data-protection`](aws-cloudwatch-data-protection/README.md) | Amazon CloudWatch Data Protection | logical | [SVG](aws-cloudwatch-data-protection/sample.svg) |
| [`aws-cloudwatch-event-event-based`](aws-cloudwatch-event-event-based/README.md) | Amazon CloudWatch Event Event Based | logical | [SVG](aws-cloudwatch-event-event-based/sample.svg) |
| [`aws-cloudwatch-event-time-based`](aws-cloudwatch-event-time-based/README.md) | Amazon CloudWatch Event Time Based | logical | [SVG](aws-cloudwatch-event-time-based/sample.svg) |
| [`aws-cloudwatch-evidently`](aws-cloudwatch-evidently/README.md) | Amazon CloudWatch Evidently | logical | [SVG](aws-cloudwatch-evidently/sample.svg) |
| [`aws-cloudwatch-logs`](aws-cloudwatch-logs/README.md) | Amazon CloudWatch Logs | logical | [SVG](aws-cloudwatch-logs/sample.svg) |
| [`aws-cloudwatch-metrics-insights`](aws-cloudwatch-metrics-insights/README.md) | Amazon CloudWatch Metrics Insights | logical | [SVG](aws-cloudwatch-metrics-insights/sample.svg) |
| [`aws-cloudwatch-rule`](aws-cloudwatch-rule/README.md) | Amazon CloudWatch Rule | logical | [SVG](aws-cloudwatch-rule/sample.svg) |
| [`aws-cloudwatch-rum`](aws-cloudwatch-rum/README.md) | Amazon CloudWatch RUM | logical | [SVG](aws-cloudwatch-rum/sample.svg) |
| [`aws-cloudwatch-synthetics`](aws-cloudwatch-synthetics/README.md) | Amazon CloudWatch Synthetics | logical | [SVG](aws-cloudwatch-synthetics/sample.svg) |
| [`aws-codeartifact`](aws-codeartifact/README.md) | AWS CodeArtifact | service | [SVG](aws-codeartifact/sample.svg) |
| [`aws-codebuild`](aws-codebuild/README.md) | AWS CodeBuild | service | [SVG](aws-codebuild/sample.svg) |
| [`aws-codecatalyst`](aws-codecatalyst/README.md) | Amazon CodeCatalyst | service | [SVG](aws-codecatalyst/sample.svg) |
| [`aws-codecommit`](aws-codecommit/README.md) | AWS CodeCommit | service | [SVG](aws-codecommit/sample.svg) |
| [`aws-codedeploy`](aws-codedeploy/README.md) | AWS CodeDeploy | service | [SVG](aws-codedeploy/sample.svg) |
| [`aws-codeguru`](aws-codeguru/README.md) | Amazon CodeGuru | service | [SVG](aws-codeguru/sample.svg) |
| [`aws-codepipeline`](aws-codepipeline/README.md) | AWS CodePipeline | service | [SVG](aws-codepipeline/sample.svg) |
| [`aws-codewhisperer`](aws-codewhisperer/README.md) | Amazon CodeWhisperer | service | [SVG](aws-codewhisperer/sample.svg) |
| [`aws-cognito`](aws-cognito/README.md) | Amazon Cognito | service | [SVG](aws-cognito/sample.svg) |
| [`aws-cold-storage-48-dark`](aws-cold-storage-48-dark/README.md) | Cold Storage 48 Dark | logical | [SVG](aws-cold-storage-48-dark/sample.svg) |
| [`aws-cold-storage-48-light`](aws-cold-storage-48-light/README.md) | Cold Storage 48 Light | logical | [SVG](aws-cold-storage-48-light/sample.svg) |
| [`aws-command-line-interface`](aws-command-line-interface/README.md) | AWS Command Line Interface | service | [SVG](aws-command-line-interface/sample.svg) |
| [`aws-comprehend`](aws-comprehend/README.md) | Amazon Comprehend | service | [SVG](aws-comprehend/sample.svg) |
| [`aws-comprehend-medical`](aws-comprehend-medical/README.md) | Amazon Comprehend Medical | service | [SVG](aws-comprehend-medical/sample.svg) |
| [`aws-compute-optimizer`](aws-compute-optimizer/README.md) | AWS Compute Optimizer | service | [SVG](aws-compute-optimizer/sample.svg) |
| [`aws-config`](aws-config/README.md) | AWS Config | service | [SVG](aws-config/sample.svg) |
| [`aws-connect`](aws-connect/README.md) | Amazon Connect | service | [SVG](aws-connect/sample.svg) |
| [`aws-console-mobile-application`](aws-console-mobile-application/README.md) | AWS Console Mobile Application | service | [SVG](aws-console-mobile-application/sample.svg) |
| [`aws-control-tower`](aws-control-tower/README.md) | AWS Control Tower | service | [SVG](aws-control-tower/sample.svg) |
| [`aws-corretto`](aws-corretto/README.md) | Amazon Corretto | service | [SVG](aws-corretto/sample.svg) |
| [`aws-cost-and-usage-report`](aws-cost-and-usage-report/README.md) | AWS Cost and Usage Report | service | [SVG](aws-cost-and-usage-report/sample.svg) |
| [`aws-cost-explorer`](aws-cost-explorer/README.md) | AWS Cost Explorer | service | [SVG](aws-cost-explorer/sample.svg) |
| [`aws-credentials-48-dark`](aws-credentials-48-dark/README.md) | Credentials 48 Dark | logical | [SVG](aws-credentials-48-dark/sample.svg) |
| [`aws-credentials-48-light`](aws-credentials-48-light/README.md) | Credentials 48 Light | logical | [SVG](aws-credentials-48-light/sample.svg) |
| [`aws-data-exchange`](aws-data-exchange/README.md) | AWS Data Exchange | service | [SVG](aws-data-exchange/sample.svg) |
| [`aws-data-exchange-for-apis`](aws-data-exchange-for-apis/README.md) | AWS Data Exchange for APIs | logical | [SVG](aws-data-exchange-for-apis/sample.svg) |
| [`aws-data-firehose`](aws-data-firehose/README.md) | Amazon Data Firehose | service | [SVG](aws-data-firehose/sample.svg) |
| [`aws-data-stream-48-dark`](aws-data-stream-48-dark/README.md) | Data Stream 48 Dark | logical | [SVG](aws-data-stream-48-dark/sample.svg) |
| [`aws-data-stream-48-light`](aws-data-stream-48-light/README.md) | Data Stream 48 Light | logical | [SVG](aws-data-stream-48-light/sample.svg) |
| [`aws-data-table-48-dark`](aws-data-table-48-dark/README.md) | Data Table 48 Dark | logical | [SVG](aws-data-table-48-dark/sample.svg) |
| [`aws-data-table-48-light`](aws-data-table-48-light/README.md) | Data Table 48 Light | logical | [SVG](aws-data-table-48-light/sample.svg) |
| [`aws-data-transfer-terminal`](aws-data-transfer-terminal/README.md) | AWS Data Transfer Terminal | service | [SVG](aws-data-transfer-terminal/sample.svg) |
| [`aws-database-48-dark`](aws-database-48-dark/README.md) | Database 48 Dark | logical | [SVG](aws-database-48-dark/sample.svg) |
| [`aws-database-48-light`](aws-database-48-light/README.md) | Database 48 Light | logical | [SVG](aws-database-48-light/sample.svg) |
| [`aws-database-migration-service`](aws-database-migration-service/README.md) | AWS Database Migration Service | service | [SVG](aws-database-migration-service/sample.svg) |
| [`aws-database-migration-service-database-migration-workflow-or-job`](aws-database-migration-service-database-migration-workflow-or-job/README.md) | AWS Database Migration Service Database migration workflow or job | logical | [SVG](aws-database-migration-service-database-migration-workflow-or-job/sample.svg) |
| [`aws-datasync`](aws-datasync/README.md) | AWS DataSync | service | [SVG](aws-datasync/sample.svg) |
| [`aws-datasync-agent`](aws-datasync-agent/README.md) | AWS Datasync Agent | logical | [SVG](aws-datasync-agent/sample.svg) |
| [`aws-datasync-discovery`](aws-datasync-discovery/README.md) | AWS DataSync Discovery | logical | [SVG](aws-datasync-discovery/sample.svg) |
| [`aws-datazone`](aws-datazone/README.md) | Amazon DataZone | service | [SVG](aws-datazone/sample.svg) |
| [`aws-datazone-business-data-catalog`](aws-datazone-business-data-catalog/README.md) | Amazon DataZone Business Data Catalog | logical | [SVG](aws-datazone-business-data-catalog/sample.svg) |
| [`aws-datazone-data-portal`](aws-datazone-data-portal/README.md) | Amazon DataZone Data Portal | logical | [SVG](aws-datazone-data-portal/sample.svg) |
| [`aws-datazone-data-projects`](aws-datazone-data-projects/README.md) | Amazon DataZone Data Projects | logical | [SVG](aws-datazone-data-projects/sample.svg) |
| [`aws-dcv`](aws-dcv/README.md) | Amazon DCV | service | [SVG](aws-dcv/sample.svg) |
| [`aws-deadline-cloud`](aws-deadline-cloud/README.md) | AWS Deadline Cloud | service | [SVG](aws-deadline-cloud/sample.svg) |
| [`aws-deep-learning-amis`](aws-deep-learning-amis/README.md) | AWS Deep Learning AMIs | service | [SVG](aws-deep-learning-amis/sample.svg) |
| [`aws-deep-learning-containers`](aws-deep-learning-containers/README.md) | AWS Deep Learning Containers | service | [SVG](aws-deep-learning-containers/sample.svg) |
| [`aws-deepcomposer`](aws-deepcomposer/README.md) | AWS DeepComposer | service | [SVG](aws-deepcomposer/sample.svg) |
| [`aws-deepracer`](aws-deepracer/README.md) | AWS DeepRacer | service | [SVG](aws-deepracer/sample.svg) |
| [`aws-detective`](aws-detective/README.md) | Amazon Detective | service | [SVG](aws-detective/sample.svg) |
| [`aws-device-farm`](aws-device-farm/README.md) | AWS Device Farm | service | [SVG](aws-device-farm/sample.svg) |
| [`aws-devops-guru`](aws-devops-guru/README.md) | Amazon DevOps Guru | service | [SVG](aws-devops-guru/sample.svg) |
| [`aws-devops-guru-insights`](aws-devops-guru-insights/README.md) | Amazon DevOps Guru Insights | logical | [SVG](aws-devops-guru-insights/sample.svg) |
| [`aws-direct-connect`](aws-direct-connect/README.md) | AWS Direct Connect | service | [SVG](aws-direct-connect/sample.svg) |
| [`aws-direct-connect-gateway`](aws-direct-connect-gateway/README.md) | AWS Direct Connect Gateway | logical | [SVG](aws-direct-connect-gateway/sample.svg) |
| [`aws-directory-service`](aws-directory-service/README.md) | AWS Directory Service | service | [SVG](aws-directory-service/sample.svg) |
| [`aws-directory-service-ad-connector`](aws-directory-service-ad-connector/README.md) | AWS Directory Service AD Connector | logical | [SVG](aws-directory-service-ad-connector/sample.svg) |
| [`aws-directory-service-aws-managed-microsoft-ad`](aws-directory-service-aws-managed-microsoft-ad/README.md) | AWS Directory Service AWS Managed Microsoft AD | logical | [SVG](aws-directory-service-aws-managed-microsoft-ad/sample.svg) |
| [`aws-directory-service-simple-ad`](aws-directory-service-simple-ad/README.md) | AWS Directory Service Simple AD | logical | [SVG](aws-directory-service-simple-ad/sample.svg) |
| [`aws-disk-48-dark`](aws-disk-48-dark/README.md) | Disk 48 Dark | logical | [SVG](aws-disk-48-dark/sample.svg) |
| [`aws-disk-48-light`](aws-disk-48-light/README.md) | Disk 48 Light | logical | [SVG](aws-disk-48-light/sample.svg) |
| [`aws-distro-for-opentelemetry`](aws-distro-for-opentelemetry/README.md) | AWS Distro for OpenTelemetry | service | [SVG](aws-distro-for-opentelemetry/sample.svg) |
| [`aws-document-48-dark`](aws-document-48-dark/README.md) | Document 48 Dark | logical | [SVG](aws-document-48-dark/sample.svg) |
| [`aws-document-48-light`](aws-document-48-light/README.md) | Document 48 Light | logical | [SVG](aws-document-48-light/sample.svg) |
| [`aws-documentdb`](aws-documentdb/README.md) | Amazon DocumentDB | service | [SVG](aws-documentdb/sample.svg) |
| [`aws-documentdb-elastic-clusters`](aws-documentdb-elastic-clusters/README.md) | Amazon DocumentDB Elastic Clusters | logical | [SVG](aws-documentdb-elastic-clusters/sample.svg) |
| [`aws-documents-48-dark`](aws-documents-48-dark/README.md) | Documents 48 Dark | logical | [SVG](aws-documents-48-dark/sample.svg) |
| [`aws-documents-48-light`](aws-documents-48-light/README.md) | Documents 48 Light | logical | [SVG](aws-documents-48-light/sample.svg) |
| [`aws-dynamodb`](aws-dynamodb/README.md) | Amazon DynamoDB | region | [SVG](aws-dynamodb/sample.svg) |
| [`aws-dynamodb-amazon-dynamodb-accelerator`](aws-dynamodb-amazon-dynamodb-accelerator/README.md) | Amazon DynamoDB Amazon DynamoDB Accelerator | logical | [SVG](aws-dynamodb-amazon-dynamodb-accelerator/sample.svg) |
| [`aws-dynamodb-attribute`](aws-dynamodb-attribute/README.md) | Amazon DynamoDB Attribute | logical | [SVG](aws-dynamodb-attribute/sample.svg) |
| [`aws-dynamodb-attributes`](aws-dynamodb-attributes/README.md) | Amazon DynamoDB Attributes | logical | [SVG](aws-dynamodb-attributes/sample.svg) |
| [`aws-dynamodb-global-secondary-index`](aws-dynamodb-global-secondary-index/README.md) | Amazon DynamoDB Global secondary index | logical | [SVG](aws-dynamodb-global-secondary-index/sample.svg) |
| [`aws-dynamodb-item`](aws-dynamodb-item/README.md) | Amazon DynamoDB Item | logical | [SVG](aws-dynamodb-item/sample.svg) |
| [`aws-dynamodb-items`](aws-dynamodb-items/README.md) | Amazon DynamoDB Items | logical | [SVG](aws-dynamodb-items/sample.svg) |
| [`aws-dynamodb-standard-access-table-class`](aws-dynamodb-standard-access-table-class/README.md) | Amazon DynamoDB Standard Access Table Class | logical | [SVG](aws-dynamodb-standard-access-table-class/sample.svg) |
| [`aws-dynamodb-standard-infrequent-access-table-class`](aws-dynamodb-standard-infrequent-access-table-class/README.md) | Amazon DynamoDB Standard Infrequent Access Table Class | logical | [SVG](aws-dynamodb-standard-infrequent-access-table-class/sample.svg) |
| [`aws-dynamodb-stream`](aws-dynamodb-stream/README.md) | Amazon DynamoDB Stream | logical | [SVG](aws-dynamodb-stream/sample.svg) |
| [`aws-dynamodb-table`](aws-dynamodb-table/README.md) | Amazon DynamoDB Table | region | [SVG](aws-dynamodb-table/sample.svg) |
| [`aws-ec2`](aws-ec2/README.md) | Amazon EC2 | subnet | [SVG](aws-ec2/sample.svg) |
| [`aws-ec2-ami`](aws-ec2-ami/README.md) | Amazon EC2 AMI | logical | [SVG](aws-ec2-ami/sample.svg) |
| [`aws-ec2-auto-scaling`](aws-ec2-auto-scaling/README.md) | Amazon EC2 Auto Scaling | service | [SVG](aws-ec2-auto-scaling/sample.svg) |
| [`aws-ec2-auto-scaling-resource`](aws-ec2-auto-scaling-resource/README.md) | Amazon EC2 Auto Scaling | logical | [SVG](aws-ec2-auto-scaling-resource/sample.svg) |
| [`aws-ec2-aws-microservice-extractor-for-net`](aws-ec2-aws-microservice-extractor-for-net/README.md) | Amazon EC2 AWS Microservice Extractor for .NET | logical | [SVG](aws-ec2-aws-microservice-extractor-for-net/sample.svg) |
| [`aws-ec2-db-instance`](aws-ec2-db-instance/README.md) | Amazon EC2 DB Instance | logical | [SVG](aws-ec2-db-instance/sample.svg) |
| [`aws-ec2-elastic-ip-address`](aws-ec2-elastic-ip-address/README.md) | Amazon EC2 Elastic IP Address | logical | [SVG](aws-ec2-elastic-ip-address/sample.svg) |
| [`aws-ec2-image-builder`](aws-ec2-image-builder/README.md) | Amazon EC2 Image Builder | service | [SVG](aws-ec2-image-builder/sample.svg) |
| [`aws-ec2-instance`](aws-ec2-instance/README.md) | Amazon EC2 Instance | subnet | [SVG](aws-ec2-instance/sample.svg) |
| [`aws-ec2-instance-with-cloudwatch`](aws-ec2-instance-with-cloudwatch/README.md) | Amazon EC2 Instance with CloudWatch | logical | [SVG](aws-ec2-instance-with-cloudwatch/sample.svg) |
| [`aws-ec2-instances`](aws-ec2-instances/README.md) | Amazon EC2 Instances | logical | [SVG](aws-ec2-instances/sample.svg) |
| [`aws-ec2-rescue`](aws-ec2-rescue/README.md) | Amazon EC2 Rescue | logical | [SVG](aws-ec2-rescue/sample.svg) |
| [`aws-ec2-spot-instance`](aws-ec2-spot-instance/README.md) | Amazon EC2 Spot Instance | logical | [SVG](aws-ec2-spot-instance/sample.svg) |
| [`aws-ecs-anywhere`](aws-ecs-anywhere/README.md) | Amazon ECS Anywhere | service | [SVG](aws-ecs-anywhere/sample.svg) |
| [`aws-efs`](aws-efs/README.md) | Amazon EFS | service | [SVG](aws-efs/sample.svg) |
| [`aws-eks-anywhere`](aws-eks-anywhere/README.md) | Amazon EKS Anywhere | service | [SVG](aws-eks-anywhere/sample.svg) |
| [`aws-eks-cloud`](aws-eks-cloud/README.md) | Amazon EKS Cloud | service | [SVG](aws-eks-cloud/sample.svg) |
| [`aws-eks-distro`](aws-eks-distro/README.md) | Amazon EKS Distro | service | [SVG](aws-eks-distro/sample.svg) |
| [`aws-elastic-beanstalk`](aws-elastic-beanstalk/README.md) | AWS Elastic Beanstalk | service | [SVG](aws-elastic-beanstalk/sample.svg) |
| [`aws-elastic-beanstalk-application`](aws-elastic-beanstalk-application/README.md) | AWS Elastic Beanstalk Application | logical | [SVG](aws-elastic-beanstalk-application/sample.svg) |
| [`aws-elastic-beanstalk-deployment`](aws-elastic-beanstalk-deployment/README.md) | AWS Elastic Beanstalk Deployment | logical | [SVG](aws-elastic-beanstalk-deployment/sample.svg) |
| [`aws-elastic-block-store`](aws-elastic-block-store/README.md) | Amazon Elastic Block Store | service | [SVG](aws-elastic-block-store/sample.svg) |
| [`aws-elastic-block-store-amazon-data-lifecycle-manager`](aws-elastic-block-store-amazon-data-lifecycle-manager/README.md) | Amazon Elastic Block Store Amazon Data Lifecycle Manager | logical | [SVG](aws-elastic-block-store-amazon-data-lifecycle-manager/sample.svg) |
| [`aws-elastic-block-store-multiple-volumes`](aws-elastic-block-store-multiple-volumes/README.md) | Amazon Elastic Block Store Multiple Volumes | logical | [SVG](aws-elastic-block-store-multiple-volumes/sample.svg) |
| [`aws-elastic-block-store-snapshot`](aws-elastic-block-store-snapshot/README.md) | Amazon Elastic Block Store Snapshot | logical | [SVG](aws-elastic-block-store-snapshot/sample.svg) |
| [`aws-elastic-block-store-volume`](aws-elastic-block-store-volume/README.md) | Amazon Elastic Block Store Volume | logical | [SVG](aws-elastic-block-store-volume/sample.svg) |
| [`aws-elastic-block-store-volume-gp3`](aws-elastic-block-store-volume-gp3/README.md) | Amazon Elastic Block Store Volume gp3 | logical | [SVG](aws-elastic-block-store-volume-gp3/sample.svg) |
| [`aws-elastic-container-registry`](aws-elastic-container-registry/README.md) | Amazon Elastic Container Registry | service | [SVG](aws-elastic-container-registry/sample.svg) |
| [`aws-elastic-container-registry-image`](aws-elastic-container-registry-image/README.md) | Amazon Elastic Container Registry Image | logical | [SVG](aws-elastic-container-registry-image/sample.svg) |
| [`aws-elastic-container-registry-registry`](aws-elastic-container-registry-registry/README.md) | Amazon Elastic Container Registry Registry | logical | [SVG](aws-elastic-container-registry-registry/sample.svg) |
| [`aws-elastic-container-service`](aws-elastic-container-service/README.md) | Amazon Elastic Container Service | service | [SVG](aws-elastic-container-service/sample.svg) |
| [`aws-elastic-container-service-container-1`](aws-elastic-container-service-container-1/README.md) | Amazon Elastic Container Service Container 1 | logical | [SVG](aws-elastic-container-service-container-1/sample.svg) |
| [`aws-elastic-container-service-container-2`](aws-elastic-container-service-container-2/README.md) | Amazon Elastic Container Service Container 2 | logical | [SVG](aws-elastic-container-service-container-2/sample.svg) |
| [`aws-elastic-container-service-container-3`](aws-elastic-container-service-container-3/README.md) | Amazon Elastic Container Service Container 3 | logical | [SVG](aws-elastic-container-service-container-3/sample.svg) |
| [`aws-elastic-container-service-copiiot-cli`](aws-elastic-container-service-copiiot-cli/README.md) | Amazon Elastic Container Service CopiIoT CLI | logical | [SVG](aws-elastic-container-service-copiiot-cli/sample.svg) |
| [`aws-elastic-container-service-ecs-service-connect`](aws-elastic-container-service-ecs-service-connect/README.md) | Amazon Elastic Container Service ECS Service Connect | logical | [SVG](aws-elastic-container-service-ecs-service-connect/sample.svg) |
| [`aws-elastic-container-service-service`](aws-elastic-container-service-service/README.md) | Amazon Elastic Container Service Service | logical | [SVG](aws-elastic-container-service-service/sample.svg) |
| [`aws-elastic-container-service-task`](aws-elastic-container-service-task/README.md) | Amazon Elastic Container Service Task | logical | [SVG](aws-elastic-container-service-task/sample.svg) |
| [`aws-elastic-disaster-recovery`](aws-elastic-disaster-recovery/README.md) | AWS Elastic Disaster Recovery | service | [SVG](aws-elastic-disaster-recovery/sample.svg) |
| [`aws-elastic-fabric-adapter`](aws-elastic-fabric-adapter/README.md) | Elastic Fabric Adapter | service | [SVG](aws-elastic-fabric-adapter/sample.svg) |
| [`aws-elastic-file-system-efs-intelligent-tiering`](aws-elastic-file-system-efs-intelligent-tiering/README.md) | Amazon Elastic File System EFS Intelligent Tiering | logical | [SVG](aws-elastic-file-system-efs-intelligent-tiering/sample.svg) |
| [`aws-elastic-file-system-efs-one-zone`](aws-elastic-file-system-efs-one-zone/README.md) | Amazon Elastic File System EFS One Zone | logical | [SVG](aws-elastic-file-system-efs-one-zone/sample.svg) |
| [`aws-elastic-file-system-efs-one-zone-infrequent-access`](aws-elastic-file-system-efs-one-zone-infrequent-access/README.md) | Amazon Elastic File System EFS One Zone Infrequent Access | logical | [SVG](aws-elastic-file-system-efs-one-zone-infrequent-access/sample.svg) |
| [`aws-elastic-file-system-efs-standard`](aws-elastic-file-system-efs-standard/README.md) | Amazon Elastic File System EFS Standard | logical | [SVG](aws-elastic-file-system-efs-standard/sample.svg) |
| [`aws-elastic-file-system-efs-standard-infrequent-access`](aws-elastic-file-system-efs-standard-infrequent-access/README.md) | Amazon Elastic File System EFS Standard Infrequent Access | logical | [SVG](aws-elastic-file-system-efs-standard-infrequent-access/sample.svg) |
| [`aws-elastic-file-system-elastic-throughput`](aws-elastic-file-system-elastic-throughput/README.md) | Amazon Elastic File System Elastic Throughput | logical | [SVG](aws-elastic-file-system-elastic-throughput/sample.svg) |
| [`aws-elastic-file-system-file-system`](aws-elastic-file-system-file-system/README.md) | Amazon Elastic File System File System | logical | [SVG](aws-elastic-file-system-file-system/sample.svg) |
| [`aws-elastic-inference`](aws-elastic-inference/README.md) | Amazon Elastic Inference | service | [SVG](aws-elastic-inference/sample.svg) |
| [`aws-elastic-kubernetes-service`](aws-elastic-kubernetes-service/README.md) | Amazon Elastic Kubernetes Service | service | [SVG](aws-elastic-kubernetes-service/sample.svg) |
| [`aws-elastic-kubernetes-service-eks-on-outposts`](aws-elastic-kubernetes-service-eks-on-outposts/README.md) | Amazon Elastic Kubernetes Service EKS on Outposts | logical | [SVG](aws-elastic-kubernetes-service-eks-on-outposts/sample.svg) |
| [`aws-elastic-load-balancing`](aws-elastic-load-balancing/README.md) | Elastic Load Balancing | service | [SVG](aws-elastic-load-balancing/sample.svg) |
| [`aws-elastic-load-balancing-application-load-balancer`](aws-elastic-load-balancing-application-load-balancer/README.md) | Elastic Load Balancing Application Load Balancer | logical | [SVG](aws-elastic-load-balancing-application-load-balancer/sample.svg) |
| [`aws-elastic-load-balancing-classic-load-balancer`](aws-elastic-load-balancing-classic-load-balancer/README.md) | Elastic Load Balancing Classic Load Balancer | logical | [SVG](aws-elastic-load-balancing-classic-load-balancer/sample.svg) |
| [`aws-elastic-load-balancing-gateway-load-balancer`](aws-elastic-load-balancing-gateway-load-balancer/README.md) | Elastic Load Balancing Gateway Load Balancer | logical | [SVG](aws-elastic-load-balancing-gateway-load-balancer/sample.svg) |
| [`aws-elastic-load-balancing-network-load-balancer`](aws-elastic-load-balancing-network-load-balancer/README.md) | Elastic Load Balancing Network Load Balancer | logical | [SVG](aws-elastic-load-balancing-network-load-balancer/sample.svg) |
| [`aws-elastic-transcoder`](aws-elastic-transcoder/README.md) | Amazon Elastic Transcoder | service | [SVG](aws-elastic-transcoder/sample.svg) |
| [`aws-elastic-vmware-service`](aws-elastic-vmware-service/README.md) | Amazon Elastic VMware Service | service | [SVG](aws-elastic-vmware-service/sample.svg) |
| [`aws-elasticache`](aws-elasticache/README.md) | Amazon ElastiCache | service | [SVG](aws-elasticache/sample.svg) |
| [`aws-elasticache-cache-node`](aws-elasticache-cache-node/README.md) | Amazon ElastiCache Cache Node | logical | [SVG](aws-elasticache-cache-node/sample.svg) |
| [`aws-elasticache-elasticache-for-memcached`](aws-elasticache-elasticache-for-memcached/README.md) | Amazon ElastiCache ElastiCache for Memcached | logical | [SVG](aws-elasticache-elasticache-for-memcached/sample.svg) |
| [`aws-elasticache-elasticache-for-redis`](aws-elasticache-elasticache-for-redis/README.md) | Amazon ElastiCache ElastiCache for Redis | logical | [SVG](aws-elasticache-elasticache-for-redis/sample.svg) |
| [`aws-elasticache-elasticache-for-valkey`](aws-elasticache-elasticache-for-valkey/README.md) | Amazon ElastiCache ElastiCache for Valkey | logical | [SVG](aws-elasticache-elasticache-for-valkey/sample.svg) |
| [`aws-elemental-appliances-software`](aws-elemental-appliances-software/README.md) | AWS Elemental Appliances & Software | service | [SVG](aws-elemental-appliances-software/sample.svg) |
| [`aws-elemental-conductor`](aws-elemental-conductor/README.md) | AWS Elemental Conductor | service | [SVG](aws-elemental-conductor/sample.svg) |
| [`aws-elemental-delta`](aws-elemental-delta/README.md) | AWS Elemental Delta | service | [SVG](aws-elemental-delta/sample.svg) |
| [`aws-elemental-link`](aws-elemental-link/README.md) | AWS Elemental Link | service | [SVG](aws-elemental-link/sample.svg) |
| [`aws-elemental-live`](aws-elemental-live/README.md) | AWS Elemental Live | service | [SVG](aws-elemental-live/sample.svg) |
| [`aws-elemental-mediaconnect`](aws-elemental-mediaconnect/README.md) | AWS Elemental MediaConnect | service | [SVG](aws-elemental-mediaconnect/sample.svg) |
| [`aws-elemental-mediaconnect-mediaconnect-gateway`](aws-elemental-mediaconnect-mediaconnect-gateway/README.md) | AWS Elemental MediaConnect MediaConnect Gateway | logical | [SVG](aws-elemental-mediaconnect-mediaconnect-gateway/sample.svg) |
| [`aws-elemental-mediaconvert`](aws-elemental-mediaconvert/README.md) | AWS Elemental MediaConvert | service | [SVG](aws-elemental-mediaconvert/sample.svg) |
| [`aws-elemental-medialive`](aws-elemental-medialive/README.md) | AWS Elemental MediaLive | service | [SVG](aws-elemental-medialive/sample.svg) |
| [`aws-elemental-mediapackage`](aws-elemental-mediapackage/README.md) | AWS Elemental MediaPackage | service | [SVG](aws-elemental-mediapackage/sample.svg) |
| [`aws-elemental-mediastore`](aws-elemental-mediastore/README.md) | AWS Elemental MediaStore | service | [SVG](aws-elemental-mediastore/sample.svg) |
| [`aws-elemental-mediatailor`](aws-elemental-mediatailor/README.md) | AWS Elemental MediaTailor | service | [SVG](aws-elemental-mediatailor/sample.svg) |
| [`aws-elemental-server`](aws-elemental-server/README.md) | AWS Elemental Server | service | [SVG](aws-elemental-server/sample.svg) |
| [`aws-email-48-dark`](aws-email-48-dark/README.md) | Email 48 Dark | logical | [SVG](aws-email-48-dark/sample.svg) |
| [`aws-email-48-light`](aws-email-48-light/README.md) | Email 48 Light | logical | [SVG](aws-email-48-light/sample.svg) |
| [`aws-emr`](aws-emr/README.md) | Amazon EMR | service | [SVG](aws-emr/sample.svg) |
| [`aws-emr-cluster`](aws-emr-cluster/README.md) | Amazon EMR Cluster | logical | [SVG](aws-emr-cluster/sample.svg) |
| [`aws-emr-emr-engine`](aws-emr-emr-engine/README.md) | Amazon EMR EMR Engine | logical | [SVG](aws-emr-emr-engine/sample.svg) |
| [`aws-emr-hdfs-cluster`](aws-emr-hdfs-cluster/README.md) | Amazon EMR HDFS Cluster | logical | [SVG](aws-emr-hdfs-cluster/sample.svg) |
| [`aws-end-user-messaging`](aws-end-user-messaging/README.md) | AWS End User Messaging | service | [SVG](aws-end-user-messaging/sample.svg) |
| [`aws-entity-resolution`](aws-entity-resolution/README.md) | AWS Entity Resolution | service | [SVG](aws-entity-resolution/sample.svg) |
| [`aws-eventbridge`](aws-eventbridge/README.md) | Amazon EventBridge | service | [SVG](aws-eventbridge/sample.svg) |
| [`aws-eventbridge-custom-event-bus`](aws-eventbridge-custom-event-bus/README.md) | Amazon EventBridge Custom Event Bus | logical | [SVG](aws-eventbridge-custom-event-bus/sample.svg) |
| [`aws-eventbridge-default-event-bus`](aws-eventbridge-default-event-bus/README.md) | Amazon EventBridge Default Event Bus | logical | [SVG](aws-eventbridge-default-event-bus/sample.svg) |
| [`aws-eventbridge-event`](aws-eventbridge-event/README.md) | Amazon EventBridge Event | logical | [SVG](aws-eventbridge-event/sample.svg) |
| [`aws-eventbridge-pipes`](aws-eventbridge-pipes/README.md) | Amazon EventBridge Pipes | logical | [SVG](aws-eventbridge-pipes/sample.svg) |
| [`aws-eventbridge-rule`](aws-eventbridge-rule/README.md) | Amazon EventBridge Rule | logical | [SVG](aws-eventbridge-rule/sample.svg) |
| [`aws-eventbridge-saas-partner-event`](aws-eventbridge-saas-partner-event/README.md) | Amazon EventBridge Saas Partner Event | logical | [SVG](aws-eventbridge-saas-partner-event/sample.svg) |
| [`aws-eventbridge-scheduler`](aws-eventbridge-scheduler/README.md) | Amazon EventBridge Scheduler | logical | [SVG](aws-eventbridge-scheduler/sample.svg) |
| [`aws-eventbridge-schema`](aws-eventbridge-schema/README.md) | Amazon EventBridge Schema | logical | [SVG](aws-eventbridge-schema/sample.svg) |
| [`aws-eventbridge-schema-registry`](aws-eventbridge-schema-registry/README.md) | Amazon EventBridge Schema Registry | logical | [SVG](aws-eventbridge-schema-registry/sample.svg) |
| [`aws-express-workflows`](aws-express-workflows/README.md) | AWS Express Workflows | service | [SVG](aws-express-workflows/sample.svg) |
| [`aws-fargate`](aws-fargate/README.md) | AWS Fargate | service | [SVG](aws-fargate/sample.svg) |
| [`aws-fault-injection-service`](aws-fault-injection-service/README.md) | AWS Fault Injection Service | service | [SVG](aws-fault-injection-service/sample.svg) |
| [`aws-file-cache`](aws-file-cache/README.md) | Amazon File Cache | service | [SVG](aws-file-cache/sample.svg) |
| [`aws-file-cache-hybrid-nfs-linked-datasets`](aws-file-cache-hybrid-nfs-linked-datasets/README.md) | Amazon File Cache Hybrid NFS linked datasets | logical | [SVG](aws-file-cache-hybrid-nfs-linked-datasets/sample.svg) |
| [`aws-file-cache-on-premises-nfs-linked-datasets`](aws-file-cache-on-premises-nfs-linked-datasets/README.md) | Amazon File Cache On premises NFS linked datasets | logical | [SVG](aws-file-cache-on-premises-nfs-linked-datasets/sample.svg) |
| [`aws-file-cache-s3-linked-datasets`](aws-file-cache-s3-linked-datasets/README.md) | Amazon File Cache S3 linked datasets | logical | [SVG](aws-file-cache-s3-linked-datasets/sample.svg) |
| [`aws-finspace`](aws-finspace/README.md) | Amazon FinSpace | service | [SVG](aws-finspace/sample.svg) |
| [`aws-firewall-48-dark`](aws-firewall-48-dark/README.md) | Firewall 48 Dark | logical | [SVG](aws-firewall-48-dark/sample.svg) |
| [`aws-firewall-48-light`](aws-firewall-48-light/README.md) | Firewall 48 Light | logical | [SVG](aws-firewall-48-light/sample.svg) |
| [`aws-firewall-manager`](aws-firewall-manager/README.md) | AWS Firewall Manager | service | [SVG](aws-firewall-manager/sample.svg) |
| [`aws-folder-48-dark`](aws-folder-48-dark/README.md) | Folder 48 Dark | logical | [SVG](aws-folder-48-dark/sample.svg) |
| [`aws-folder-48-light`](aws-folder-48-light/README.md) | Folder 48 Light | logical | [SVG](aws-folder-48-light/sample.svg) |
| [`aws-folders-48-dark`](aws-folders-48-dark/README.md) | Folders 48 Dark | logical | [SVG](aws-folders-48-dark/sample.svg) |
| [`aws-folders-48-light`](aws-folders-48-light/README.md) | Folders 48 Light | logical | [SVG](aws-folders-48-light/sample.svg) |
| [`aws-forecast`](aws-forecast/README.md) | Amazon Forecast | service | [SVG](aws-forecast/sample.svg) |
| [`aws-forums-48-dark`](aws-forums-48-dark/README.md) | Forums 48 Dark | logical | [SVG](aws-forums-48-dark/sample.svg) |
| [`aws-forums-48-light`](aws-forums-48-light/README.md) | Forums 48 Light | logical | [SVG](aws-forums-48-light/sample.svg) |
| [`aws-fraud-detector`](aws-fraud-detector/README.md) | Amazon Fraud Detector | service | [SVG](aws-fraud-detector/sample.svg) |
| [`aws-freertos`](aws-freertos/README.md) | FreeRTOS | service | [SVG](aws-freertos/sample.svg) |
| [`aws-fsx`](aws-fsx/README.md) | Amazon FSx | service | [SVG](aws-fsx/sample.svg) |
| [`aws-fsx-for-lustre`](aws-fsx-for-lustre/README.md) | Amazon FSx for Lustre | service | [SVG](aws-fsx-for-lustre/sample.svg) |
| [`aws-fsx-for-netapp-ontap`](aws-fsx-for-netapp-ontap/README.md) | Amazon FSx for NetApp ONTAP | service | [SVG](aws-fsx-for-netapp-ontap/sample.svg) |
| [`aws-fsx-for-openzfs`](aws-fsx-for-openzfs/README.md) | Amazon FSx for OpenZFS | service | [SVG](aws-fsx-for-openzfs/sample.svg) |
| [`aws-fsx-for-wfs`](aws-fsx-for-wfs/README.md) | Amazon FSx for WFS | service | [SVG](aws-fsx-for-wfs/sample.svg) |
| [`aws-gamelift-servers`](aws-gamelift-servers/README.md) | Amazon GameLift Servers | service | [SVG](aws-gamelift-servers/sample.svg) |
| [`aws-gamelift-streams`](aws-gamelift-streams/README.md) | Amazon GameLift Streams | service | [SVG](aws-gamelift-streams/sample.svg) |
| [`aws-gear-48-dark`](aws-gear-48-dark/README.md) | Gear 48 Dark | logical | [SVG](aws-gear-48-dark/sample.svg) |
| [`aws-gear-48-light`](aws-gear-48-light/README.md) | Gear 48 Light | logical | [SVG](aws-gear-48-light/sample.svg) |
| [`aws-generic-application-48-dark`](aws-generic-application-48-dark/README.md) | Generic Application 48 Dark | logical | [SVG](aws-generic-application-48-dark/sample.svg) |
| [`aws-generic-application-48-light`](aws-generic-application-48-light/README.md) | Generic Application 48 Light | logical | [SVG](aws-generic-application-48-light/sample.svg) |
| [`aws-git-repository-48-dark`](aws-git-repository-48-dark/README.md) | Git Repository 48 Dark | logical | [SVG](aws-git-repository-48-dark/sample.svg) |
| [`aws-git-repository-48-light`](aws-git-repository-48-light/README.md) | Git Repository 48 Light | logical | [SVG](aws-git-repository-48-light/sample.svg) |
| [`aws-global-accelerator`](aws-global-accelerator/README.md) | AWS Global Accelerator | service | [SVG](aws-global-accelerator/sample.svg) |
| [`aws-globe-48-dark`](aws-globe-48-dark/README.md) | Globe 48 Dark | logical | [SVG](aws-globe-48-dark/sample.svg) |
| [`aws-globe-48-light`](aws-globe-48-light/README.md) | Globe 48 Light | logical | [SVG](aws-globe-48-light/sample.svg) |
| [`aws-glue`](aws-glue/README.md) | AWS Glue | service | [SVG](aws-glue/sample.svg) |
| [`aws-glue-aws-glue-for-ray`](aws-glue-aws-glue-for-ray/README.md) | AWS Glue AWS Glue for Ray | logical | [SVG](aws-glue-aws-glue-for-ray/sample.svg) |
| [`aws-glue-crawler`](aws-glue-crawler/README.md) | AWS Glue Crawler | logical | [SVG](aws-glue-crawler/sample.svg) |
| [`aws-glue-data-catalog`](aws-glue-data-catalog/README.md) | AWS Glue Data Catalog | logical | [SVG](aws-glue-data-catalog/sample.svg) |
| [`aws-glue-data-quality`](aws-glue-data-quality/README.md) | AWS Glue Data Quality | logical | [SVG](aws-glue-data-quality/sample.svg) |
| [`aws-glue-databrew`](aws-glue-databrew/README.md) | AWS Glue DataBrew | service | [SVG](aws-glue-databrew/sample.svg) |
| [`aws-ground-station`](aws-ground-station/README.md) | AWS Ground Station | service | [SVG](aws-ground-station/sample.svg) |
| [`aws-guardduty`](aws-guardduty/README.md) | Amazon GuardDuty | service | [SVG](aws-guardduty/sample.svg) |
| [`aws-health-dashboard`](aws-health-dashboard/README.md) | AWS Health Dashboard | service | [SVG](aws-health-dashboard/sample.svg) |
| [`aws-healthimaging`](aws-healthimaging/README.md) | AWS HealthImaging | service | [SVG](aws-healthimaging/sample.svg) |
| [`aws-healthlake`](aws-healthlake/README.md) | AWS HealthLake | service | [SVG](aws-healthlake/sample.svg) |
| [`aws-healthomics`](aws-healthomics/README.md) | AWS HealthOmics | service | [SVG](aws-healthomics/sample.svg) |
| [`aws-healthscribe`](aws-healthscribe/README.md) | AWS HealthScribe | service | [SVG](aws-healthscribe/sample.svg) |
| [`aws-iam-identity-center`](aws-iam-identity-center/README.md) | AWS IAM Identity Center | service | [SVG](aws-iam-identity-center/sample.svg) |
| [`aws-identity-access-management-add-on`](aws-identity-access-management-add-on/README.md) | AWS Identity Access Management Add on | logical | [SVG](aws-identity-access-management-add-on/sample.svg) |
| [`aws-identity-access-management-aws-sts`](aws-identity-access-management-aws-sts/README.md) | AWS Identity Access Management AWS STS | logical | [SVG](aws-identity-access-management-aws-sts/sample.svg) |
| [`aws-identity-access-management-aws-sts-alternate`](aws-identity-access-management-aws-sts-alternate/README.md) | AWS Identity Access Management AWS STS Alternate | logical | [SVG](aws-identity-access-management-aws-sts-alternate/sample.svg) |
| [`aws-identity-access-management-data-encryption-key`](aws-identity-access-management-data-encryption-key/README.md) | AWS Identity Access Management Data Encryption Key | logical | [SVG](aws-identity-access-management-data-encryption-key/sample.svg) |
| [`aws-identity-access-management-encrypted-data`](aws-identity-access-management-encrypted-data/README.md) | AWS Identity Access Management Encrypted Data | logical | [SVG](aws-identity-access-management-encrypted-data/sample.svg) |
| [`aws-identity-access-management-iam-access-analyzer`](aws-identity-access-management-iam-access-analyzer/README.md) | AWS Identity Access Management IAM Access Analyzer | logical | [SVG](aws-identity-access-management-iam-access-analyzer/sample.svg) |
| [`aws-identity-access-management-iam-roles-anywhere`](aws-identity-access-management-iam-roles-anywhere/README.md) | AWS Identity Access Management IAM Roles Anywhere | logical | [SVG](aws-identity-access-management-iam-roles-anywhere/sample.svg) |
| [`aws-identity-access-management-long-term-security-credential`](aws-identity-access-management-long-term-security-credential/README.md) | AWS Identity Access Management Long Term Security Credential | logical | [SVG](aws-identity-access-management-long-term-security-credential/sample.svg) |
| [`aws-identity-access-management-mfa-token`](aws-identity-access-management-mfa-token/README.md) | AWS Identity Access Management MFA Token | logical | [SVG](aws-identity-access-management-mfa-token/sample.svg) |
| [`aws-identity-access-management-permissions`](aws-identity-access-management-permissions/README.md) | AWS Identity Access Management Permissions | logical | [SVG](aws-identity-access-management-permissions/sample.svg) |
| [`aws-identity-access-management-role`](aws-identity-access-management-role/README.md) | AWS Identity Access Management Role | logical | [SVG](aws-identity-access-management-role/sample.svg) |
| [`aws-identity-access-management-temporary-security-credential`](aws-identity-access-management-temporary-security-credential/README.md) | AWS Identity Access Management Temporary Security Credential | logical | [SVG](aws-identity-access-management-temporary-security-credential/sample.svg) |
| [`aws-identity-and-access-management`](aws-identity-and-access-management/README.md) | AWS Identity and Access Management | global | [SVG](aws-identity-and-access-management/sample.svg) |
| [`aws-infrastructure-composer`](aws-infrastructure-composer/README.md) | AWS Infrastructure Composer | service | [SVG](aws-infrastructure-composer/sample.svg) |
| [`aws-inspector`](aws-inspector/README.md) | Amazon Inspector | service | [SVG](aws-inspector/sample.svg) |
| [`aws-inspector-agent`](aws-inspector-agent/README.md) | Amazon Inspector Agent | logical | [SVG](aws-inspector-agent/sample.svg) |
| [`aws-interactive-video-service`](aws-interactive-video-service/README.md) | Amazon Interactive Video Service | service | [SVG](aws-interactive-video-service/sample.svg) |
| [`aws-internet-48-dark`](aws-internet-48-dark/README.md) | Internet 48 Dark | logical | [SVG](aws-internet-48-dark/sample.svg) |
| [`aws-internet-48-light`](aws-internet-48-light/README.md) | Internet 48 Light | logical | [SVG](aws-internet-48-light/sample.svg) |
| [`aws-internet-alt1-48-dark`](aws-internet-alt1-48-dark/README.md) | Internet alt1 48 Dark | logical | [SVG](aws-internet-alt1-48-dark/sample.svg) |
| [`aws-internet-alt1-48-light`](aws-internet-alt1-48-light/README.md) | Internet alt1 48 Light | logical | [SVG](aws-internet-alt1-48-light/sample.svg) |
| [`aws-internet-alt2-48-dark`](aws-internet-alt2-48-dark/README.md) | Internet alt2 48 Dark | logical | [SVG](aws-internet-alt2-48-dark/sample.svg) |
| [`aws-internet-alt2-48-light`](aws-internet-alt2-48-light/README.md) | Internet alt2 48 Light | logical | [SVG](aws-internet-alt2-48-light/sample.svg) |
| [`aws-iot-action`](aws-iot-action/README.md) | AWS IoT Action | logical | [SVG](aws-iot-action/sample.svg) |
| [`aws-iot-actuator`](aws-iot-actuator/README.md) | AWS IoT Actuator | logical | [SVG](aws-iot-actuator/sample.svg) |
| [`aws-iot-alexa-enabled-device`](aws-iot-alexa-enabled-device/README.md) | AWS IoT Alexa Enabled Device | logical | [SVG](aws-iot-alexa-enabled-device/sample.svg) |
| [`aws-iot-alexa-skill`](aws-iot-alexa-skill/README.md) | AWS IoT Alexa Skill | logical | [SVG](aws-iot-alexa-skill/sample.svg) |
| [`aws-iot-alexa-voice-service`](aws-iot-alexa-voice-service/README.md) | AWS IoT Alexa Voice Service | logical | [SVG](aws-iot-alexa-voice-service/sample.svg) |
| [`aws-iot-analytics`](aws-iot-analytics/README.md) | AWS IoT Analytics | service | [SVG](aws-iot-analytics/sample.svg) |
| [`aws-iot-analytics-channel`](aws-iot-analytics-channel/README.md) | AWS IoT Analytics Channel | logical | [SVG](aws-iot-analytics-channel/sample.svg) |
| [`aws-iot-analytics-data-store`](aws-iot-analytics-data-store/README.md) | AWS IoT Analytics Data Store | logical | [SVG](aws-iot-analytics-data-store/sample.svg) |
| [`aws-iot-analytics-dataset`](aws-iot-analytics-dataset/README.md) | AWS IoT Analytics Dataset | logical | [SVG](aws-iot-analytics-dataset/sample.svg) |
| [`aws-iot-analytics-notebook`](aws-iot-analytics-notebook/README.md) | AWS IoT Analytics Notebook | logical | [SVG](aws-iot-analytics-notebook/sample.svg) |
| [`aws-iot-analytics-pipeline`](aws-iot-analytics-pipeline/README.md) | AWS IoT Analytics Pipeline | logical | [SVG](aws-iot-analytics-pipeline/sample.svg) |
| [`aws-iot-button`](aws-iot-button/README.md) | AWS IoT Button | service | [SVG](aws-iot-button/sample.svg) |
| [`aws-iot-certificate`](aws-iot-certificate/README.md) | AWS IoT Certificate | logical | [SVG](aws-iot-certificate/sample.svg) |
| [`aws-iot-core`](aws-iot-core/README.md) | AWS IoT Core | service | [SVG](aws-iot-core/sample.svg) |
| [`aws-iot-core-device-advisor`](aws-iot-core-device-advisor/README.md) | AWS IoT Core Device Advisor | logical | [SVG](aws-iot-core-device-advisor/sample.svg) |
| [`aws-iot-core-device-location`](aws-iot-core-device-location/README.md) | AWS IoT Core Device Location | logical | [SVG](aws-iot-core-device-location/sample.svg) |
| [`aws-iot-desired-state`](aws-iot-desired-state/README.md) | AWS IoT Desired State | logical | [SVG](aws-iot-desired-state/sample.svg) |
| [`aws-iot-device-defender`](aws-iot-device-defender/README.md) | AWS IoT Device Defender | service | [SVG](aws-iot-device-defender/sample.svg) |
| [`aws-iot-device-defender-iot-device-jobs`](aws-iot-device-defender-iot-device-jobs/README.md) | AWS IoT Device Defender IoT Device Jobs | logical | [SVG](aws-iot-device-defender-iot-device-jobs/sample.svg) |
| [`aws-iot-device-gateway`](aws-iot-device-gateway/README.md) | AWS IoT Device Gateway | logical | [SVG](aws-iot-device-gateway/sample.svg) |
| [`aws-iot-device-management`](aws-iot-device-management/README.md) | AWS IoT Device Management | service | [SVG](aws-iot-device-management/sample.svg) |
| [`aws-iot-device-management-fleet-hub`](aws-iot-device-management-fleet-hub/README.md) | AWS IoT Device Management Fleet Hub | logical | [SVG](aws-iot-device-management-fleet-hub/sample.svg) |
| [`aws-iot-device-tester`](aws-iot-device-tester/README.md) | AWS IoT Device Tester | logical | [SVG](aws-iot-device-tester/sample.svg) |
| [`aws-iot-echo`](aws-iot-echo/README.md) | AWS IoT Echo | logical | [SVG](aws-iot-echo/sample.svg) |
| [`aws-iot-events`](aws-iot-events/README.md) | AWS IoT Events | service | [SVG](aws-iot-events/sample.svg) |
| [`aws-iot-expresslink`](aws-iot-expresslink/README.md) | AWS IoT ExpressLink | service | [SVG](aws-iot-expresslink/sample.svg) |
| [`aws-iot-fire-tv`](aws-iot-fire-tv/README.md) | AWS IoT Fire TV | logical | [SVG](aws-iot-fire-tv/sample.svg) |
| [`aws-iot-fire-tv-stick`](aws-iot-fire-tv-stick/README.md) | AWS IoT Fire TV Stick | logical | [SVG](aws-iot-fire-tv-stick/sample.svg) |
| [`aws-iot-fleetwise`](aws-iot-fleetwise/README.md) | AWS IoT FleetWise | service | [SVG](aws-iot-fleetwise/sample.svg) |
| [`aws-iot-greengrass`](aws-iot-greengrass/README.md) | AWS IoT Greengrass | edge | [SVG](aws-iot-greengrass/sample.svg) |
| [`aws-iot-greengrass-artifact`](aws-iot-greengrass-artifact/README.md) | AWS IoT Greengrass Artifact | logical | [SVG](aws-iot-greengrass-artifact/sample.svg) |
| [`aws-iot-greengrass-component`](aws-iot-greengrass-component/README.md) | AWS IoT Greengrass Component | logical | [SVG](aws-iot-greengrass-component/sample.svg) |
| [`aws-iot-greengrass-component-machine-learning`](aws-iot-greengrass-component-machine-learning/README.md) | AWS IoT Greengrass Component Machine Learning | logical | [SVG](aws-iot-greengrass-component-machine-learning/sample.svg) |
| [`aws-iot-greengrass-component-nucleus`](aws-iot-greengrass-component-nucleus/README.md) | AWS IoT Greengrass Component Nucleus | logical | [SVG](aws-iot-greengrass-component-nucleus/sample.svg) |
| [`aws-iot-greengrass-component-private`](aws-iot-greengrass-component-private/README.md) | AWS IoT Greengrass Component Private | logical | [SVG](aws-iot-greengrass-component-private/sample.svg) |
| [`aws-iot-greengrass-component-public`](aws-iot-greengrass-component-public/README.md) | AWS IoT Greengrass Component Public | logical | [SVG](aws-iot-greengrass-component-public/sample.svg) |
| [`aws-iot-greengrass-connector`](aws-iot-greengrass-connector/README.md) | AWS IoT Greengrass Connector | logical | [SVG](aws-iot-greengrass-connector/sample.svg) |
| [`aws-iot-greengrass-deployment`](aws-iot-greengrass-deployment/README.md) | AWS IoT Greengrass Deployment | edge | [SVG](aws-iot-greengrass-deployment/sample.svg) |
| [`aws-iot-greengrass-interprocess-communication`](aws-iot-greengrass-interprocess-communication/README.md) | AWS IoT Greengrass Interprocess Communication | logical | [SVG](aws-iot-greengrass-interprocess-communication/sample.svg) |
| [`aws-iot-greengrass-protocol`](aws-iot-greengrass-protocol/README.md) | AWS IoT Greengrass Protocol | logical | [SVG](aws-iot-greengrass-protocol/sample.svg) |
| [`aws-iot-greengrass-recipe`](aws-iot-greengrass-recipe/README.md) | AWS IoT Greengrass Recipe | logical | [SVG](aws-iot-greengrass-recipe/sample.svg) |
| [`aws-iot-greengrass-service`](aws-iot-greengrass-service/README.md) | AWS IoT Greengrass | service | [SVG](aws-iot-greengrass-service/sample.svg) |
| [`aws-iot-greengrass-stream-manager`](aws-iot-greengrass-stream-manager/README.md) | AWS IoT Greengrass Stream Manager | logical | [SVG](aws-iot-greengrass-stream-manager/sample.svg) |
| [`aws-iot-hardware-board`](aws-iot-hardware-board/README.md) | AWS IoT Hardware Board | logical | [SVG](aws-iot-hardware-board/sample.svg) |
| [`aws-iot-http-protocol`](aws-iot-http-protocol/README.md) | AWS IoT HTTP Protocol | logical | [SVG](aws-iot-http-protocol/sample.svg) |
| [`aws-iot-http2-protocol`](aws-iot-http2-protocol/README.md) | AWS IoT HTTP2 Protocol | logical | [SVG](aws-iot-http2-protocol/sample.svg) |
| [`aws-iot-lambda-function`](aws-iot-lambda-function/README.md) | AWS IoT Lambda Function | logical | [SVG](aws-iot-lambda-function/sample.svg) |
| [`aws-iot-lorawan-protocol`](aws-iot-lorawan-protocol/README.md) | AWS IoT LoRaWAN Protocol | logical | [SVG](aws-iot-lorawan-protocol/sample.svg) |
| [`aws-iot-mqtt-protocol`](aws-iot-mqtt-protocol/README.md) | AWS IoT MQTT Protocol | logical | [SVG](aws-iot-mqtt-protocol/sample.svg) |
| [`aws-iot-over-air-update`](aws-iot-over-air-update/README.md) | AWS IoT Over Air Update | logical | [SVG](aws-iot-over-air-update/sample.svg) |
| [`aws-iot-policy`](aws-iot-policy/README.md) | AWS IoT Policy | logical | [SVG](aws-iot-policy/sample.svg) |
| [`aws-iot-reported-state`](aws-iot-reported-state/README.md) | AWS IoT Reported State | logical | [SVG](aws-iot-reported-state/sample.svg) |
| [`aws-iot-rule`](aws-iot-rule/README.md) | AWS IoT Rule | logical | [SVG](aws-iot-rule/sample.svg) |
| [`aws-iot-sailboat`](aws-iot-sailboat/README.md) | AWS IoT Sailboat | logical | [SVG](aws-iot-sailboat/sample.svg) |
| [`aws-iot-sensor`](aws-iot-sensor/README.md) | AWS IoT Sensor | logical | [SVG](aws-iot-sensor/sample.svg) |
| [`aws-iot-servo`](aws-iot-servo/README.md) | AWS IoT Servo | logical | [SVG](aws-iot-servo/sample.svg) |
| [`aws-iot-shadow`](aws-iot-shadow/README.md) | AWS IoT Shadow | logical | [SVG](aws-iot-shadow/sample.svg) |
| [`aws-iot-simulator`](aws-iot-simulator/README.md) | AWS IoT Simulator | logical | [SVG](aws-iot-simulator/sample.svg) |
| [`aws-iot-sitewise`](aws-iot-sitewise/README.md) | AWS IoT SiteWise | service | [SVG](aws-iot-sitewise/sample.svg) |
| [`aws-iot-sitewise-asset`](aws-iot-sitewise-asset/README.md) | AWS IoT SiteWise Asset | logical | [SVG](aws-iot-sitewise-asset/sample.svg) |
| [`aws-iot-sitewise-asset-hierarchy`](aws-iot-sitewise-asset-hierarchy/README.md) | AWS IoT SiteWise Asset Hierarchy | logical | [SVG](aws-iot-sitewise-asset-hierarchy/sample.svg) |
| [`aws-iot-sitewise-asset-model`](aws-iot-sitewise-asset-model/README.md) | AWS IoT SiteWise Asset Model | logical | [SVG](aws-iot-sitewise-asset-model/sample.svg) |
| [`aws-iot-sitewise-asset-properties`](aws-iot-sitewise-asset-properties/README.md) | AWS IoT SiteWise Asset Properties | logical | [SVG](aws-iot-sitewise-asset-properties/sample.svg) |
| [`aws-iot-sitewise-data-streams`](aws-iot-sitewise-data-streams/README.md) | AWS IoT SiteWise Data Streams | logical | [SVG](aws-iot-sitewise-data-streams/sample.svg) |
| [`aws-iot-thing-bank`](aws-iot-thing-bank/README.md) | AWS IoT Thing Bank | logical | [SVG](aws-iot-thing-bank/sample.svg) |
| [`aws-iot-thing-bicycle`](aws-iot-thing-bicycle/README.md) | AWS IoT Thing Bicycle | logical | [SVG](aws-iot-thing-bicycle/sample.svg) |
| [`aws-iot-thing-camera`](aws-iot-thing-camera/README.md) | AWS IoT Thing Camera | logical | [SVG](aws-iot-thing-camera/sample.svg) |
| [`aws-iot-thing-car`](aws-iot-thing-car/README.md) | AWS IoT Thing Car | logical | [SVG](aws-iot-thing-car/sample.svg) |
| [`aws-iot-thing-cart`](aws-iot-thing-cart/README.md) | AWS IoT Thing Cart | logical | [SVG](aws-iot-thing-cart/sample.svg) |
| [`aws-iot-thing-coffee-pot`](aws-iot-thing-coffee-pot/README.md) | AWS IoT Thing Coffee Pot | logical | [SVG](aws-iot-thing-coffee-pot/sample.svg) |
| [`aws-iot-thing-door-lock`](aws-iot-thing-door-lock/README.md) | AWS IoT Thing Door Lock | logical | [SVG](aws-iot-thing-door-lock/sample.svg) |
| [`aws-iot-thing-factory`](aws-iot-thing-factory/README.md) | AWS IoT Thing Factory | logical | [SVG](aws-iot-thing-factory/sample.svg) |
| [`aws-iot-thing-freertos-device`](aws-iot-thing-freertos-device/README.md) | AWS IoT Thing FreeRTOS Device | logical | [SVG](aws-iot-thing-freertos-device/sample.svg) |
| [`aws-iot-thing-generic`](aws-iot-thing-generic/README.md) | AWS IoT Thing Generic | logical | [SVG](aws-iot-thing-generic/sample.svg) |
| [`aws-iot-thing-house`](aws-iot-thing-house/README.md) | AWS IoT Thing House | logical | [SVG](aws-iot-thing-house/sample.svg) |
| [`aws-iot-thing-humidity-sensor`](aws-iot-thing-humidity-sensor/README.md) | AWS IoT Thing Humidity Sensor | logical | [SVG](aws-iot-thing-humidity-sensor/sample.svg) |
| [`aws-iot-thing-industrial-pc`](aws-iot-thing-industrial-pc/README.md) | AWS IoT Thing Industrial PC | logical | [SVG](aws-iot-thing-industrial-pc/sample.svg) |
| [`aws-iot-thing-lightbulb`](aws-iot-thing-lightbulb/README.md) | AWS IoT Thing Lightbulb | logical | [SVG](aws-iot-thing-lightbulb/sample.svg) |
| [`aws-iot-thing-medical-emergency`](aws-iot-thing-medical-emergency/README.md) | AWS IoT Thing Medical Emergency | logical | [SVG](aws-iot-thing-medical-emergency/sample.svg) |
| [`aws-iot-thing-plc`](aws-iot-thing-plc/README.md) | AWS IoT Thing PLC | logical | [SVG](aws-iot-thing-plc/sample.svg) |
| [`aws-iot-thing-police-emergency`](aws-iot-thing-police-emergency/README.md) | AWS IoT Thing Police Emergency | logical | [SVG](aws-iot-thing-police-emergency/sample.svg) |
| [`aws-iot-thing-relay`](aws-iot-thing-relay/README.md) | AWS IoT Thing Relay | logical | [SVG](aws-iot-thing-relay/sample.svg) |
| [`aws-iot-thing-stacklight`](aws-iot-thing-stacklight/README.md) | AWS IoT Thing Stacklight | logical | [SVG](aws-iot-thing-stacklight/sample.svg) |
| [`aws-iot-thing-temperature-humidity-sensor`](aws-iot-thing-temperature-humidity-sensor/README.md) | AWS IoT Thing Temperature Humidity Sensor | logical | [SVG](aws-iot-thing-temperature-humidity-sensor/sample.svg) |
| [`aws-iot-thing-temperature-sensor`](aws-iot-thing-temperature-sensor/README.md) | AWS IoT Thing Temperature Sensor | logical | [SVG](aws-iot-thing-temperature-sensor/sample.svg) |
| [`aws-iot-thing-temperature-vibration-sensor`](aws-iot-thing-temperature-vibration-sensor/README.md) | AWS IoT Thing Temperature Vibration Sensor | logical | [SVG](aws-iot-thing-temperature-vibration-sensor/sample.svg) |
| [`aws-iot-thing-thermostat`](aws-iot-thing-thermostat/README.md) | AWS IoT Thing Thermostat | logical | [SVG](aws-iot-thing-thermostat/sample.svg) |
| [`aws-iot-thing-travel`](aws-iot-thing-travel/README.md) | AWS IoT Thing Travel | logical | [SVG](aws-iot-thing-travel/sample.svg) |
| [`aws-iot-thing-utility`](aws-iot-thing-utility/README.md) | AWS IoT Thing Utility | logical | [SVG](aws-iot-thing-utility/sample.svg) |
| [`aws-iot-thing-vibration-sensor`](aws-iot-thing-vibration-sensor/README.md) | AWS IoT Thing Vibration Sensor | logical | [SVG](aws-iot-thing-vibration-sensor/sample.svg) |
| [`aws-iot-thing-windfarm`](aws-iot-thing-windfarm/README.md) | AWS IoT Thing Windfarm | logical | [SVG](aws-iot-thing-windfarm/sample.svg) |
| [`aws-iot-topic`](aws-iot-topic/README.md) | AWS IoT Topic | logical | [SVG](aws-iot-topic/sample.svg) |
| [`aws-iot-twinmaker`](aws-iot-twinmaker/README.md) | AWS IoT TwinMaker | service | [SVG](aws-iot-twinmaker/sample.svg) |
| [`aws-iq`](aws-iq/README.md) | AWS IQ | service | [SVG](aws-iq/sample.svg) |
| [`aws-json-script-48-dark`](aws-json-script-48-dark/README.md) | JSON Script 48 Dark | logical | [SVG](aws-json-script-48-dark/sample.svg) |
| [`aws-json-script-48-light`](aws-json-script-48-light/README.md) | JSON Script 48 Light | logical | [SVG](aws-json-script-48-light/sample.svg) |
| [`aws-kendra`](aws-kendra/README.md) | Amazon Kendra | service | [SVG](aws-kendra/sample.svg) |
| [`aws-key-management-service`](aws-key-management-service/README.md) | AWS Key Management Service | service | [SVG](aws-key-management-service/sample.svg) |
| [`aws-key-management-service-external-key-store`](aws-key-management-service-external-key-store/README.md) | AWS Key Management Service External Key Store | logical | [SVG](aws-key-management-service-external-key-store/sample.svg) |
| [`aws-keyspaces`](aws-keyspaces/README.md) | Amazon Keyspaces | service | [SVG](aws-keyspaces/sample.svg) |
| [`aws-kinesis`](aws-kinesis/README.md) | Amazon Kinesis | service | [SVG](aws-kinesis/sample.svg) |
| [`aws-kinesis-data-streams`](aws-kinesis-data-streams/README.md) | Amazon Kinesis Data Streams | service | [SVG](aws-kinesis-data-streams/sample.svg) |
| [`aws-kinesis-video-streams`](aws-kinesis-video-streams/README.md) | Amazon Kinesis Video Streams | service | [SVG](aws-kinesis-video-streams/sample.svg) |
| [`aws-lake-formation`](aws-lake-formation/README.md) | AWS Lake Formation | service | [SVG](aws-lake-formation/sample.svg) |
| [`aws-lake-formation-data-lake`](aws-lake-formation-data-lake/README.md) | AWS Lake Formation Data Lake | logical | [SVG](aws-lake-formation-data-lake/sample.svg) |
| [`aws-lambda`](aws-lambda/README.md) | AWS Lambda | region | [SVG](aws-lambda/sample.svg) |
| [`aws-lambda-lambda-function`](aws-lambda-lambda-function/README.md) | AWS Lambda Lambda Function | region | [SVG](aws-lambda-lambda-function/sample.svg) |
| [`aws-launch-wizard`](aws-launch-wizard/README.md) | AWS Launch Wizard | service | [SVG](aws-launch-wizard/sample.svg) |
| [`aws-lex`](aws-lex/README.md) | Amazon Lex | service | [SVG](aws-lex/sample.svg) |
| [`aws-license-manager`](aws-license-manager/README.md) | AWS License Manager | service | [SVG](aws-license-manager/sample.svg) |
| [`aws-license-manager-application-discovery`](aws-license-manager-application-discovery/README.md) | AWS License Manager Application Discovery | logical | [SVG](aws-license-manager-application-discovery/sample.svg) |
| [`aws-license-manager-license-blending`](aws-license-manager-license-blending/README.md) | AWS License Manager License Blending | logical | [SVG](aws-license-manager-license-blending/sample.svg) |
| [`aws-lightsail`](aws-lightsail/README.md) | Amazon Lightsail | service | [SVG](aws-lightsail/sample.svg) |
| [`aws-lightsail-for-research`](aws-lightsail-for-research/README.md) | Amazon Lightsail for Research | service | [SVG](aws-lightsail-for-research/sample.svg) |
| [`aws-local-zones`](aws-local-zones/README.md) | AWS Local Zones | service | [SVG](aws-local-zones/sample.svg) |
| [`aws-location-service`](aws-location-service/README.md) | Amazon Location Service | service | [SVG](aws-location-service/sample.svg) |
| [`aws-location-service-geofence`](aws-location-service-geofence/README.md) | Amazon Location Service Geofence | logical | [SVG](aws-location-service-geofence/sample.svg) |
| [`aws-location-service-map`](aws-location-service-map/README.md) | Amazon Location Service Map  | logical | [SVG](aws-location-service-map/sample.svg) |
| [`aws-location-service-place`](aws-location-service-place/README.md) | Amazon Location Service Place | logical | [SVG](aws-location-service-place/sample.svg) |
| [`aws-location-service-routes`](aws-location-service-routes/README.md) | Amazon Location Service Routes | logical | [SVG](aws-location-service-routes/sample.svg) |
| [`aws-location-service-track`](aws-location-service-track/README.md) | Amazon Location Service Track  | logical | [SVG](aws-location-service-track/sample.svg) |
| [`aws-logs-48-dark`](aws-logs-48-dark/README.md) | Logs 48 Dark | logical | [SVG](aws-logs-48-dark/sample.svg) |
| [`aws-logs-48-light`](aws-logs-48-light/README.md) | Logs 48 Light | logical | [SVG](aws-logs-48-light/sample.svg) |
| [`aws-lookout-for-equipment`](aws-lookout-for-equipment/README.md) | Amazon Lookout for Equipment | service | [SVG](aws-lookout-for-equipment/sample.svg) |
| [`aws-lookout-for-metrics`](aws-lookout-for-metrics/README.md) | Amazon Lookout for Metrics | service | [SVG](aws-lookout-for-metrics/sample.svg) |
| [`aws-lookout-for-vision`](aws-lookout-for-vision/README.md) | Amazon Lookout for Vision | service | [SVG](aws-lookout-for-vision/sample.svg) |
| [`aws-macie`](aws-macie/README.md) | Amazon Macie | service | [SVG](aws-macie/sample.svg) |
| [`aws-magnifying-glass-48-dark`](aws-magnifying-glass-48-dark/README.md) | Magnifying Glass 48 Dark | logical | [SVG](aws-magnifying-glass-48-dark/sample.svg) |
| [`aws-magnifying-glass-48-light`](aws-magnifying-glass-48-light/README.md) | Magnifying Glass 48 Light | logical | [SVG](aws-magnifying-glass-48-light/sample.svg) |
| [`aws-mainframe-modernization`](aws-mainframe-modernization/README.md) | AWS Mainframe Modernization | service | [SVG](aws-mainframe-modernization/sample.svg) |
| [`aws-mainframe-modernization-analyzer`](aws-mainframe-modernization-analyzer/README.md) | AWS Mainframe Modernization Analyzer | logical | [SVG](aws-mainframe-modernization-analyzer/sample.svg) |
| [`aws-mainframe-modernization-compiler`](aws-mainframe-modernization-compiler/README.md) | AWS Mainframe Modernization Compiler | logical | [SVG](aws-mainframe-modernization-compiler/sample.svg) |
| [`aws-mainframe-modernization-converter`](aws-mainframe-modernization-converter/README.md) | AWS Mainframe Modernization Converter | logical | [SVG](aws-mainframe-modernization-converter/sample.svg) |
| [`aws-mainframe-modernization-developer`](aws-mainframe-modernization-developer/README.md) | AWS Mainframe Modernization Developer | logical | [SVG](aws-mainframe-modernization-developer/sample.svg) |
| [`aws-mainframe-modernization-runtime`](aws-mainframe-modernization-runtime/README.md) | AWS Mainframe Modernization Runtime | logical | [SVG](aws-mainframe-modernization-runtime/sample.svg) |
| [`aws-managed-blockchain`](aws-managed-blockchain/README.md) | Amazon Managed Blockchain | service | [SVG](aws-managed-blockchain/sample.svg) |
| [`aws-managed-blockchain-blockchain`](aws-managed-blockchain-blockchain/README.md) | Amazon Managed Blockchain Blockchain | logical | [SVG](aws-managed-blockchain-blockchain/sample.svg) |
| [`aws-managed-grafana`](aws-managed-grafana/README.md) | Amazon Managed Grafana | service | [SVG](aws-managed-grafana/sample.svg) |
| [`aws-managed-service-for-apache-flink`](aws-managed-service-for-apache-flink/README.md) | Amazon Managed Service for Apache Flink | service | [SVG](aws-managed-service-for-apache-flink/sample.svg) |
| [`aws-managed-service-for-prometheus`](aws-managed-service-for-prometheus/README.md) | Amazon Managed Service for Prometheus | service | [SVG](aws-managed-service-for-prometheus/sample.svg) |
| [`aws-managed-services`](aws-managed-services/README.md) | AWS Managed Services | service | [SVG](aws-managed-services/sample.svg) |
| [`aws-managed-streaming-for-apache-kafka`](aws-managed-streaming-for-apache-kafka/README.md) | Amazon Managed Streaming for Apache Kafka | service | [SVG](aws-managed-streaming-for-apache-kafka/sample.svg) |
| [`aws-managed-workflows-for-apache-airflow`](aws-managed-workflows-for-apache-airflow/README.md) | Amazon Managed Workflows for Apache Airflow | service | [SVG](aws-managed-workflows-for-apache-airflow/sample.svg) |
| [`aws-management-console`](aws-management-console/README.md) | AWS Management Console | service | [SVG](aws-management-console/sample.svg) |
| [`aws-management-console-48-dark`](aws-management-console-48-dark/README.md) | AWS Management Console 48 Dark | logical | [SVG](aws-management-console-48-dark/sample.svg) |
| [`aws-management-console-48-light`](aws-management-console-48-light/README.md) | AWS Management Console 48 Light | logical | [SVG](aws-management-console-48-light/sample.svg) |
| [`aws-marketplace-dark`](aws-marketplace-dark/README.md) | AWS Marketplace Dark | service | [SVG](aws-marketplace-dark/sample.svg) |
| [`aws-marketplace-light`](aws-marketplace-light/README.md) | AWS Marketplace Light | service | [SVG](aws-marketplace-light/sample.svg) |
| [`aws-memorydb`](aws-memorydb/README.md) | Amazon MemoryDB | service | [SVG](aws-memorydb/sample.svg) |
| [`aws-metrics-48-dark`](aws-metrics-48-dark/README.md) | Metrics 48 Dark | logical | [SVG](aws-metrics-48-dark/sample.svg) |
| [`aws-metrics-48-light`](aws-metrics-48-light/README.md) | Metrics 48 Light | logical | [SVG](aws-metrics-48-light/sample.svg) |
| [`aws-migration-evaluator`](aws-migration-evaluator/README.md) | AWS Migration Evaluator | service | [SVG](aws-migration-evaluator/sample.svg) |
| [`aws-migration-hub`](aws-migration-hub/README.md) | AWS Migration Hub | service | [SVG](aws-migration-hub/sample.svg) |
| [`aws-migration-hub-refactor-spaces-applications`](aws-migration-hub-refactor-spaces-applications/README.md) | AWS Migration Hub Refactor Spaces Applications | logical | [SVG](aws-migration-hub-refactor-spaces-applications/sample.svg) |
| [`aws-migration-hub-refactor-spaces-environments`](aws-migration-hub-refactor-spaces-environments/README.md) | AWS Migration Hub Refactor Spaces Environments | logical | [SVG](aws-migration-hub-refactor-spaces-environments/sample.svg) |
| [`aws-migration-hub-refactor-spaces-services`](aws-migration-hub-refactor-spaces-services/README.md) | AWS Migration Hub Refactor Spaces Services | logical | [SVG](aws-migration-hub-refactor-spaces-services/sample.svg) |
| [`aws-mobile-client-48-dark`](aws-mobile-client-48-dark/README.md) | Mobile client 48 Dark | logical | [SVG](aws-mobile-client-48-dark/sample.svg) |
| [`aws-mobile-client-48-light`](aws-mobile-client-48-light/README.md) | Mobile client 48 Light | logical | [SVG](aws-mobile-client-48-light/sample.svg) |
| [`aws-monitron`](aws-monitron/README.md) | Amazon Monitron | service | [SVG](aws-monitron/sample.svg) |
| [`aws-mq`](aws-mq/README.md) | Amazon MQ | service | [SVG](aws-mq/sample.svg) |
| [`aws-mq-broker`](aws-mq-broker/README.md) | Amazon MQ Broker | logical | [SVG](aws-mq-broker/sample.svg) |
| [`aws-msk-amazon-msk-connect`](aws-msk-amazon-msk-connect/README.md) | Amazon MSK Amazon MSK Connect | logical | [SVG](aws-msk-amazon-msk-connect/sample.svg) |
| [`aws-multimedia-48-dark`](aws-multimedia-48-dark/README.md) | Multimedia 48 Dark | logical | [SVG](aws-multimedia-48-dark/sample.svg) |
| [`aws-multimedia-48-light`](aws-multimedia-48-light/README.md) | Multimedia 48 Light | logical | [SVG](aws-multimedia-48-light/sample.svg) |
| [`aws-neptune`](aws-neptune/README.md) | Amazon Neptune | service | [SVG](aws-neptune/sample.svg) |
| [`aws-network-firewall`](aws-network-firewall/README.md) | AWS Network Firewall | service | [SVG](aws-network-firewall/sample.svg) |
| [`aws-network-firewall-endpoints`](aws-network-firewall-endpoints/README.md) | AWS Network Firewall Endpoints | logical | [SVG](aws-network-firewall-endpoints/sample.svg) |
| [`aws-neuron`](aws-neuron/README.md) | AWS Neuron | service | [SVG](aws-neuron/sample.svg) |
| [`aws-nice-enginframe`](aws-nice-enginframe/README.md) | NICE EnginFrame | service | [SVG](aws-nice-enginframe/sample.svg) |
| [`aws-nitro-enclaves`](aws-nitro-enclaves/README.md) | AWS Nitro Enclaves | service | [SVG](aws-nitro-enclaves/sample.svg) |
| [`aws-nova`](aws-nova/README.md) | Amazon Nova | service | [SVG](aws-nova/sample.svg) |
| [`aws-office-building-48-dark`](aws-office-building-48-dark/README.md) | Office building 48 Dark | logical | [SVG](aws-office-building-48-dark/sample.svg) |
| [`aws-office-building-48-light`](aws-office-building-48-light/README.md) | Office building 48 Light | logical | [SVG](aws-office-building-48-light/sample.svg) |
| [`aws-open-3d-engine`](aws-open-3d-engine/README.md) | Open 3D Engine | service | [SVG](aws-open-3d-engine/sample.svg) |
| [`aws-opensearch-service`](aws-opensearch-service/README.md) | Amazon OpenSearch Service | service | [SVG](aws-opensearch-service/sample.svg) |
| [`aws-opensearch-service-cluster-administrator-node`](aws-opensearch-service-cluster-administrator-node/README.md) | Amazon OpenSearch Service Cluster Administrator Node | logical | [SVG](aws-opensearch-service-cluster-administrator-node/sample.svg) |
| [`aws-opensearch-service-data-node`](aws-opensearch-service-data-node/README.md) | Amazon OpenSearch Service Data Node | logical | [SVG](aws-opensearch-service-data-node/sample.svg) |
| [`aws-opensearch-service-index`](aws-opensearch-service-index/README.md) | Amazon OpenSearch Service Index | logical | [SVG](aws-opensearch-service-index/sample.svg) |
| [`aws-opensearch-service-observability`](aws-opensearch-service-observability/README.md) | Amazon OpenSearch Service Observability | logical | [SVG](aws-opensearch-service-observability/sample.svg) |
| [`aws-opensearch-service-opensearch-dashboards`](aws-opensearch-service-opensearch-dashboards/README.md) | Amazon OpenSearch Service OpenSearch Dashboards | logical | [SVG](aws-opensearch-service-opensearch-dashboards/sample.svg) |
| [`aws-opensearch-service-opensearch-ingestion`](aws-opensearch-service-opensearch-ingestion/README.md) | Amazon OpenSearch Service OpenSearch Ingestion | logical | [SVG](aws-opensearch-service-opensearch-ingestion/sample.svg) |
| [`aws-opensearch-service-traces`](aws-opensearch-service-traces/README.md) | Amazon OpenSearch Service Traces | logical | [SVG](aws-opensearch-service-traces/sample.svg) |
| [`aws-opensearch-service-ultrawarm-node`](aws-opensearch-service-ultrawarm-node/README.md) | Amazon OpenSearch Service UltraWarm Node | logical | [SVG](aws-opensearch-service-ultrawarm-node/sample.svg) |
| [`aws-oracle-database-at-aws`](aws-oracle-database-at-aws/README.md) | Oracle Database at AWS | service | [SVG](aws-oracle-database-at-aws/sample.svg) |
| [`aws-organizations`](aws-organizations/README.md) | AWS Organizations | service | [SVG](aws-organizations/sample.svg) |
| [`aws-organizations-account`](aws-organizations-account/README.md) | AWS Organizations Account | logical | [SVG](aws-organizations-account/sample.svg) |
| [`aws-organizations-management-account`](aws-organizations-management-account/README.md) | AWS Organizations Management Account | logical | [SVG](aws-organizations-management-account/sample.svg) |
| [`aws-organizations-organizational-unit`](aws-organizations-organizational-unit/README.md) | AWS Organizations Organizational Unit | logical | [SVG](aws-organizations-organizational-unit/sample.svg) |
| [`aws-outposts-family`](aws-outposts-family/README.md) | AWS Outposts family | service | [SVG](aws-outposts-family/sample.svg) |
| [`aws-outposts-rack`](aws-outposts-rack/README.md) | AWS Outposts rack | service | [SVG](aws-outposts-rack/sample.svg) |
| [`aws-outposts-servers`](aws-outposts-servers/README.md) | AWS Outposts servers | service | [SVG](aws-outposts-servers/sample.svg) |
| [`aws-panorama`](aws-panorama/README.md) | AWS Panorama | service | [SVG](aws-panorama/sample.svg) |
| [`aws-parallel-cluster`](aws-parallel-cluster/README.md) | AWS Parallel Cluster | service | [SVG](aws-parallel-cluster/sample.svg) |
| [`aws-parallel-computing-service`](aws-parallel-computing-service/README.md) | AWS Parallel Computing Service | service | [SVG](aws-parallel-computing-service/sample.svg) |
| [`aws-payment-cryptography`](aws-payment-cryptography/README.md) | AWS Payment Cryptography | service | [SVG](aws-payment-cryptography/sample.svg) |
| [`aws-personalize`](aws-personalize/README.md) | Amazon Personalize | service | [SVG](aws-personalize/sample.svg) |
| [`aws-pinpoint`](aws-pinpoint/README.md) | Amazon Pinpoint | service | [SVG](aws-pinpoint/sample.svg) |
| [`aws-pinpoint-apis`](aws-pinpoint-apis/README.md) | Amazon Pinpoint APIs | service | [SVG](aws-pinpoint-apis/sample.svg) |
| [`aws-pinpoint-journey`](aws-pinpoint-journey/README.md) | Amazon Pinpoint Journey | logical | [SVG](aws-pinpoint-journey/sample.svg) |
| [`aws-polly`](aws-polly/README.md) | Amazon Polly | service | [SVG](aws-polly/sample.svg) |
| [`aws-private-5g`](aws-private-5g/README.md) | AWS Private 5G | service | [SVG](aws-private-5g/sample.svg) |
| [`aws-private-certificate-authority`](aws-private-certificate-authority/README.md) | AWS Private Certificate Authority | service | [SVG](aws-private-certificate-authority/sample.svg) |
| [`aws-privatelink`](aws-privatelink/README.md) | AWS PrivateLink | service | [SVG](aws-privatelink/sample.svg) |
| [`aws-professional-services`](aws-professional-services/README.md) | AWS Professional Services | service | [SVG](aws-professional-services/sample.svg) |
| [`aws-programming-language-48-dark`](aws-programming-language-48-dark/README.md) | Programming Language 48 Dark | logical | [SVG](aws-programming-language-48-dark/sample.svg) |
| [`aws-programming-language-48-light`](aws-programming-language-48-light/README.md) | Programming Language 48 Light | logical | [SVG](aws-programming-language-48-light/sample.svg) |
| [`aws-proton`](aws-proton/README.md) | AWS Proton | service | [SVG](aws-proton/sample.svg) |
| [`aws-pytorch-on-aws`](aws-pytorch-on-aws/README.md) | PyTorch on AWS | service | [SVG](aws-pytorch-on-aws/sample.svg) |
| [`aws-q`](aws-q/README.md) | Amazon Q | service | [SVG](aws-q/sample.svg) |
| [`aws-quantum-ledger-database`](aws-quantum-ledger-database/README.md) | Amazon Quantum Ledger Database | service | [SVG](aws-quantum-ledger-database/sample.svg) |
| [`aws-question-48-dark`](aws-question-48-dark/README.md) | Question 48 Dark | logical | [SVG](aws-question-48-dark/sample.svg) |
| [`aws-question-48-light`](aws-question-48-light/README.md) | Question 48 Light | logical | [SVG](aws-question-48-light/sample.svg) |
| [`aws-quicksight`](aws-quicksight/README.md) | Amazon QuickSight | service | [SVG](aws-quicksight/sample.svg) |
| [`aws-quicksight-paginated-reports`](aws-quicksight-paginated-reports/README.md) | Amazon Quicksight Paginated Reports | logical | [SVG](aws-quicksight-paginated-reports/sample.svg) |
| [`aws-rds`](aws-rds/README.md) | Amazon RDS | region | [SVG](aws-rds/sample.svg) |
| [`aws-rds-blue-green-deployments`](aws-rds-blue-green-deployments/README.md) | Amazon RDS Blue Green Deployments | logical | [SVG](aws-rds-blue-green-deployments/sample.svg) |
| [`aws-rds-multi-az`](aws-rds-multi-az/README.md) | Amazon RDS Multi AZ | logical | [SVG](aws-rds-multi-az/sample.svg) |
| [`aws-rds-multi-az-db-cluster`](aws-rds-multi-az-db-cluster/README.md) | Amazon RDS Multi AZ DB Cluster | logical | [SVG](aws-rds-multi-az-db-cluster/sample.svg) |
| [`aws-rds-optimized-writes`](aws-rds-optimized-writes/README.md) | Amazon RDS Optimized Writes | logical | [SVG](aws-rds-optimized-writes/sample.svg) |
| [`aws-rds-proxy-instance`](aws-rds-proxy-instance/README.md) | Amazon RDS Proxy Instance | logical | [SVG](aws-rds-proxy-instance/sample.svg) |
| [`aws-rds-proxy-instance-alternate`](aws-rds-proxy-instance-alternate/README.md) | Amazon RDS Proxy Instance Alternate | logical | [SVG](aws-rds-proxy-instance-alternate/sample.svg) |
| [`aws-rds-trusted-language-extensions-for-postgresql`](aws-rds-trusted-language-extensions-for-postgresql/README.md) | Amazon RDS Trusted Language Extensions for PostgreSQL | logical | [SVG](aws-rds-trusted-language-extensions-for-postgresql/sample.svg) |
| [`aws-recover-48-dark`](aws-recover-48-dark/README.md) | Recover 48 Dark | logical | [SVG](aws-recover-48-dark/sample.svg) |
| [`aws-recover-48-light`](aws-recover-48-light/README.md) | Recover 48 Light | logical | [SVG](aws-recover-48-light/sample.svg) |
| [`aws-red-hat-openshift-service-on-aws`](aws-red-hat-openshift-service-on-aws/README.md) | Red Hat OpenShift Service on AWS | service | [SVG](aws-red-hat-openshift-service-on-aws/sample.svg) |
| [`aws-redshift`](aws-redshift/README.md) | Amazon Redshift | service | [SVG](aws-redshift/sample.svg) |
| [`aws-redshift-auto-copy`](aws-redshift-auto-copy/README.md) | Amazon Redshift Auto copy | logical | [SVG](aws-redshift-auto-copy/sample.svg) |
| [`aws-redshift-data-sharing-governance`](aws-redshift-data-sharing-governance/README.md) | Amazon Redshift Data Sharing Governance | logical | [SVG](aws-redshift-data-sharing-governance/sample.svg) |
| [`aws-redshift-dense-compute-node`](aws-redshift-dense-compute-node/README.md) | Amazon Redshift Dense Compute Node | logical | [SVG](aws-redshift-dense-compute-node/sample.svg) |
| [`aws-redshift-dense-storage-node`](aws-redshift-dense-storage-node/README.md) | Amazon Redshift Dense Storage Node | logical | [SVG](aws-redshift-dense-storage-node/sample.svg) |
| [`aws-redshift-ml`](aws-redshift-ml/README.md) | Amazon Redshift ML | logical | [SVG](aws-redshift-ml/sample.svg) |
| [`aws-redshift-query-editor-v2-0`](aws-redshift-query-editor-v2-0/README.md) | Amazon Redshift Query Editor v2.0 | logical | [SVG](aws-redshift-query-editor-v2-0/sample.svg) |
| [`aws-redshift-ra3`](aws-redshift-ra3/README.md) | Amazon Redshift RA3 | logical | [SVG](aws-redshift-ra3/sample.svg) |
| [`aws-redshift-streaming-ingestion`](aws-redshift-streaming-ingestion/README.md) | Amazon Redshift Streaming Ingestion | logical | [SVG](aws-redshift-streaming-ingestion/sample.svg) |
| [`aws-rekognition`](aws-rekognition/README.md) | Amazon Rekognition | service | [SVG](aws-rekognition/sample.svg) |
| [`aws-rekognition-image`](aws-rekognition-image/README.md) | Amazon Rekognition Image | logical | [SVG](aws-rekognition-image/sample.svg) |
| [`aws-rekognition-video`](aws-rekognition-video/README.md) | Amazon Rekognition Video | logical | [SVG](aws-rekognition-video/sample.svg) |
| [`aws-repost`](aws-repost/README.md) | AWS rePost | service | [SVG](aws-repost/sample.svg) |
| [`aws-repost-private`](aws-repost-private/README.md) | AWS rePost Private | service | [SVG](aws-repost-private/sample.svg) |
| [`aws-reserved-instance-reporting`](aws-reserved-instance-reporting/README.md) | Reserved Instance Reporting | service | [SVG](aws-reserved-instance-reporting/sample.svg) |
| [`aws-resilience-hub`](aws-resilience-hub/README.md) | AWS Resilience Hub | service | [SVG](aws-resilience-hub/sample.svg) |
| [`aws-resource-access-manager`](aws-resource-access-manager/README.md) | AWS Resource Access Manager | service | [SVG](aws-resource-access-manager/sample.svg) |
| [`aws-resource-explorer`](aws-resource-explorer/README.md) | AWS Resource Explorer | service | [SVG](aws-resource-explorer/sample.svg) |
| [`aws-route-53`](aws-route-53/README.md) | Amazon Route 53 | global | [SVG](aws-route-53/sample.svg) |
| [`aws-route-53-hosted-zone`](aws-route-53-hosted-zone/README.md) | Amazon Route 53 Hosted Zone | logical | [SVG](aws-route-53-hosted-zone/sample.svg) |
| [`aws-route-53-readiness-checks`](aws-route-53-readiness-checks/README.md) | Amazon Route 53 Readiness Checks | logical | [SVG](aws-route-53-readiness-checks/sample.svg) |
| [`aws-route-53-resolver`](aws-route-53-resolver/README.md) | Amazon Route 53 Resolver | logical | [SVG](aws-route-53-resolver/sample.svg) |
| [`aws-route-53-resolver-dns-firewall`](aws-route-53-resolver-dns-firewall/README.md) | Amazon Route 53 Resolver DNS Firewall | logical | [SVG](aws-route-53-resolver-dns-firewall/sample.svg) |
| [`aws-route-53-resolver-query-logging`](aws-route-53-resolver-query-logging/README.md) | Amazon Route 53 Resolver Query Logging | logical | [SVG](aws-route-53-resolver-query-logging/sample.svg) |
| [`aws-route-53-route-table`](aws-route-53-route-table/README.md) | Amazon Route 53 Route Table | logical | [SVG](aws-route-53-route-table/sample.svg) |
| [`aws-route-53-routing-controls`](aws-route-53-routing-controls/README.md) | Amazon Route 53 Routing Controls | logical | [SVG](aws-route-53-routing-controls/sample.svg) |
| [`aws-s3`](aws-s3/README.md) | Amazon Simple Storage Service | region | [SVG](aws-s3/sample.svg) |
| [`aws-s3-bucket`](aws-s3-bucket/README.md) | Amazon Simple Storage Service Bucket | region | [SVG](aws-s3-bucket/sample.svg) |
| [`aws-s3-bucket-with-objects`](aws-s3-bucket-with-objects/README.md) | Amazon Simple Storage Service Bucket With Objects | logical | [SVG](aws-s3-bucket-with-objects/sample.svg) |
| [`aws-s3-directory-bucket`](aws-s3-directory-bucket/README.md) | Amazon Simple Storage Service Directory bucket | logical | [SVG](aws-s3-directory-bucket/sample.svg) |
| [`aws-s3-general-access-points`](aws-s3-general-access-points/README.md) | Amazon Simple Storage Service General Access Points | logical | [SVG](aws-s3-general-access-points/sample.svg) |
| [`aws-s3-glacier`](aws-s3-glacier/README.md) | Amazon Simple Storage Service Glacier | service | [SVG](aws-s3-glacier/sample.svg) |
| [`aws-s3-glacier-archive`](aws-s3-glacier-archive/README.md) | Amazon Simple Storage Service Glacier Archive | logical | [SVG](aws-s3-glacier-archive/sample.svg) |
| [`aws-s3-glacier-vault`](aws-s3-glacier-vault/README.md) | Amazon Simple Storage Service Glacier Vault | logical | [SVG](aws-s3-glacier-vault/sample.svg) |
| [`aws-s3-object`](aws-s3-object/README.md) | Amazon Simple Storage Service Object | logical | [SVG](aws-s3-object/sample.svg) |
| [`aws-s3-on-outposts`](aws-s3-on-outposts/README.md) | Amazon S3 on Outposts | service | [SVG](aws-s3-on-outposts/sample.svg) |
| [`aws-s3-s3-batch-operations`](aws-s3-s3-batch-operations/README.md) | Amazon Simple Storage Service S3 Batch Operations | logical | [SVG](aws-s3-s3-batch-operations/sample.svg) |
| [`aws-s3-s3-express-one-zone`](aws-s3-s3-express-one-zone/README.md) | Amazon Simple Storage Service S3 Express One Zone | logical | [SVG](aws-s3-s3-express-one-zone/sample.svg) |
| [`aws-s3-s3-glacier-deep-archive`](aws-s3-s3-glacier-deep-archive/README.md) | Amazon Simple Storage Service S3 Glacier Deep Archive | logical | [SVG](aws-s3-s3-glacier-deep-archive/sample.svg) |
| [`aws-s3-s3-glacier-flexible-retrieval`](aws-s3-s3-glacier-flexible-retrieval/README.md) | Amazon Simple Storage Service S3 Glacier Flexible Retrieval | logical | [SVG](aws-s3-s3-glacier-flexible-retrieval/sample.svg) |
| [`aws-s3-s3-glacier-instant-retrieval`](aws-s3-s3-glacier-instant-retrieval/README.md) | Amazon Simple Storage Service S3 Glacier Instant Retrieval | logical | [SVG](aws-s3-s3-glacier-instant-retrieval/sample.svg) |
| [`aws-s3-s3-intelligent-tiering`](aws-s3-s3-intelligent-tiering/README.md) | Amazon Simple Storage Service S3 Intelligent Tiering | logical | [SVG](aws-s3-s3-intelligent-tiering/sample.svg) |
| [`aws-s3-s3-multi-region-access-points`](aws-s3-s3-multi-region-access-points/README.md) | Amazon Simple Storage Service S3 Multi Region Access Points | logical | [SVG](aws-s3-s3-multi-region-access-points/sample.svg) |
| [`aws-s3-s3-object-lambda`](aws-s3-s3-object-lambda/README.md) | Amazon Simple Storage Service S3 Object Lambda | logical | [SVG](aws-s3-s3-object-lambda/sample.svg) |
| [`aws-s3-s3-object-lambda-access-points`](aws-s3-s3-object-lambda-access-points/README.md) | Amazon Simple Storage Service S3 Object Lambda Access Points | logical | [SVG](aws-s3-s3-object-lambda-access-points/sample.svg) |
| [`aws-s3-s3-object-lock`](aws-s3-s3-object-lock/README.md) | Amazon Simple Storage Service S3 Object Lock | logical | [SVG](aws-s3-s3-object-lock/sample.svg) |
| [`aws-s3-s3-on-outposts`](aws-s3-s3-on-outposts/README.md) | Amazon Simple Storage Service S3 On Outposts | logical | [SVG](aws-s3-s3-on-outposts/sample.svg) |
| [`aws-s3-s3-one-zone-ia`](aws-s3-s3-one-zone-ia/README.md) | Amazon Simple Storage Service S3 One Zone IA | logical | [SVG](aws-s3-s3-one-zone-ia/sample.svg) |
| [`aws-s3-s3-replication`](aws-s3-s3-replication/README.md) | Amazon Simple Storage Service S3 Replication | logical | [SVG](aws-s3-s3-replication/sample.svg) |
| [`aws-s3-s3-replication-time-control`](aws-s3-s3-replication-time-control/README.md) | Amazon Simple Storage Service S3 Replication Time Control | logical | [SVG](aws-s3-s3-replication-time-control/sample.svg) |
| [`aws-s3-s3-select`](aws-s3-s3-select/README.md) | Amazon Simple Storage Service S3 Select | logical | [SVG](aws-s3-s3-select/sample.svg) |
| [`aws-s3-s3-standard`](aws-s3-s3-standard/README.md) | Amazon Simple Storage Service S3 Standard | logical | [SVG](aws-s3-s3-standard/sample.svg) |
| [`aws-s3-s3-standard-ia`](aws-s3-s3-standard-ia/README.md) | Amazon Simple Storage Service S3 Standard IA | logical | [SVG](aws-s3-s3-standard-ia/sample.svg) |
| [`aws-s3-s3-storage-lens`](aws-s3-s3-storage-lens/README.md) | Amazon Simple Storage Service S3 Storage Lens | logical | [SVG](aws-s3-s3-storage-lens/sample.svg) |
| [`aws-s3-s3-tables`](aws-s3-s3-tables/README.md) | Amazon Simple Storage Service S3 Tables | logical | [SVG](aws-s3-s3-tables/sample.svg) |
| [`aws-s3-s3-vectors`](aws-s3-s3-vectors/README.md) | Amazon Simple Storage Service S3 Vectors | logical | [SVG](aws-s3-s3-vectors/sample.svg) |
| [`aws-s3-vpc-access-points`](aws-s3-vpc-access-points/README.md) | Amazon Simple Storage Service VPC Access Points | logical | [SVG](aws-s3-vpc-access-points/sample.svg) |
| [`aws-sagemaker`](aws-sagemaker/README.md) | Amazon SageMaker | service | [SVG](aws-sagemaker/sample.svg) |
| [`aws-sagemaker-ai`](aws-sagemaker-ai/README.md) | Amazon SageMaker AI | service | [SVG](aws-sagemaker-ai/sample.svg) |
| [`aws-sagemaker-ai-canvas`](aws-sagemaker-ai-canvas/README.md) | Amazon SageMaker AI Canvas | logical | [SVG](aws-sagemaker-ai-canvas/sample.svg) |
| [`aws-sagemaker-ai-geospatial-ml`](aws-sagemaker-ai-geospatial-ml/README.md) | Amazon SageMaker AI Geospatial ML | logical | [SVG](aws-sagemaker-ai-geospatial-ml/sample.svg) |
| [`aws-sagemaker-ai-model`](aws-sagemaker-ai-model/README.md) | Amazon SageMaker AI Model | logical | [SVG](aws-sagemaker-ai-model/sample.svg) |
| [`aws-sagemaker-ai-notebook`](aws-sagemaker-ai-notebook/README.md) | Amazon SageMaker AI Notebook | logical | [SVG](aws-sagemaker-ai-notebook/sample.svg) |
| [`aws-sagemaker-ai-shadow-testing`](aws-sagemaker-ai-shadow-testing/README.md) | Amazon SageMaker AI Shadow Testing | logical | [SVG](aws-sagemaker-ai-shadow-testing/sample.svg) |
| [`aws-sagemaker-ai-train`](aws-sagemaker-ai-train/README.md) | Amazon SageMaker AI Train | logical | [SVG](aws-sagemaker-ai-train/sample.svg) |
| [`aws-sagemaker-ground-truth`](aws-sagemaker-ground-truth/README.md) | Amazon SageMaker Ground Truth | service | [SVG](aws-sagemaker-ground-truth/sample.svg) |
| [`aws-sagemaker-studio-lab`](aws-sagemaker-studio-lab/README.md) | Amazon SageMaker Studio Lab | service | [SVG](aws-sagemaker-studio-lab/sample.svg) |
| [`aws-saml-token-48-dark`](aws-saml-token-48-dark/README.md) | SAML token 48 Dark | logical | [SVG](aws-saml-token-48-dark/sample.svg) |
| [`aws-saml-token-48-light`](aws-saml-token-48-light/README.md) | SAML token 48 Light | logical | [SVG](aws-saml-token-48-light/sample.svg) |
| [`aws-savings-plans`](aws-savings-plans/README.md) | Savings Plans | service | [SVG](aws-savings-plans/sample.svg) |
| [`aws-sdk-48-dark`](aws-sdk-48-dark/README.md) | SDK 48 Dark | logical | [SVG](aws-sdk-48-dark/sample.svg) |
| [`aws-sdk-48-light`](aws-sdk-48-light/README.md) | SDK 48 Light | logical | [SVG](aws-sdk-48-light/sample.svg) |
| [`aws-secrets-manager`](aws-secrets-manager/README.md) | AWS Secrets Manager | service | [SVG](aws-secrets-manager/sample.svg) |
| [`aws-security-hub`](aws-security-hub/README.md) | AWS Security Hub | service | [SVG](aws-security-hub/sample.svg) |
| [`aws-security-hub-finding`](aws-security-hub-finding/README.md) | AWS Security Hub Finding | logical | [SVG](aws-security-hub-finding/sample.svg) |
| [`aws-security-incident-response`](aws-security-incident-response/README.md) | AWS Security Incident Response | service | [SVG](aws-security-incident-response/sample.svg) |
| [`aws-security-lake`](aws-security-lake/README.md) | Amazon Security Lake | service | [SVG](aws-security-lake/sample.svg) |
| [`aws-server-48-dark`](aws-server-48-dark/README.md) | Server 48 Dark | logical | [SVG](aws-server-48-dark/sample.svg) |
| [`aws-server-48-light`](aws-server-48-light/README.md) | Server 48 Light | logical | [SVG](aws-server-48-light/sample.svg) |
| [`aws-serverless-application-repository`](aws-serverless-application-repository/README.md) | AWS Serverless Application Repository | service | [SVG](aws-serverless-application-repository/sample.svg) |
| [`aws-servers-48-dark`](aws-servers-48-dark/README.md) | Servers 48 Dark | logical | [SVG](aws-servers-48-dark/sample.svg) |
| [`aws-servers-48-light`](aws-servers-48-light/README.md) | Servers 48 Light | logical | [SVG](aws-servers-48-light/sample.svg) |
| [`aws-service-catalog`](aws-service-catalog/README.md) | AWS Service Catalog | service | [SVG](aws-service-catalog/sample.svg) |
| [`aws-service-management-connector`](aws-service-management-connector/README.md) | AWS Service Management Connector | service | [SVG](aws-service-management-connector/sample.svg) |
| [`aws-shield`](aws-shield/README.md) | AWS Shield | service | [SVG](aws-shield/sample.svg) |
| [`aws-shield-48-dark`](aws-shield-48-dark/README.md) | Shield 48 Dark | logical | [SVG](aws-shield-48-dark/sample.svg) |
| [`aws-shield-48-light`](aws-shield-48-light/README.md) | Shield 48 Light | logical | [SVG](aws-shield-48-light/sample.svg) |
| [`aws-shield-aws-shield-advanced`](aws-shield-aws-shield-advanced/README.md) | AWS Shield AWS Shield Advanced | logical | [SVG](aws-shield-aws-shield-advanced/sample.svg) |
| [`aws-signer`](aws-signer/README.md) | AWS Signer | service | [SVG](aws-signer/sample.svg) |
| [`aws-simple-email-service`](aws-simple-email-service/README.md) | Amazon Simple Email Service | service | [SVG](aws-simple-email-service/sample.svg) |
| [`aws-simple-email-service-email`](aws-simple-email-service-email/README.md) | Amazon Simple Email Service Email | logical | [SVG](aws-simple-email-service-email/sample.svg) |
| [`aws-simspace-weaver`](aws-simspace-weaver/README.md) | AWS SimSpace Weaver | service | [SVG](aws-simspace-weaver/sample.svg) |
| [`aws-site-to-site-vpn`](aws-site-to-site-vpn/README.md) | AWS Site to Site VPN | service | [SVG](aws-site-to-site-vpn/sample.svg) |
| [`aws-snowball`](aws-snowball/README.md) | AWS Snowball | service | [SVG](aws-snowball/sample.svg) |
| [`aws-snowball-edge`](aws-snowball-edge/README.md) | AWS Snowball Edge | service | [SVG](aws-snowball-edge/sample.svg) |
| [`aws-snowball-snowball-import-export`](aws-snowball-snowball-import-export/README.md) | AWS Snowball Snowball Import Export | logical | [SVG](aws-snowball-snowball-import-export/sample.svg) |
| [`aws-sns`](aws-sns/README.md) | Amazon Simple Notification Service | region | [SVG](aws-sns/sample.svg) |
| [`aws-sns-email-notification`](aws-sns-email-notification/README.md) | Amazon Simple Notification Service Email Notification | logical | [SVG](aws-sns-email-notification/sample.svg) |
| [`aws-sns-http-notification`](aws-sns-http-notification/README.md) | Amazon Simple Notification Service HTTP Notification | logical | [SVG](aws-sns-http-notification/sample.svg) |
| [`aws-sns-topic`](aws-sns-topic/README.md) | Amazon Simple Notification Service Topic | region | [SVG](aws-sns-topic/sample.svg) |
| [`aws-source-code-48-dark`](aws-source-code-48-dark/README.md) | Source Code 48 Dark | logical | [SVG](aws-source-code-48-dark/sample.svg) |
| [`aws-source-code-48-light`](aws-source-code-48-light/README.md) | Source Code 48 Light | logical | [SVG](aws-source-code-48-light/sample.svg) |
| [`aws-sqs`](aws-sqs/README.md) | Amazon Simple Queue Service | region | [SVG](aws-sqs/sample.svg) |
| [`aws-sqs-message`](aws-sqs-message/README.md) | Amazon Simple Queue Service Message | logical | [SVG](aws-sqs-message/sample.svg) |
| [`aws-sqs-queue`](aws-sqs-queue/README.md) | Amazon Simple Queue Service Queue | region | [SVG](aws-sqs-queue/sample.svg) |
| [`aws-ssl-padlock-48-dark`](aws-ssl-padlock-48-dark/README.md) | SSL padlock 48 Dark | logical | [SVG](aws-ssl-padlock-48-dark/sample.svg) |
| [`aws-ssl-padlock-48-light`](aws-ssl-padlock-48-light/README.md) | SSL padlock 48 Light | logical | [SVG](aws-ssl-padlock-48-light/sample.svg) |
| [`aws-step-functions`](aws-step-functions/README.md) | AWS Step Functions | service | [SVG](aws-step-functions/sample.svg) |
| [`aws-step-functions-workflow`](aws-step-functions-workflow/README.md) | AWS Step Functions workflow | region | [SVG](aws-step-functions-workflow/sample.svg) |
| [`aws-storage-gateway`](aws-storage-gateway/README.md) | AWS Storage Gateway | service | [SVG](aws-storage-gateway/sample.svg) |
| [`aws-storage-gateway-amazon-fsx-file-gateway`](aws-storage-gateway-amazon-fsx-file-gateway/README.md) | AWS Storage Gateway Amazon FSx File Gateway | logical | [SVG](aws-storage-gateway-amazon-fsx-file-gateway/sample.svg) |
| [`aws-storage-gateway-amazon-s3-file-gateway`](aws-storage-gateway-amazon-s3-file-gateway/README.md) | AWS Storage Gateway Amazon S3 File Gateway | logical | [SVG](aws-storage-gateway-amazon-s3-file-gateway/sample.svg) |
| [`aws-storage-gateway-cached-volume`](aws-storage-gateway-cached-volume/README.md) | AWS Storage Gateway Cached Volume | logical | [SVG](aws-storage-gateway-cached-volume/sample.svg) |
| [`aws-storage-gateway-file-gateway`](aws-storage-gateway-file-gateway/README.md) | AWS Storage Gateway File Gateway | logical | [SVG](aws-storage-gateway-file-gateway/sample.svg) |
| [`aws-storage-gateway-noncached-volume`](aws-storage-gateway-noncached-volume/README.md) | AWS Storage Gateway Noncached Volume | logical | [SVG](aws-storage-gateway-noncached-volume/sample.svg) |
| [`aws-storage-gateway-tape-gateway`](aws-storage-gateway-tape-gateway/README.md) | AWS Storage Gateway Tape Gateway | logical | [SVG](aws-storage-gateway-tape-gateway/sample.svg) |
| [`aws-storage-gateway-virtual-tape-library`](aws-storage-gateway-virtual-tape-library/README.md) | AWS Storage Gateway Virtual Tape Library | logical | [SVG](aws-storage-gateway-virtual-tape-library/sample.svg) |
| [`aws-storage-gateway-volume-gateway`](aws-storage-gateway-volume-gateway/README.md) | AWS Storage Gateway Volume Gateway | logical | [SVG](aws-storage-gateway-volume-gateway/sample.svg) |
| [`aws-supply-chain`](aws-supply-chain/README.md) | AWS Supply Chain | service | [SVG](aws-supply-chain/sample.svg) |
| [`aws-support`](aws-support/README.md) | AWS Support | service | [SVG](aws-support/sample.svg) |
| [`aws-systems-manager`](aws-systems-manager/README.md) | AWS Systems Manager | service | [SVG](aws-systems-manager/sample.svg) |
| [`aws-systems-manager-application-manager`](aws-systems-manager-application-manager/README.md) | AWS Systems Manager Application Manager | logical | [SVG](aws-systems-manager-application-manager/sample.svg) |
| [`aws-systems-manager-automation`](aws-systems-manager-automation/README.md) | AWS Systems Manager Automation | logical | [SVG](aws-systems-manager-automation/sample.svg) |
| [`aws-systems-manager-change-calendar`](aws-systems-manager-change-calendar/README.md) | AWS Systems Manager Change Calendar | logical | [SVG](aws-systems-manager-change-calendar/sample.svg) |
| [`aws-systems-manager-change-manager`](aws-systems-manager-change-manager/README.md) | AWS Systems Manager Change Manager | logical | [SVG](aws-systems-manager-change-manager/sample.svg) |
| [`aws-systems-manager-compliance`](aws-systems-manager-compliance/README.md) | AWS Systems Manager Compliance | logical | [SVG](aws-systems-manager-compliance/sample.svg) |
| [`aws-systems-manager-distributor`](aws-systems-manager-distributor/README.md) | AWS Systems Manager Distributor | logical | [SVG](aws-systems-manager-distributor/sample.svg) |
| [`aws-systems-manager-documents`](aws-systems-manager-documents/README.md) | AWS Systems Manager Documents | logical | [SVG](aws-systems-manager-documents/sample.svg) |
| [`aws-systems-manager-incident-manager`](aws-systems-manager-incident-manager/README.md) | AWS Systems Manager Incident Manager | logical | [SVG](aws-systems-manager-incident-manager/sample.svg) |
| [`aws-systems-manager-inventory`](aws-systems-manager-inventory/README.md) | AWS Systems Manager Inventory | logical | [SVG](aws-systems-manager-inventory/sample.svg) |
| [`aws-systems-manager-maintenance-windows`](aws-systems-manager-maintenance-windows/README.md) | AWS Systems Manager Maintenance Windows | logical | [SVG](aws-systems-manager-maintenance-windows/sample.svg) |
| [`aws-systems-manager-opscenter`](aws-systems-manager-opscenter/README.md) | AWS Systems Manager OpsCenter | logical | [SVG](aws-systems-manager-opscenter/sample.svg) |
| [`aws-systems-manager-parameter-store`](aws-systems-manager-parameter-store/README.md) | AWS Systems Manager Parameter Store | logical | [SVG](aws-systems-manager-parameter-store/sample.svg) |
| [`aws-systems-manager-patch-manager`](aws-systems-manager-patch-manager/README.md) | AWS Systems Manager Patch Manager | logical | [SVG](aws-systems-manager-patch-manager/sample.svg) |
| [`aws-systems-manager-run-command`](aws-systems-manager-run-command/README.md) | AWS Systems Manager Run Command | logical | [SVG](aws-systems-manager-run-command/sample.svg) |
| [`aws-systems-manager-session-manager`](aws-systems-manager-session-manager/README.md) | AWS Systems Manager Session Manager | logical | [SVG](aws-systems-manager-session-manager/sample.svg) |
| [`aws-systems-manager-state-manager`](aws-systems-manager-state-manager/README.md) | AWS Systems Manager State Manager | logical | [SVG](aws-systems-manager-state-manager/sample.svg) |
| [`aws-tape-storage-48-dark`](aws-tape-storage-48-dark/README.md) | Tape storage 48 Dark | logical | [SVG](aws-tape-storage-48-dark/sample.svg) |
| [`aws-tape-storage-48-light`](aws-tape-storage-48-light/README.md) | Tape storage 48 Light | logical | [SVG](aws-tape-storage-48-light/sample.svg) |
| [`aws-telco-network-builder`](aws-telco-network-builder/README.md) | AWS Telco Network Builder | service | [SVG](aws-telco-network-builder/sample.svg) |
| [`aws-tensorflow-on-aws`](aws-tensorflow-on-aws/README.md) | TensorFlow on AWS | service | [SVG](aws-tensorflow-on-aws/sample.svg) |
| [`aws-textract`](aws-textract/README.md) | Amazon Textract | service | [SVG](aws-textract/sample.svg) |
| [`aws-textract-analyze-lending`](aws-textract-analyze-lending/README.md) | Amazon Textract Analyze Lending | logical | [SVG](aws-textract-analyze-lending/sample.svg) |
| [`aws-thinkbox-deadline`](aws-thinkbox-deadline/README.md) | AWS Thinkbox Deadline | service | [SVG](aws-thinkbox-deadline/sample.svg) |
| [`aws-thinkbox-frost`](aws-thinkbox-frost/README.md) | AWS Thinkbox Frost | service | [SVG](aws-thinkbox-frost/sample.svg) |
| [`aws-thinkbox-krakatoa`](aws-thinkbox-krakatoa/README.md) | AWS Thinkbox Krakatoa | service | [SVG](aws-thinkbox-krakatoa/sample.svg) |
| [`aws-thinkbox-sequoia`](aws-thinkbox-sequoia/README.md) | AWS Thinkbox Sequoia | service | [SVG](aws-thinkbox-sequoia/sample.svg) |
| [`aws-thinkbox-stoke`](aws-thinkbox-stoke/README.md) | AWS Thinkbox Stoke | service | [SVG](aws-thinkbox-stoke/sample.svg) |
| [`aws-thinkbox-xmesh`](aws-thinkbox-xmesh/README.md) | AWS Thinkbox XMesh | service | [SVG](aws-thinkbox-xmesh/sample.svg) |
| [`aws-timestream`](aws-timestream/README.md) | Amazon Timestream | service | [SVG](aws-timestream/sample.svg) |
| [`aws-toolkit-48-dark`](aws-toolkit-48-dark/README.md) | Toolkit 48 Dark | logical | [SVG](aws-toolkit-48-dark/sample.svg) |
| [`aws-toolkit-48-light`](aws-toolkit-48-light/README.md) | Toolkit 48 Light | logical | [SVG](aws-toolkit-48-light/sample.svg) |
| [`aws-tools-and-sdks`](aws-tools-and-sdks/README.md) | AWS Tools and SDKs | service | [SVG](aws-tools-and-sdks/sample.svg) |
| [`aws-training-certification`](aws-training-certification/README.md) | AWS Training Certification | service | [SVG](aws-training-certification/sample.svg) |
| [`aws-transcribe`](aws-transcribe/README.md) | Amazon Transcribe | service | [SVG](aws-transcribe/sample.svg) |
| [`aws-transfer-family`](aws-transfer-family/README.md) | AWS Transfer Family | service | [SVG](aws-transfer-family/sample.svg) |
| [`aws-transfer-family-aws-as2`](aws-transfer-family-aws-as2/README.md) | AWS Transfer Family AWS AS2 | logical | [SVG](aws-transfer-family-aws-as2/sample.svg) |
| [`aws-transfer-family-aws-ftp`](aws-transfer-family-aws-ftp/README.md) | AWS Transfer Family AWS FTP | logical | [SVG](aws-transfer-family-aws-ftp/sample.svg) |
| [`aws-transfer-family-aws-ftps`](aws-transfer-family-aws-ftps/README.md) | AWS Transfer Family AWS FTPS | logical | [SVG](aws-transfer-family-aws-ftps/sample.svg) |
| [`aws-transfer-family-aws-sftp`](aws-transfer-family-aws-sftp/README.md) | AWS Transfer Family AWS SFTP | logical | [SVG](aws-transfer-family-aws-sftp/sample.svg) |
| [`aws-transform`](aws-transform/README.md) | AWS Transform | service | [SVG](aws-transform/sample.svg) |
| [`aws-transit-gateway`](aws-transit-gateway/README.md) | AWS Transit Gateway | service | [SVG](aws-transit-gateway/sample.svg) |
| [`aws-transit-gateway-attachment`](aws-transit-gateway-attachment/README.md) | AWS Transit Gateway Attachment | logical | [SVG](aws-transit-gateway-attachment/sample.svg) |
| [`aws-translate`](aws-translate/README.md) | Amazon Translate | service | [SVG](aws-translate/sample.svg) |
| [`aws-trusted-advisor`](aws-trusted-advisor/README.md) | AWS Trusted Advisor | service | [SVG](aws-trusted-advisor/sample.svg) |
| [`aws-trusted-advisor-checklist`](aws-trusted-advisor-checklist/README.md) | AWS Trusted Advisor Checklist | logical | [SVG](aws-trusted-advisor-checklist/sample.svg) |
| [`aws-trusted-advisor-checklist-cost`](aws-trusted-advisor-checklist-cost/README.md) | AWS Trusted Advisor Checklist Cost | logical | [SVG](aws-trusted-advisor-checklist-cost/sample.svg) |
| [`aws-trusted-advisor-checklist-fault-tolerant`](aws-trusted-advisor-checklist-fault-tolerant/README.md) | AWS Trusted Advisor Checklist Fault Tolerant | logical | [SVG](aws-trusted-advisor-checklist-fault-tolerant/sample.svg) |
| [`aws-trusted-advisor-checklist-performance`](aws-trusted-advisor-checklist-performance/README.md) | AWS Trusted Advisor Checklist Performance | logical | [SVG](aws-trusted-advisor-checklist-performance/sample.svg) |
| [`aws-trusted-advisor-checklist-security`](aws-trusted-advisor-checklist-security/README.md) | AWS Trusted Advisor Checklist Security | logical | [SVG](aws-trusted-advisor-checklist-security/sample.svg) |
| [`aws-user-48-dark`](aws-user-48-dark/README.md) | User 48 Dark | logical | [SVG](aws-user-48-dark/sample.svg) |
| [`aws-user-48-light`](aws-user-48-light/README.md) | User 48 Light | logical | [SVG](aws-user-48-light/sample.svg) |
| [`aws-user-notifications`](aws-user-notifications/README.md) | AWS User Notifications | service | [SVG](aws-user-notifications/sample.svg) |
| [`aws-users-48-dark`](aws-users-48-dark/README.md) | Users 48 Dark | logical | [SVG](aws-users-48-dark/sample.svg) |
| [`aws-users-48-light`](aws-users-48-light/README.md) | Users 48 Light | logical | [SVG](aws-users-48-light/sample.svg) |
| [`aws-verified-access`](aws-verified-access/README.md) | AWS Verified Access | service | [SVG](aws-verified-access/sample.svg) |
| [`aws-verified-permissions`](aws-verified-permissions/README.md) | Amazon Verified Permissions | service | [SVG](aws-verified-permissions/sample.svg) |
| [`aws-virtual-private-cloud`](aws-virtual-private-cloud/README.md) | Amazon Virtual Private Cloud | service | [SVG](aws-virtual-private-cloud/sample.svg) |
| [`aws-vpc-carrier-gateway`](aws-vpc-carrier-gateway/README.md) | Amazon VPC Carrier Gateway | logical | [SVG](aws-vpc-carrier-gateway/sample.svg) |
| [`aws-vpc-customer-gateway`](aws-vpc-customer-gateway/README.md) | Amazon VPC Customer Gateway | external | [SVG](aws-vpc-customer-gateway/sample.svg) |
| [`aws-vpc-elastic-network-adapter`](aws-vpc-elastic-network-adapter/README.md) | Amazon VPC Elastic Network Adapter | subnet | [SVG](aws-vpc-elastic-network-adapter/sample.svg) |
| [`aws-vpc-elastic-network-interface`](aws-vpc-elastic-network-interface/README.md) | Amazon VPC Elastic Network Interface | subnet | [SVG](aws-vpc-elastic-network-interface/sample.svg) |
| [`aws-vpc-flow-logs`](aws-vpc-flow-logs/README.md) | Amazon VPC Flow Logs | logical | [SVG](aws-vpc-flow-logs/sample.svg) |
| [`aws-vpc-internet-gateway`](aws-vpc-internet-gateway/README.md) | Amazon VPC Internet Gateway | vpc-boundary | [SVG](aws-vpc-internet-gateway/sample.svg) |
| [`aws-vpc-lattice`](aws-vpc-lattice/README.md) | Amazon VPC Lattice | service | [SVG](aws-vpc-lattice/sample.svg) |
| [`aws-vpc-nat-gateway`](aws-vpc-nat-gateway/README.md) | Amazon VPC NAT Gateway | subnet | [SVG](aws-vpc-nat-gateway/sample.svg) |
| [`aws-vpc-network-access-analyzer`](aws-vpc-network-access-analyzer/README.md) | Amazon VPC Network Access Analyzer | logical | [SVG](aws-vpc-network-access-analyzer/sample.svg) |
| [`aws-vpc-network-access-control-list`](aws-vpc-network-access-control-list/README.md) | Amazon VPC Network Access Control List | logical | [SVG](aws-vpc-network-access-control-list/sample.svg) |
| [`aws-vpc-peering-connection`](aws-vpc-peering-connection/README.md) | Amazon VPC Peering Connection | logical | [SVG](aws-vpc-peering-connection/sample.svg) |
| [`aws-vpc-reachability-analyzer`](aws-vpc-reachability-analyzer/README.md) | Amazon VPC Reachability Analyzer | logical | [SVG](aws-vpc-reachability-analyzer/sample.svg) |
| [`aws-vpc-router`](aws-vpc-router/README.md) | Amazon VPC Router | logical | [SVG](aws-vpc-router/sample.svg) |
| [`aws-vpc-traffic-mirroring`](aws-vpc-traffic-mirroring/README.md) | Amazon VPC Traffic Mirroring | logical | [SVG](aws-vpc-traffic-mirroring/sample.svg) |
| [`aws-vpc-virtual-private-cloud-vpc`](aws-vpc-virtual-private-cloud-vpc/README.md) | Amazon VPC Virtual private cloud VPC | logical | [SVG](aws-vpc-virtual-private-cloud-vpc/sample.svg) |
| [`aws-vpc-vpn-connection`](aws-vpc-vpn-connection/README.md) | Amazon VPC VPN Connection | logical | [SVG](aws-vpc-vpn-connection/sample.svg) |
| [`aws-vpc-vpn-gateway`](aws-vpc-vpn-gateway/README.md) | Amazon VPC VPN Gateway | vpc-boundary | [SVG](aws-vpc-vpn-gateway/sample.svg) |
| [`aws-waf`](aws-waf/README.md) | AWS WAF | service | [SVG](aws-waf/sample.svg) |
| [`aws-waf-bad-bot`](aws-waf-bad-bot/README.md) | AWS WAF Bad Bot | logical | [SVG](aws-waf-bad-bot/sample.svg) |
| [`aws-waf-bot`](aws-waf-bot/README.md) | AWS WAF Bot | logical | [SVG](aws-waf-bot/sample.svg) |
| [`aws-waf-bot-control`](aws-waf-bot-control/README.md) | AWS WAF Bot Control | logical | [SVG](aws-waf-bot-control/sample.svg) |
| [`aws-waf-filtering-rule`](aws-waf-filtering-rule/README.md) | AWS WAF Filtering Rule | logical | [SVG](aws-waf-filtering-rule/sample.svg) |
| [`aws-waf-labels`](aws-waf-labels/README.md) | AWS WAF Labels | logical | [SVG](aws-waf-labels/sample.svg) |
| [`aws-waf-managed-rule`](aws-waf-managed-rule/README.md) | AWS WAF Managed Rule | logical | [SVG](aws-waf-managed-rule/sample.svg) |
| [`aws-waf-rule`](aws-waf-rule/README.md) | AWS WAF Rule | logical | [SVG](aws-waf-rule/sample.svg) |
| [`aws-wavelength`](aws-wavelength/README.md) | AWS Wavelength | service | [SVG](aws-wavelength/sample.svg) |
| [`aws-well-architected-tool`](aws-well-architected-tool/README.md) | AWS Well Architected Tool | service | [SVG](aws-well-architected-tool/sample.svg) |
| [`aws-wickr`](aws-wickr/README.md) | AWS Wickr | service | [SVG](aws-wickr/sample.svg) |
| [`aws-workdocs`](aws-workdocs/README.md) | Amazon WorkDocs | service | [SVG](aws-workdocs/sample.svg) |
| [`aws-workdocs-sdk`](aws-workdocs-sdk/README.md) | Amazon WorkDocs SDK | service | [SVG](aws-workdocs-sdk/sample.svg) |
| [`aws-workmail`](aws-workmail/README.md) | Amazon WorkMail | service | [SVG](aws-workmail/sample.svg) |
| [`aws-workspaces-family`](aws-workspaces-family/README.md) | Amazon WorkSpaces Family | service | [SVG](aws-workspaces-family/sample.svg) |
| [`aws-workspaces-family-amazon-workspaces`](aws-workspaces-family-amazon-workspaces/README.md) | Amazon WorkSpaces Family Amazon WorkSpaces | logical | [SVG](aws-workspaces-family-amazon-workspaces/sample.svg) |
| [`aws-workspaces-family-amazon-workspaces-core`](aws-workspaces-family-amazon-workspaces-core/README.md) | Amazon WorkSpaces Family Amazon WorkSpaces Core | logical | [SVG](aws-workspaces-family-amazon-workspaces-core/sample.svg) |
| [`aws-workspaces-family-amazon-workspaces-secure-browser`](aws-workspaces-family-amazon-workspaces-secure-browser/README.md) | Amazon WorkSpaces Family Amazon WorkSpaces Secure Browser | logical | [SVG](aws-workspaces-family-amazon-workspaces-secure-browser/sample.svg) |
| [`aws-x-ray`](aws-x-ray/README.md) | AWS X Ray | service | [SVG](aws-x-ray/sample.svg) |
| [`corporate-data-center`](corporate-data-center/README.md) | Corporate data center | external | [SVG](corporate-data-center/sample.svg) |
| [`ec2-instance-contents`](ec2-instance-contents/README.md) | EC2 instance contents | host | [SVG](ec2-instance-contents/sample.svg) |
| [`elastic-beanstalk-container`](elastic-beanstalk-container/README.md) | Elastic Beanstalk container | region | [SVG](elastic-beanstalk-container/sample.svg) |
| [`generic-group`](generic-group/README.md) | Generic group | logical | [SVG](generic-group/sample.svg) |
| [`private-subnet`](private-subnet/README.md) | Private subnet | subnet | [SVG](private-subnet/sample.svg) |
| [`public-subnet`](public-subnet/README.md) | Public subnet | subnet | [SVG](public-subnet/sample.svg) |
| [`region`](region/README.md) | Region | region | [SVG](region/sample.svg) |
| [`security-group`](security-group/README.md) | Security group | vpc | [SVG](security-group/sample.svg) |
| [`server-contents`](server-contents/README.md) | Server contents | host | [SVG](server-contents/sample.svg) |
| [`spot-fleet`](spot-fleet/README.md) | Spot Fleet | region | [SVG](spot-fleet/sample.svg) |
| [`vpc`](vpc/README.md) | VPC | vpc | [SVG](vpc/sample.svg) |
| [`vpc-endpoint`](vpc-endpoint/README.md) | Amazon VPC Endpoints | vpc-boundary | [SVG](vpc-endpoint/sample.svg) |
