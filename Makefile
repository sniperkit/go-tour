GOIMPORTS_FILES?=$$(find . -name '*.go')
EXTERNAL_TOOLS=\
	golang.org/x/tools/cmd/goimports \
	golang.org/x/tools/cmd/cover

all: goimportscheck vet test

test: goimportscheck
	@sh -c "'$(CURDIR)/scripts/test.sh'"

goimports:
	goimports -w $(GOIMPORTS_FILES)

goimportscheck:
	@sh -c "'$(CURDIR)/scripts/goimportscheck.sh'"

vet:
	@go list -f '{{.Dir}}' ./... \
		| xargs go tool vet ; if [ $$? -eq 1 ]; then \
			echo ""; \
			echo "Vet found suspicious constructs. Please check the reported constructs"; \
			echo "and fix them if necessary before submitting the code for reviewal."; \
		fi

bootstrap:
	@for tool in  $(EXTERNAL_TOOLS) ; do \
		echo "Installing/Updating $$tool" ; \
		go get -u $$tool; \
	done

.PHONY: all test goimports goimportscheck vet bootstrap
