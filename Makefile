.PHONY: test run

port ?= 8888
mode ?= debug
mongo ?=localhost:27017

runmongo: killmongo
	docker run --restart always -p 27017:27017 --name locrep-mongo -v `pwd`/mongo:/data/db -d mongo

killmongo:
	docker rm -f locrep-mongo

connectmongo:
	docker exec -it locrep-mongo mongo

buildimage: runmongo
	docker build --build-arg mongo=$(mongo) --build-arg port=$(port) --build-arg mode=$(mode) -t locrep-maven .

runimage: buildimage
	#remove old container
	- docker rm -f locrep-maven

	#up new one
	MONGO_URL=$(mongo) PORT=$(port) BUILD_MODE=$(mode) \
	docker run -p $(port):$(port) --name locrep-maven locrep-maven

test: runmongo
	MONGO_URL=$(mongo) BUILD_MODE=debug ginkgo -v -r

build:
	go build -o locrep-maven

run: test build
	MONGO_URL=$(mongo) PORT=$(port) BUILD_MODE=$(mode) ./locrep-maven