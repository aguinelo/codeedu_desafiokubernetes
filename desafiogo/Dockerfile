FROM golang:1.14-alpine as builder

WORKDIR /desafiogo

COPY src/app/main.go main.go
COPY src/app/main_test.go main_test.go

ENV CGO_ENABLED=0

RUN go build 

EXPOSE 8000

FROM scratch

COPY --from=builder /desafiogo .

ENTRYPOINT ["./desafiogo"]