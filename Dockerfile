FROM golang:1.19-alpine as builder

ENV GO111MODULE=on

WORKDIR /app
COPY . .

RUN apk --no-cache add git alpine-sdk build-base gcc

RUN go get

RUN go build -o sms-gateway .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/sms-gateway .
COPY .env .
RUN ls -al
ENTRYPOINT ["./sms-gateway"]
CMD ["serve"]