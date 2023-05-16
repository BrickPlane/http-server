local env = import 'env.libsonnet';
local NAMES = env.NAMES;

{
  new(): (
    {
      apiVersion: 'v1',
      kind: 'ServiceAccount',
      metadata: {
        name: 'boints-sa',
        namespace: NAMES['namespace-name'],
      },
    }
  ),
}
