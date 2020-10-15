
GO = go
TARGETS = mkdir true false yes whoami uptime pwd base64 base32 head


.PHONY: all
all: $(TARGETS)

$(TARGETS):
	cd cmd/$@ && $(GO) build .

.PHONY: clean
clean: $(TARGETS:%=clean-%)
	$(GO) clean

$(TARGETS:%=clean-%):
	cd cmd/$(@:clean-%=%) && $(GO) clean
