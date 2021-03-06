---
theme : "white"
transition: "zoom"
highlightTheme: "darkula"
---
# REST API in Go

This is a sample Go project that illustrates a REST service. The code example
is a walkthru from [a Traversy Media video on
YouTube](https://www.youtube.com/watch?v=SonwZ6MF5BE). The application
includes a simple REST application with the usual crate, read, update, and
delete (CRUD) functions.

<span class="fragment"><small>Created by [David Deal](http://github.com/dealako) / [@dealako](http://twitter.com/dealako)</small></fragment>

---

## Dependencies

To build and run the application, the `mux` router package and the `logrus` logging package must be installed.

The `Makefile` in the root folder includes a target to install these:

```bash
# Install dependencies
make install-deps
```

Otherwise, you can install them manually:

```bash
# Manually install dependencies
go get -u github.com/gorilla/mux
go get -u github.com/sirupsen/logrus
go get -u github.com/spf13/cobra
go get -u gopkg.in/alexcesaro/statsd.v2
```

---

## Code Organization

The models for the sample REST application are in the `models` folder which
consist of a Book, Author and Address objects.

<span class="fragment">The `utils` folder contains the command line ascii artwork.</span>

<span class="fragment">The main business logic is in the `cmd` folder.</span>

---

## Build

To build the application, run one of the following commands:

```bash
# Using make - will include additional metadata in the binary such as BUILD_TIME, VERSION, git commit/branch, and app name
make

or

make restapi
```

```bash
# Using go build which doesn't include additional metadata
go build
```

This will generate a binary for the current computer architecture.

---

## Testing

To execute the unit tests, run:

```bash
make test
```

---

## Running

To run the application, run the following command. The default port is `8000`:

```bash
./restapi
```

or specify a different HTTP port:

```bash
./restapi -p 8080
```

Once the application is running, connect on port 8000 with a web browser,
your favorite REST client such as [Postman](https://www.getpostman.com/), or
other HTTP clients such as [cURL](https://curl.haxx.se/) or even
[resty](https://github.com/micha/resty).

---

## Docker

### Docker Build

The `Makefile` has a target to build a docker image.

```bash
make docker
```

---

### Docker Run

To run the docker image, simply run:

```bash
docker run -it -p 8000:8000 dealako/restapi:<git hash>
```

for example:

```bash
docker run -it -p 8000:8000 dealako/restapi:38d8dff
```

---

### Docker Publish

To publish the docker image to hub.docker.com under the dealako project, run:

```bash
make docker-push
```

You will need permissions to push the docker image unless you change the image tag/path.

---

## API - REST Endpoints

| Method | REST Endpoint   | Description                              |
|:-------|:----------------|:-----------------------------------------|
| GET    | /api/books      | Retrives all the books                   |
| GET    | /api/books/{id} | Retrievs a specific book by ID           |
| POST   | /api/books      | Adds a book                              |
| PUT    | /api/books/{id} | Updates a book based on the book ID      |
| DELETE | /api/books/{id} | Deletes a specific book based on the ID  |

---