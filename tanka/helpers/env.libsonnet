{
  LABEL: '<LABEL_NAME>',
  BRANCH_NAME: '<BRANCH_NAME>',
  STAGE_ID: '<STAGE_ID>',
  BASE_NAME: if self.BRANCH_NAME == 'main' then self.LABEL else self.LABEL + '-' + self.STAGE_ID,
  BASE_URL: 'playtowinapps.com',
  VERSION: 'v-<VERSION_NUMBER>',
  IMAGE_TAG: '<IMAGE_TAG>',
  REGISTRY_BASE: 'max3014/',
  NAMES: {
    'certificate-name': $.BASE_NAME + '-cert',
    'certificate-secret': $.BASE_NAME + '-secret',
    'namespace-name': $.BASE_NAME,
    'service-account': 'boints-sa',
    'vault-folder-name': if $.BRANCH_NAME == 'main' then 'boints' else $.BASE_NAME,
  },
  SERVICES: [
    {
      name: 'auth',
      protocol: 'http2',
      exposed: true,
      externalPort: 8080,
      internalPort: 8080,
      limitCPU: '400m',
      limitRAM: '384Mi',
      requestCPU: '100m',
      requestRAM: '256Mi',
      replicas: if $.BRANCH_NAME == 'main' then 8 else 1,
      command: 'cat /vault/secrets/* >> /vault/env && source /vault/env && npm run start:prod',
      annotations: $.ANNOTATIONS[self.name],
    },
    {
      name: 'appsflyer-provider',
      protocol: 'http2',
      exposed: true,
      externalPort: 8080,
      internalPort: 8080,
      limitCPU: '400m',
      limitRAM: '384Mi',
      requestCPU: '100m',
      requestRAM: '256Mi',
      replicas: if $.BRANCH_NAME == 'main' then 2 else 1,
      command: 'cat /vault/secrets/* >> /vault/env && source /vault/env && npm run start:prod',
      annotations: $.ANNOTATIONS[self.name],
    },
    {
      name: 'balance',
      protocol: 'http2',
      exposed: true,
      externalPort: 8080,
      internalPort: 8080,
      limitCPU: '400m',
      limitRAM: '384Mi',
      requestCPU: '100m',
      requestRAM: '256Mi',
      replicas: if $.BRANCH_NAME == 'main' then 8 else 1,
      command: 'cat /vault/secrets/* >> /vault/env && source /vault/env && npm run start:prod',
      annotations: $.ANNOTATIONS[self.name],
    },
    {
      name: 'earning',
      protocol: 'http',
      exposed: true,
      externalPort: 8080,
      internalPort: 8080,
      limitCPU: '400m',
      limitRAM: '384Mi',
      requestCPU: '100m',
      requestRAM: '256Mi',
      replicas: if $.BRANCH_NAME == 'main' then 8 else 1,
      command: 'cat /vault/secrets/* >> /vault/env && source /vault/env && npm run start:prod',
      annotations: $.ANNOTATIONS[self.name],
    },
    {
      name: 'leaderboard',
      protocol: 'http2',
      exposed: true,
      externalPort: 8080,
      internalPort: 8080,
      limitCPU: '400m',
      limitRAM: '384Mi',
      requestCPU: '100m',
      requestRAM: '256Mi',
      replicas: if $.BRANCH_NAME == 'main' then 4 else 1,
      command: 'cat /vault/secrets/* >> /vault/env && source /vault/env && npm run start:prod',
      annotations: $.ANNOTATIONS[self.name],
    },
    {
      name: 'reward',
      protocol: 'http2',
      exposed: true,
      externalPort: 8080,
      internalPort: 8080,
      limitCPU: '400m',
      limitRAM: '384Mi',
      requestCPU: '100m',
      requestRAM: '256Mi',
      replicas: if $.BRANCH_NAME == 'main' then 2 else 1,
      command: 'cat /vault/secrets/* >> /vault/env && source /vault/env && npm run start:prod',
      annotations: $.ANNOTATIONS[self.name],
    },
    {
      name: 'schedule',
      protocol: 'http2',
      exposed: false,
      externalPort: 8080,
      internalPort: 8080,
      limitCPU: '400m',
      limitRAM: '384Mi',
      requestCPU: '100m',
      requestRAM: '256Mi',
      replicas: if $.BRANCH_NAME == 'main' then 1 else 1,
      command: 'cat /vault/secrets/* >> /vault/env && source /vault/env && npm run start:prod',
      annotations: $.ANNOTATIONS[self.name],
    },
    {
      name: 'sentinel',
      protocol: 'http',
      exposed: false,
      externalPort: 9464,
      internalPort: 9464,
      limitCPU: '400m',
      limitRAM: '384Mi',
      requestCPU: '100m',
      requestRAM: '256Mi',
      replicas: if $.BRANCH_NAME == 'main' then 1 else 1,
      command: 'cat /vault/secrets/* >> /vault/env && source /vault/env && npm run start:prod',
      annotations: $.ANNOTATIONS[self.name],
    },
    {
      name: 'stats',
      protocol: 'http2',
      exposed: true,
      externalPort: 8080,
      internalPort: 8080,
      limitCPU: '400m',
      limitRAM: '384Mi',
      requestCPU: '200m',
      requestRAM: '256Mi',
      replicas: if $.BRANCH_NAME == 'main' then 6 else 1,
      command: 'cat /vault/secrets/* >> /vault/env && source /vault/env && npm run start:prod',
      annotations: $.ANNOTATIONS[self.name],
    },
    {
      name: 'cashout',
      protocol: 'http2',
      exposed: true,
      externalPort: 8080,
      internalPort: 8080,
      limitCPU: '400m',
      limitRAM: '384Mi',
      requestCPU: '100m',
      requestRAM: '256Mi',
      replicas: if $.BRANCH_NAME == 'main' then 2 else 1,
      command: 'cat /vault/secrets/* >> /vault/env && source /vault/env && npm run start:prod',
      annotations: $.ANNOTATIONS[self.name],
    },
    {
      name: 'referral',
      protocol: 'http2',
      exposed: true,
      externalPort: 8080,
      internalPort: 8080,
      limitCPU: '400m',
      limitRAM: '384Mi',
      requestCPU: '100m',
      requestRAM: '256Mi',
      replicas: if $.BRANCH_NAME == 'main' then 2 else 1,
      command: 'cat /vault/secrets/* >> /vault/env && source /vault/env && npm run start:prod',
      annotations: $.ANNOTATIONS[self.name],
    },
    {
      name: 'tag',
      protocol: 'http2',
      exposed: true,
      externalPort: 8080,
      internalPort: 8080,
      limitCPU: '400m',
      limitRAM: '384Mi',
      requestCPU: '100m',
      requestRAM: '256Mi',
      replicas: if $.BRANCH_NAME == 'main' then 2 else 1,
      command: 'cat /vault/secrets/* >> /vault/env && source /vault/env && npm run start:prod',
      annotations: $.ANNOTATIONS[self.name],
    },
    {
      name: 'double-boints',
      protocol: 'http2',
      exposed: true,
      externalPort: 8080,
      internalPort: 8080,
      limitCPU: '400m',
      limitRAM: '384Mi',
      requestCPU: '100m',
      requestRAM: '256Mi',
      replicas: if $.BRANCH_NAME == 'main' then 2 else 1,
      command: 'cat /vault/secrets/* >> /vault/env && source /vault/env && npm run start:prod',
      annotations: $.ANNOTATIONS[self.name],
    },
    {
      name: 'orchestrator',
      protocol: 'http2',
      exposed: false,
      externalPort: 8080,
      internalPort: 8080,
      limitCPU: '400m',
      limitRAM: '384Mi',
      requestCPU: '100m',
      requestRAM: '256Mi',
      replicas: if $.BRANCH_NAME == 'main' then 2 else 1,
      command: 'cat /vault/secrets/* >> /vault/env && source /vault/env && npm run start:prod',
      annotations: $.ANNOTATIONS[self.name],
    },
    {
      name: 'google-play',
      protocol: 'http',
      exposed: false,
      externalPort: 8080,
      internalPort: 8080,
      limitCPU: '400m',
      limitRAM: '384Mi',
      requestCPU: '100m',
      requestRAM: '256Mi',
      replicas: if $.BRANCH_NAME == 'main' then 2 else 1,
      command: 'cat /vault/secrets/* >> /vault/env && source /vault/env && npm run start:prod',
      annotations: $.ANNOTATIONS[self.name],
    },
    {
      name: 'offerwall',
      protocol: 'http2',
      exposed: true,
      externalPort: 8080,
      internalPort: 8080,
      limitCPU: '400m',
      limitRAM: '384Mi',
      requestCPU: '100m',
      requestRAM: '256Mi',
      replicas: if $.BRANCH_NAME == 'main' then 2 else 1,
      command: 'cat /vault/secrets/* >> /vault/env && source /vault/env && npm run start:prod',
      annotations: $.ANNOTATIONS[self.name]
    },
    {
      name: 'offerwall-fe',
      protocol: 'http',
      exposed: true,
      externalPort: 8080,
      internalPort: 8080,
      limitCPU: '400m',
      limitRAM: '384Mi',
      requestCPU: '100m',
      requestRAM: '256Mi',
      replicas: if $.BRANCH_NAME == 'main' then 4 else 1,
      command: 'cat /vault/secrets/* >> /vault/env && source /vault/env && npm run start:prod',
      annotations: $.ANNOTATIONS[self.name],
    },
    {
      name: 'bff',
      protocol: 'http',
      exposed: true,
      externalPort: 8080,
      internalPort: 8080,
      limitCPU: '400m',
      limitRAM: '384Mi',
      requestCPU: '100m',
      requestRAM: '256Mi',
      replicas: if $.BRANCH_NAME == 'main' then 4 else 1,
      command: 'cat /vault/secrets/* >> /vault/env && source /vault/env && npm run start:prod',
      annotations: $.ANNOTATIONS[self.name],
    }
  ],
  ANNOTATIONS: {
    'appsflyer-provider': VAULT.base + VAULT.common + VAULT.postgres + VAULT.nats + VAULT.redis,
    auth: VAULT.base + VAULT.auth + VAULT.common + VAULT.postgres + VAULT.nats + VAULT.redis,
    balance: VAULT.base + VAULT.balance + VAULT.common + VAULT.postgres_without_db_name + VAULT.nats + VAULT.redis,
    cashout: VAULT.base + VAULT.cashout + VAULT.appsflyer + VAULT.paypal + VAULT.pushwoosh + VAULT.common + VAULT.postgres_without_db_name + VAULT.nats + VAULT.redis,
    referral: VAULT.base + VAULT.referral + VAULT.common + VAULT.postgres_without_db_name + VAULT.nats + VAULT.redis,
    tag: VAULT.base + VAULT.tag + VAULT.common + VAULT.postgres_without_db_name + VAULT.nats + VAULT.redis,
    'double-boints': VAULT.base + VAULT['double-boints'] + VAULT.common + VAULT.postgres_without_db_name + VAULT.nats + VAULT.redis,
    earning: VAULT.base + VAULT.earning + VAULT.common + VAULT.postgres_without_db_name + VAULT.nats + VAULT.redis,
    leaderboard: VAULT.base + VAULT.common + VAULT.postgres + VAULT.nats + VAULT.redis,
    schedule: VAULT.base + VAULT.common + VAULT.postgres,
    stats: VAULT.base + VAULT.stats + VAULT.common + VAULT.appsflyer + VAULT.paypal + VAULT.pushwoosh + VAULT.postgres_without_db_name + VAULT.nats + VAULT.redis,
    sentinel: VAULT.base + VAULT.common + VAULT.nats + VAULT.redis,
    reward: VAULT.base + VAULT.reward + VAULT.common + VAULT.postgres_without_db_name + VAULT.nats + VAULT.redis,
    orchestrator: VAULT.base + VAULT.common + VAULT.postgres_without_db_name + VAULT.nats + VAULT.redis,
    offerwall: VAULT.base + VAULT.common + VAULT.postgres_without_db_name + VAULT.nats + VAULT.redis,
    'offerwall-fe': VAULT.base + VAULT.common + VAULT.postgres_without_db_name + VAULT.nats + VAULT.redis,
    bff: VAULT.base + VAULT.common + VAULT.postgres_without_db_name + VAULT.nats + VAULT.redis,
    'google-play': VAULT.base + VAULT.common + VAULT.postgres_without_db_name + VAULT.nats + VAULT.redis,
  },

  local VAULT = {
    base: {
      'vault.hashicorp.com/agent-init-first': 'true',
      'vault.hashicorp.com/agent-inject': 'true',
      'vault.hashicorp.com/role': 'boints-role',
      'vault.hashicorp.com/tls-skip-verify': 'true',
      'vault.hashicorp.com/agent-pre-populate-only': 'true',
    },
    common: {
      'vault.hashicorp.com/agent-inject-secret-common': 'secret/%s/boints-common' % $.NAMES['vault-folder-name'],
      'vault.hashicorp.com/agent-inject-template-common': |||
        {{ with secret "secret/%s/boints-common" -}}
          export SERVICE_DEVICE_ID="{{ .Data.SERVICE_DEVICE_ID }}"
          export ADMIN_DEVICE_ID="{{ .Data.ADMIN_DEVICE_ID }}"
          export PLAYER_DEVICE_ID="{{ .Data.PLAYER_DEVICE_ID }}"
          export GRPC_CLIENTS_SSL="{{ .Data.GRPC_CLIENTS_SSL }}"
          export ENVIRONMENT="{{ .Data.ENVIRONMENT }}"
          export IP_QUALITY_KEY="{{ .Data.IP_QUALITY_KEY }}"
          export IP_QUALITY_URL="{{ .Data.IP_QUALITY_URL }}"
        {{- end }}
      ||| % $.NAMES['vault-folder-name'],
    },
    postgres: {
      'vault.hashicorp.com/agent-inject-secret-postgres': 'secret/%s/postgres-secret' % $.NAMES['vault-folder-name'],
      'vault.hashicorp.com/agent-inject-template-postgres': |||
        {{ with secret "secret/%s/postgres-secret" -}}
          export DB_USER_WR="{{ .Data.ADMIN_USER }}"
          export DB_USER_RO="{{ .Data.ADMIN_USER }}"
          export DB_PASS_WR="{{ .Data.ADMIN_PASSWORD }}"
          export DB_PASS_RO="{{ .Data.ADMIN_PASSWORD }}"
          export DB_PORT_WR="{{ .Data.DB_PORT }}"
          export DB_PORT_RO="{{ .Data.DB_PORT }}"
          export DB_NAME_WR="{{ .Data.DB_NAME }}"
          export DB_NAME_RO="{{ .Data.DB_NAME }}"
          export DB_HOST_WR="{{ .Data.DB_HOST_WR }}"
          export DB_HOST_RO="{{ .Data.DB_HOST_RO }}"
          export DB_CONNECTION_LIMIT="{{ .Data.DB_CONNECTION_LIMIT }}"
          export DB_IDLE_CONNECTIONS_COUNT="{{ .Data.DB_IDLE_CONNECTIONS_COUNT }}"
          export DB_IDLE_IN_TX_SESSION_TIMEOUT_MILLIS="{{ .Data.DB_IDLE_IN_TX_SESSION_TIMEOUT_MILLIS }}"
        {{- end }}
      ||| % $.NAMES['vault-folder-name'],
    },
    postgres_without_db_name: {
      'vault.hashicorp.com/agent-inject-secret-postgres': 'secret/%s/postgres-secret' % $.NAMES['vault-folder-name'],
      'vault.hashicorp.com/agent-inject-template-postgres': |||
        {{ with secret "secret/%s/postgres-secret" -}}
          export DATABASE_LOGGING="{{ .Data.DATABASE_LOGGING }}"
          export DATABASE_RUN_MIGRATIONS="{{ .Data.DATABASE_RUN_MIGRATIONS }}"
          export DB_USER_WR="{{ .Data.ADMIN_USER }}"
          export DB_USER_RO="{{ .Data.ADMIN_USER }}"
          export DB_PASS_WR="{{ .Data.ADMIN_PASSWORD }}"
          export DB_PASS_RO="{{ .Data.ADMIN_PASSWORD }}"
          export DB_PORT_WR="{{ .Data.DB_PORT }}"
          export DB_PORT_RO="{{ .Data.DB_PORT }}"
          export DB_HOST_WR="{{ .Data.DB_HOST_WR }}"
          export DB_HOST_RO="{{ .Data.DB_HOST_RO }}"
          export DB_CONNECTION_LIMIT="{{ .Data.DB_CONNECTION_LIMIT }}"
          export DB_IDLE_CONNECTIONS_COUNT="{{ .Data.DB_IDLE_CONNECTIONS_COUNT }}"
          export DB_IDLE_IN_TX_SESSION_TIMEOUT_MILLIS="{{ .Data.DB_IDLE_IN_TX_SESSION_TIMEOUT_MILLIS }}"
        {{- end }}
      ||| % $.NAMES['vault-folder-name'],
    },
    nats: {
      'vault.hashicorp.com/agent-inject-secret-nats': 'secret/%s/nats' % $.NAMES['vault-folder-name'],
      'vault.hashicorp.com/agent-inject-template-nats': |||
        {{ with secret "secret/%s/nats" -}}
          export NATS_URL="{{ .Data.NATS_URL }}"
          export NATS_SERVER_URL="{{ .Data.NATS_SERVER_URL }}"
          export NATS_ACK_WAIT="{{ .Data.NATS_ACK_WAIT }}"
        {{- end }}
      ||| % $.NAMES['vault-folder-name'],
    },
    redis: {
      'vault.hashicorp.com/agent-inject-secret-redis': 'secret/%s/redis' % $.NAMES['vault-folder-name'],
      'vault.hashicorp.com/agent-inject-template-redis': |||
        {{ with secret "secret/%s/redis" -}}
          export REDIS_URL_WR="{{ .Data.REDIS_URL_WR }}"
          export REDIS_URL_RO="{{ .Data.REDIS_URL_RO }}"
          export REDIS_PASSWORD="{{ .Data.REDIS_KEY }}"
        {{- end }}
      ||| % $.NAMES['vault-folder-name'],
    },
    pushwoosh: {
      'vault.hashicorp.com/agent-inject-secret-pushwoosh': 'secret/%s/pushwoosh' % $.NAMES['vault-folder-name'],
      'vault.hashicorp.com/agent-inject-template-pushwoosh': |||
        {{ with secret "secret/%s/pushwoosh" -}}
          export PUSHWOOSH_URL="{{ .Data.PUSHWOOSH_URL }}"
          export PUSHWOOSH_TOKEN="{{ .Data.PUSHWOOSH_TOKEN }}"
          export PUSHWOOSH_APP_ID="{{ .Data.PUSHWOOSH_APP_ID }}"
        {{- end }}
      |||,
    },
    paypal: {
      'vault.hashicorp.com/agent-inject-secret-paypal': 'secret/%s/paypal' % $.NAMES['vault-folder-name'],
      'vault.hashicorp.com/agent-inject-template-paypal': |||
        {{ with secret "secret/%s/paypal" -}}
          export PAYPAL_URL="{{ .Data.PAYPAL_URL }}"
          export PAYPAL_USERNAME="{{ .Data.PAYPAL_USERNAME }}"
          export PAYPAL_PASSWORD="{{ .Data.PAYPAL_PASSWORD }}"
        {{- end }}
      ||| % $.NAMES['vault-folder-name'],
    },
    appsflyer: {
      'vault.hashicorp.com/agent-inject-secret-appsflyer': 'secret/%s/appsflyer' % $.NAMES['vault-folder-name'],
      'vault.hashicorp.com/agent-inject-template-appsflyer': |||
        {{ with secret "secret/%s/appsflyer" -}}
          export APPSFLYER_URL="{{ .Data.APPSFLYER_URL }}"
          export APPSFLYER_AUTHENTICATION_KEY="{{ .Data.APPSFLYER_AUTHENTICATION_KEY }}"
        {{- end }}
      ||| % $.NAMES['vault-folder-name'],
    },
    auth: {
      'vault.hashicorp.com/agent-inject-secret-auth': 'secret/%s/boints-auth' % $.NAMES['vault-folder-name'],
      'vault.hashicorp.com/agent-inject-template-auth': |||
        {{ with secret "secret/%s/boints-auth" -}}
          export JWT_SECRET="{{ .Data.JWT_SECRET }}"
          export MAX_COUNT_USER_WITH_SAME_IP="{{ .Data.MAX_COUNT_USER_WITH_SAME_IP }}"
        {{- end }}
      ||| % $.NAMES['vault-folder-name'],
    },
    balance: {
      'vault.hashicorp.com/agent-inject-secret-balance': 'secret/%s/boints-balance' % $.NAMES['vault-folder-name'],
      'vault.hashicorp.com/agent-inject-template-balance': |||
        {{ with secret "secret/%s/boints-balance" -}}
          export NATS_DELIVER_GROUP="{{ .Data.NATS_DELIVER_GROUP }}"
          export DB_NAME_WR="{{ .Data.DB_NAME }}"
          export DB_NAME_RO="{{ .Data.DB_NAME }}"
        {{- end }}
      ||| % $.NAMES['vault-folder-name'],
    },
    cashout: {
      'vault.hashicorp.com/agent-inject-secret-cashout': 'secret/%s/boints-cashout' % $.NAMES['vault-folder-name'],
      'vault.hashicorp.com/agent-inject-template-cashout': |||
        {{ with secret "secret/%s/boints-cashout" -}}
          export NATS_DELIVER_GROUP="{{ .Data.NATS_DELIVER_GROUP }}"
          export DB_NAME_WR="{{ .Data.DB_NAME }}"
          export DB_NAME_RO="{{ .Data.DB_NAME }}"
        {{- end }}
      ||| % $.NAMES['vault-folder-name'],
    },
    referral: {
      'vault.hashicorp.com/agent-inject-secret-referral': 'secret/%s/boints-referral' % $.NAMES['vault-folder-name'],
      'vault.hashicorp.com/agent-inject-template-referral': |||
        {{ with secret "secret/%s/boints-referral" -}}
          export NATS_DELIVER_GROUP="{{ .Data.NATS_DELIVER_GROUP }}"
          export DB_NAME_WR="{{ .Data.DB_NAME }}"
          export DB_NAME_RO="{{ .Data.DB_NAME }}"
        {{- end }}
      ||| % $.NAMES['vault-folder-name'],
    },
    'double-boints': {
      'vault.hashicorp.com/agent-inject-secret-double-boints': 'secret/%s/boints-double-boints' % $.NAMES['vault-folder-name'],
      'vault.hashicorp.com/agent-inject-template-double-boints': |||
        {{ with secret "secret/%s/boints-double-boints" -}}
          export NATS_DELIVER_GROUP="{{ .Data.NATS_DELIVER_GROUP }}"
          export DB_NAME_WR="{{ .Data.DB_NAME }}"
          export DB_NAME_RO="{{ .Data.DB_NAME }}"
        {{- end }}
      ||| % $.NAMES['vault-folder-name'],
    },
    tag: {
      'vault.hashicorp.com/agent-inject-secret-tag': 'secret/%s/boints-tag' % $.NAMES['vault-folder-name'],
      'vault.hashicorp.com/agent-inject-template-tag': |||
        {{ with secret "secret/%s/boints-tag" -}}
          export NATS_DELIVER_GROUP="{{ .Data.NATS_DELIVER_GROUP }}"
          export DB_NAME_WR="{{ .Data.DB_NAME }}"
          export DB_NAME_RO="{{ .Data.DB_NAME }}"
        {{- end }}
      ||| % $.NAMES['vault-folder-name'],
    },
    stats: {
      'vault.hashicorp.com/agent-inject-secret-stats': 'secret/%s/boints-stats' % $.NAMES['vault-folder-name'],
      'vault.hashicorp.com/agent-inject-template-stats': |||
        {{ with secret "secret/%s/boints-stats" -}}
          export NATS_DELIVER_GROUP="{{ .Data.NATS_DELIVER_GROUP }}"
          export DB_NAME_WR="{{ .Data.DB_NAME }}"
          export DB_NAME_RO="{{ .Data.DB_NAME }}"
        {{- end }}
      ||| % $.NAMES['vault-folder-name'],
    },
    earning: {
      'vault.hashicorp.com/agent-inject-secret-earning': 'secret/%s/boints-earning' % $.NAMES['vault-folder-name'],
      'vault.hashicorp.com/agent-inject-template-earning': |||
        {{ with secret "secret/%s/boints-earning" -}}
          export TAPJOY_SECRET_KEY="{{ .Data.TAPJOY_SECRET_KEY }}"
          export FYBER_SECRET_KEY="{{ .Data.FYBER_SECRET_KEY }}"
          export IRON_SOURCE_SECRET_KEY="{{ .Data.IRON_SOURCE_SECRET_KEY }}"
          export POLLFISH_SECRET_KEY="{{ .Data.POLLFISH_SECRET_KEY }}"
          export REVU_SECRET_KEY="{{ .Data.REVU_SECRET_KEY }}"
          export MAX_SECRET_KEY="{{ .Data.MAX_SECRET_KEY }}"
          export ADJOE_SECRET_KEY="{{ .Data.ADJOE_SECRET_KEY }}"
          export AYE_SECRET_KEY="{{ .Data.AYE_SECRET_KEY }}"
          export ADGEM_TEST_APP_SECRET_KEY="{{ .Data.ADGEM_TEST_APP_SECRET_KEY }}"
          export ADGEM_PROD_APP_SECRET_KEY="{{ .Data.ADGEM_PROD_APP_SECRET_KEY }}"
          export INBRAIN_SECRET_KEY="{{ .Data.INBRAIN_SECRET_KEY }}"
          export TEST_SECRET_KEY="{{ .Data.TEST_SECRET_KEY }}"
          export INBRAIN_SECRET_KEY="{{ .Data.INBRAIN_SECRET_KEY }}"
          export IOS_INBRAIN_SECRET_KEY="{{ .Data.IOS_INBRAIN_SECRET_KEY }}"
          export TEST_HRS_SECRET_KEY="{{ .Data.TEST_HRS_SECRET_KEY }}"
          export HRS_APPSAMURAI_SECRET_KEY="{{ .Data.HRS_APPSAMURAI_SECRET_KEY }}"
        {{- end }}
      ||| % $.NAMES['vault-folder-name'],
    },
    reward: {
      'vault.hashicorp.com/agent-inject-secret-reward': 'secret/%s/boints-reward' % $.NAMES['vault-folder-name'],
      'vault.hashicorp.com/agent-inject-template-reward': |||
        {{ with secret "secret/%s/boints-reward" -}}
          export GRPC_SERVER_PORT="{{ .Data.GRPC_SERVER_PORT }}"
          export NATS_DELIVER_GROUP="{{ .Data.NATS_DELIVER_GROUP }}"
          export DB_NAME_WR="{{ .Data.DB_NAME }}"
          export DB_NAME_RO="{{ .Data.DB_NAME }}"
        {{- end }}
      ||| % $.NAMES['vault-folder-name'],
    },
  },
}
