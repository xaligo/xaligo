// Public listener/target-group terminology, with ALB and NLB kept distinct.
export function loadBalancerParameters(tag) {
  const alb = tag === 'aws-elastic-load-balancing-application-load-balancer';
  const nlb = tag === 'aws-elastic-load-balancing-network-load-balancer';
  if (!alb && !nlb) return undefined;
  const text = (name, description, example, type = 'text') => ({name, description, example, type, values: []});
  const enumeration = (name, description, values) => ({name, description, example: values[0], type: 'enum', values});
  return [
    enumeration('scheme', 'Load balancer reachability annotation', ['internet-facing','internal']),
    enumeration('ip-address-type', 'Load balancer address family, not target IP', alb ? ['dualstack','ipv4','dualstack-without-public-ipv4'] : ['dualstack','ipv4']),
    text('subnets', 'Subnet/AZ or subnet-mapping summary', 'subnet-a, subnet-b'),
    text('security-groups', 'Security group references', 'sg-example'),
    enumeration('listener-protocol', 'One illustrated listener; additional listeners use separate components', alb ? ['HTTPS','HTTP'] : ['TLS','TCP','UDP','TCP_UDP','QUIC','TCP_QUIC']),
    text('listener-port', 'Listener port (1..65535)', '443', 'port'),
    text('certificate', 'Server certificate reference; not a client CA bundle', 'server-certificate'),
    text('tls-policy', 'Named security policy annotation; availability is not validated', 'selected-policy'),
    ...(alb ? [enumeration('mutual-tls-mode', 'ALB client certificate mode; passthrough forwards HTTP headers, not raw TLS', ['verify','off','passthrough']), text('trust-store', 'Client CA trust-store reference for ALB verify mode', 'client-ca-store'), text('listener-rules', 'Priority / condition / action summary', '10: host api.example.test -> forward')] : [text('alpn-policy', 'NLB TLS listener ALPN policy annotation', 'HTTP2Preferred'), text('client-ip-preservation', 'Target-group client IP preservation annotation', 'false', 'boolean')]),
    text('target-group', 'Forward action target-group reference', 'application-targets'),
    enumeration('target-type', 'Target registration type', alb ? ['ip','instance','lambda'] : ['ip','instance','alb']),
    enumeration('target-protocol', 'Backend protocol annotation, independent of the client connection', alb ? ['HTTP','HTTPS'] : ['TLS','TCP','UDP','TCP_UDP','QUIC','TCP_QUIC']),
    text('target-port', 'Backend target port (1..65535); omitted for Lambda', alb ? '8080' : '8443', 'port'),
    text('targets', 'Registered instance IDs or IP:port references, not load-balancer IPs', alb ? '10.0.2.10:8080' : '10.0.2.20:8443'),
    text('health-check', 'Protocol / port / path summary', alb ? 'HTTP:8080 /health' : 'TCP:8443'),
  ];
}
