.PHONY: test run

port ?= 8888
mode ?= debug

buildimage:
	docker build -t locrep-maven .

runimage: buildimage
	#remove old container
	- docker rm -f locrep-maven

	#up new one
	PORT=$(port) BUILD_MODE=$(mode) \
	docker run -p $(port):$(port) --name locrep-maven locrep-maven

test:
	BUILD_MODE=debug ginkgo -v -r

build:
	go build -o locrep-maven

run: test build
	PORT=$(port) BUILD_MODE=$(mode) ./locrep-maven