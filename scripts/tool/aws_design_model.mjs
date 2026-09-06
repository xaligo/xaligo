import {namespaces, primaryTypes, exactTypes, apiAliases, productGuides, retiredProducts, lifecycleSource} from './aws_research_mapping.mjs';

export const normalize = value => String(value ?? '').toLowerCase().replace(/amazon|aws|[^a-z0-9]/g, '');
const shortType = type => type.split('::').at(-1);
const plural = value => normalize(value).replace(/s$/, '');
const metadataFields = /^(Tags|Name|Description|.*Name|.*Description)$/;
const priority = ['Type','VPCEndpointType','VpcId','CidrBlock','AvailabilityZone','SubnetId','SubnetIds','SecurityGroupIds','ServiceName','PrivateDnsEnabled','RouteTableIds','Protocol','Port','IpAddressType','InstanceType','ImageId','Engine','DBInstanceClass','MultiAZ','DBClusterIdentifier','BucketName','VersioningConfiguration','BucketEncryption','PublicAccessBlockConfiguration','LifecycleConfiguration','ReplicationConfiguration','DeliveryStreamType','KinesisStreamSourceConfiguration','ExtendedS3DestinationConfiguration','HttpEndpointDestinationConfiguration','TableName','KeySchema','AttributeDefinitions','BillingMode','ProvisionedThroughput','Runtime','Handler','Code','Role','MemorySize','Timeout','Architectures','QueueName','FifoQueue','RedrivePolicy','TopicName','StateMachineType','Definition','ClusterName','TaskDefinition','LaunchType','DesiredCount','ServiceToken'];
const fieldRank = name => priority.includes(name) ? priority.indexOf(name) : 1000;
export const orderedFields = fields => [...fields].sort((a, b) => fieldRank(a.name) - fieldRank(b.name) || Number(metadataFields.test(a.name)) - Number(metadataFields.test(b.name)) || Number(b.required) - Number(a.required) || a.name.localeCompare(b.name, 'en'));

export function diagramOperationRank(name, primary) {
  const focus = plural(shortType(primary ?? ''));
  const operation = plural(name.replace(/^(Create|Delete|Describe|Modify|Update|Get|Put|List|Run|Start|Stop|Terminate|Invoke|Send)/, ''));
  if (focus && operation === focus) return 0;
  if (focus && normalize(name).includes(focus)) return 1;
  return 2;
}

export function buildDesigns(catalog, cfn, api) {
  const services = catalog.filter(d => d.kind === 'service').sort((a, b) => b.tag.length - a.tag.length);
  const resourceNames = Object.keys(cfn.models).filter(name => /^AWS::[^:]+::[^.]+$/.test(name));
  const namespaceNames = [...new Set(resourceNames.map(name => name.split('::')[1]))];
  const byCategory = category => services.filter(d => d.category === category).sort((a, b) => a.tag.localeCompare(b.tag, 'en'));
  return catalog.map(d => {
    const ownerAlias = Object.entries({'aws-elastic-file-system-':'aws-efs','aws-identity-access-management-':'aws-identity-and-access-management','aws-iot-greengrass-':'aws-iot-greengrass-service','aws-iot-':'aws-iot-core','aws-msk-':'aws-managed-streaming-for-apache-kafka'}).find(([prefix]) => d.kind === 'resource' && d.tag.startsWith(prefix));
    const owner = d.kind === 'service' ? d : services.find(s => d.tag.startsWith(s.tag + '-')) ?? (ownerAlias && catalog.find(s => s.tag === ownerAlias[1]))
      ?? (d.tag.startsWith('aws-vpc-') || d.tag === 'vpc-endpoint' || ['vpc','private-subnet','public-subnet','security-group'].includes(d.tag) ? catalog.find(s => s.tag === 'aws-virtual-private-cloud') : undefined);
    const key = (owner?.tag ?? d.tag).replace(/^aws-/, '');
    const ns = namespaces[key] ?? namespaceNames.filter(name => normalize(name) === normalize(key));
    const types = resourceNames.filter(name => ns.includes(name.split('::')[1]));
    const apis = [...(apiAliases[key] ?? []), ...Object.entries(api.services).filter(([name, s]) => [name, s.id, s.name].some(value => [key, ...ns].some(n => normalize(value) === normalize(n)))).map(([name]) => name)];
    const validAPIs = [...new Set(apis.filter(name => api.services[name]))];
    const suffix = owner && d.tag !== owner.tag ? d.tag.slice(owner.tag.length + 1) : '';
    const exact = (exactTypes[d.tag] ?? []).map(name => 'AWS::' + name).filter(name => cfn.models[name]);
    const matching = suffix ? types.filter(name => plural(shortType(name)) === plural(suffix)) : [];
    let primary = exact[0] ?? matching[0] ?? types.find(name => shortType(name) === primaryTypes[name.split('::')[1]]) ?? types[0];
    // A variant of an instance icon is an instance, not a DB cluster or table.
    if (!exact.length && /^(aws-rds-|aws-aurora-).*instance/.test(d.tag)) primary = 'AWS::RDS::DBInstance';
    const researchKind = exact.length || matching.length ? 'resource-schema' : primary ? 'service-context' : validAPIs.length ? 'api-context' : productGuides[key] ? 'product-guide' : d.kind === 'category' ? 'category' : d.group ? 'group' : d.category === 'General Icons' ? 'symbol' : 'unmapped';
    const guide = productGuides[key];
    const fields = primary ? orderedFields(cfn.models[primary].fields) : [];
    const nested = primary ? fields.flatMap(field => {
      const type = field.type.replace(/^(List|Map)</, '').replace(/>$/, '');
      const target = primary.split('.')[0] + '.' + type;
      return cfn.models[target] ? [{field: field.name, model: target}] : [];
    }) : [];
    const apiOperations = validAPIs.flatMap(service => api.services[service].operations.map(operation => ({service, ...operation}))).sort((a, b) => {
      const rank = name => /^(Create|Put|Start|Run|Invoke|Detect|Synthesize|Send|Batch)/.test(name) ? 0 : /^(Update|Get|Describe|List)/.test(name) ? 1 : 2;
      return diagramOperationRank(a.name, primary) - diagramOperationRank(b.name, primary) || rank(a.name) - rank(b.name) || a.name.localeCompare(b.name, 'en');
    });
    const retired = retiredProducts.find(slug => key === slug || key.startsWith(slug + '-'));
    return {
      tag: d.tag, focus: d.name, kind: researchKind, owner: owner?.tag,
      primary, resources: [...new Set([...exact, ...types])], nested,
      related: types.filter(type => type !== primary && fields.some(field => plural(shortType(type)) === plural(field.name.replace(/(Arns?|Ids?)$/, '')))),
      apis: validAPIs, operations: apiOperations.map(o => ({service:o.service, name:o.name})),
      fields, guide, peers: byCategory(d.category).filter(s => s.tag !== d.tag).map(s => s.tag),
      sources: [...new Set([primary && cfn.models[primary].documentation, ...validAPIs.map(name => api.services[name].source), guide?.source, retired && lifecycleSource, (!primary && !validAPIs.length && !guide) && 'https://aws.amazon.com/architecture/icons/'].filter(Boolean))],
      lifecycle: retired ? 'retired' : key === 'codewhisperer' ? 'renamed-q-developer' : 'check-service-guide',
    };
  });
}
