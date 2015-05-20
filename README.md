# Example application using the `nut` vendoring tool

* Import paths are rewritten
* Users can continue to use the `go build` command to build their application.

# To build this application:

```
mkdir vendornut
cd vendornut
export GOPATH=`pwd`
go get github.com/jbuberel/example_nut
go build github.com/jbuberel/example_nut/nutserver/
```

The binary will be named `nutserver`

