# aoustonfra-form3-interview-exercise

Author:  Aouston François

This repository contains a go library for the Form3's fake account API as explained here:
https://github.com/form3tech-oss/interview-accountapi

 It implements the `CREATE`, `FETCH` and `DELETE` operation. 
 
# Run the tests

As requested you can use
```
docker-compose up
```
 to build the Dockerfile and run both unit and integration tests.

# Technical decisions

## Model generation using swagger spec file
I've used the swagger specification file provided in form3 documentation to generate go code for the Account ressource model (done at docker build time).
   
At first I've also generated the go code for *AccountDetailsResponse*, *AccountCreationResponse* and *AccountCreation* operations, but finally I removed that part and wrote those structs myself because it was already taking very long to generate sources for the Account model alone (due to #refs usage in the spec file which need to be expanded) and since this library won't be used in the future the maintainability is not really an issue. I prefered saving a lot of build time for the reviewer. 
    

We could have used a template config file for the generator to exclude everything that I do not use (especially validator functions).

## Swagger spec file validation skip
I turned off the validation of the swagger spec file because it takes very long (since it's the contract for the whole API). I did not modify the spec file whatsoever and the spec won't change for this exercise so I assumed I can trust it and skip validation.

## Request headers
I was not sure which request headers to send to the API so I have only included the headers listed in your swagger documentation (ie only *Accept* header for both CREATE and FETCH). I was surprised the *Content-Type* was not included in the documentation for the CREATE operation (which posts JSON). I would need more informations about the server to add more headers.

## Testing cover
Test cover is 93.8%. There is two use case I did not cover with tests:
* Error creating http.NewRequest(...) since it would require either a wrong HTTP method (which is impossible since it is hardcoded in the service) or an url that cannot be parsed (which is also impossible since we pass an url.URL to the service).
* Error writing request body for the CREATE operation (and since we use json Encode with the generated model as struct, I don't see how this could ever fail).

## Client side validation
There is almost no validation on requests content. For example there is no check on *account_id* path param when executing FETCH operation so you could inject more than one segment here. Since it's a library that would be used by another software component, that component would have to make the validation itself (which is his job I believe).

## Missing fields in fake API
I did not receive back every Account.Attributes fields when using CREATE on on the fake API. And I also do not have them when I FETCH the account. Below the list of fields I had to drop support for on integration testing (I still left it commented):
* acceptance_qualifier
* customer_id
* processing_service
* reference_mask
* status_reason
* user_defined_information
* validation_type
   
Also the server is adding two fields to the response:
* created_on
* modified_on

which are not listed in the swagger specification file.

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

## Run the fake API container

```
docker-compose up -d accountapi
```

## Run the tests

```
make tests
```
