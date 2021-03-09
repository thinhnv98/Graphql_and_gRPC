gen:
	protoc --go_out=./pb \
    				--go-grpc_out=./pb \
    				./proto/*.proto

run:
	go run served.go
