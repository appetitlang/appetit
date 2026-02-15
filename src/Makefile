version = 1
dev_version = "true"

OS = $(shell uname)

# Thanks to https://stackoverflow.com/a/60413199 and
# Thanks to https://github.com/memkind/memkind/issues/33#issuecomment-540614162
ifeq ($(OS), Darwin)
	MAKEFLAGS := --jobs=$(shell sysctl -n hw.logicalcpu)
else
	MAKEFLAGS := --jobs=$(shell nproc)
endif

DPKG_ARCH = $(shell dpkg --print-architecture)

# If the dev_version is true, don't add a build date so that it remains "-development"
ifeq ($(dev_version), "false")
	# Thanks to https://www.digitalocean.com/community/tutorials/using-ldflags-to-set-version-information-for-go-applications
	LD_FLAGS=-s -w -X 'appetit/parser.BuildDate=$(shell date)'
else
	LD_FLAGS=-s -w
endif

# Build all binaries
all: clean freebsd linux macos netbsd openbsd windows

# Clean up old builds/packaging and the Go cache
clean:
	@echo "[Clean] Cleaning up caches and ../dist/ directories"
	@go clean -cache
	@rm -rf ../dist/

# Build for FreeBSD
freebsd: clean
	@echo "\033[31m[FreeBSD] Building x86_64 binary\033[0m"
	@env GOOS=freebsd GOARCH=amd64 go build -ldflags="$(LD_FLAGS)" -o ../dist/appetit-freebsd-amd64-$(version)
	@echo "\033[31m[FreeBSD] Building armv6 binary\033[0m"
	@env GOOS=freebsd GOARM=6 GOARCH=arm go build -ldflags="$(LD_FLAGS)" -o ../dist/appetit-freebsd-arm-$(version)
	@echo "\033[31m[FreeBSD] Building arm64 binary\033[0m"
	@env GOOS=freebsd GOARCH=arm64 go build -ldflags="$(LD_FLAGS)" -o ../dist/appetit-freebsd-arm64-$(version)

# Install the application
install: me
	@echo "\033[37m:: Installing binary to /usr/local/bin/...\033[0m"
	@sudo mv ../dist/appetit /usr/local/bin/
	@echo "\033[31mRemoving ../dist/ directory...\033[0m"
	@rm -rf ../dist/

# Build for Linux
linux: clean
	@echo "\033[33m[Linux] Building x86_64 binary\033[0m"
	@env GOOS=linux GOARCH=amd64 go build -ldflags="$(LD_FLAGS)" -o ../dist/appetit-linux-amd64-$(version)
	@echo "\033[33m[Linux] Building armv6 binary\033[0m"
	@env GOOS=linux GOARM=6 GOARCH=arm go build -ldflags="$(LD_FLAGS)" -o ../dist/appetit-linux-arm-$(version)
	@echo "\033[33m[Linux] Building arm64 binary\033[0m"
	@env GOOS=linux GOARCH=arm64 go build -ldflags="$(LD_FLAGS)" -o ../dist/appetit-linux-arm64-$(version)

# Build for macOS
macos: clean
	@echo "\033[36m[macOS/Darwin] Building x86_64 binary\033[0m"
	@env GOOS=darwin GOARCH=amd64 go build -ldflags="$(LD_FLAGS)" -o ../dist/appetit-darwin-amd64-$(version)
	@echo "\033[36m[macOS/Darwin] Building arm64 binary\033[0m"
	@env GOOS=darwin GOARCH=arm64 go build -ldflags="$(LD_FLAGS)" -o ../dist/appetit-darwin-arm64-$(version)
	@echo "\033[36m[macOS/Darwin] Building universal binary\033[0m"
	@lipo -create -output ../dist/appetit-darwin-universal-$(version) ../dist/appetit-darwin-amd64-$(version) ../dist/appetit-darwin-arm64-$(version)

# Build for the platform you are on
me: clean
	@echo ":: Building binary for your platform"
	@go build -ldflags="$(LD_FLAGS)" -o ../dist/appetit

