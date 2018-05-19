FROM golang:1.9.2-alpine3.6 AS build

# Install tools required to build the project
# We need to run `docker build --no-cache .` to update those dependencies
RUN apk add --no-cache --update git make
RUN go get github.com/golang/dep/cmd/dep

# Gopkg.toml and Gopkg.lock lists project dependencies
# These layers are only re-built when Gopkg files are updated
COPY Gopkg.lock Gopkg.toml /go/src/github.com/dealako/restapi/
WORKDIR /go/src/github.com/dealako/restapi/
# Install library dependencies
RUN dep ensure -vendor-only

# Copy all project and build it
# This layer is rebuilt when ever a file has changed in the project directory
COPY . /go/src/github.com/dealako/restapi/
RUN make all
COPY restapi-linux-amd64 /bin/restapi

# This results in a single layer image
FROM scratch
COPY --from=build /bin/restapi /bin/restapi
ENTRYPOINT ["/bin/restapi"]
CMD ["-p", "8000"]