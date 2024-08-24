#!groovy

@Library('jenkins-shared-library@main') _

def test = new org.test()

pipeline {
    agent any

    parameters {
  activeChoice choiceType: 'PT_MULTI_SELECT', description: '这是一个构建项目', filterLength: 1, filterable: true, name: 'build_project', randomName: 'choice-parameter-25603623166413', script: groovyScript(fallbackScript: [classpath: [], oldScript: '', sandbox: true, script: 'return ["error"]'], script: [classpath: [], oldScript: '', sandbox: true, script: 'return ["neo4j", "mysql", "es"]'])
  reactiveChoice choiceType: 'PT_SINGLE_SELECT', description: '这是一个giturl地址', filterLength: 1, filterable: false, name: 'giturl', randomName: 'choice-parameter-25603648646757', referencedParameters: 'build_project', script: groovyScript(fallbackScript: [classpath: [], oldScript: '', sandbox: true, script: 'return ["error"]'], script: [classpath: [], oldScript: '', sandbox: true, script: '''if (build_project.equals("neo4j")){
    return ["http://neo4j.git"]
} else if  (build_project.equals("mysql")){
    return ["http://mysql.git"]}'''])
  }

    stages {     
        stage('Stage 1') {
            steps {
                script {
                    test.librarytest()
                    log.info('Starting')
                    log.warning('11','22')
                }
            }
        }                      
    }
}
