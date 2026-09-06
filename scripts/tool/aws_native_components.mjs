// Compact, editable V2 components. Their geometry belongs to usc/aws, not XAL.
const wrap = frame => `<xaligo version="2">\n<frames>\n${frame}\n</frames>\n</xaligo>\n`;
export function nativeLoadBalancerSample(tag) {
  const alb = tag.endsWith('-application-load-balancer');
  const listeners = alb ? [
    '<aws-listener id="http" protocol="HTTP" port="80" target-group="web" />',
    '<aws-listener id="https" protocol="HTTPS" port="443" certificate="server-cert" target-group="api" />',
    '<aws-listener id="admin" protocol="HTTPS" port="8443" mtls="verify" certificate="server-cert" trust-store="client-ca" target-group="admin" />',
  ] : [
    '<aws-listener id="tcp" protocol="TCP" port="80" target-group="web" />',
    '<aws-listener id="tls" protocol="TLS" port="443" certificate="server-cert" backend-tls="true" backend-mtls="false" target-group="secure" />',
    '<aws-listener id="passthrough" protocol="TCP" port="8443" backend-tls="true" backend-mtls="true" target-group="private" />',
  ];
  return wrap(`  <frame id="sample" title="${alb ? 'ALB' : 'NLB'} / listeners" width="640" height="320">
    <${tag} id="component" domain="${alb ? 'api' : 'nlb'}.example.test" margin="24">
      ${listeners.join('\n      ')}
    </${tag}>
  </frame>`);
}

export function nativeLoadBalancerFlows(tag) {
  const alb = tag.endsWith('-application-load-balancer');
  return (alb ? ['verify'] : ['termination', 'passthrough']).map(mode => {
    const passthrough = mode === 'passthrough';
    const protocol = alb ? 'HTTPS' : passthrough ? 'TCP' : 'TLS';
    const backendPort = alb ? 8080 : passthrough ? 443 : 8443;
    const settings = alb ? 'mtls="verify" certificate="server-cert" trust-store="client-ca"' : passthrough ? 'backend-tls="true" backend-mtls="true"' : 'certificate="server-cert" backend-tls="true" backend-mtls="false"';
    return {mode, source: wrap(`  <frame id="flow-${mode}" title="${alb ? 'ALB' : 'NLB'} / ${mode}" width="860" height="420" layout="absolute">
    <rectangle id="client" x="24" y="${passthrough ? 125 : 135}" width="140" height="90" title="Client" fill="#f8fbff" stroke="#c7d8f4" />
    <${tag} id="component" domain="${alb ? 'api' : 'nlb'}.example.test" x="190" y="40">
      <aws-listener id="listener" protocol="${protocol}" port="443" ${settings} target-group="app" />
    </${tag}>
    <rectangle id="target-group" x="580" y="${passthrough ? 125 : 135}" width="240" height="90" title="Target group / app&#10;${alb ? 'HTTP' : passthrough ? 'TCP' : 'TLS'} :${backendPort} · ip" fill="#fff7ed" stroke="#fb923c" />
    <rectangle id="target" x="580" y="280" width="240" height="96" title="10.0.2.20 :${backendPort}${passthrough ? '&#10;TLS + mTLS' : ''}" fill="#f0fdf4" stroke="#86efac" />
    <connections>
      <connection src="client" dst="listener" kind="traffic" src-side="right" dst-side="left" color="#2563eb" label="443" />
      <connection src="listener" dst="target-group" kind="traffic" src-side="right" dst-side="left" color="#2563eb" />
      <connection src="target-group" dst="target" kind="traffic" src-side="bottom" dst-side="top" color="#2563eb" label="${backendPort}" />
    </connections>
  </frame>`)};
  });
}
