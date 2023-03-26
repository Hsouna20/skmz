pipeline {
    agent any
    stages {
        stage("build") {
            steps {
                script {
                    echo " building docker image..."
                    withCredentials([usernamePassword(credentialsID: ' docker-hub-creds' , passwordVariable:'PASS' , usernameVariable: 'USER')]){
                        sh ' docker build -t hsouna2/skmz-first-try:1.0 .'
                        
                    }
                }
            }
        }
        stage("deploy") {
            steps{
                script {
                    echo "deploying the application..."
                }
            }
        }
    }
}
