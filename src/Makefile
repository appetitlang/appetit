version = 1

OS = $(shell uname)

# Thanks to https://stackoverflow.com/a/60413199 and
# Thanks to https://github.com/memkind/memkind/issues/33#issuecomment-540614162
ifeq ($(OS), Darwin)
	MAKEFLAGS := --jobs=$(shell sysctl -n hw.logicalcpu)
else
	MAKEFLAGS := --jobs=$(shell nproc)
endif

# Thanks to https://www.digitalocean.com/community/tutorials/using-ldflags-to-set-version-information-for-go-applications
LD_FLAGS=-s -w -X 'main.BuildDate=$(shell date)'

# Build all binaries
all: freebsd linux macos netbsd openbsd windows

# Clean up old builds/packaging and the Go cache
clean:
	@echo "[Clean] Cleaning up caches and dist/ directories"
	@go clean -cache
	@rm -rf dist/

# Build for FreeBSD
freebsd: clean
	@echo "\033[31m[FreeBSD] Building x86_64 binary\033[0m"
	@env GOOS=freebsd GOARCH=amd64 go build -ldflags="$(LD_FLAGS)" -o dist/appetit-freebsd-amd64-$(version)
	@echo "\033[31m[FreeBSD] Building armv6 binary\033[0m"
	@env GOOS=freebsd GOARM=6 GOARCH=arm go build -ldflags="$(LD_FLAGS)" -o dist/appetit-freebsd-arm-$(version)
	@echo "\033[31m[FreeBSD] Building arm64 binary\033[0m"
	@env GOOS=freebsd GOARCH=arm64 go build -ldflags="$(LD_FLAGS)" -o dist/appetit-freebsd-arm64-$(version)

# Build for Linux
linux: clean
	@echo "\033[33m[Linux] Building x86_64 binary\033[0m"
	@env GOOS=linux GOARCH=amd64 go build -ldflags="$(LD_FLAGS)" -o dist/appetit-linux-amd64-$(version)
	@echo "\033[33m[Linux] Building armv6 binary\033[0m"
	@env GOOS=linux GOARM=6 GOARCH=arm go build -ldflags="$(LD_FLAGS)" -o dist/appetit-linux-arm-$(version)
	@echo "\033[33m[Linux] Building arm64 binary\033[0m"
	@env GOOS=linux GOARCH=arm64 go build -ldflags="$(LD_FLAGS)" -o dist/appetit-linux-arm64-$(version)

# Build for macOS
macos: clean
	@echo "\033[36m[macOS/Darwin] Building x86_64 binary\033[0m"
	@env GOOS=darwin GOARCH=amd64 go build -ldflags="$(LD_FLAGS)" -o dist/appetit-darwin-amd64-$(version)
	@echo "\033[36m[macOS/Darwin] Building arm64 binary\033[0m"
	@env GOOS=darwin GOARCH=arm64 go build -ldflags="$(LD_FLAGS)" -o dist/appetit-darwin-arm64-$(version)
	@echo "\033[36m[macOS/Darwin] Building universal binary\033[0m"
	@lipo -create -output dist/appetit-darwin-universal-$(version) dist/appetit-darwin-amd64-$(version) dist/appetit-darwin-arm64-$(version)

# Build for NetBSD
netbsd: clean
	@echo "\033[35m[NetBSD] Building x86_64 binary\033[0m"
	@env GOOS=netbsd GOARCH=amd64 go build -ldflags="$(LD_FLAGS)" -o dist/appetit-netbsd-amd64-$(version)
	@echo "\033[35m[NetBSD] Building armv6 binary\033[0m"
	@env GOOS=netbsd GOARM=6 GOARCH=arm go build -ldflags="$(LD_FLAGS)" -o dist/appetit-netbsd-arm-$(version)
	@echo "\033[35m[NetBSD] Building arm64 binary\033[0m"
	@env GOOS=netbsd GOARCH=arm64 go build -ldflags="$(LD_FLAGS)" -o dist/appetit-netbsd-arm64-$(version)

# Build for OpenBSD
openbsd: clean
	@echo "\033[32m[OpenBSD] Building x86_64 binary\033[0m"
	@env GOOS=openbsd GOARCH=amd64 go build -ldflags="$(LD_FLAGS)" -o dist/appetit-openbsd-amd64-$(version)
	@echo "\033[32m[OpenBSD] Building armv6 binary\033[0m"
	@env GOOS=openbsd GOARM=6 GOARCH=arm go build -ldflags="$(LD_FLAGS)" -o dist/appetit-openbsd-arm-$(version)
	@echo "\033[32m[OpenBSD] Building arm64 binary\033[0m"
	@env GOOS=openbsd GOARCH=arm64 go build -ldflags="$(LD_FLAGS)" -o dist/appetit-openbsd-arm64-$(version)

