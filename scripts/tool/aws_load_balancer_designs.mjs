// Deliberate functional examples, not a projection of the shared ELBv2 schema:
// the API schema alone cannot tell ALB and NLB listener capabilities apart.
export const loadBalancerDesigns = {
  'aws-elastic-load-balancing-application-load-balancer': {
    name: 'ALB — HTTPS / mTLS verify',
    summary: ['Layer 7 / HTTP routing', 'Scheme: internet-facing', 'IP address type: dualstack', 'Subnets: two Availability Zones', 'Security groups: client HTTPS + target access'],
    cards: [
      ['HTTPS listener', 'Protocol = HTTPS', 'Port = 443', 'Certificates = server certificate ARN', 'SslPolicy = selected TLS policy', 'MutualAuthentication.Mode = verify'],
      ['Rules / actions', 'Priority = 10; host = api.example.test', 'Path = /api/*; action = forward', 'Target-group weight = 100', 'Default action = fixed-response 404', 'Optional: redirect / authentication'],
      ['Target group', 'TargetType = ip', 'Protocol = HTTP; Port = 8080', 'ProtocolVersion = HTTP1', 'Health check = HTTP /health', 'Success codes = 200; target = 10.0.2.10'],
      ['Trust store / client authentication', 'CA bundle = S3 bucket + object key', 'Revocation = optional CRL object', 'TrustStoreArn -> HTTPS listener', 'Verify mode validates client chain', 'One trust store per secure listener'],
      ['Network / target registration', 'Listener address != target IP', 'Targets can be instance / ip / lambda', 'Security-group rules cover app + health', 'Cross-zone / stickiness / draining', 'TLS to targets is a separate decision'],
      ['Modes are not interchangeable', 'verify: ALB checks client certificate', 'passthrough: certificate in HTTP headers', 'off: no client-certificate authentication', 'ALB passthrough is NOT raw TLS relay', 'TCP passthrough belongs in the NLB view'],
    ],
    sources: [
      'https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html',
      'https://docs.aws.amazon.com/elasticloadbalancing/latest/application/listener-rules.html',
      'https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-target-groups.html',
      'https://docs.aws.amazon.com/elasticloadbalancing/latest/application/mutual-authentication.html',
      'https://docs.aws.amazon.com/elasticloadbalancing/latest/application/configuring-mtls-with-elb.html',
    ],
  },
  'aws-elastic-load-balancing-network-load-balancer': {
    name: 'NLB — TLS termination / TCP passthrough',
    summary: ['Layer 4 / connection forwarding', 'Scheme: internet-facing', 'IP address type: dualstack', 'Subnet mapping: one address per AZ', 'EIP allocation or private address, as applicable'],
    cards: [
      ['TLS listener — termination example', 'Protocol = TLS; Port = 443', 'Certificates = server certificate ARN', 'SslPolicy = selected TLS policy', 'AlpnPolicy = chosen application policy', 'Forward action -> TLS target group'],
      ['TCP listener — alternate example', 'Protocol = TCP; Port = 443', 'Encrypted bytes pass through NLB', 'No certificate / trust store at NLB', 'Server certificate lives at the target', 'Backend may implement mutual TLS'],
      ['Target group / health', 'TargetType = ip', 'Protocol = TLS; Port = 8443', 'Target = 10.0.2.20:8443', 'Health check = TCP', 'Target type choices: instance / ip / alb'],
      ['Addressing / network', 'SubnetMappings: SubnetId + AllocationId', 'PrivateIPv4Address / IPv6Address', 'Security groups: ingress + target access', 'Client IP preservation / proxy protocol v2', 'Cross-zone / deregistration delay'],
      ['Listener and rule capabilities', 'TCP, TLS, UDP, TCP_UDP', 'QUIC, TCP_QUIC (current guide)', 'Port range = 1..65535', 'Weighted target-group forwarding', 'Dualstack source-IP rules: IP version'],
      ['Authentication boundary', 'NLB TLS listeners do NOT terminate mTLS', 'No ALB-style TrustStoreArn on NLB', 'TCP passthrough keeps backend mTLS', 'TLS target-group policy is separate', 'Do not combine TLS and TCP TGs on TLS'],
    ],
    sources: [
      'https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-listeners.html',
      'https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html',
      'https://docs.aws.amazon.com/elasticloadbalancing/latest/network/tls-listener-certificates.html',
      'https://docs.aws.amazon.com/elasticloadbalancing/latest/network/network-load-balancers.html',
    ],
  },
};

