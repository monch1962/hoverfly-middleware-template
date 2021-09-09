# hoverfly-middleware-template
Sample Hoverfly middleware template written in Go 

This is a Go template for creating Hoverfly (https://hoverfly.io) middleware

The only code you should need to change is the `transform` function in `main.go`. In this template it contains some sample values that you should replace with your own middleware logic

Once you've finished implementing your middleware logic, you should run `go build` which will create a file named `middleware`. The `middleware` file is the file you want to use as middleware in your Hoverfly stub.

Note that `middleware` runs as a standalone file, with no external dependencies

## TODO

- Add some sort of test cases that make sense
- Create a Dockerfile to allow middleware to be built without having to install Go compiler on your laptop
- Add instructions on how to incorporate into the standard Hoverfly Docker image
- Add instructions on how to run your middleware inside the Hoverfly Docker image