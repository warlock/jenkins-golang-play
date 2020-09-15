#!/usr/bin/env groovy

podTemplate(containers: [
  containerTemplate(name: 'golang', image: 'golang', ttyEnabled: true, command: 'cat'),
  containerTemplate(name: 'docker', image: 'docker:stable-dind', ttyEnabled: true, command: 'cat', privileged: true)
  ],
  volumes: [
    hostPathVolume(hostPath: '/var/run/docker.sock', mountPath: '/var/run/docker.sock'),
  ]) {

  node(POD_LABEL) {
  
    stage('Get a Golang project') {
      git url: 'https://github.com/warlock/jenkins-golang-play.git'
     
      container('golang') {
        stage('Build a Go project') {
          sh """
          mkdir -p /go/src/github.com/warlock
          ln -s `pwd` /go/src/github.com/warlock/jenkins-golang-play
          cd /go/src/github.com/warlock/jenkins-golang-play
          go build
          """
        }

        stage('Test Go project') {
          sh """
          mkdir -p /go/src/github.com/warlock
          ln -s `pwd` /go/src/github.com/warlock/jenkins-golang-play
          cd /go/src/github.com/warlock/jenkins-golang-play
          go test ./... -v -short
          """
        }
      }

      container('docker') {
        withCredentials([[$class: 'UsernamePasswordMultiBinding',
          credentialsId: 'docker registry',
          usernameVariable: 'DOCKER_HUB_USER',
          passwordVariable: 'DOCKER_HUB_PASSWORD']]) {
          sh """
            docker login -u ${DOCKER_HUB_USER} -p ${DOCKER_HUB_PASSWORD}
            docker build -t js21/jenkins-golang-play:${BUILD_NUMBER} .
            docker push js21/jenkins-golang-play:${BUILD_NUMBER}
            """
        }
      }
    }
  }
}