# Build for Windows
windows: clean
	@echo "\033[34m[Windows] Building x86_64 binary\033[0m"
	@env GOOS=windows GOARCH=amd64 go build -ldflags="$(LD_FLAGS)" -o dist/appetit-windows-amd64-$(version).exe
	@echo "\033[34m[Windows] Building arm64 binary\033[0m"
	@env GOOS=windows GOARCH=arm64 go build -ldflags="$(LD_FLAGS)" -o dist/appetit-windows-arm64-$(version).exe

# Build for the platform you are on
me: clean
	@echo ":: Building binary for your platform"
	@go build -ldflags="$(LD_FLAGS)" -o dist/appetit

# Install the application
install: clean
	@echo "\033[37m:: Building release binary...\033[0m"
	@go build -ldflags="$(LD_FLAGS)" -o dist/appetit
	@echo "\033[37m:: Installing binary to /usr/local/bin/...\033[0m"
	@sudo mv dist/appetit /usr/local/bin/
	@echo "\033[31mRemoving dist/ directory...\033[0m"
	@rm -rf dist/

# This is for releases as it builds all the binaries, creates release archives, and a source archive
release: all
	@echo "\n\033[31m[Release: FreeBSD] Making release archive\033[0m"
	@tar -cJf dist/appetit-freebsd-v$(version).tar.xz dist/appetit-freebsd-amd64-$(version) dist/appetit-freebsd-arm64-$(version) dist/appetit-freebsd-arm-$(version)
	@rm -rf dist/appetit-freebsd-amd64-$(version) dist/appetit-freebsd-arm64-$(version) dist/appetit-freebsd-arm-$(version)

	@echo "\033[33m[Release: Linux] Making release archive\033[0m"
	@tar -cJf dist/appetit-linux-v$(version).tar.xz dist/appetit-linux-amd64-$(version) dist/appetit-linux-arm64-$(version) dist/appetit-linux-arm-$(version)
	@rm -rf dist/appetit-linux-amd64-$(version) dist/appetit-linux-arm64-$(version) dist/appetit-linux-arm-$(version)

	@echo "\033[36m[Release: macOS] Making release archive\033[0m"
	@tar -cJf dist/appetit-macos-v$(version).tar.xz dist/appetit-darwin-amd64-$(version) dist/appetit-darwin-arm64-$(version) dist/appetit-darwin-universal-$(version)
	@rm -rf dist/appetit-darwin-amd64-$(version) dist/appetit-darwin-arm64-$(version) dist/appetit-darwin-universal-$(version)

	@echo "\033[35m[Release: NetBSD] Making release archive\033[0m"
	@tar -cJf dist/appetit-netbsd-v$(version).tar.xz dist/appetit-netbsd-amd64-$(version) dist/appetit-netbsd-arm64-$(version) dist/appetit-netbsd-arm-$(version)
	@rm -rf dist/appetit-netbsd-amd64-$(version) dist/appetit-netbsd-arm64-$(version) dist/appetit-netbsd-arm-$(version)

	@echo "\033[32m[Release: OpenBSD] Making release archive\033[0m"
	@tar -cJf dist/appetit-openbsd-v$(version).tar.xz dist/appetit-openbsd-amd64-$(version) dist/appetit-openbsd-arm64-$(version) dist/appetit-openbsd-arm-$(version)
	@rm -rf dist/appetit-openbsd-amd64-$(version) dist/appetit-openbsd-arm64-$(version) dist/appetit-openbsd-arm-$(version)

	@echo "\033[34m[Release: Windows] Making release archive\033[0m"
	@zip -q dist/appetit-windows-v$(version).zip dist/appetit-windows-amd64-$(version).exe dist/appetit-windows-arm64-$(version).exe
	@rm -rf dist/appetit-windows-amd64-$(version).exe dist/appetit-windows-arm64-$(version).exe

	@$(MAKE) source

# Make a source archive
source:
	@echo "\033[37m[Source] Making archive\033[0m"
	@mkdir -p dist/
	@tar --exclude=dist/ -cJf dist/appetit-src-v$(version).tar.xz *

# Run the Go unit tests
test:
	@go test ./...