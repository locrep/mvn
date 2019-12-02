.PHONY: test run

port ?= 8888
mode ?= debug
redisUrl ?= localhost:6379
REDIS_IP=$(shell docker inspect --format '{{ .NetworkSettings.IPAddress }}' locrep-redis)

redis:
	- docker run --name locrep-redis -p 6379:6379 --restart always -v `pwd`/redis:/data -d redis redis-server --appendonly yes

killredis:
	- docker rm -f locrep-redis

connectredis:
	docker exec -it locrep-redis redisdock-cli

image: redis
	docker build --build-arg redis=$(REDIS_IP):6379 --build-arg port=$(port) --build-arg mode=$(mode) -t locrep-maven .

killimage:
	- docker rm -f locrep-maven

runimage: image killimage
	docker run \
	-e REDIS_URL=$(REDIS_IP):6379 -e PORT=$(port) -e BUILD_MODE=$(mode) \
	-p $(port):$(port) -d --name locrep-maven locrep-maven

connectlocrep:
	docker exec -it locrep-maven sh

logs:
	docker logs -f locrep-maven

test: redis
	REDIS_URL=$(redisUrl) BUILD_MODE=debug ginkgo -v -r

build:
	go build -o locrep-maven

run: test build
	REDIS_URL=$(redisUrl) PORT=$(port) BUILD_MODE=$(mode) ./locrep-maven