FROM golang:1.17 AS build
WORKDIR /go/src
COPY . .

ENV CGO_ENABLED=0
RUN go mod download
RUN go build -o build/backend cmd/main.go

FROM scratch AS runtime
COPY --from=build /go/src/build/backend .
EXPOSE 8080/tcp
ENTRYPOINT ["./backend"]