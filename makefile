
GO = go
TARGETS = arch b2sum base32 base64 basename cat
TARGETS += chgrp chmod chown chroot cksum cp csplit cut
TARGETS += date dd dirname
TARGETS += echo factor false head hostid hostname kill logname md5sum
TARGETS += mkdir mkfifo mktemp pwd stat tail true uname uptime
TARGETS += who whoami yes
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
