FROM golang:alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY ./ ./

RUN swag init -g delivery/funcs.go

RUN CGO_ENABLED=0 GOOS=linux go build -o /backend main.go

FROM scratch

COPY --from=builder /backend /backend
COPY --from=builder /app/times.ttf /times.ttf

EXPOSE 8080

ENTRYPOINT ["/backend"]
