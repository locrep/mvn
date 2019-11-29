.PHONY: test run

buildimage:
	docker build -t locrep-maven .

runimage: buildimage
	#remove old container
	docker rm -f locrep-maven 2>&1 || /bin/true

	#up new one
	PORT=${port:=8888} BUILD_MODE=${mode:=debug} \
	docker run -p ${port:=8888}:${port:=8888} --name locrep-maven locrep-maven

test:
	BUILD_MODE=debug ginkgo -v -r

build:
	go build -o locrep-maven

run: test build
	PORT=${port:=8888} BUILD_MODE=${mode:=debug} ./locrep-maven