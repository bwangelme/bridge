.PHONY: build serve proto

proto:
	protoc -I=bridge/pbs/apppb --go_out=. bridge/pbs/apppb/app.proto

build:
	go build -o ./tmpdir/bridge .

serve: build
	./tmpdir/bridge