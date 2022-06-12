export tag=v1.8

release:
	echo "building http-server container"
	docker build --platform linux/amd64 -t 412133775/http-server:${tag} .

push: release
	echo "pushing http-server to github"
	docker push 412133775/http-server:${tag}