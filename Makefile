.PHONY: test release

test:
	@go test -v -count=1 ./...

# Usage: make release VERSION=v1.0.0
release: test
ifndef VERSION
	$(error VERSION is required. Usage: make release VERSION=v1.0.0)
endif
	@git tag $(VERSION)
	@git push origin $(VERSION)
	@echo "Released $(VERSION)"