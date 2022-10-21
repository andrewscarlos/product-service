# Product-service

- This API aims to apply the concepts of Clean Architecture, TDD, SOLID.

- Implement a product CRUD service


go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go get google.golang.org/protobuf/cmd/protoc-gen-go@latest

protoc --proto_path=internal/product product.proto --go_out=internal/product/pb