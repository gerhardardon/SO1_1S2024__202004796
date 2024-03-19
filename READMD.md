# Plataforma de Monitoreo y Señales a Procesos 🖥️🛰️

## Sistema Operativo 🛠️
- Ubuntu Server 22.04
- VirtualBox 

## Backend Go 📚
Se escogio fiber y go debido a su rendimiento y facilidad de uso, ademas que lo vuelve una API ligera 

Para la creacion de esta API se utilizo:
- Fiber V2
- GoRutines

Se leyeron los modulos insertados con anterioridad en la carpeta "/proc" de neustro SO y se manejaron los datos para devolver los siguientes valores:

| Modulo    | Valor de retorno             | Endpoint       |
|-----------|------------------------------|----------------|
| RAM       | int porcentaje libre         | "api/ram"      |
| CPU       | int porcentaje libre         | "api/cpu"      |
| Procesos  | json procesos con sus hijos  | "api/proclist" |

Ademas se debe mencionar que el manejo del Json que retorna los procesos con sus hijos desde el módulo se maneja con los siguientes structs para su facilidad:

```
type Info struct {
	Processes []Process `json:"Processes"`
	Running   int       `json:"running"`
	Sleeping  int       `json:"sleeping"`
	Zombie    int       `json:"zombie"`
	Stopped   int       `json:"stopped"`
	Total     int       `json:"total"`
}

type Process struct {
	Pid   int            `json:"pid"`
	Name  string         `json:"name"`
	User  int            `json:"user"`
	State int            `json:"state"`
	Ram   int            `json:"ram"`
	Rss   int            `json:"rss"`
	Child []ChildProcess `json:"child"`
}

type ChildProcess struct {
	Pid      int    `json:"pid"`
	Name     string `json:"name"`
	State    int    `json:"state"`
	PidPadre int    `json:"pidPadre"`
	RssChild int    `json:"rssChild"`
	ChildUID int    `json:"childUID"`
}
```

## Web UI 📱  
Para la parte de UI se utilizaron las siguientes herramientas:
- React
- Vite
- Bootstrap
- Chart-js-2
- Graphviz-react

La elección de Vite para un proyecto con React puede ofrecer varias ventajas:

- Rápido tiempo de inicio: Vite utiliza el esquema de módulos nativos de JavaScript (ESM) y la compilación en tiempo real para proporcionar un tiempo de inicio rápido. Esto significa que tus aplicaciones React se cargarán y recargarán más rápidamente durante el desarrollo.

- Desarrollo instantáneo: Vite ofrece una experiencia de desarrollo instantáneo, lo que significa que los cambios en tu código se reflejarán casi al instante en el navegador, sin necesidad de recargar la página. Esto acelera el proceso de desarrollo y mejora la productividad.
Soporte para TypeScript y JSX: Vite es compatible con TypeScript y JSX de manera nativa, lo que te permite aprovechar las ventajas de estas tecnologías sin necesidad de configuración adicional.

Ademas la creacion con
```
create-vite frontend --template react
npm run dev
```
es mucho mas rapida

## Uso de modulos 🧬🧬
En C para Linux, los módulos son fragmentos de código que pueden ser cargados y descargados en el kernel del sistema operativo en tiempo de ejecución. Estos módulos extienden las funcionalidades del kernel sin necesidad de recompilarlo por completo. Permiten agregar características específicas, controladores de dispositivos, sistemas de archivos, entre otros, sin tener que modificar el núcleo del sistema operativo. Esto proporciona flexibilidad y modularidad al sistema Linux, ya que los módulos pueden ser cargados y descargados según sea necesario sin afectar al funcionamiento general del sistema.

En este caso en específico se utilizaron los siguientes:
- Modulo Ram: Este módulo es el encargado de leer la ram total y la libre del sistema
- Modulo CPU: Este modulo lee la utilizacion del cpu y devuelve su porcetaje
- Modulo Proc: El modulo de procesos lee todos los procesos del sistema y los ordena segun su PID, ademas este modulo obtiene todos los child processes (si los hay) y los enlaza con su proceso padre para llevar todo el tracking de estos.

### Notas
Estos módulos estan escritos en C y varian dependiendo de las necesidades solicitas, adeas están acompañados de de un Mkefile que proporciona lo necesario para su uso.

Pasos para instalar los modulos:
- cd a la carpeta del modulo
- cmd "make"
- cmd "Sudo insmod nombre_del_modulo.ko"

Para esto se necesita el secure boot desactivado.

