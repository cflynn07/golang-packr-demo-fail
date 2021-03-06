FROM cflynnus/golang-1.12.0-packr2 as builder
WORKDIR /go/src/golang-packr-demo
COPY . .
RUN packr2
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./main .

FROM alpine:latest
WORKDIR /root
COPY --from=builder /go/src/golang-packr-demo/main ./main
CMD ["./main"]
