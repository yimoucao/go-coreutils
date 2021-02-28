
GO = go
TARGETS = $(shell ls ./cmd)
OUTDIR = _out

.PHONY: all
all: $(TARGETS)

$(TARGETS): | out-dir
	$(GO) build -o $(OUTDIR)/ ./cmd/$@

.PHONY: out-dir
out-dir:
	mkdir -p $(OUTDIR)

.PHONY: clean
clean: $(TARGETS:%=clean-%)
	$(GO) clean

$(TARGETS:%=clean-%):
	rm $(OUTDIR)/$(@:clean-%=%) || true
	cd cmd/$(@:clean-%=%) && $(GO) clean
