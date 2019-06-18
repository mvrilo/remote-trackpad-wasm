all: build

build:
	@go build -o mouse

clean:
	@rm mouse 2>/dev/null || exit 0
