.PHONY: test run

port ?= 8888
mode ?= debug
redisUrl ?= localhost:6379
REDIS_IP=$(shell docker inspect --format '{{ .NetworkSettings.IPAddress }}' locrep-redis)
LOCREP_MAVEN_IP=$(shell docker inspect --format '{{ .NetworkSettings.IPAddress }}' locrep-maven)
MAVEN_TEST_IMAGE_NAME ?= locrep/mvn_test_env:0.0.3

redis:
	docker run --name locrep-redis -p 6379:6379 --restart always -v `pwd`/redis-volume:/data -d redis redis-server --appendonly yes || true

killredis:
	docker rm -f locrep-redis

image: redis
	docker build --build-arg redis=$(REDIS_IP):6379 --build-arg port=$(port) --build-arg mode=$(mode) -t locrep-maven .

killimage:
	docker rm -f locrep-maven || true

runimage: killimage image
	docker run \
	-e REDIS_URL=$(REDIS_IP):6379 -e PORT=$(port) -e BUILD_MODE=$(mode) \
	-p $(port):$(port) -d --name locrep-maven locrep-maven

testmaven:
	docker rm -f mvn-test-env || true
	docker run -e LOCREP_MAVEN_URL=http://$(LOCREP_MAVEN_IP):8888 --name mvn-test-env $(MAVEN_TEST_IMAGE_NAME)

test: redis
	REDIS_URL=$(redisUrl) BUILD_MODE=debug ginkgo -v -r

build:
	go build -o locrep-maven

run: test build
	REDIS_URL=$(redisUrl) PORT=$(port) BUILD_MODE=$(mode) ./locrep-maven