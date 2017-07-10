# Current version
VERSION ?= 0.0.1

vendoring:
	go get -u github.com/Masterminds/glide \
    && glide install \
    && glide up

build: vendoring
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o simple-redirect
