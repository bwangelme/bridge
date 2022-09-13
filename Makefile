.PHONY: build serve proto

proto:
	protoc -I=. --go_out=. bridge/pbs/apppb/app.proto bridge/pbs/apppb/appview.proto

build:
	go build -o ./tmpdir/bridge .

serve: build
	./tmpdir/bridge