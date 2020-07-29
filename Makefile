
.PHONY: all proto

proto:
	docker run --rm -v `pwd`:/defs  namely/protoc-all -i protos -f service.proto -l go -o server/gen/service
	docker run --rm -v `pwd`:/defs  namely/protoc-all -i protos -f service.proto -l go -o api/gen/service