# Example application using the `nut` vendoring tool

* Import paths are rewritten.
* Users can continue to use the `go build` command to build their application.

# To build this application:

You do not need to install `nut` to build.
```
mkdir vendornut
cd vendornut
export GOPATH=`pwd`
go get github.com/jbuberel/example_nut
```

The binary will be named `nutserver`

# To deploy this app to [Google App Engine](https://cloud.google.com/appengine/):

* Install the [Google Cloud SDK](https://cloud.google.com/sdk/).
* Create a project in the [Google Cloud Console](https://console.developers.google.com/project).

First, make sure you're authenticated to your Google Clould account and have your default project ID set:
```
gcloud auth login
gcloud config set project <your-project-name>
```
Next, you'll need to download and install our `aedeploy` program, which will make deployment much easier:
```
go get -u google.golang.org/appengine/cmd/aedeploy
```
Make sure that the `aedeploy` command is in your command `$PATH`, and then run the following command from within the directory that contains the `app.yaml` file:
```
$ aedeploy gcloud --quiet preview app deploy --version myapp ./app.yaml --remote
$ curl myapp.<your-project-name>.appspot.com
```
