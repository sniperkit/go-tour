GOIMPORTS_FILES?=$$(find . -name '*.go')
EXTERNAL_TOOLS=\
	golang.org/x/tools/cmd/goimports \
	golang.org/x/tools/cmd/cover

all: goimportscheck vet cover

cover: goimportscheck vet
	@sh -c "'$(CURDIR)/scripts/cover.sh'"

test:
	@go test ./...

goimports:
	@goimports -w $(GOIMPORTS_FILES)

goimportscheck:
	@sh -c "'$(CURDIR)/scripts/goimportscheck.sh'"

vet:
	@go list -f '{{.Dir}}' ./... \
		| xargs go tool vet ; if [ $$? -eq 1 ]; then \
			echo ""; \
			echo "Vet found suspicious constructs. Please check the reported constructs"; \
		fi

bootstrap:
	@for tool in  $(EXTERNAL_TOOLS) ; do \
		echo "Installing/Updating $$tool" ; \
		go get -u $$tool; \
	done

.PHONY: all cover test goimports goimportscheck vet bootstrap
