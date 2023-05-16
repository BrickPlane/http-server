local env = import 'env.libsonnet';
local NAMES = env.NAMES;

{
  new(
    service,
  ): (
    {
      apiVersion: 'networking.istio.io/v1alpha3',
      kind: 'DestinationRule',
      metadata: {
        name: NAMES['namespace-name'] + '-' + service.name + '-dr',
        namespace: NAMES['namespace-name'],
      },
      spec: {
        host: service.name,
        trafficPolicy: {
          tls: {
            mode: 'DISABLE',
          },
        },
      },
    }
  ),
}
