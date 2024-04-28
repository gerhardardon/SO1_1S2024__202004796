import json

from random import randrange
from locust import HttpUser, task, between

debug = True

## Función para imprimir mensajes de debug
def printDebug(msg):
    if debug:
        print(msg)

## Clase para leer los datos del archivo ------------------------------------------------------
class Reader():
    ## array de datos json
    def __init__(self) -> None:
        self.datos = []

    ## Función para cargar los datos del archivo
    def load(self):
        print("-leyendo data")
        try:
            with open("datos.json", "r") as data_file:
                self.datos = json.loads(data_file.read())
        except Exception as error:
            print(f'-err: {error}')

    ## Función para obtener un valor aleatorio del archivo
    def pickRandom(self):
        length = len(self.datos)

        if ( length > 0 ):
            random_index = randrange(0, length - 1) if length > 1 else 0
            return self.datos.pop(random_index)
        else:
            print("-err: json vacio")
            return None


## Uso de Locust -------------------------------------------------------------------------------
class MessageTraffic(HttpUser):
    wait_time = between(0.1, 0.9)
    reader = Reader()
    reader.load()

    def on_start(self):
        print("-Locust: Inicio de envío de tráfico")

    @task
    def PostMessage(self):
        random_data = self.reader.pickRandom()

        if ( random_data is not None ):
            data_to_send = json.dumps(random_data)
            printDebug(data_to_send)
            self.client.post("/", json=random_data)
        else:
            print("-Locust: Envío finalizado")
            self.stop(True)