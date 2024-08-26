#!groovy

@Library('jenkins-shared-library@main') _

def tools = new org.color()
def checkout = new org.checkout()

pipeline {

  agent {
    kubernetes {
      retries 2
      yamlFile 'KubernetesPod.yaml'   
    }
  }

  parameters {
    choice choices: ['main', 'pre', 'test'], name: 'branch_name'
  }

  options {
    timestamps()      
    parallelsAlwaysFailFast()
    timeout(time: 600, unit: 'SECONDS') 
    disableConcurrentBuilds(abortPrevious: true) 
    buildDiscarder(logRotator(numToKeepStr: '30'))
    skipDefaultCheckout() 
  }

  environment {
    String year = new Date().format("yyyy") 
    String month = new Date().format("MMdd") 
    String day = new Date().format("HHmm") 
    String second = new Date().format("ss")         
    giturl = "http://10.0.7.30/golang/go.git" 
    images_head = "registry.cn-hangzhou.aliyuncs.com/tool-bucket/tool"   
    ImageTag = "${images_head}:${BUILD_USER}-${branch_name}-${year}${month}${day}${second}-${BUILD_TAG}"           
  }

  post {
    failure {
      script {
        manager.addShortText("${BUILD_NUMBER}次构建! 构建失败!!!") 
      }
    }
    success {
      script {
        manager.addShortText("${BUILD_NUMBER}次构建! 构建成功!!!")
      }               
    }                                  
    aborted {
      script {
        manager.addShortText("${BUILD_NUMBER}次构建!构建取消!!!")
      }
    }
    always {
      script {
        manager.addShortText("构建用户: ${BUILD_USER}")
      }
    }           
  } 

  stages {
    stage('1.克隆代码') {
      steps {
        container('git') {
          script {
            tools.PrintMessage("1.克隆代码","blue")   
            checkout.scm(branch_name,giturl)   
          }
        }
      }
    }        
    stage('2.SonarQube扫描') {
      steps {
        container('sonar') {
          script {
            tools.PrintMessage("2.SonarQube扫描","blue") 
            withSonarQubeEnv('sonarqube') {
              sh '''
                sonar-scanner \
                -Dsonar.projectKey=test-go \
                -Dsonar.projectName=test-go \
                -Dsonar.projectVersion=test-go-${BUILD_NUMBER} \
                -Dsonar.ws.timeout=30 \
                -Dsonar.sources=. \
                -Dsonar.sourceEncoding=UTF-8
                sleep 3
              '''
            }            
          }
        }
      }
    }    
    stage("3.质量阀控制"){
      steps {
        container('sonar') {
          script {      
            tools.PrintMessage("3.质量阀控制","blue")
            timeout(time: 10, unit: 'SECONDS') { 
              def qg = waitForQualityGate('sonarqube') 
              if (qg.status != 'OK') {
                error "未通过Sonarqube的代码质量阈检查，请及时修改！failure: ${qg.status}"
              }
            }
          }
        }
      }
    }    
    stage('4、构建镜像') {
      steps {
        container('kaniko') {
          script {
            tools.PrintMessage("4.构建镜像","blue")
            sh """
              /kaniko/executor --dockerfile=Dockerfile \
              --destination=${ImageTag} \
              --context=. \
              --cache-copy-layers \
              --cache=true \
              --cache-repo=${images_head}
            """
          }
        }
      }
    }       
    stage('5.服务部署') {  
      steps {  
        container('kubectl') {  
          script {
            tools.PrintMessage("5.服务部署","blue")
            sh """
              sed -i "s#image: .*#image: ${ImageTag}#" deploy.yaml
              kubectl apply -f deploy.yaml 
              # kubectl create secret generic kubeconfig-secret --from-file=config=/root/.kube/config 
            """            
          }
        }  
      }  
    }     
    stage('6.MeterShpere接口测试') {
      steps {
        script {
          tools.PrintMessage("6.MeterShpere接口测试","blue")
          meterSphere method: 'testPlan',
          mode: 'serial', 
          msAccessKey: 'OYPWJwNk9vi6JNiF', 
          msEndpoint: 'http://10.0.7.27:8081/', 
          msSecretKey: '7D0UlTZXXTZMt0fD', 
          openMode: 'auth', 
          projectId: 'a6ad182f-f6f1-44de-bdc7-3e6aef7c9e84', 
          projectName: '', 
          projectType: 'projectId', 
          resourcePoolId: '237d98d4-40b9-11ee-a07c-0242ac1e0a09', 
          testCaseId: '', testCaseName: '', 
          testPlanId: '99355526-981f-4927-bd6d-b37e54eb00c1', 
          testPlanName: '', 
          workspaceId: '4d144486-dc2a-4282-acd6-9b449869eeae'                    
        } 
      }  
    }   
  }   
}
