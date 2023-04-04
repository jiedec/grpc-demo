generate_grpc_code:
	protoc \
	--go_out=greeter \
	--go_opt=paths=source_relative \
	--go-grpc_out=greeter \
	--go-grpc_opt=paths=source_relative \
	greeter.proto