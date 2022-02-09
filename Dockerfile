FROM golang:1.17-buster AS build

WORKDIR /go/src/app
COPY . .

RUN make

FROM debian:buster-slim

COPY --from=build /go/src/app/bin/nycmesh-tool /nycmesh-tool

ENTRYPOINT ["/nycmesh-tool"]
CMD ["-h"]
