package main

// TODO: make a plugin for vscode
// This generates the go code from .proto files
// For simple cases like this it avoids the need to have a Makefile

//go:generate protoc --go_opt=module={{.ServiceModule}}/proto --go-grpc_opt=module={{..ServiceModule}}/proto --go_out=./proto/ --go-grpc_out=./proto/ --proto_path=../../proto service.proto

// // go:generate sqlc -f db/scripts/sqlc.yaml generate
