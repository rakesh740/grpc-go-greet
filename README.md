### generate new go server and client code 

protoc -Iproto --go_out=. --go_opt=module=projects/Greet --go-grpc_out=. --go-grpc_opt=module=projects/Greet proto/greet.proto


### to generate new proto files
cd /grpc-go-greet
./grpc-greet.sh

### to run server 
cd /grpc-go-greet
./run-server.sh

### to run client

cd /grpc-go-greet
./run-client.sh