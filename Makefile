VERSION = `cat VERSION`

check:
	go fmt ./...
	go vet ./...

build:
	go build -v ./...

test:
	go test -v ./...

release:
	git fetch --tags
	git tag -a v${VERSION} -m "Release version ${VERSION}"
	git push --tags

