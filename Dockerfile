# base image for building the binary
FROM golang:1.11-alpine AS base
# $GOPATH is /go
COPY . /go/src/github.com/ironarachne/namegen
# the output is located in the working directory without fileextension
# binary path: /go/main
RUN go build /go/src/github.com/ironarachne/namegen/cmd/namegen/main.go

# add the binary to an empty image
FROM scratch
# copy from build-image
COPY --from=base /go/main /namegen
# set namegen as
ENTRYPOINT ["/namegen"]
