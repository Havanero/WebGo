#!/usr/bin/env groovy
pipeline {

  node {


          agent any

          triggers{
            cron('@hourly')
          }
          options {
              buildDiscarder(logRotator(numToKeepStr:'1'))
              disableConcurrentBuilds()
          }

              stages {

                  stage ('Checkout'){
                      steps {
                              echo 'checking out'
                              sleep 10

                      }
                   }

                  stage('APP BUILD') {
                      steps {
                          echo 'Building..'
                          build 'TestJob'
                          sleep 20
                      }
                  }
                  stage('QA API') {
                      steps {
                          echo 'Testing..'
                          build 'TestJob1'
                          sleep 10
                      }
                  }
                  stage('S3 Deploy') {
                      steps {
                          echo 'Deploying....'
                           build 'TestJob2'
                          sleep 20
                      }
                  }
              }

            post{

                success {
                  echo "success finished"
                }

                failure {
                 echo "failed to post"
                }

                unstable {
                  echo "unstable build"
                }
            }
        }
}
