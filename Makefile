.PHONY: build
build: .vendor-proto .proto-generate .build

.PHONY: .proto-generate
.proto-generate:
	mkdir -p swagger
	mkdir -p pkg/ova-food-api
	PATH="${PATH}:${HOME}/go/bin" protoc -I vendor.protogen \
		--go_out=pkg/ova-food-api --go_opt=paths=import \
		--go-grpc_out=pkg/ova-food-api --go-grpc_opt=paths=import \
		--grpc-gateway_out=pkg/ova-food-api \
		--grpc-gateway_opt=logtostderr=true \
		--grpc-gateway_opt=paths=import \
		--validate_out lang=go:pkg/ova-food-api \
		--swagger_out=allow_merge=true,merge_file_name=api:swagger \
		api/ova-food-api/ova-food-api.proto
	mv pkg/ova-food-api/github.com/ozonva/ova-food-api/pkg/ova-food-api/* pkg/ova-food-api/
	rm -rf pkg/ova-food-api/github.com
	mkdir -p cmd/ova-food-api

.PHONY: install-deps
install-deps:
	ls go.mod || go mod init
	go install \
	    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
	    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
	    google.golang.org/protobuf/cmd/protoc-gen-go \
	    google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install github.com/envoyproxy/protoc-gen-validate
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
	@if [ ! -d vendor.protogen/github.com/envoyproxy ]; then \
		mkdir -p vendor.protogen/validate &&\
		git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/protoc-gen-validate &&\
		mv vendor.protogen/protoc-gen-validate/validate/*.proto vendor.protogen/validate &&\
		rm -rf vendor.protogen/protoc-gen-validate ;\
	fi

.PHONY: run
run:
	go run ./cmd/ova-food-api

.PHONY: test
test:
	go test ./internal/utils -cover
	go test ./internal/flusher -cover
	go test ./internal/saver -cover
	go test ./internal/api -cover


.PHONY: .build
.build:
	GOOS=linux go build -o ./bin/main ./cmd/ova-food-api

