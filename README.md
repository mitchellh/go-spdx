# go-spdx [![Godoc](https://godoc.org/github.com/mitchellh/go-spdx?status.svg)](https://godoc.org/github.com/mitchellh/go-spdx)

go-spdx is a Go library for listing and looking up licenses using
[SPDX IDs](https://spdx.org/licenses/). SPDX IDs are an unambiguous way
to reference a specific software license.

This library does not implement the SPDX document format. SPDX document
parsing and printing are provided by other libraries, including a library
[in the official spdx organization](https://github.com/spdx/tools-go). This
library instead provides the ability to look up licenses via SPDX IDs.
