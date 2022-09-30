# example-wire

[Wire] is a code generation tool that automatically generates code about dependency injection, the generated code would be built by compile-time.
You do not have additional effort to write bulk codes of error-check for the initialization of dependencies, Wire helps you complete that. 

## How to build & run
### Build
```shell
$ go build ./...
```
### Run
```shell
$ export CACHE_SIZE=10; export DB_URL="test.com"; # Config loads environment variables
$ ./example-wire
2022/10/01 03:59:27 serve
```

## How to generate `wire_gen.go`
### Install wire
```shell
$ go install github.com/google/wire/cmd/wire@latest
```
### Generate code
```shell
$ wire
```

[Wire]: https://github.com/google/wire