FROM golang:latest as build

RUN go version
ENV GOPATH=/

WORKDIR /stat
COPY ./ ./

RUN go mod download


# Build the binary
RUN go build -o statistic ./cmd/apiserver/main.go
#FROM alpine
#WORKDIR /stat
#COPY  ./configs ./
#COPY --from=build /stat/statistic ./
#COPY  ./wait_for_postgres.sh ./
RUN apt-get update
RUN apt-get -y install postgresql-client
RUN chmod +x wait_for_postgres.sh
ENTRYPOINT ["./statistic"]