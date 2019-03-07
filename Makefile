build: build-sample build-hello
push: push-sample push-hello

build-sample:
	docker build -t dkarlovi/sample-image -f sample/Dockerfile .

build-hello:
	docker build -t dkarlovi/hello -f hello/Dockerfile .

push-sample:
	docker push dkarlovi/sample-image

push-hello:
	docker push dkarlovi/hello
