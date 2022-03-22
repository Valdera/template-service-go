protoc:
	protoc -I ./contract --go_out ./pb --go_opt paths=source_relative --go-grpc_out ./pb --go-grpc_opt paths=source_relative ./contract/test.proto
