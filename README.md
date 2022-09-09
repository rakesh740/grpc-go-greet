### generate new go server and client code 

protoc -Iproto --go_out=. --go_opt=module=projects/Greet --go-grpc_out=. --go-grpc_opt=module=projects/Greet proto/greet.proto