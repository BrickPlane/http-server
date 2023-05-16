local deployment = import '../../helpers/deployment.libsonnet';
local env = import '../../helpers/env.libsonnet';
local service = import '../../helpers/service.libsonnet';
local SERVICES = env.SERVICES;
local ANNOTATIONS = env.ANNOTATIONS;

{
  payload: [
    [
      deployment.new(SERVICE),
      service.new(SERVICE),
    ]
    for SERVICE in SERVICES
  ],
}
