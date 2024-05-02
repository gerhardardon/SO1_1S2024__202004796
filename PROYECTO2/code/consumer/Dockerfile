# Dockerfile para el servicio consumidor
FROM golang:1.16-alpine
WORKDIR /app
COPY . .
RUN go build -v -o consumer .
CMD ["./consumer"]
