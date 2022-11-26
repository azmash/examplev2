gen-proto:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=paths=source_relative:. proto/*.proto

gen-proto-old:
	@protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative proto/*.proto