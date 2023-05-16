local env = import 'env.libsonnet';
local NAMES = env.NAMES;

{
  new(
    hosts,
  ): (
    {
      apiVersion: 'cert-manager.io/v1',
      kind: 'Certificate',
      metadata: {
        name: NAMES['certificate-name'],
        namespace: 'istio-system',
      },
      spec: {
        secretName: NAMES['certificate-secret'],
        duration: '2160h',
        renewBefore: '360h',
        isCA: false,
        privateKey: {
          algorithm: 'RSA',
          encoding: 'PKCS1',
          size: 2048,
        },
        usages: [
          'server auth',
          'client auth',
        ],
        dnsNames: hosts,
        issuerRef: {
          name: 'letsencrypt-prod-cluster',
          kind: 'ClusterIssuer',
          group: 'cert-manager.io',
        },
      },
    }
  ),
}
