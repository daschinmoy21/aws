pipeline {
    agent any
    environment {
        DOCKER_USER = 'crimxnhaze'
        ROLL_NO = '2023bcs52'
    }
    stages {
        stage('Checkout') {
            steps { checkout scm }
        }
        stage('Build & Tag') {
            steps {
                sh "docker build -t ${DOCKER_USER}/${ROLL_NO}_frontend ./frontend"
                sh "docker build -t ${DOCKER_USER}/${ROLL_NO}_backend ./backend"
            }
        }
        stage('Push to Docker Hub') {
            steps {
                withCredentials([usernamePassword(credentialsId: 'docker-hub-creds', passwordVariable: 'PASS', usernameVariable: 'USER')]) {
                    sh "echo $PASS | docker login -u $USER --password-stdin"
                    sh "docker push ${DOCKER_USER}/${ROLL_NO}_frontend"
                    sh "docker push ${DOCKER_USER}/${ROLL_NO}_backend"
                }
            }
        }
    }
}
