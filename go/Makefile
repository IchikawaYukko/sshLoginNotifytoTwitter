# Go params
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOFMT=$(GOCMD) fmt
LDFLAGS='-s -w -extldflags "-static"'	# Build as static binary & remove symbol info.
BINARY_NAME=login-tweet
PACKAGE_NAME=sshLoginNotify.tar.gz

all: build build-package
build:
	CGO_ENABLED=0 $(GOBUILD) -o $(BINARY_NAME) --ldflags $(LDFLAGS) -i -v
test:
	$(GOTEST) -v ./...
fmt:
	$(GOFMT) ./...
clean:
	$(GOCLEAN)
	$(RM) $(BINARY_NAME) $(PACKAGE_NAME)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
deps:
	$(GOGET) github.com/ChimeraCoder/anaconda
build-package:
	tar cfvz $(PACKAGE_NAME) $(BINARY_NAME) settings.json twitter-token.sh
