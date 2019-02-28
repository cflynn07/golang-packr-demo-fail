golang-packr-demo-fail
======================

A demonstration of [gobuffalo/packr](https://github.com/gobuffalo/packr). This
is a simple website with static content like .js, .css, .jpg files. Packr
allows for creation of a single binary, bundled with static files, that
services a website. It also allows static files to be loaded from disk during
development and testing.

This project uses [skaffold](https://github.com/GoogleContainerTools/skaffold)
for local development with kubernetes. With skaffold static files can be
updated in a running container and only edits to .go files will trigger a
rebuild.

Development Instructions
------------------------
After checking out this repository locally:
```
$ skaffold dev
```
Changes to static files will be pushed to the running container and changes to
.go files will trigger a rebuild.

Production Build Instructions
-----------------------------
This will product an alpine image with a single built binary that contains all
of this website's static files.
```
$ go build -t golang-packr-demo .
// optionally run the image
$ go run --name golang-packr-demo -e PORT 9000 -p 9000:9000 --rm golang-packr-demo
Listening port 9000
```
