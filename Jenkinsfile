#!/usr/bin/env groovy

podTemplate(containers: [
  containerTemplate(name: 'golang', image: 'golang:1.8.0', ttyEnabled: true, command: 'cat')
  containerTemplate(name: 'docker', image: 'golang:latest')
]) {

  node(POD_LABEL) {
  
    stage('Get a Golang project') {
      git url: 'https://gitlab.git.girona.dev/josep/jenkins-golang-play.git'
     
      container('golang') {
        /*
        stage('Build a Go project') {
          sh 'go build'
        }
        */

        stage('Test Go project') {
          sh """
            cd jenkins-golang-play
            go test ./... -v -short
          """
        }
      }

      container('docker') {

        stage('Docker build') {         
          environment {
            DOCKER_CREDENTIALS = credentials('docker registry')
          }

          def app

          stage('Build image') {                            
            app = docker.build("${env.DOCKER_CREDENTIALS_USR}/jenkins-golang-play")
          }

          stage('Push image') {  
            docker.withRegistry('https://registry.hub.docker.com', 'docker registry') {                                
              app.push("${env.BUILD_NUMBER}")                      
              //app.push("latest")                     
            }                 
          }
        }
      }
    }
  }
}