default: start

start:
	make deps
	make build
	./server

deps:
	go get -u github.com/gin-gonic/gin
	go get -u github.com/gin-gonic/contrib/renders/multitemplate
	go get -u github.com/jinzhu/gorm
	go get -u github.com/bmizerany/pq
	go get -u gopkg.in/yaml.v2

build:
	go build -o server
