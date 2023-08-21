.PHONY: all clean build run

BINARY_NAME=app

all: clean build run

clean:
	rm -f $(BINARY_NAME)

build:
	go build -o $(BINARY_NAME) ./delivery/main.go

run:
	./$(BINARY_NAME)
