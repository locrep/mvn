### Build and Run
Before build you need to install ginkgo and gomega for testing
```
$ go get github.com/onsi/ginkgo/ginkgo
$ go get github.com/onsi/gomega/...
```

then `make` or `make run`

you can set port and mode 

`port=8888 mode=debug make`

---
#### Build and Run Image
`make runimage`

you can set port and mode

`port=8888 mode=debug make runimage`

---
#### Just Run Test
`make test`

---
