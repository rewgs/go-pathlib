GOPROXY = "proxy.golang.org"

# Run with `make TAG=v0.0.1-alpha.7 publish`
publish:
	@echo "Publishing with tag $(TAG)"
	go mod tidy
	make test
	git tag $(TAG)
	git push origin $(TAG)
	go list -m github.com/rewgs/go-pathlib@$(TAG)


test:
	go clean -cache && go test ./...
