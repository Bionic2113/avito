FROM golang:1.20.4-alpine3.18 as builder
COPY go.mod go.sum /app/
WORKDIR /app
RUN go mod download
COPY . /app

RUN go build -o app ./cmd/app/main.go


FROM alpine as main
WORKDIR /
COPY --from=builder /app .
CMD [ "./app" ]
