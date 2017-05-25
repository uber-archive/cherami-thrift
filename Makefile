.PHONY: bins clean
PROJECT_ROOT=github.com/uber/cherami-thrift
THRIFT_GENDIR=.generated

export PATH := $(GOPATH)/bin:$(PATH)

THRIFT_DIR = idl/cherami
THRIFT_SERVER_DIR = idl/cherami_server

THRIFT_SRCS = $(THRIFT_SERVER_DIR)/metadata.thrift \
	      $(THRIFT_SERVER_DIR)/controller.thrift \
	      $(THRIFT_SERVER_DIR)/admin.thrift \
	      $(THRIFT_SERVER_DIR)/store.thrift \
	      $(THRIFT_SERVER_DIR)/replicator.thrift \
	      $(THRIFT_SERVER_DIR)/shared.thrift \
	      $(THRIFT_DIR)/cherami.thrift

export GO15VENDOREXPERIMENT=1
NOVENDOR = $(shell GO15VENDOREXPERIMENT=1 glide novendor)

export PATH := $(GOPATH)/bin:$(PATH)

THRIFT_GEN=$(GOPATH)/bin/thrift-gen

define thriftrule
THRIFT_GEN_SRC += $(THRIFT_GENDIR)/go/$1/tchan-$1.go

$(THRIFT_GENDIR)/go/$1/tchan-$1.go:: $2 $(THRIFT_GEN)
	@mkdir -p $(THRIFT_GENDIR)/go
	$(ECHO_V)$(THRIFT_GEN) --generateThrift --packagePrefix $(PROJECT_ROOT)/$(THRIFT_GENDIR)/go/ --inputFile $2 --outputDir $(THRIFT_GENDIR)/go \
		$(foreach template,$(THRIFT_TEMPLATES), --template $(template))
endef

$(foreach tsrc,$(THRIFT_SRCS),$(eval $(call \
	thriftrule,$(basename $(notdir \
	$(shell echo $(tsrc) | tr A-Z a-z))),$(tsrc))))

thriftc: $(THRIFT_GEN_SRC)

bins:   thriftc
	git status --porcelain | awk '{print $$2;}' | grep .go | xargs -t -n 1 ./prepend.sh        

clean:
	rm -rf .generated
