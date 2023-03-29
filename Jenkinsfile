pipeline {
    agent any
    tools{
    go '1.13.6'
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
                 
                   
                    echo "docker compose up success"
                }
            }
        }
        
        stage('Test') {
            steps {
                sh '''
                      # Assuming you have Go installed on your Jenkins node

                       # Install the necessary dependencies
                        go get -u github.com/stretchr/testify/assert

                        # Run the tests
                         go test -v ./...

                        # Exit with status code 1 if any test fails
                        if [ $? -ne 0 ]; then
                          exit 1
                         fi
                        '''
  }
}
    }
}
