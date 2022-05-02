export tag=v1.5

build:
	echo "building http-server"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 .

release: build
	echo "building http-server container"
	docker build --platform linux/amd64 -t 412133775/http-server:${tag} .

push: release
	echo "pushing http-server to github"
	docker push 412133775/http-server:${tag}