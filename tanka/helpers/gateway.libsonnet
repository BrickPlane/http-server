local env = import 'env.libsonnet';
local NAMES = env.NAMES;

{
  new(
    hosts,
  ): (
    {
      apiVersion: 'networking.istio.io/v1alpha3',
      kind: 'Gateway',
      metadata: {
        name: NAMES['namespace-name'] + '-gateway',
        namespace: NAMES['namespace-name'],
      },
      spec: {
        selector: {
          istio: 'ingressgateway',
        },
        servers: [
          {
            port: {
              number: 443,
              name: 'https',
              protocol: 'HTTPS',
            },
            tls: {
              mode: 'SIMPLE',
              credentialName: NAMES['certificate-secret'],
            },
            hosts: hosts,
          },
          {
            port: {
              number: 80,
              name: 'http',
              protocol: 'HTTP',
            },
            hosts: hosts,
          },
        ],
      },
    }
  ),
}
