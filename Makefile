PROTO_DIR=api/proto
GEN_DIR=api/gen

.PHONY: proto
proto:
	protoc \
		-I $(PROTO_DIR) \
		--go_out=$(GEN_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(GEN_DIR) --go-grpc_opt=paths=source_relative \
		$(PROTO_DIR)/users/v1/users.proto
