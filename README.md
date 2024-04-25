
# CONTEO DE VOTOS CON K8S 📚

Gerhard Benjamin Ardon Valdez
202004796
  

## Introduccion 📕

En este proyecto, se tiene como objetivo principal implementar un sistema de votaciones para un

concurso de bandas de música guatemalteca; el propósito de este es enviar tráfico por medio de

archivos con votaciones creadas hacia distintos servicios (grpc y wasm) que van a encolar cada uno

de los datos enviados, así mismo se tendrán ciertos consumidores a la escucha del sistema de colas

para enviar datos a una base de datos en Redis; estos datos se verán en dashboards en tiempo real.

También se tiene una base de datos de Mongodb para guardar los logs, los cuales serán

consultados por medio de una aplicación web.

  

## Objetivos 📗

- Implementar un sistema distribuido con microservicios en kubernetes.

- Encolar distintos servicios con sistemas de mensajerías.

- Utilizar Grafana como interfaz gráfica de dashboards.

- Utilizar redis para el conteo de votos

- Manejar Dockerfiles y archivos .yaml

  

## Tenologias usadas 📘

### K8S 💻
![enter image description here](https://kubernetes.io/images/docs/Container_Evolution.svg)
Kubernetes es una plataforma de código abierto para gestionar y orquestar contenedores, proporcionando herramientas para desplegar, escalar y gestionar aplicaciones de manera eficiente y automatizada en cualquier entorno de infraestructura. Ofrece características como escalabilidad automática, autoreparación, despliegues sin tiempo de inactividad y gestión de recursos, facilitando el desarrollo y la operación de aplicaciones en contenedores a gran escala.

A continuacion se mostraran algunos comandos para la creacion de clusteres en google cloud

Para configurar el gcloud en tu maquina usa (esocge tu usuario, tu proyecto y zona horaria de preferencia):
```
gcloud init
```
Para crear un cluster de kubernetes primero debes habilitar el servicio desde la consola y luego:
```
gcloud container clusters create NOMBRE_DEL_CLUSTER --num-nodes=NUMERO_DE_NODOS --zone=ZONA
```
Despues debes obtener tus credenciales para conectarte desde tu consola de linux con kubectl de la siguiente forma:
```
gcloud container clusters get-credentials NOMBRE_DEL_CLUSTER --zone=ZONA
```
Por ultimo revisa tu conexion con el cluster:
```
kubectl get nodes
```

Pero como creo los deployments y services de mi cluster? 
Para crear los deployments primero se deben crear las imagenes de Docker de cada uno de los servicios que se utilizaran, ademas de subirlos a dockerhub y crear los archivos .yaml (para mas informacion busca el respectivo tema en esta documentacion)

Para aplicar un manifiesto .yaml de tu servicio en tu cluster 
```
kubectl apply -f <tu-manifiesto.yaml>
```
Para ver los pods creados:
```
kubectl get pods
```
Para ver los servicios creaods junto a sus puertos
```
kubectl get services
```
Los archivos yaml del autoscalling se deben subir de la misma manera, sin embargo para los ingress se deben seguir pasos especiales (buscar en este manual)
Informacion extra:
Para ver los logs de tus pods de kubernetes:
```
kubectl logs <nombre-del-pod>
```
Para eliminar uno o multiples pods:
```
kubectl deletepods <nombre-del-pod-1><nombre-del-pod-2><nombre-del-pod-n>
```

### Docker y dockerhub 🐋
![enter image description here](https://www.docker.com/wp-content/uploads/2023/10/products-hub-1-hero-1.svg)
Docker es una plataforma que simplifica el desarrollo, empaquetado y ejecución de aplicaciones en contenedores. Los contenedores son unidades de software que incluyen todo lo necesario para ejecutar una aplicación de forma independiente. DockerHub es un servicio en la nube que actúa como repositorio para contenedores Docker, permitiendo a los desarrolladores compartir y distribuir sus aplicaciones.

Para crear una imagen primero situate en en la rama del Dockerfile y ejecuta:

    docker build -t TU_USUARIO/TU_IMAGEN:TAG .
Inicia sesion en dockerhub:

    docker login
Empuja tu imagen hacia dockerhub:

    docker push TU_USUARIO/TU_IMAGEN:TAG
Luego de eato, tu imagen debera verse entre tus repositorios de dockerhub, si no le pusiste tag se configurara con :latest 

Estas imagenes son necesarias ya que deben ser referenciadas en los deployments.yaml con los que funciona kubernetes.
Para esto te dejo un ejemplo en donde se referencia una imagen en dockerhub de nombre "gerhardardon/go.client:v1"
```
kind: Deployment
metadata:
  annotations:
    kompose.cmd: kompose convert -f docker-compose.yaml
    kompose.version: 1.33.0 (3ce457399)
  labels:
    io.kompose.service: grpc-client
  name: grpc-client
spec:
  replicas: 1
  selector:
    matchLabels:
      io.kompose.service: grpc-client
  template:
    metadata:
      annotations:
        kompose.cmd: kompose convert -f docker-compose.yaml
        kompose.version: 1.33.0 (3ce457399)
      labels:
        io.kompose.network/code-default: "true"
        io.kompose.service: grpc-client
    spec:
      containers:
        - image: gerhardardon/grpc-client:v1
          name: grpc-client
          ports:
            - containerPort: 3000
              hostPort: 3000
              protocol: TCP
      restartPolicy: Always
status: {}
```
Con esto estaras listo para modificar tus servicios y deployments de kubernetes

### Kompose 🐳🚢
![enter image description here](https://miro.medium.com/v2/resize:fit:720/format:webp/1*GNquY8yJP55Hs8ub4g-b0w.jpeg)
Kompose es una herramienta que permite convertir archivos de definición de Docker Compose en archivos de configuración de Kubernetes. Simplifica el proceso de migración de aplicaciones basadas en Docker Compose a entornos de Kubernetes, permitiendo a los usuarios aprovechar las ventajas de la orquestación de contenedores en Kubernetes sin tener que volver a escribir la configuración desde cero.

Algunos comandos básicos de Kompose son:

-   `kompose convert`: Convierte un archivo de definición de Docker Compose en archivos de configuración de Kubernetes.
-   `kompose up`: Crea y despliega recursos de Kubernetes basados en el archivo de definición de Docker Compose.
-   `kompose down`: Elimina los recursos de Kubernetes creados por `kompose up`.
-   `kompose ps`: Muestra el estado de los recursos de Kubernetes creados por `kompose up`.

Estos son solo algunos de los comandos más comunes. Kompose ofrece una variedad de opciones y configuraciones adicionales que pueden ser útiles según las necesidades específicas de tu proyecto.
### Kubernetes Engine 🚢
![enter image description here](https://storage.googleapis.com/gweb-cloudblog-publish/images/gke-ui-ga-5dfkp.max-700x700.PNG)
Google Kubernetes Engine (GKE) es un servicio de Google Cloud Platform que facilita la implementación y gestión de aplicaciones en contenedores utilizando Kubernetes. Con GKE, los desarrolladores pueden ejecutar aplicaciones en contenedores de manera eficiente sin preocuparse por la gestión de la infraestructura subyacente. GKE proporciona herramientas para escalar automáticamente las aplicaciones, garantizar la alta disponibilidad y gestionar la seguridad, todo ello de manera integrada con otros servicios de Google Cloud Platform. En resumen, GKE simplifica la administración de clústeres de Kubernetes, permitiendo a los equipos de desarrollo centrarse en crear aplicaciones.

## Descripcion de despliegues usados 📖
Estos servicios se pueden ver utilizando

    kubectl get svc
- **go-client** 🔵:  este cliente en Golang es el encargado de desencolar los datos de kafka, ademas de llevar el conteo de votos en redis y guardar los logs de votos en la db de Mongo.
Este cliente esta configurado con autoscalling de 2 a 5 pods con el 30% de rendimiento en cada uno.

    `kubectl get hpa`

- **grpc-server** 🔵: El servido de grpc en Go es el encargado de recibir los datos desde el client y usar la libreria de kafka para encolar los datos y que estos sean leidos por el client.
- **grpc-client** 🔵: Este servicio recibe todos los datos de Locust y los envia hacia el server por medio de HTTP2 lo cual lo hace muy rapido y confiable, este es el unico servicio que esta vinculado con el loadBalancer para ser ingresado por este desde el exterior
- **grafana** 📊: Para conectarse desde el exterior del cluster con su external-IP y poder ver graficos de los votos guardados en redis.
- **mongo** 🍃: Para guardar los logs de los votos, se utilizo una db no relacional ya que solo necesitamos ver la informacion como un "registro"
- **redis** 📈: DB de gran velocidad ya que usa registros para la informacio, en esta se guardan todos los contadores de las votaciones 
- **kafka** ⚪: Es un encolador que permite el envio de mensajes utilizando distintos Brokers y Topicos para acceder a estos
- **zookeper** 🦁: Apache ZooKeeper es un servicio centralizado de gestión de configuraciones, sincronización y descubrimiento distribuido para sistemas distribuidos. Proporciona un conjunto de características esenciales para sistemas distribuidos, como la sincronización de procesos, la elección de líderes, la gestión de bloqueos y la gestión de configuraciones.


## Conclusiones 📙
- GKE es una herramienta simple de usar y facil de configurar para tener clusteres de kubernetes en la nube
- K8S es un orquestador de contenedores facil de usar, sin embargo se necesita tener bastante conocimiento de otros servicios (docker, dockerhub, linea de comandos, yamls)
- Kompose es una herramienta bastante util a la hora de la creacion de .yaml, ya que funciona agafrrando como base un docker-compose.yaml el cual es mas facil de configurar
(**IMPORTANE:** RECORDAR HACER LOS CAMBIOS RESPECTIVOS A LOS .YAML GENERADOS CON ESTA HERRAMIENTA)
- LA configuarcion de kubectl y gcloud en el ordenador es fundamental pero bastante sencilla con la documentacion adecuada
- Utilizar LoadBalancer en todos los servecios que estaran expuestos hacia afuera del cluster por facilidad
