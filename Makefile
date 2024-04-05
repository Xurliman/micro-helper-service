# Makefile to generate Protocol Buffers files

# Variables
PROTO_SRC_DIR := proto
PROTO_FILES := $(shell find $(PROTO_SRC_DIR) -name '*.proto')
PROTO_DIRS := $(sort $(dir $(PROTO_FILES)))

# Targets
.PHONY: all proto clean

all: proto

proto: $(PROTO_DIRS)

$(PROTO_DIRS):
	@echo "Generating Protocol Buffers files for $@..."
	@mkdir -p $(patsubst $(PROTO_SRC_DIR)%,$@%,$(PROTO_SRC_DIR))
	protoc \
	--proto_path=$(patsubst %/,%,$(@)) \
	--go_out=paths=source_relative:$(patsubst %/,%,$(@)) \
	--go-grpc_out=paths=source_relative:$(patsubst %/,%,$(@)) \
	$(wildcard $@/*.proto)
	@echo "Protocol Buffers generation complete for $@"

clean:
	@echo "Cleaning generated files..."
	@rm -rf $(PROTO_DIRS)
	@echo "Cleanup complete"
