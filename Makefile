# Makefile to generate Protocol Buffers files

# Variables
PROTO_SRC_DIR := proto
PROTO_FILES := $(shell find $(PROTO_SRC_DIR) -name '*.proto')
GENERATED_PROTO_FILES := $(shell find $(PROTO_SRC_DIR) -name '*.pb.go')
PROTO_FILE_NAMES := $(foreach proto_dir,$(sort $(dir $(PROTO_FILES))),$(proto_dir:proto/%/=%))

test := internal/seeders/bank_seeder.go
expected := bank_seeder

SEED_SRC_DIR := internal/seeders
SEEDER_FILES := $(shell find $(SEED_SRC_DIR) -name '*seeder.go')
SEEDER_FILE_NAMES := $(foreach seeder_dir,$(sort $(dir $(SEEDER_FILES))),$(seeder_dir:$(SEED_SRC_DIR)/%.go=%))


#Targets
.PHONY:all proto clean migrate seed

all: proto

proto: $(PROTO_FILE_NAMES)
	@echo $^

$(PROTO_FILE_NAMES):
	@echo "Generating Protocol Buffer files for $(@)..."
	protoc \
	--proto_path=proto/ \
    --go_out=proto/ \
    --go_opt=paths=source_relative proto/$(@)/$(@).proto \
    --go-grpc_out=proto/ \
    --go-grpc_opt=paths=source_relative proto/$(@)/$(@).proto

run:
	@go run cmd/grpc-clean/main.go

migrate:
	@go run cmd/grpc-clean/main.go migrate

seed:
	@echo $^
#	@go run cmd/grpc-clean/main.go seed

test:
	@go test /var/www/banking-microservice/pkg/v1/handler/tests

clean:
	@echo "Cleaning generated files..."
	@rm -rf $(GENERATED_PROTO_FILES)
	@echo "Cleanup complete"