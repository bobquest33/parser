.DEFAULT: build

VERSION := 0.1.0
TARGET  := reapub-parser
RELEASE := $(TARGET)-v$(VERSION).tar.gz
LDFLAGS := -ldflags "-X main.version $(VERSION)"

build: $(TARGET)

test:
	@go test ./...

$(TARGET):
	@go build -o $(TARGET) $(LDFLAGS)

run: build
	@./${TARGET}

.PHONY: $(TARGET) test
