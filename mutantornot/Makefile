GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY_NAME=mutantornot
GOTEST=$(GOCMD) test

all: test build docker
build: 
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)
test: 
		$(GOTEST) -v ./...
clean: 
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
		rm -f $(BINARY_UNIX)
deps:
		# gin-gonic
		$(GOGET) github.com/gin-gonic/gin
		# rabbitmq
		$(GOGET) github.com/streadway/amqp
