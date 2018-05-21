FROM golang:1.10 as build

WORKDIR /go/src/github.com/faheetah/registry-pattern

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0

COPY . .

RUN go build -ldflags "-s -w" registry-pattern.go

FROM scratch
EXPOSE 1323
COPY --from=build /go/src/github.com/faheetah/registry-pattern/registry-pattern /registry-pattern
COPY registry-pattern.toml /registry-pattern.toml
CMD ["/registry-pattern"]
