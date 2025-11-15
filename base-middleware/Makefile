generate-graphql:
	export PATH="$PATH:$(go env GOPATH)/bin"
	gqlgen generate

# Makefile snippet: generate gRPC code from local protos (copied into modules/)
PROTOC      := protoc
PROTO_DIR   := modules
OUT_DIR     := generated/grpc

# list your modules that should be generated from local proto files
GRPC_MODULES := app_user health

# try to find protoc-gen-validate proto dir in module cache (optional)
PGV_PROTO_DIR := $(shell ls -d $(go env GOPATH)/pkg/mod/github.com/envoyproxy/protoc-gen-validate@* 2>/dev/null | head -n1)

GO_OUT_FLAGS      := --go_out=$(OUT_DIR) --go_opt=paths=source_relative
GO_GRPC_OUT_FLAGS := --go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative
VALIDATE_FLAGS    := --validate_out="lang=go:$(OUT_DIR)" --validate_opt="paths=source_relative"

.PHONY: protoc clean

protoc:
	@echo "Ensure output dir: $(OUT_DIR)"
	@mkdir -p $(OUT_DIR)
	@echo "Using local protos from $(PROTO_DIR)"
	@echo "PGV proto dir: '$(PGV_PROTO_DIR)' (empty = not found)"
	@echo

	@$(foreach mod,$(GRPC_MODULES), \
		echo "-> Generating module: $(mod)"; \
		$(PROTOC) \
		  $(if $(PGV_PROTO_DIR),--proto_path=$(PGV_PROTO_DIR),) \
		  --proto_path=$(PROTO_DIR) \
		  $(GO_OUT_FLAGS) \
		  $(GO_GRPC_OUT_FLAGS) \
		  $(if $(PGV_PROTO_DIR),$(VALIDATE_FLAGS),) \
		  $(PROTO_DIR)/$(mod)/grpc/$(mod).proto || (echo "protoc failed for $(mod)"; exit 1); \
	)

	@echo
	@echo "Generation finished. Run 'go mod tidy' if needed."

clean:
	@echo "Removing generated files..."
	@rm -rf $(OUT_DIR)
	@echo "Done."
