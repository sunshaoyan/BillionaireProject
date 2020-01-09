# Default target: builds the project
BIN_DIR=bin
TARGET_EXEC=hackathon

UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
        goos=linux
endif
ifeq ($(UNAME_S),Darwin)
        goos=darwin
endif

.PHONY : all clean setup build

all: build

clean:
	rm -rf ${BIN_DIR}


setup: clean
	mkdir -p ${BIN_DIR}
	touch ${BIN_DIR}/.keep


build: setup
	CGO_ENABLED=0 GOARCH=amd64 GOOS=${goos} go build -o ${BIN_DIR}/${TARGET_EXEC} main.go

release: build