# Build for NetBSD
netbsd: clean
	@echo "\033[35m[NetBSD] Building x86_64 binary\033[0m"
	@env GOOS=netbsd GOARCH=amd64 go build -ldflags="$(LD_FLAGS)" -o ../dist/appetit-netbsd-amd64-$(version)
	@echo "\033[35m[NetBSD] Building armv6 binary\033[0m"
	@env GOOS=netbsd GOARM=6 GOARCH=arm go build -ldflags="$(LD_FLAGS)" -o ../dist/appetit-netbsd-arm-$(version)
	@echo "\033[35m[NetBSD] Building arm64 binary\033[0m"
	@env GOOS=netbsd GOARCH=arm64 go build -ldflags="$(LD_FLAGS)" -o ../dist/appetit-netbsd-arm64-$(version)

# Build for OpenBSD
openbsd: clean
	@echo "\033[32m[OpenBSD] Building x86_64 binary\033[0m"
	@env GOOS=openbsd GOARCH=amd64 go build -ldflags="$(LD_FLAGS)" -o ../dist/appetit-openbsd-amd64-$(version)
	@echo "\033[32m[OpenBSD] Building armv6 binary\033[0m"
	@env GOOS=openbsd GOARM=6 GOARCH=arm go build -ldflags="$(LD_FLAGS)" -o ../dist/appetit-openbsd-arm-$(version)
	@echo "\033[32m[OpenBSD] Building arm64 binary\033[0m"
	@env GOOS=openbsd GOARCH=arm64 go build -ldflags="$(LD_FLAGS)" -o ../dist/appetit-openbsd-arm64-$(version)

# dpkg build for Debian based systems
pkg_deb: me
	@echo "\033[37m:: Making ../dist/appetit_$(version).deb for Debian $(DPKG_ARCH)...\033[0m"
	@mkdir -p ../dist/appetit_$(version)/DEBIAN
	@mkdir -p ../dist/appetit_$(version)/usr/local/bin
	@cp ../dist/appetit ../dist/appetit_$(version)/usr/local/bin
	@cp ../packaging/control ../dist/appetit_$(version)/DEBIAN/control
	@echo "\nArchitecture: $(DPKG_ARCH)" >> ../dist/appetit_$(version)/DEBIAN/control
	@dpkg-deb --root-owner-group --build ../dist/appetit_$(version)
	@mv ../dist/appetit_$(version).deb ../dist/appetit_$(version).$(DPKG_ARCH).deb
	@rm -rf ../dist/appetit_$(version)/
	@rm -rf ../dist/appetit

# pkg build for macOS
pkg_macos: macos
	@echo "\033[37m:: Making ../dist/appetit_$(version).pkg for macOS (universal binary)...\033[0m"
	@mkdir -p ../dist/appetit_$(version)/usr/local/bin
	@mv ../dist/appetit-darwin-universal-$(version) ../dist/appetit_$(version)/usr/local/bin/appetit
	@pkgbuild --root ../dist/appetit_$(version)/usr/local/bin --install-location /usr/local/bin --identifier com.appetitlang.appetit --version $(version) ../dist/appetit.pkg  > /dev/null 2>&1
	@cp ../LICENCE ../dist/; cp ../packaging/macos_distribution.xml ../dist/
	@cd ../dist/; productbuild --quiet --distribution macos_distribution.xml --resources . --package-path appetit.pkg appetit_$(version).pkg
	@rm -rf ../dist/appetit_$(version)/; rm -rf ../dist/appetit.pkg; rm -rf ../dist/macos_distribution.xml; rm -rf ../dist/LICENCE

# This is for releases as it builds all the binaries and a source archive
release: all
	@$(MAKE) source

# Make a source archive
source:
	@echo "\033[37m[Source] Making archive\033[0m"
	@cp ../LICENCE .
	@mkdir -p ../dist/
	@tar --exclude=../dist/ -cJf ../dist/appetit-src-v$(version).tar.xz *
	@rm -f LICENCE

# Run the Go unit tests
test:
	@go clean -testcache
	@go test ./...

# Build for Windows
windows: clean
	@echo "\033[34m[Windows] Building x86_64 binary\033[0m"
	@env GOOS=windows GOARCH=amd64 go build -ldflags="$(LD_FLAGS)" -o ../dist/appetit-windows-amd64-$(version).exe
	@echo "\033[34m[Windows] Building arm64 binary\033[0m"
	@env GOOS=windows GOARCH=arm64 go build -ldflags="$(LD_FLAGS)" -o ../dist/appetit-windows-arm64-$(version).exe