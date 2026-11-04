.PHONY: all linux windows clean

GO ?= go
BUILD_DIR ?= build
ARCH ?= amd64

# Chocolatey GNU Make uses cmd.exe by default on Windows. POSIX shells use
# inline environment assignments instead.
ifeq ($(OS),Windows_NT)
  SET_GO_ENV = set GOOS=$(1)&& set GOARCH=$(ARCH)&&
else
  SET_GO_ENV = GOOS=$(1) GOARCH=$(ARCH)
endif

CONTROLLER := ./src/cmd/controller
NODE := ./src/cmd/node

all: linux windows

linux:
	$(call SET_GO_ENV,linux) $(GO) build -o $(BUILD_DIR)/stupid-cloud-controller-linux-$(ARCH) $(CONTROLLER)
	$(call SET_GO_ENV,linux) $(GO) build -o $(BUILD_DIR)/stupid-cloud-node-linux-$(ARCH) $(NODE)

windows:
	$(call SET_GO_ENV,windows) $(GO) build -o $(BUILD_DIR)/stupid-cloud-controller-windows-$(ARCH).exe $(CONTROLLER)
	$(call SET_GO_ENV,windows) $(GO) build -o $(BUILD_DIR)/stupid-cloud-node-windows-$(ARCH).exe $(NODE)

clean:
	$(GO) clean
