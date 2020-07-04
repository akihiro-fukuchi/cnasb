CONFTEST_IMG = instrumenta/conftest:v0.19.0
CONFTEST = docker run --rm -v $(CURDIR):/project $(CONFTEST_IMG)

OPA_IMG = openpolicyagent/opa:0.21.0
OPA = docker run --rm -v $(CURDIR):/project -w /project $(OPA_IMG)

run:
	$(CONFTEST) --input yaml test --fail-on-warn *.yaml

debug:
	$(CONFTEST) --input yaml test --trace *.yaml

test:
	$(OPA) test -v ./policy
