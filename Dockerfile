FROM registry.docker.ir/library/golang:1.19 AS build-stage

WORKDIR /app

COPY . .
RUN go get -v -d ./...
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /simple-file-service -v github.com/xoltawn/simple-file-storage

FROM build-stage AS run-test-stage
RUN go test -v ./...

FROM registry.docker.ir/library/alpine:3 AS build-release-stage

WORKDIR /

COPY --from=build-stage /simple-file-service /simple-file-service

EXPOSE 8080

USER 1001:1001

ENTRYPOINT ["/simple-file-service"]