## Docker Images y Compose 🐋
La creación de imágenes en Docker se realiza mediante la elaboración de un Dockerfile que contiene instrucciones para construir la imagen, como la definición del sistema base y la instalación de dependencias. Luego, se ejecuta el comando docker build con el contexto de construcción adecuado, que incluye los archivos necesarios. Este comando procesa el Dockerfile y genera la imagen resultante. Opcionalmente, la imagen puede etiquetarse con un nombre y una versión específica utilizando docker tag, y luego puede ser compartida a través de docker push en un registro de Docker público o privado para su uso y distribución.

Estas imagenes están subidas en nuestro repositorio de DockerHub para su descarga

Ademas se utilizo el compose para la comunicacion entre contenedores Docker:
Docker Compose es una herramienta que permite definir y ejecutar aplicaciones Docker multi-contenedor de manera sencilla mediante un archivo YAML llamado docker-compose.yml. En este archivo, se especifican los servicios de la aplicación, cada uno con su propia configuración, como la imagen de Docker a utilizar, las variables de entorno, los volúmenes, los puertos expuestos y cualquier otra configuración necesaria. Una vez definido el archivo docker-compose.yml, se puede utilizar el comando docker-compose up para iniciar todos los contenedores definidos en el archivo, y docker-compose down para detenerlos y eliminarlos cuando ya no sean necesarios. Docker Compose simplifica enormemente la gestión de aplicaciones complejas con múltiples contenedores, facilitando su desarrollo, despliegue y mantenimiento.

Para descargar una imagen de Docker Hub, puedes utilizar el comando docker pull
```
docker pull nombre_de_la_imagen[:tag]
```
nombre_de_la_imagen es el nombre de la imagen que deseas descargar.[tag] es una etiqueta opcional que especifica la versión de la imagen. Si no se proporciona, se asume la etiqueta latest.

Por ejemplo, si deseas descargar la imagen oficial de NGINX desde Docker Hub, puedes ejecutar el siguiente comando:
```
docker pull nginx
```

Ademas Docker Compose es una herramienta que simplifica la gestión de aplicaciones Docker multi-contenedor. Permite definir y ejecutar aplicaciones con múltiples servicios, cada uno en su propio contenedor, utilizando un archivo YAML llamado docker-compose.yml. En este archivo, puedes especificar la configuración de cada servicio, incluyendo la imagen Docker a utilizar, las variables de entorno, los volúmenes, los puertos expuestos y cualquier otra configuración necesaria. Algunos comandos básicos de Docker Compose incluyen:

 ```
    docker-compose up: Inicia todos los contenedores definidos en el archivo docker-compose.yml.
    docker-compose down: Detiene y elimina todos los contenedores definidos en el archivo docker-compose.yml.
    docker-compose build: Construye o reconstruye los contenedores definidos en el archivo docker-compose.yml.
    docker-compose start: Inicia los contenedores definidos en el archivo docker-compose.yml.
    docker-compose stop: Detiene los contenedores definidos en el archivo docker-compose.yml.
    docker-compose ps: Muestra el estado de los contenedores definidos en el archivo docker-compose.yml.
    docker-compose logs: Muestra los logs de los contenedores definidos en el archivo docker-compose.yml.
```

## NGINX 🔧
NGINX es un servidor web y proxy inverso de alto rendimiento que se utiliza comúnmente para servir contenido estático, balanceo de carga, proxying y como servidor de caché. Se puede utilizar de varias maneras, como:

    Servir sitios web estáticos: NGINX puede servir archivos estáticos como HTML, CSS, JavaScript, imágenes, etc. Simplemente configura NGINX para apuntar al directorio raíz donde se encuentran los archivos y los servirá a los clientes que soliciten esos recursos.

    Proxy inverso: NGINX se puede configurar como un proxy inverso que redirige las solicitudes a diferentes servidores basándose en ciertos criterios, como la URL o el nombre de host. Esto es útil para equilibrar la carga entre varios servidores backend o para ocultar la infraestructura de backend a los clientes.

    Servidor de caché: NGINX puede actuar como un servidor de caché, almacenando en caché las respuestas de los servidores backend para servirlas rápidamente a los clientes. Esto reduce la carga en los servidores backend y mejora el rendimiento del sitio web.

```
nginx: Inicia el servidor NGINX. 

nginx -s stop: Detiene el servidor NGINX de manera segura.

nginx -s reload: Recarga la configuración del servidor NGINX sin detenerlo, lo que permite aplicar cambios en la configuración sin interrumpir el servicio.

nginx -t: Prueba la configuración de NGINX en busca de errores de sintaxis antes de aplicarla.
```