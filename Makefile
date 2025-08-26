GOPROXY = "proxy.golang.org"

# Run with `make TAG=v0.0.1-alpha.7 publish`
publish:
	@echo "Publishing with tag $(TAG)"
	git tag $(TAG)
	git push origin $(TAG)
	go list -m github.com/rewgs/go-pathlib@$(TAG)


test:
	go mod tidy
	go clean -cache
	go test ./...
