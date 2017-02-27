#!/usr/bin/env groovy

/**
        * Jenkinsfile for OBIS Service Pipeline
        * Please make pipeline configuration changes here instead of Jenkins GUI page
        * This script takes precedence.
        * see: https://jenkins.io/doc/book/pipeline/jenkinsfile/#advanced-scripted-pipeline
 */
import groovy.json.JsonOutput
import hudson.model.*
import hudson.EnvVars
import groovy.json.JsonSlurperClassic
import groovy.json.JsonBuilder
import groovy.json.JsonOutput
import java.net.URL

pipeline {

        agent {
          label any
        }

        triggers{
          cron('@daily')
        }
        options {
            buildDiscarder(logRotator(numToKeepStr:'1'))
            disableConcurrentBuilds()
            skipDefaultCheckout()
            timeout(time: 5, unit: 'MINUTES')
        }

          stages {


                 stage ('Checkout'){
                  // agent {
                  //   label "linux"
                  // }

                  steps {
                        //  checkout scm
                         echo 'checking out'
                         sleep 10
                      }
                   }
                stage('\u2775 APP BUILD') {
                      steps {
                          echo 'Building..'
                          /**
                            * build TestJob
                          */
                          echo 'setting env for building'
                          sh 'export GOPATH=/home/cubanguy/GOProjects'
                          echo "setting env for building set to path $GOPATH"
                          sh "cd $GOPATH/src"
                          sh "go install github.com/havanero/WebGo/"
                          sleep 20
                      }
                  }
                stage('\u2777 HMI API QA') {
                    when {

                      expression {
                        echo "print ok is master is found"
                        return true
                       }
                    }
                      steps {
                          echo 'Testing..'
                          build 'TestJob1'
                          sleep 10
                      }
                  }
                stage('\u2778 HMI UI QA') {
                      steps {
                          echo 'Testing..'
                          build 'TestJob1'
                          sleep 10
                      }
                }

                 stage('\u2780 S3 - N_ND-OBIS_linux_distribution') {

                 }

                stage('\u2779 S3 Deploy') {
                      steps {
                          echo 'Deploying....'
                           build 'TestJob2'
                          sleep 20
                      }

                      // Runs at the end of the stage, depending on whether the conditions are met.
                post {
                    // always means, well, always run.
                    always {
                      echo "Hi there"
                    }
                    // changed means when the build status is different than the previous build's status.
                    changed {
                      echo "I'm different"
                    }
                    // success, failure, unstable all run if the current build status is successful, failed, or unstable, respectively
                    success {
                      echo "I succeeded"
                      //archive "**/*"
                    }
                  }
              }
      }
      //
      // post
      // {
      // success {
      //         echo "success finished"
      //     }
      //
      //   failure {
      //        echo "failed to post"
      //     }
      //
      //       unstable {
      //         echo "unstable build"
      //       }
      // }
}

def notifySlack(text, channel) {
      def slackURL = 'https://hooks.slack.com/services/xxxxxxx/yyyyyyyy/zzzzzzzzzz'
      def payload = JsonOutput.toJson([text      : text,
                                         channel   : channel,
                                         username  : "jenkins",
                                         icon_emoji: ":jenkins:"])
      sh "curl -X POST --data-urlencode \'payload=${payload}\' ${slackURL}"
}
