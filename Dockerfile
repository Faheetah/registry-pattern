FROM golang:1.10 as build

WORKDIR /go/src/github.com/faheetah/repository-pattern

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0

COPY . .

RUN go build -ldflags "-s -w" repository-pattern.go

FROM scratch
EXPOSE 1323
COPY --from=build /go/src/github.com/faheetah/repository-pattern/repository-pattern /repository-pattern
COPY example.toml /repository-pattern.toml
CMD ["/repository-pattern"]
