FROM golang:1.19.3-alpine as build

WORKDIR /go/build
COPY . .
RUN go build -o simplehttp ./

FROM alpine:3.17 as release
WORKDIR /app

COPY --from=build /go/build/simplehttp ./

CMD ["/app/simplehttp"]