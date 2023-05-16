local deployment = (import '1.21/main.libsonnet').apps.v1.deployment;
local container = (import '1.21/main.libsonnet').core.v1.container;
local env = import 'env.libsonnet';
local LABEL = env.LABEL;
local REGISTRY_BASE = env.REGISTRY_BASE;
local VERSION = env.VERSION;
local BRANCH_NAME = env.BRANCH_NAME;
local NAMES = env.NAMES;
local IMAGE_TAG = env.IMAGE_TAG;

{
  new(
    service,
  ): (
    deployment.new(
      name=service.name,
      replicas=service.replicas,
      containers=[
        container.new(
          name=service.name,
          image=REGISTRY_BASE + LABEL + '-' + service.name + ':' + IMAGE_TAG,
        )
        + container.withCommand(
          '/bin/sh'
        )
        + container.withArgs([
          '-c',
          service.command,
        ])
        + container.withImagePullPolicy(
          'Always'
        )
        + container.withPorts([{
          containerPort: service.internalPort,
          protocol: 'TCP',
        }])
        + container.resources.withRequests({
          memory: service.requestRAM,
          cpu: service.requestCPU,
        })
        + container.resources.withLimits({
          memory: service.limitRAM,
          cpu: service.limitCPU,
        })
        + container.withTerminationMessagePath(
          '/dev/termination-log',
        )
        + container.withTerminationMessagePolicy(
          'File',
        )
        + container.withVolumeMounts([{
          mountPath: '/usr/src/app/log',
          name: 'log-vol',
        }]),
      ]
    )
    + deployment.metadata.withLabels({
      app: NAMES['namespace-name'],
      version: VERSION,
    })
    + deployment.metadata.withNamespace(
      NAMES['namespace-name']
    )
    + deployment.spec.selector.withMatchLabels({
      app: NAMES['namespace-name'],
      svc: service.name,
    })
    + deployment.spec.strategy.rollingUpdate.withMaxSurge(
      '25%',
    )
    + deployment.spec.strategy.rollingUpdate.withMaxUnavailable(
      '25%',
    )
    + deployment.spec.strategy.withType(
      'RollingUpdate',
    )
    + deployment.spec.withProgressDeadlineSeconds(
      600,
    )
    + deployment.spec.withRevisionHistoryLimit(
      10,
    )
    + deployment.spec.template.metadata.withLabels({
      app: NAMES['namespace-name'],
      version: VERSION,
      svc: service.name,
    })
    + deployment.spec.template.metadata.withAnnotations(
      service.annotations,
    )
    + deployment.spec.template.spec.affinity.podAffinity.withPreferredDuringSchedulingIgnoredDuringExecution([{
      weight: 100,
      podAffinityTerm: {
        labelSelector: {
          matchExpressions: [{
            key: 'svc',
            operator: 'In',
            values: [
              service.name,
            ],
          }],
        },
        topologyKey: 'kubernetes.io/hostname',
      },
    }])
    + deployment.spec.template.spec.withDnsPolicy(
      'ClusterFirst',
    )
    + deployment.spec.template.spec.withServiceAccount(
      NAMES['service-account'],
    )
    + deployment.spec.template.spec.withRestartPolicy(
      'Always',
    )
    + deployment.spec.template.spec.withVolumes([{
      hostPath: {
        path: '/var/log' + NAMES['namespace-name'],
        type: 'DirectoryOrCreate',
      },
      name: 'log-vol',
    }])
  ),
}
