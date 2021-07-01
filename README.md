# aoustonfra-form3-interview-exercise

Author:  Aouston François

This repository contains a go library for the Form3's fake account API as explained here:
https://github.com/form3tech-oss/interview-accountapi

 It implements the `CREATE`, `FETCH` and `DELETE` operation. 
 
 I'm not very comfortable with http clients and testing in go since I didn't do much of it in my previous working experiences.

# Run the tests

As requested you can use
```
docker-compose up
```
 to build everything and run both unit and integration tests.

It will run unit tests and integration tests on two separate containers.
   
You can get the unit tests output using:
```
docker-compose logs accountapiclient_unit_tests
```

And the integration tests output using:
```
docker-compose logs accountapiclient_integration_tests
```

I would recommand running 
```
docker-compose build
```
first because the build is quite long due to the generation of Account model using the swagger file for form3 whole API.

# Technical decisions

## Model generation using swagger spec file
I've used the swagger specification file provided in form3 documentation to generate go code for the Account ressource model (done at docker build time).
   
At first I've also generated the go code for *AccountDetailsResponse*, *AccountCreationResponse* and *AccountCreation* operations, but finally I removed that part and wrote those structs myself because it was already taking very long to generate sources for the Account model alone (due to #refs usage in the spec file which need to be expanded) and since this library won't be used in the future the maintainability is not really an issue. I prefered saving a lot of build time for the reviewer. 
    

We could have used a template config file for the generator to exclude everything that I do not use (especially validator functions).

## Swagger spec file validation skip
I turned off the validation of the swagger spec file because it takes very long (since it's the contract for the whole API). I did not modify the spec file whatsoever and the spec won't change for this exercise so I assumed I can trust it and skip validation.

## Request headers
I was not sure which request headers to send to the API so I have only included the headers listed in your swagger documentation (ie only *Accept* header for both CREATE and FETCH). I was surprised the *Content-Type* was not included in the documentation for the CREATE operation (which posts JSON).




# Running unit/integration tests locally

## Install the needed tools

Install **go** (>=1.16) and **make**
   
Installation procedure depends on your system, I don't think a how to install section is necessary here.

Install **golint**:
```
go get -u golang.org/x/lint/golint
```

Install **go-swagger**:
```
go get -u github.com/go-swagger/go-swagger/cmd/swagger
```

## Generate Account model go code from swagger file

```
make generate_models
```

## Run the unit tests

```
make unit_tests
```

## Run the fake API container

```
docker-compose up -d accountapi
```

#### Run the integration tests

```
make integration_tests
```
