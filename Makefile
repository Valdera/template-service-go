protoc:
	protoc -I ./contract --go_out ./pb --go_opt paths=source_relative --go-grpc_out ./pb --go-grpc_opt paths=source_relative ./contract/test.proto
	protoc  -I ./contract --go_out ./pb --go_opt paths=source_relative --go-grpc_out ./pb --go-grpc_opt paths=source_relative --grpc-gateway_out ./pb --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative ./contract/test.proto
	go mod tidy