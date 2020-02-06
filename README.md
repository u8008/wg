# wg
`wg` is a Golang module (toolkit) containing packages to facilitate the software development.
  - As the 4th edition developed on top of the other 3 similar libraries for C & C++, it is relatively mature to be used in the production code
  - It has NO dependency on any external packages other than the built-in ones shipped with [Golang](https://golang.org/dl/), so you are free of analyzing threats for the packages in the dependency-chain. It's a huge relief when using it in your production code
## Usage
Managed by the `go` sub-commands shipped with Golang, so nothing is special or unfamiliar here
  - Install with `go get github.com/u8008/wg/...`
  - Import with `import github.com/u8008/wg/<pkg>`
  - Upgrade with `go get -u github.com/u8008/wg/...`
  - Uninstall with `rm -f $GOPATH/src/wg`
  - Read docs with `go doc -all github.com/u8008/wg/<pkg>`
## Packages
### test
  - A wrapper of the built-in `testing` package to facilitate the unit-testing
  - [Testing with test.Ok()](../../wiki/test.Ok)
## License
[The MIT License](LICENSE)
