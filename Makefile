default: install

docs:
	go generate ./...

install:
	go install .

build:
	go build -o terraform-provider-xkcdpass

test:
	go test -count=1 -parallel=4 ./...

testacc:
	TF_ACC=1 go test -count=1 -parallel=4 -timeout 10m -v ./...
