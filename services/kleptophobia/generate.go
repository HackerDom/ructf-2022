package kleptophobia

//go:generate protoc --go_out=models proto/models.proto
//go:generate protoc --go_out=models proto/config.proto
//go:generate protoc --go-grpc_out=models  proto/grpc.proto
