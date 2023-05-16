local certificate = import '../../helpers/certificate.libsonnet';
local destinationRule = import '../../helpers/destination-rule.libsonnet';
local env = import '../../helpers/env.libsonnet';
local gateway = import '../../helpers/gateway.libsonnet';
local namespace = import '../../helpers/namespace.libsonnet';
local serviceAccount = import '../../helpers/service-account.libsonnet';
local virtualService = import '../../helpers/virtual-service.libsonnet';
local NAMES = env.NAMES;
local BASE_URL = env.BASE_URL;
local SERVICES = env.SERVICES;
local BRANCH_NAME = env.BRANCH_NAME;

{
  HOSTS_OBJ:: {
    [SERVICE.name]: if BRANCH_NAME == 'main' then '%s-%s.%s' % ['boints', SERVICE.name, BASE_URL] else '%s-%s.%s' % [NAMES['namespace-name'], SERVICE.name, BASE_URL]
    for SERVICE in SERVICES
    if SERVICE.exposed == true
  },

  HOSTS_ARR:: [
    self.HOSTS_OBJ[SERVICE.name]
    for SERVICE in SERVICES
    if SERVICE.exposed == true
  ],

  payload: [
    [
      virtualService.new(SERVICE, self.HOSTS_OBJ[SERVICE.name]),
      destinationRule.new(SERVICE),
    ]
    for SERVICE in SERVICES
    if SERVICE.exposed == true
  ],

  namespace: namespace.new(),
  'service-account': serviceAccount.new(),
  gateway: gateway.new(self.HOSTS_ARR),
  certificate: certificate.new(self.HOSTS_ARR),
}