// Each relationship is explicitly typed. Configuration arrows must never be
// presented as request traffic, and the two NLB alternatives never coexist on
// port 443 in one load balancer.
export function loadBalancerFlow(tag, {xml, rectangle}) {
  const alb = tag.endsWith('-application-load-balancer');
  const node = (id, title, lines, width = 260, height = 160) => rectangle(id, [title, ...lines], width, height);
  const edge = (src, dst, kind, label) => `<connection src="${src}" dst="${dst}" kind="${kind}" color="${kind === 'traffic' ? '#2563eb' : '#7c3aed'}" stroke-style="${kind === 'traffic' ? 'solid' : 'dashed'}" label="${xml(label)}" />`;
  const frames = [];
  for (const mode of alb ? ['verify'] : ['termination', 'passthrough']) {
    const prefix = `flow-${mode}`;
    const id = s => prefix + '-' + s;
    const listener = alb ? ['HTTPS :443', 'mTLS mode: verify'] : mode === 'termination' ? ['TLS :443', 'TLS terminates at NLB'] : ['TCP :443', 'Encrypted bytes pass through'];
    const target = alb ? ['HTTP :8080', '10.0.2.10', 'Health check: HTTP /health'] : mode === 'termination' ? ['TLS :8443', '10.0.2.20', 'Health check: TCP'] : ['TCP :443', '10.0.2.20', 'Backend TLS + mTLS verification'];
    const sources = alb ? ['Client CA bundle / CRL in S3', 'TrustStoreArn -> HTTPS listener', 'Configuration, not traffic'] : mode === 'termination' ? ['ACM / IAM certificate ARN', 'TLS policy + ALPN policy', 'Configuration, not traffic'] : ['Backend certificate + client CA', 'Configured on target application', 'No NLB trust-store attachment'];
    const refTarget = !alb && mode === 'passthrough' ? 'target' : 'listener';
    frames.push({mode, source: `  <frame id="${prefix}" title="${xml((alb ? 'ALB' : 'NLB') + ' — ' + mode + ' / 関係図')}" width="1500" height="880">
    <col width="1360" height="720" margin="32" gap="64">
      <row height="160" gap="40">
        <${tag} id="${id('lb')}" label="${alb ? 'ALB' : 'NLB'} / ${mode}" label-width="260" show-details="false" listener-protocol="${alb ? 'HTTPS' : mode === 'termination' ? 'TLS' : 'TCP'}" listener-port="443" target-type="ip" target-group="${id('tg')}"${alb ? ` mutual-tls-mode="verify" trust-store="${id('config')}"` : ''} />
        ${node(id('config'), 'Configuration reference', sources, 520)}
        ${node(id('net'), 'Network identity', ['Load-balancer address / DNS', 'Subnet mappings / AZs', 'Not a backend target address'], 420)}
      </row>
      <row height="180" gap="52">
        ${node(id('client'), 'Client', alb ? ['Server TLS validation', 'Client certificate + key'] : ['Connect to load-balancer DNS', 'TLS endpoint depends on mode'])}
        ${node(id('listener'), 'Listener', listener)}
        ${node(id('tg'), 'Target group', ['Forward action', ...target.slice(0, 1), 'TargetType = ip'])}
        ${node(id('target'), 'Registered target', target)}
      </row>
      ${node(id('legend'), 'Reading the diagram', ['Solid traffic edges: client -> listener -> target group -> target', 'Configuration/reference edges: certificate / trust settings; not request traffic', alb ? 'Host/path rules choose a target group. CA trust and the ALB server certificate are different objects.' : 'This frame is one alternative. The other frame uses a different load balancer / configuration.'], 1280, 150)}
    </col>
    <connections>
      ${edge(id('client'), id('listener'), 'traffic', '443')}
      ${edge(id('listener'), id('tg'), 'traffic', '')}
      ${edge(id('tg'), id('target'), 'traffic', alb ? '8080' : mode === 'termination' ? '8443' : '443')}
      ${edge(id('config'), id(refTarget), 'connection', 'ref')}
    </connections>
  </frame>`});
  }
  return frames;
}
