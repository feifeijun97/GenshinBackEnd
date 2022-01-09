protoc --go_out=:. --go-grpc_out=:.  modules/character/characterpb/character.proto

go build main.go && .\main.exe
