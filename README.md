# REST API in Go

This is a sample Go project that illustrates a REST service.  The code example is a walkthru from [a Traversy Media video on YouTube](https://www.youtube.com/watch?v=SonwZ6MF5BE). The application includes a simple REST application with the usual crate, read, update, and delete (CRUD) functions.

The models for the sample REST application are in the `models` folder which consist of a Book, Author and Address objects.

## Dependencies

To build and run the application, the `mux` router package and the `logrus` logging package must be installed.

The `Makefile` in the root folder includes a target to install these:

```bash
# Install dependencies
make install-deps
```

Otherwise, you can install them manually:

```bash
# Install mux router
go get -u github.com/gorilla/mux
go get -u github.com/sirupsen/logrus
```

## Build

To build the application, run one of the following commands:

```bash
# Using make - will include additional metadata in the binary such as BUILD_TIME, VERSION, git commit/branch, and app name
make
```

```bash
# Using go build which doesn't include additional metadata
go build
```

This will generate a binary for the current computer architecture.

## Running

To run the application, run the following command:

```bash
./restapi
```

Once the application is running, connect on port 8000 with a web browser,
your favorite REST client such as [Postman](https://www.getpostman.com/), or
other HTTP clients such as [cURL](https://curl.haxx.se/) or even
[resty](https://github.com/micha/resty).