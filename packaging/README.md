## Packaging

This directory houses materials needed to make native package installers.

- control: this is the [control](https://www.debian.org/doc/debian-policy/ch-controlfields.html) file for making Debian (deb) packages. This will default to making a deb package for the current architecture.
- macos_distribution.xml: this is the distribution.xml file needed by productbuild on macOS to formalise the name and LICENCE. By default, this file is part of a universal binary build.