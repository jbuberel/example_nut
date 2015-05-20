# Example application using the `nut` vendoring tool

* Import paths are rewritten.
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

# To deploy this app to [Google App Engine](https://cloud.google.com/appengine/):

* Install the [Google Cloud SDK](https://cloud.google.com/sdk/).
* Create a project in the [Google Cloud Console](https://console.developers.google.com/project).
* `gcloud auth login`
* `gcloud config set project <your-project-name>`
* `go get -u google.golang.org/appengine/cmd/aedeploy`

```
$ cd nutserver
$ aedeploy gcloud --quiet preview app deploy --version myapp ./app.yaml --remote
$ curl myapp.<your-project-name>.appspot.com
```
