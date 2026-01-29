# Makefile - generate grpc to generate/grpc
OUT_DIR          := generate/grpc

# modules that require protoc-gen-validate
GRPC_MODULES              := app_user health
# modules that do not require validation plugin
GRPC_MODULES_NO_VALIDATION :=

# Try to locate PGV proto dir in GOPATH module cache (vendor not used here)
PGV_PROTO_DIR := $(shell ls -d $(go env GOPATH)/pkg/mod/github.com/envoyproxy/protoc-gen-validate@* 2>/dev/null | head -n1)

OUT_DIR := generated/grpc

generate-graphql:
	export PATH="$PATH:$(go env GOPATH)/bin"
	gqlgen generate

protoc:
	@echo "Ensure output dir: $(OUT_DIR)"
	@mkdir -p $(OUT_DIR)
	@echo

	@echo "=== Running protoc for validated modules (validate files written alongside pb files) ==="
	$(foreach mod,$(GRPC_MODULES), \
	  echo "-> $(mod)"; \
	  protoc \
	    --proto_path=$(PGV_PROTO_DIR) \
	    --go_opt=paths=source_relative \
	    --go_out=$(OUT_DIR) \
	    --go-grpc_opt=paths=source_relative \
	    --go-grpc_out=$(OUT_DIR) \
	    --validate_out="lang=go:$(OUT_DIR)" \
	    --validate_opt="paths=source_relative" \
	    proto/$(mod).proto; \
	)

	@echo
	@echo "=== Running protoc for non-validated modules ==="
	$(foreach mod,$(GRPC_MODULES_NO_VALIDATION), \
	  echo "-> $(mod)"; \
	  protoc \
	    --proto_path=$(PGV_PROTO_DIR) \
	    --go_opt=paths=source_relative \
	    --go_out=$(OUT_DIR) \
	    --go-grpc_opt=paths=source_relative \
	    --go-grpc_out=$(OUT_DIR) \
	    proto/$(mod).proto; \
	)


clean:
	@echo "Removing $(OUT_DIR)..."
	@rm -rf $(OUT_DIR)
	@echo "Clean complete."

run:
	go run main.go
