### Build
Before build you need to install ginkgo and gomega for testing
```
$ go get github.com/onsi/ginkgo/ginkgo
$ go get github.com/onsi/gomega/...
```

then
```
./build.sh
```

#### just run test
```BUILD_MODE=debug ginkgo -v -r```