# https://blog.jgc.org/2011/07/gnu-make-recursive-wildcard-function.html
rwildcard = $(foreach d,$(wildcard $(addsuffix *,$(1))),$(call rwildcard,$(d)/,$(2)) $(filter $(subst *,%,$(2)),$(d)))

PROTODIR=policies
PROTO_SOURCES := $(patsubst %.proto,%.pb.go,$(call rwildcard,$(PROTODIR),*.proto))
PGRPC_SOURCES := $(patsubst %.proto,%_grpc.pb.go,$(call rwildcard,$(PROTODIR),*.proto))

.PHONY: all
all: protobuf

.PHONY: protobuf
protobuf: $(PROTO_SOURCES)

%.pb.go: %.proto
	$(info $(PROTO_SOURCES))
	protoc --go_out=paths=source_relative:. $^
	protoc --go-grpc_out=paths=source_relative:. $^

.PHONY: protoclean
protoclean:
	rm -f $(PROTO_SOURCES)
	rm -f $(PGRPC_SOURCES)

.PHONY: clean
clean: protoclean