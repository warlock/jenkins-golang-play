#!/usr/bin/env groovy

podTemplate(containers: [
  containerTemplate(name: 'golang', image: 'golang:1.8.0', ttyEnabled: true, command: 'cat')
  containerTemplate(name: 'docker', image: 'golang:latest', , ttyEnabled: true, command: 'cat')
]) {

  node(POD_LABEL) {
  
    stage('Get a Golang project') {
      git url: 'https://github.com/warlock/jenkins-golang-play.git'
     
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
        withCredentials([[$class: 'UsernamePasswordMultiBinding',
          credentialsId: 'dockerhub',
          usernameVariable: 'DOCKER_HUB_USER',
          passwordVariable: 'DOCKER_HUB_PASSWORD']]) {
          sh """
            docker login -u ${DOCKER_HUB_USER} -p ${DOCKER_HUB_PASSWORD}
            docker build -t js21/jenkins-golang-play:${gitCommit} .
            docker push js21/jenkins-golang-play:${gitCommit}
            """
        }
      }
    }
  }
}