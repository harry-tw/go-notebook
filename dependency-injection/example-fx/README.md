# example-fx
[Fx] is dependency injection library that makes you construct dependencies easily in runtime.

See [Fx Example] for more guidance.

## How to build & run
### Build
```shell
$ go build ./...
```
### Run
```shell
$ export CACHE_SIZE=10; export DB_URL="test.com"; # Config loads environment variables
$ ./example-fx
2022/10/01 03:35:19 OnStart hook
2022/10/01 03:35:19 serve
2022/10/01 03:35:19 OnStop hook
2022/10/01 03:35:19 Stop
```

[Fx]: https://github.com/uber-go/fx
[Fx Example]: https://github.com/uber-go/fx/blob/master/example_test.go