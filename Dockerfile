# Use a mutli-stage build pipeline to generate the executable
FROM golang:1.11

ARG VERSION="development"

ENV GO_PATH="/go"

ADD . /src
WORKDIR /src

ENV CGO_ENABLED=0
ENV GOOS=linux
RUN go build -o bin/rb-ldap -a -installsuffix cgo -ldflags "-s -X main.version=$VERSION" github.com/redbrick/rb-ldap/cmd/rb-ldap

# Build the actual container
FROM alpine:latest

RUN apk add --update ca-certificates

COPY --from=0 /src/bin/rb-ldap /bin/rb-ldap

LABEL VERSION=$VERSION

WORKDIR /bin
ENTRYPOINT ["/bin/rb-ldap"]
