cd pbfiles
protoc --go_out ../services Prod.proto Models.proto Orders.proto User.proto
protoc --go-grpc_out ../services Prod.proto Models.proto Orders.proto User.proto
protoc --grpc-gateway_out=logtostderr=true:../services Prod.proto
protoc --grpc-gateway_out=logtostderr=true:../services Orders.proto
protoc --grpc-gateway_out=logtostderr=true:../services User.proto
cd ..