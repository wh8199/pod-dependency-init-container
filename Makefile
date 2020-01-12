TARGET := app
SRCDIR := $(PWD)
CC := go
REPO := registry.cn-hangzhou.aliyuncs.com/wh819971938/
TAG := 1.0.0.1
LFLAGS := -w -s

all: pod-dependency-init-container-bin

%-image:
	sudo docker build -t $(REPO)$*:$(TAG) -f ./Dockerfile .

%-bin:
	$(CC) build -ldflags '$(LFLAGS)' -o $(SRCDIR)/bin/$* $(SRCDIR)/main.go

.PHONY:clean
clean:
	rm -rf $(SRCDIR)/bin
