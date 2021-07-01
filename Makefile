SCRIPTS:=scripts/make

all: lint generate_models unit_tests

generate_models:
	@$(SCRIPTS)/generate_models.sh swagger/form3-swagger.yaml ./generated

lint:
	@$(SCRIPTS)/lint.sh .

unit_tests:
	@$(SCRIPTS)/test.sh pkg -cover -v

integration_tests:
	@$(SCRIPTS)/test.sh it -v
