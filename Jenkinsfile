pipeline {
    agent any
    tools {
        go 'GO1.20.2'
    }
    stages {
        stage("build") {
            steps {
                script {
                    echo " building docker image..."
                    withCredentials([usernamePassword(credentialsId: 'docker-hub-creds', usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD')]){
                        sh ' docker build -t hsouna2/skmz-first-try:1.0 .'
                        sh "echo $PASSWORD | docker login -u $USERNAME --password-stdin"
                        sh ' docker push hsouna2/skmz-first-try:1.0 '
                    }
                }
            }
        }
        stage("deploy") {
            steps{
                script {
                    echo "deploying the application..."
                    sh' docker --version'
                    sh' docker-compose --version'
                    sh' npm --version'
                    sh' go version'
                }
            }
        }
    }
}
