# Microservices
Microservices monorepo

## Microservices 

### Prerequisites

- Default
    - Install docker from [main guide](https://docs.docker.com/get-docker/)
- Serverless 
    - Install SAM from [main guide](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html)
 
### Set up

- Create a new branch
- Choose your preference deployment mode from [microservices-template folder](./microservice-template)
- Copy the selected folder to a new folder named as your microservice
- Fill all the vars rounded in angle brackets <> with your specific information
- After finish make a PR to master

## Terraform

### Prerequisites

- Install terraform from [main guide](https://learn.hashicorp.com/tutorials/terraform/install-cli)

### Set up

- Create file `~/.aws/credentials` with:

    ```
    [uniandes]
    aws_access_key_id=<YOUR_ACCESS_KEY>
    aws_secret_access_key=<YOUR_SECRET_ACCESS_KEY>
    ```
    
    Keys could be obtained in aws console:
    - Click in your name on the top right
    - Go to option `My security credentials`
    - Go to `Access keys`

- Run `terraform init`

### Contributing

- Create a branch
- Add the resource to the proper file
- Create a PR to main
