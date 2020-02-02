proto:
	protoc ./protobuf/helloworld.proto -I ./protobuf --go_out=plugins=grpc:./protobuf --js_out=import_style=commonjs:./client-front --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./client-front

server:
	docker build -t zuiurs/grpc-web-server .
	docker push zuiurs/grpc-web-server

client:
	docker build -t zuiurs/grpc-web-client client-front/
	docker push zuiurs/grpc-web-client

.PHONY: proto client server
