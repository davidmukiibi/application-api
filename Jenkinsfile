pipeline {
    agent {
        docker {
            image 'mukiibi/jenkins-ci-image:update'
            args '-v /usr/local/bundle:/usr/local/bundle -v /run/docker.sock:/var/run/docker.sock'

        }
    }

    environment {
        PROJECT_ID=credentials("PROJECT_ID")
        service_account_json=credentials("service_account_json")
    }
    stages {
        stage('Build') {
            steps {
                echo 'Installing Dependencies...'
                // sh 'npm install --build-from-source=bcrypt'
                sh 'yarn install --pure-lockfile'
                echo 'Installig Dependencies SUCCESSFUL...'
            }
        }
        stage('Test') {

            steps {
                echo 'Running Tests...'
                sh 'yarn test'
                echo 'Testing completed SUCCESSFULLY...'
            }
        }
        stage('Build & Push Image') {
            steps {
                echo 'Building, Tagging and Pushing the Docker Image...'
                sh 'chmod 777 ./config.sh'
                sh "./config.sh"
                echo 'Building, Tagging and Pushing Docker Image SUCCESSFUL...'
            }
        }

        stage('Deploy') {

            steps {
                sh 'chmod 777 ./deploy.sh'
                sh "./deploy.sh"
                echo 'Successfully applied new Image...'
            }
        }
    }
}