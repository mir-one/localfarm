SHELL = /bin/bash
NODE_BINDIR = ./node_modules/.bin
export PATH := $(NODE_BINDIR):$(PATH)
LOGNAME ?= $(shell logname)

# adding the name of the user's login name to the template file, so that
# on a multi-user system several users can run this without interference
TEMPLATE_POT = ./languages/template.pot

# Where to find input files (it can be multiple paths).
INPUT_FILES = ./resources/js

# Where to write the files generated by this makefile.
OUTPUT_DIR = ./languages

# Available locales for the app.
LOCALES = en_GB id_ID hu_HU ru_RU

# Name of the generated .po files for each available locale.
LOCALE_FILES ?= $(patsubst %,$(OUTPUT_DIR)/locale/app.po, $(LOCALES))

GETTEXT_SOURCES ?= $($(INPUT_FILES) 2> /dev/null)

.PHONY: all cover clean clean-osx clean-linux-arm clean-linux-amd64 clean-win64 \
	osx linux-amd64 linux-arm windows fetch-dep run osxcross.bin \
  cleantranslations makemessages translations

all: osx linux-amd64 linux-arm windows

clean:
	@[ -f localfarm.osx.amd64 ] && rm -f localfarm.osx.amd64 || true
	@[ -f localfarm.linux.arm ] && rm -f localfarm.linux.arm || true
	@[ -f localfarm.linux.amd64 ] && rm -f localfarm.linux.amd64 || true
	@[ -f localfarm.win.amd64.exe ] && rm -f localfarm.win.amd64.exe || true

clean-osx: localfarm.osx.amd64
	rm -rf $^

clean-linux-arm: localfarm.linux.arm
	rm -rf $^

clean-linux-amd64: localfarm.linux.amd64
	rm -rf $^

clean-win64: localfarm.win.amd64.exe
	rm -rf $^

localfarm.osx.amd64: main.go
	CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -ldflags '-s -w' -o $@
	file $@

osx: localfarm.osx.amd64

osxcross: main.go
	CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 \
		CC=o64-clang	\
		CXX=o64-clang++ \
		go build -ldflags '-s -w' -o localfarm.osx.amd64
	file localfarm.osx.amd64

localfarm.linux.arm: main.go
	CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=7 \
		CC=arm-linux-gnueabihf-gcc	\
		CXX=arm-linux-gnueabihf-g++ \
		go build -ldflags '-s -w' -o $@
	file $@

linux-arm: localfarm.linux.arm

localfarm.linux.amd64: main.go
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -o $@
	file $@

linux-amd64: localfarm.linux.amd64

localfarm.windows.amd64.exe: main.go
	CGO_ENABLED=1 GOOS=windows GOARCH=amd64 \
		CC=x86_64-w64-mingw32-gcc \
		CXX=x86_64-w64-mingw32-g++ \
		go build -ldflags '-s -w' -o $@
	file $@

windows: localfarm.windows.amd64.exe

fetch-dep: Gopkg.toml Gopke.lock
	dep ensure

run: main.go
	go run $^

# Translation makefile
cleantranslations:
	rm -f $(TEMPLATE_POT) $(OUTPUT_DIR)/translations.json

makemessages: $(TEMPLATE_POT)

translations: ./$(OUTPUT_DIR)/translations.json

# Create a main .pot template, then generate .po files for each available language.
# Thanx to Systematic: https://github.com/Polyconseil/systematic/blob/866d5a/mk/main.mk#L167-L183
$(TEMPLATE_POT): $(GETTEXT_SOURCES)
# Extract gettext strings from templates files and create a POT dictionary template.
	node node_modules/vue-webpack-gettext/extract --output ./$(OUTPUT_DIR)/template.pot --src ./resources/js/
# Generate .po files for each available language.
	@for lang in $(LOCALES); do \
		export PO_FILE=$(OUTPUT_DIR)/locale/$$lang/LC_MESSAGES/app.po; \
		mkdir -p $$(dirname $$PO_FILE); \
		if [ -f $$PO_FILE ]; then  \
			echo "msgmerge --update $$PO_FILE $@"; \
			msgmerge --lang=$$lang --update $$PO_FILE $@ || break ;\
		else \
			msginit --no-translator --locale=$$lang --input=$@ --output-file=$$PO_FILE || break ; \
			msgattrib --no-wrap --no-obsolete -o $$PO_FILE $$PO_FILE || break; \
		fi; \
	done;

$(OUTPUT_DIR)/translations.json: $(LOCALE_FILES)
	mkdir -p $(OUTPUT_DIR)
	gettext-compile --output $@ $(LOCALE_FILES)
