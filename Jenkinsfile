pipeline {

  options {
    ansiColor('xterm')
  }

  agent none
   
  environment {
    BRANCH_NAME = BRANCH_NAME.toLowerCase().replace("/", "-")
    JOB_NAME = "${JOB_NAME.split('/')[0]}"
  }

  stages {
    stage('Build & Push Images') {
      parallel {
        stage('http-project') {
          agent {
            kubernetes {
              yamlFile 'builders/builder-http-project.yaml'
            }
          }
          steps {
            container('buildkitd-http-project') {
              script {
                sh '''
                buildctl build \
                  --output type=image,name=docker.io/max3014/${JOB_NAME}:${BRANCH_NAME}-${BUILD_NUMBER},push=true \
                  --export-cache mode=max,type=registry,ref=docker.io/max3014/${JOB_NAME}:${BRANCH_NAME}-${BUILD_NUMBER}-buildcache \
                  --import-cache type=registry,ref=docker.io/max3014/${JOB_NAME}:${BRANCH_NAME}-${BUILD_NUMBER}-buildcache \
                  --frontend dockerfile.v0 \
                  --opt platform=linux/amd64 \
                  --local context=./ \
                  --local dockerfile=./ \
                  --opt filename=./Dockerfile
               '''
              }
            }
          }
        }
      }
    }
    
    // stage('Pre-deploy') {    
    //   parallel {
    //     stage('Istio Network Setup') {
    //       environment {
    //         IMAGE_TAG = "${env.BRANCH_NAME_MAIN != null ? env.BRANCH_NAME_MAIN : env.BRANCH_NAME }-${env.BUILD_NUMBER_MAIN != null ? env.BUILD_NUMBER_MAIN : env.BUILD_NUMBER}"
    //         STAGE_ID = "${env.STAGE_ID}"
    //       }
    //       agent {
    //         kubernetes {
    //           yamlFile 'builders/tanka.yaml'
    //         }
    //       }
    //       steps {
    //         container('kubectl-tanka') {
    //           withCredentials([file(credentialsId: 'kubeconfig', variable: 'KUBECONFIG')]) {
    //             sh 'sed -i "s/<VERSION_NUMBER>/${BUILD_NUMBER}/" tanka/helpers/env.libsonnet'
    //             sh 'sed -i "s/<IMAGE_TAG>/${IMAGE_TAG}/" tanka/helpers/env.libsonnet'
    //             sh 'sed -i "s/<LABEL_NAME>/${JOB_NAME}/" tanka/helpers/env.libsonnet'
    //             sh 'sed -i "s/<BRANCH_NAME>/${BRANCH_NAME}/" tanka/helpers/env.libsonnet'
    //             sh 'sed -i "s/<STAGE_ID>/${STAGE_ID}/" tanka/helpers/env.libsonnet'
    //             sh 'cd tanka && jb install && tk env set environments/initial-deploy --server=$(kubectl config view --minify --output jsonpath="{.clusters[*].cluster.server}") && tk apply environments/initial-deploy/main.jsonnet --auto-approve "always"'
    //           }
    //         }
    //       }
    //     }
    //   }
    // }

    // stage('Regular Deploy') {   
    //   environment {
    //     IMAGE_TAG = "${env.BRANCH_NAME_MAIN != null ? env.BRANCH_NAME_MAIN : env.BRANCH_NAME }-${env.BUILD_NUMBER_MAIN != null ? env.BUILD_NUMBER_MAIN : env.BUILD_NUMBER}"
    //     STAGE_ID = "${env.STAGE_ID}"
    //   }
    //   agent {
    //     kubernetes {
    //       yamlFile 'builders/tanka.yaml'
    //     }
    //   }
    //   steps {
    //     container('kubectl-tanka') {
    //       withCredentials([file(credentialsId: 'kubeconfig', variable: 'KUBECONFIG')]) {
    //         sh 'sed -i "s/<VERSION_NUMBER>/${BUILD_NUMBER}/" tanka/helpers/env.libsonnet'
    //         sh 'sed -i "s/<IMAGE_TAG>/${IMAGE_TAG}/" tanka/helpers/env.libsonnet'
    //         sh 'sed -i "s/<LABEL_NAME>/${JOB_NAME}/" tanka/helpers/env.libsonnet'
    //         sh 'sed -i "s/<BRANCH_NAME>/${BRANCH_NAME}/" tanka/helpers/env.libsonnet'
    //         sh 'sed -i "s/<STAGE_ID>/${STAGE_ID}/" tanka/helpers/env.libsonnet'
    //         sh 'cd tanka && jb install && tk env set environments/regular-deploy --server=$(kubectl config view --minify --output jsonpath="{.clusters[*].cluster.server}") && tk apply environments/regular-deploy/main.jsonnet --auto-approve "always"'
    //       }
    //     }
    //   }
    // } 
  }
}


