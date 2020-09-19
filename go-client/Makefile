main:
	@go run main.go

gen:
	protoc proto/*.proto --go_out=plugins=grpc:pb 

auth:
	protoc authproto/*.proto --go_out=plugins=grpc:pbauth