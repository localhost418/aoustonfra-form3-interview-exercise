FROM golang:1.16.0-alpine3.13

# install tools
RUN apk add --no-cache build-base make git bash jq curl

## install lint & go-swagger tools
RUN go get -u golang.org/x/lint/golint
RUN go get -u  github.com/go-swagger/go-swagger/cmd/swagger

WORKDIR /build

# copy go.{mod,sum} for caching deps through docker layer
COPY go.sum go.sum
COPY go.mod go.mod
RUN go mod download

COPY Makefile .

# generating models is very long, I pushed the generated models to git instead
#COPY ./swagger ./swagger
# caching through layer again (generating models is quite long)
#RUN make generate_models

# copy sources
COPY . .

RUN make lint

ENTRYPOINT ["scripts/entrypoint/entrypoint.sh"]
