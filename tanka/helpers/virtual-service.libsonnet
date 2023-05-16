local env = import 'env.libsonnet';
local NAMES = env.NAMES;

{
  new(
    service,
    host,
  ): (
    {
      apiVersion: 'networking.istio.io/v1beta1',
      kind: 'VirtualService',
      metadata: {
        name: NAMES['namespace-name'] + '-' + service.name + '-vsvc',
        namespace: NAMES['namespace-name'],
      },
      spec: {
        gateways: [
          NAMES['namespace-name'] + '-gateway',
        ],
        hosts: [
          host,
        ],
        http: [
          {
            corsPolicy: {
              allowCredentials: true,
              allowHeaders: [
                'Authorization',
                'keep-alive',
                'user-agent',
                'cache-control',
                'content-type',
                'content-transfer-encoding',
                'custom-header-1',
                'x-accept-content-transfer-encoding',
                'x-accept-response-streaming',
                'x-user-agent',
                'x-http2-port-web',
                'x-grpc-web',
                'http2-port-timeout',
              ],
              allowMethods: [
                'POST',
                'GET',
                'PUT',
                'DELETE',
                'OPTIONS',
              ],
              allowOrigins: [
                {
                  exact: '*',
                },
              ],
              maxAge: '24h',
            },
            route: [
              {
                destination: {
                  host: service.name,
                  port: {
                    number: service.externalPort,
                  },
                },
              },
            ],
          },
        ],
      },
    }
  ),
}
