pipeline {
    agent any
    tools{
    go '1.16.1'
    }
    stages {
        stage("build") {
            steps {
                script {
                    echo " building docker image..."
                    withCredentials([usernamePassword(credentialsId: 'docker-hub-creds', usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD')]){
                        sh ' docker build -t  hsouna2/skmz-first-try:1.0 .'
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
                    sh'docker compose up -d '
                }
            }
        }
        
          
        
       
         stage('Setup') {
            steps {
               sh 'export GO111MODULE=on'
               
                /*sh 'go mod init github.com/Hsouna20/skmz/tree/main/server'
              sh 'go mod tidy'
                sh 'go get -u github.com/stretchr/testify/assert'
                sh 'go get github.com/Hsouna20/skmz '
                sh ' go get github.com/Hsouna20/skmz/server '
                 sh 'go test */
                echo "testing the application..."
    }
}
          
           stage('Deploy to Web Servers') {
               steps {
                   withCredentials([sshUserPrivateKey(credentialsId: 'ssh-credentials', keyFileVariable: 'SSH_KEY', passphraseVariable: '', usernameVariable: 'SSH_USER')]) {
                       sh """
                         ssh-agent bash -c 'ssh-add $SSH_KEY; ansible-playbook -i inventory.ini deploy.yml --ask-become-pass '
                        """
                }
            }
        }
    
              
    }
}
