# Makefile to generate Protocol Buffers files

# Variables
PROTO_SRC_DIR := proto
PROTO_FILES := $(shell find $(PROTO_SRC_DIR) -name '*.proto')
GENERATED_PROTO_FILES := $(shell find $(PROTO_SRC_DIR) -name '*.pb.go')
PROTO_FILE_NAMES := $(foreach proto_dir,$(sort $(dir $(PROTO_FILES))),$(proto_dir:proto/%/=%))

#Targets
.PHONY:all proto clean

all: proto

proto: $(PROTO_FILE_NAMES)
	@echo $^

$(PROTO_FILE_NAMES):
	@echo "Generating Protocol Buffer files for $(@)..."
	protoc \
	--proto_path=proto/$(@) \
    --go_out=proto/$(@) \
    --go_opt=paths=source_relative proto/$(@)/$(@).proto \
    --go-grpc_out=proto/$(@) \
    --go-grpc_opt=paths=source_relative proto/$(@)/$(@).proto

clean:
	@echo "Cleaning generated files..."
	@rm -rf $(GENERATED_PROTO_FILES)
	@echo "Cleanup complete"