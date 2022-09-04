.PHONY: build serve

build:
	go build -o ./tmpdir/bridge .

serve: build
	./tmpdir/bridge