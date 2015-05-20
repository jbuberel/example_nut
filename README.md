# Example application using the `nut` vendoring tool

* Import paths are rewritten
* Users can continue to use the `go build` command to build their application.

# To build this application:

You do not need to install `nut` to build.
```
mkdir vendornut
cd vendornut
export GOPATH=`pwd`
go build github.com/jbuberel/example_nut/nutserver/
```

The binary will be named `nutserver`

