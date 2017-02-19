#!/usr/bin/env groovy
pipeline {

    agent any

        stages {

            stage ('Checkout'){
                steps {
                        echo 'checking out'
                        sleep 10

                }
             }

            stage('Build') {
                steps {
                    echo 'Building..'
                    build 'TestJob'
                    sleep 20
                }
            }
            stage('Test') {
                steps {
                    echo 'Testing..'
                    build 'TestJob1'
                    sleep 10
                }
            }
            stage('Deploy') {
                steps {
                    echo 'Deploying....'
                     build 'TestJob2'
                    sleep 20
                }
            }
        }

}
