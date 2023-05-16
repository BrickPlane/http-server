local svc = (import '1.21/main.libsonnet').core.v1.service;
local env = import 'env.libsonnet';
local LABEL = env.LABEL;
local VERSION = env.VERSION;
local NAMES = env.NAMES;

{
  new(
    service,
  ): (
    svc.new(
      name=service.name,
      selector={
        app: NAMES['namespace-name'],
        version: VERSION,
        svc: service.name,
      },
      ports=[{
        name: if service.protocol == 'http2' then 'grpc-web' else 'http-web',
        port: service.externalPort,
        targetPort: service.internalPort,
        protocol: 'TCP',
        appProtocol: service.protocol,
      }]
    )
    + svc.metadata.withNamespace(
      NAMES['namespace-name']
    )
    + svc.metadata.withLabels({
      app: NAMES['namespace-name'],
      version: VERSION,
      svc: service.name,
    })
  ),
}
