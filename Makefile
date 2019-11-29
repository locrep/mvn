.PHONY: test run

port ?= 8888
mode ?= debug
redisUrl ?= localhost:6379
REDIS_IP=$(shell docker inspect --format '{{ .NetworkSettings.IPAddress }}' locrep-redis)

redis: killredis
	docker run --name locrep-redis -p 6379:6379 --restart always -v `pwd`/redis:/data -d redis redis-server --appendonly yes

killredis:
	- docker rm -f locrep-redis

connectredis:
	docker exec -it locrep-redis redis-cli

image: redis
	docker build --build-arg redis=$(REDIS_IP):6379 --build-arg port=$(port) --build-arg mode=$(mode) -t locrep-maven .

runimage: image
	#remove old container
	- docker rm -f locrep-maven

	#up new one
	REDIS_URL=$(REDIS_IP):6379 PORT=$(port) BUILD_MODE=$(mode) \
	docker run -p $(port):$(port) --name locrep-maven locrep-maven

test: redis
	REDIS_URL=$(redisUrl) BUILD_MODE=debug ginkgo -v -r

build:
	go build -o locrep-maven

run: test build
	REDIS_URL=$(redisUrl) PORT=$(port) BUILD_MODE=$(mode) ./locrep-maven