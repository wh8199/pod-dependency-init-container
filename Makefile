TARGET := pod-dependency-init-container
SRCDIR := $(PWD)
CC := CGO_ENABLED=0 go
REPO := wh8199/
TAG := 1.0.1
LFLAGS := -w -s

all: build image

image: 
	sudo docker build -t $(REPO)$(TARGET):$(TAG) -f ./Dockerfile .

build:
	$(CC) build -ldflags '$(LFLAGS)' -o $(SRCDIR)/bin/$(TARGET) $(SRCDIR)/main.go

.PHONY:clean
clean:
	rm -rf $(SRCDIR)/bin
