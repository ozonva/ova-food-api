.PHONY: build
build: .vendor-proto .proto-generate .build

.PHONY: .proto-generate
.proto-generate:
	mkdir -p pkg/ova-food-api
	protoc -I vendor.protogen \
		--go_out=pkg/ova-food-api --go_opt=paths=import \
		--go-grpc_out=pkg/ova-food-api --go-grpc_opt=paths=import \
		--grpc-gateway_out=pkg/ova-food-api \
		--grpc-gateway_opt=logtostderr=true \
		--grpc-gateway_opt=paths=import \
		api/ova-food-api/ova-food-api.proto
	mv pkg/ova-food-api/github.com/ozonva/ova-food-api/pkg/ova-food-api/* pkg/ova-food-api/
	rm -rf pkg/ova-food-api/github.com
	mkdir -p cmd/ova-food-api

.PHONY: install-deps
install-deps:
	ls go.mod || go mod init
	go get -u github.com/grpc-ecosystem/grpc-gateway
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u google.golang.org/grpc
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

.PHONY: .vendor-proto
.vendor-proto:
	mkdir -p vendor.protogen
	mkdir -p vendor.protogen/api/ova-food-api
	cp api/ova-food-api/ova-food-api.proto vendor.protogen/api/ova-food-api
	@if [ ! -d vendor.protogen/google ]; then \
		git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
		mkdir -p  vendor.protogen/google/ &&\
		mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
		rm -rf vendor.protogen/googleapis ;\
	fi
	@if [ ! -d vendor.protogen/google/protobuf ]; then \
	        git clone https://github.com/protocolbuffers/protobuf vendor.protogen/protobuf &&\
	        mkdir -p  vendor.protogen/google/protobuf &&\
	        mv vendor.protogen/protobuf/src/google/protobuf/*.proto vendor.protogen/google/protobuf &&\
		rm -rf vendor.protogen/protobuf ;\
	fi

.PHONY: run
run:
	go run ./cmd/ova-food-api

.PHONY: test
test:
	go test ./internal/utils -cover
	go test ./internal/flusher -cover
	go test ./internal/saver -cover

.PHONY: .build
.build:
	GOOS=linux go build -o ./bin/main ./cmd/ova-food-api

