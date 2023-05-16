local env = import 'env.libsonnet';
local NAMES = env.NAMES;

{
  new(): (
    {
      apiVersion: 'v1',
      kind: 'Namespace',
      metadata: {
        name: NAMES['namespace-name'],
        labels: {
          'istio-injection': 'enabled',
          monitoring: 'prometheus',
        },
      },
    }
  ),
}
