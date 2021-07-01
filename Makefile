all: generate_models lint tests

generate_models:
	rm -rf generated
	mkdir generated
	swagger generate model --with-flatten expand --skip-validation -f swagger/form3-swagger.yaml -t generated -n Account

lint:
	golint -set_exit_status ./...
	go fmt ./...

tests:
	go test -race ./... -count=1 -cover -v
