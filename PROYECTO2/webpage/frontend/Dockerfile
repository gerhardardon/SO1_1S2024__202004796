# Usa la imagen oficial de Node como base
FROM node:14-alpine as build-stage

# Establece el directorio de trabajo en /app
WORKDIR /app

# Copia el archivo package.json al directorio de trabajo
COPY package*.json ./

# Instala las dependencias
RUN npm install

# Copia el resto de los archivos del proyecto al directorio de trabajo
COPY . .

# Compila la aplicación para producción
RUN npm run build

# Etapa de producción
FROM nginx:alpine

# Copia los archivos compilados desde la etapa de compilación
COPY --from=build-stage /app/dist /usr/share/nginx/html

# Configura el puerto en el que escucha el servidor Nginx
EXPOSE 80

# Comando para iniciar Nginx en primer plano cuando se inicie el contenedor
CMD ["nginx", "-g", "daemon off;"]
