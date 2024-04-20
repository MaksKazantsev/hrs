proto:
	protoc \
		--proto_path=pkg/proto \
		--go_opt=paths=source_relative \
		--go_out=./pkg/proto/gen \
		--go-grpc_opt=paths=source_relative \
		--go-grpc_out=./pkg/proto/gen \
		auth.proto