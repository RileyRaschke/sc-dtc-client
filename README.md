
# sc-dtc-client

A `go` client implementation for DTC Protocol servers

**Status:** Early Initial Development (~15% complete)

#### Build
```
go build
./sc-dtc-client --genconfig
# populate config as needed, baseline may work for localhost
./sc-dtc-client
```

#### Recompile protocol buffers

Running `./scripts/protoc_compile.sh` should create/update/restore `dtc/DTCProtocol.pb.go`. This isn't necessary for build, but will be for updates for the `.proto` DTC spec

