
GO = go
TARGETS = arch basename cat date dirname echo uname mkdir tail true false yes who whoami uptime pwd base64 base32 head


.PHONY: all
all: $(TARGETS)

$(TARGETS):
	cd cmd/$@ && $(GO) build .

.PHONY: clean
clean: $(TARGETS:%=clean-%)
	$(GO) clean

$(TARGETS:%=clean-%):
	cd cmd/$(@:clean-%=%) && $(GO) clean
