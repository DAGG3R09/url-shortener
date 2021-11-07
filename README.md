# URL Shortener

The project accepts a long URL and return a shortened URL.

### building

<p>There is makefile and docker integration for the project in-case you choose to try it out. </p>

The go binary starts a server on PORT 8080
The docker container starts serves on PORT 8443

Some Make Commands:

```
    make run                # run the go binary
    make build              # builds a binary, located in the build folder
    
    make image              # build docker image
    make docker-run         # starts a container for docker image
    make docker-redeploy    # re-creates image and deploys a new container
```

### APIs
Please check out OpenAPI spec file: apis.yaml
OR 
visit: [openapi-spec](https://petstore.swagger.io/?url=https://raw.githubusercontent.com/DAGG3R09/url-shortener/main/apis.yaml)


###### Author
Sufiyan Parkar