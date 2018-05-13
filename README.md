# REST API in Go

This is a sample Go project that illustrates a REST service.  The code example is a walkthru from [a Traversy Media video on YouTube](https://www.youtube.com/watch?v=SonwZ6MF5BE). The application includes a simple REST application with the usual crate, read, update, and delete (CRUD) functions.

The models for the sample REST application are in the `models` folder which consist of a Book, Author and Address objects.

## Dependencies

To build and run the application, the `mux` router package must be installed.

```bash
# Install mux router
go get -u github.com/gorilla/mux
```

## Build

To build the application, run the following command:

```bash
go build
```

This will generate a binary for the current computer architecture.

## Running

To run the application, run the following command:

```bash
./restapi
```

Once the application is running, connect on port 8000 with a web browser, your favorite REST client such as Postman, or other HTTP client.