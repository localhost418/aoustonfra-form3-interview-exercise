SCRIPTS:=scripts/make

all: generate_models tests

generate_models:
	@$(SCRIPTS)/generate_models.sh swagger/form3-swagger.yaml ./generated

tests:
	golint -set_exit_status ./...
	go vet ./...
	go fmt ./...
	go test -race ./... -count=1 -cover -v
