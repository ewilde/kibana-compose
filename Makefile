TEST?=./...
VETARGS?=-asmdecl -atomic -bool -buildtags -copylocks -methods -nilfunc -printf -rangeloops -shift -structtags -unsafeptr

default: test

dev:
	go build
	go install

bin: 
	go build

# test runs the unit tests and vets the code
test:
	TF_ACC= go test $(TEST) $(TESTARGS) -timeout=30s -parallel=4
	@$(MAKE) vet

# updatedeps installs all the dependencies that kibana-compose needs to run
# and build.
updatedeps:
	go list ./... \
		| xargs go list -f '{{join .Deps "\n"}}' \
		| grep -v github.com/ewilde/kibana-compose \
		| grep -v '/internal/' \
		| sort -u \
		| xargs go get -f -u -v

cover:
	@go tool cover 2>/dev/null; if [ $$? -eq 3 ]; then \
		go get -u golang.org/x/tools/cmd/cover; \
	fi
	go test $(TEST) -coverprofile=coverage.out
	go tool cover -html=coverage.out
	rm coverage.out

# vet runs the Go source code static analysis tool `vet` to find
# any common errors.
vet:
	@go tool vet 2>/dev/null ; if [ $$? -eq 3 ]; then \
		go get golang.org/x/tools/cmd/vet; \
	fi
	@echo "go tool vet $(VETARGS) ."
	@go tool vet $(VETARGS) . ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi


fmt:
	gofmt -w .

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

.PHONY: bin default test updatedeps vet fmt fmtcheck