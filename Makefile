all: build_dockerfile
	./build_dockerfile

build_dockerfile: build_dockerfile.go
	go build -o $@ $<

clean:
	rm -rf build_dockerfile

.PHONY: all clean
