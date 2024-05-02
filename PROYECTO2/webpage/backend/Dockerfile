# Utiliza la imagen oficial de Golang como base
FROM golang:1.17-alpine as build-stage

# Establece el directorio de trabajo en /app
WORKDIR /app

# Copia el código fuente de la aplicación al directorio de trabajo
COPY . .

# Compila la aplicación Go
RUN go build -o app

# Etapa de producción
FROM alpine:latest

# Copia el ejecutable compilado desde la etapa de compilación
COPY --from=build-stage /app/app /app/app

# Instala las dependencias necesarias para ejecutar la aplicación (por ejemplo, MongoDB driver)
RUN apk add --no-cache ca-certificates

# Establece el directorio de trabajo en /app
WORKDIR /app

# Comando para iniciar la aplicación cuando se inicie el contenedor
CMD ["./app"]
