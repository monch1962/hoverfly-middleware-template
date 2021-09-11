# hoverfly-middleware-template
Sample Hoverfly middleware template written in Go 

This is a Go template for creating Hoverfly (https://hoverfly.io) middleware

**The only code you should need to change is the `transform` function in `middleware.go`, and write any functional tests you like in `middleware_test.go` to confirm your `transform` middleware is doing what it should. This template contains some sample code that you should replace with your own test cases & middleware logic**

Once you've finished implementing your middleware logic, you should run `go build` which will create a file named `middleware`. The `middleware` file is the file you want to use as middleware in your Hoverfly stub.

Note that `middleware` runs as a standalone file, with no external dependencies

## What is Hoverfly middleware?

Hoverfly is a highly capable stub engine for HTTP traffic. While it is highly capable, there are occasions when the logic you want to implement in your stub configurations is beyond what Hoverfly can do natively.

In this case, you can implement your own "middleware" layer to add functionality that the Hoverfly engine doesn't natively support. For example, you might want to implement middleware that hits a separate URL to grab a token, then send that token back in your response. Hoverfly doesn't support this functionality natively, but it can be implemented in middleware.

## Why write Hoverfly middleware in Go rather than Python or JavaScript?

A big advantage of using Go is that it creates statically-linked, standalone executable files. This means that your Hoverfly middleware lives in a single file that can be loaded into a container running Hoverfly, rather than requiring Python or JavaScript runtimes plus various libraries to be loaded alongside the middleware.

This makes it easier to use the standard Hoverfly Docker image at https://hub.docker.com/r/spectolabs/hoverfly/ when you need to implement middleware, rather than build your own bespoke Hoverfly Docker image containing Python or Javascript runtimes, plus keep that image updated. You can load your middleware via a volume mount, then load it at runtime when the container starts.

At an enterprise level, this gives you far better control over your software supply chain. You can copy the standard Hoverfly Docker image into a Docker registry controlled by the enterprise, rather than build your own Hoverfly image containing all the files required to run your (Python or JavaScript) middleware then keep it updated over time. When the standard Hoverfly Docker image is updated, you can simply update the copy in your enterprise's Docker registry at your convenience and keep working.

## Getting started

### Case 1: you've got a Go compiler installed

#### Testing

I tend to use Test-Driven Development (TDD) so I've made it easy to implement. You should be able to copy/paste/rename/change the `TestPayload1` test case in `middleware_test.go` to construct however many test cases you need to satisfy yourself that your middleware is functionally OK.

When you've finished, you can run

`$ go test`

and all the test cases in `*_test.go` files will execute.

Once you're written your test cases, you can start changing your middleware to get all the tests to pass

#### Building middleware

To compile a version of your middleware that will run on your own system

`$ go build -o middleware middleware.go`

This will create a `./middleware/middleware` executable file that you can run from your command line

To build a version that will run inside a Docker container

`$ GOOS=linux GOARCH=amd64 go build -o middleware middleware.go`

This will also create a `./middleware/middleware` executable file, but this one will run under linux/amd64 (i.e. Linux or a Docker container). You can mount the `./middleware` directory as a Docker volume for the Hoverfly container image if you choose to do so

### Case 2: you don't have a Go compiler installed, but you can run Docker

## TODO

- ~~Add some sort of test cases that make sense~~
- Create a Dockerfile to allow middleware to be built without having to install Go compiler on your laptop
- Add instructions on how to incorporate into the standard Hoverfly Docker image
- Add instructions on how to run your middleware inside the Hoverfly Docker image