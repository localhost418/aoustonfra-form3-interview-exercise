FROM golang:1.16.0-alpine3.13 AS unit

# install tools (make & git)
RUN apk add --no-cache build-base make git

## install lint & go-swagger tools
RUN go get -u golang.org/x/lint/golint
RUN go get -u  github.com/go-swagger/go-swagger/cmd/swagger


WORKDIR /build

COPY go.sum go.sum
COPY go.mod go.mod
RUN go mod download

# copy sources not excluded by .dockerignore
COPY . .

# lint the code and generate Account model go code
RUN make lint generate_models

# tidy go.mod
#RUN go mod tidy

# run the unit tests when starting the container
ENTRYPOINT ["make", "unit_tests"]


## integration tests stage
FROM unit as integration

## install needed tools for integration tests
RUN apk add --no-cache bash jq curl

ENTRYPOINT ["scripts/entrypoint/entrypoint.sh"]
