INSTALL_DIR?=/usr/local/bin
GO_FILES=$(shell find -name '*.go')

clean:
	rm -rf textwrap go.mod go.sum

go.mod:
	go mod init git.sr.ht/~dricottone/textwrap

textwrap: go.mod $(GO_FILES)
	go build

build: textwrap

install: textwrap
	install -m755 textwrap $(INSTALL_DIR)/textwrap

uninstall:
	cd $(INSTALL_DIR) && rm -f textwrap

.PHONY: clean build install uninstall
