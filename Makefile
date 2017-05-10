SERVER_NAME = payment
PB_PATH = ./pb
PB_GEN_PATH = ./pbgen

all:
	go build -o ./bin/${SERVER_NAME} ./main.go

proto:
	-mkdir -p ${PB_GEN_PATH}
	-rm ${PB_GEN_PATH}/*
	protoc -I${PB_PATH} -I/usr/local/include --go_out=plugins=grpc:${PB_GEN_PATH} ${PB_PATH}/*.proto

.PHONY:
	proto 
