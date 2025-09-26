# Appetit
Appetit is a solution to a problem that really doesn't need to be solved but I wanted to solve it anyway. In short, Appetit is a simple scripting language specifically designed to help with managing a system. It's like a shell scripting language that is way less powerful but has a nicer (human readable) syntax that works across platforms.

At the end of the day, this is really a project being used to learn Go, less so a project that has the end goal of being viable for anything serious. In light of that, I'm happy to take on board any suggestions but note that this project is a hobby project of mine first and foremost so having fun is the most important focus here.

Started: 11/05/2025.
First public release: 24/09/2025.

Homepage: https://bryanabsmith.com/.

### Should I use this?
Probably not. You should not expect this to be reliable and in any shape that even approximates stable. In light of that, you are **strongly encouraged to run this in a virtual machine or on a machine where data loss is acceptable**.


## Building

### App
This is just a regular Go application so simply do a quick `go build` in the `src/` directory and you'll get a compiled binary that isn't optimised. If you want an optimised release build, execute the following:

    go build -ldflags="-s -w -X 'main.BuildDate=$(date)'" -o dist/appetit

#### Makefiles
There is a conventional Makefile for non-Windows systems and a `Make.ps1` which is a Makefile of sorts for Windows.

If you want this to work across platforms, you can run `make` in the `src/` directory and you will get x86_64 and ARM64 builds for macOS, Linux, Windows, NetBSD, FreeBSD, and OpenBSD. You can also just run `make [insert os]` to make builds for one of the supported operating systems.

The `Make.ps1` file is really just to build Windows builds for now through PowerShell. In short, it's an effort to replace the `Makefile` with something more Windows friendly.

#### Installers
With [fpm](https://fpm.readthedocs.io/en) installed, you can build installers like so from the src directory after `make me` has been run:

**macOS**

    fpm -t osxpkg -p appetit.pkg

**Linux (deb)**

    fpm -t deb -p appetit.deb

**Linux (rpm)**

    fpm -t rpm --rpm-os linux -a arm64 -p appetit.rpm [for ARM64 builds]

    fpm -t rpm --rpm-os linux -p appetit.rpm [for x86_64 builds]


### Testing
Tests are rather haphazardly written so far but there are an increasing number coming. To run what's available, run `go test ./...` in the root of the `src/` directory.

You can also run `make test`.

#### Testing Machines
The language is tested on the following platforms.
- macOS 26 [this is the primary development platform]
- Fedora 42
- FreeBSD 14.3
- Windows 11

The version of Go used is whatever is available via Homebrew on macOS. See [here](https://formulae.brew.sh/formula/go#default) for more info.


### Visual Studio Code Extension
This is a simple extension for Visual Studio Code that adds some basic syntax highlighting and snippet support to the editor.

#### Making the Extension
If vsce isn't installed, get it first:

    npm install -g @vscode/vsce

Run the following:

    mkdir -p dist/
	cd src/
    vsce package --allow-missing-repository
	mv appetit*.vsix ../dist/

There's also a Makefile available. Simply run `make` to clean up any lingering artefacts (like an old build) and package a fresh version which is output to `dist/`.


#### Installing the Extension
In Visual Studio Code:

1. Open the Command Palette
2. Select "Extensions: Install from VSIX..."
3. Select the VSIX file in dist/
4. Profit


## Using
Using the app is as simple as invoking it with the name of the script:

    appetit [script name.apt]

### Flags
There are a handful of flags:

| Flag | Description |
|----|----|
| -allowexec | Allow execution of system commands. This defaults to disabled. If you use the `execute` command. |
| -create | Pass a file name to create a template script. Eg: `-create=~/Desktop/test.apt |
| -dev | Prints out information relevant for development of the interpreter itself. |
| -docs | Serves up a local copy of some lightweight documentation via a web server. |
| -timer | Time the execution of the script. |
| -verbose | Output details about steps when certain actions are performed but don't normally have output. Defaults to disabled. |
| -version | Outputs the version number of the interpreter. |

## Language Syntax and Functionality
The documentation is available in one of two places:
1. [The project's homepage](https://bryanabsmith.com).
2. Running the interpreter with the `-docs` flag.


## Licence

### Everything but the art/icons/ folder
Copyright 2025 Bryan Smith.

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the “Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

### The art/icons/ folder
The icon bases here are from the [Tango Project](https://commons.wikimedia.org/wiki/Tango_icons) which kindly released their icons into the public domain. As a result, the icons for this project are also released into the public domain.