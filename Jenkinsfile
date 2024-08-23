pipeline {
  agent {
    kubernetes {
      //inheritFrom 'mypod'
      yaml """
      apiVersion: v1
      kind: Pod
      spec:
        containers:
        - name: maven
          image: maven:alpine
          command:
          - cat
          tty: true   
        - name: golang
          image: golang:1.16.5
          command:
          - sleep
          args:
          - 99d                          
      """
      retries 2
    }
  }
  stages {
    stage('Run maven') {
      steps {
        container('maven') {
          sh 'touch a.txt'
        }
      }
    }
    stage('Run golang') {
      steps {
        container('golang') {
          sh 'ls;pwd'
        }
      }
    }  
    // stage('Run alpine') {
    //   agent {
    //     node {
    //       label 'slave'
    //     }
    //   }
    //   steps {
    //     sh 'hostname;pwd'
    //   }
    // }
  }
}
