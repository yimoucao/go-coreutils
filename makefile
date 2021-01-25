
GO = go
TARGETS = arch basename cat date dirname echo uname mkdir tail true false yes who whoami uptime pwd base64 base32 head factor
OUTDIR = ./

.PHONY: all
all: $(TARGETS)

$(TARGETS): | out-dir
	$(GO) build -o $(OUTDIR) ./cmd/$@

.PHONY: out-dir
out-dir:
	mkdir -p $(OUTDIR)

.PHONY: clean
clean: $(TARGETS:%=clean-%)
	$(GO) clean

$(TARGETS:%=clean-%):
	rm $(OUTDIR)/$(@:clean-%=%) || true
	cd cmd/$(@:clean-%=%) && $(GO) clean
