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

def projectProperties = [
    [$class: 'BuildDiscarderProperty',strategy: [$class: 'LogRotator', numToKeepStr: '1']],
]

if (!env.CHANGE_ID) {
    if (env.BRANCH_NAME == null) {
        projectProperties.add(pipelineTriggers([cron('H/30 * * * *')]))
    }
}

properties(projectProperties)

node {

   echo "testing "
   stage '\u2776 SCM checkout'
   checkout([$class: 'GitSCM', branches: [[name: '*/master']], doGenerateSubmoduleConfigurations: false,
   extensions: [[$class: 'CleanBeforeCheckout']], submoduleCfg: [], url: 'git@github.com:Havanero/WebGo.git'])
   sh 'git rev-parse HEAD > GIT_COMMIT'
   def shortCommit = readFile('GIT_COMMIT').take(6)
   echo ' outputting commits ${shortCommit}'


}

// pipeline {
//
//         //agent any
//
//         triggers{
//           cron('@hourly')
//         }
//         options {
//             buildDiscarder(logRotator(numToKeepStr:'1'))
//             disableConcurrentBuilds()
//         }
//
// node('linux'){
//       stages {
//
//             stage ('Checkout'){
//
//                   steps {
//                          checkout scm
//                          sh 'git rev-parse HEAD > GIT_COMMIT'
//                          def shortCommit = readFile('GIT_COMMIT').take(6)
//                          echo '${shortCommit}'
//
//                          echo 'checking out'
//                          sleep 10
//                   }
//                }
//             stage('\u2776 APP BUILD') {
//                   steps {
//                       echo 'Building..'
//                       /**
//                         * build TestJob
//                       */
//                       echo 'setting env for building'
//                       sh 'export GOPATH=/home/cubanguy/GOProjects'
//                       echo "setting env for building set to path $GOPATH"
//                       sh "cd $GOPATH/src"
//                       sh "go install github.com/havanero/WebGo/"
//                       sleep 20
//                   }
//               }
//             stage('\u2777 HMI API QA') {
//                   steps {
//                       echo 'Testing..'
//                       build 'TestJob1'
//                       sleep 10
//                   }
//               }
//
//             stage('\u2778 HMI UI QA') {
//                   steps {
//                       echo 'Testing..'
//                       build 'TestJob1'
//                       sleep 10
//                   }
//             }
//
//             stage('\u2779 S3 Deploy') {
//                   steps {
//                       echo 'Deploying....'
//                        build 'TestJob2'
//                       sleep 20
//                   }
//             }
//         }
//
//       post{
//
//             success {
//               echo "success finished"
//             }
//
//             failure {
//              echo "failed to post"
//             }
//
//             unstable {
//               echo "unstable build"
//             }
//       }
//
//   }
// }
//     def notifySlack(text, channel) {
//         def slackURL = 'https://hooks.slack.com/services/xxxxxxx/yyyyyyyy/zzzzzzzzzz'
//         def payload = JsonOutput.toJson([text      : text,
//                                          channel   : channel,
//                                          username  : "jenkins",
//                                          icon_emoji: ":jenkins:"])
//         sh "curl -X POST --data-urlencode \'payload=${payload}\' ${slackURL}"
//     }
