<template>
  <div class="container">
    <h1 class="title">SO1-PROYECTO2-202004796</h1>
    <textarea v-model="resultados" rows="20" cols="50" autocorrect="off" class="textarea"></textarea>
  </div>
</template>

<script>
export default {
  data() {
    return {
      resultados: '' // Inicialmente vacío
    };
  },
  created() {
    // Llamamos al método fetchData cada 5 segundos
    setInterval(this.fetchData, 5000);
  },
  methods: {
    async fetchData() {
      try {
        // Limpiamos el contenido antiguo del textarea
        this.resultados = '';

        const response = await fetch('http://localhost:8081/records'); // Reemplaza 'URL_DE_TU_API' con la URL real de tu API
        const data = await response.json();
        
        // Convertimos los datos a JSON y los agregamos al textarea, uno por fila
        this.resultados = data.map(item => JSON.stringify(item)).join('\n');
      } catch (error) {
        console.error('Error al obtener datos:', error);
      }
    }
  }
}
</script>

<style>
body {
  background-color: #333;
  color: #ccc;
}

.container {
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
}

.title {
  font-size: 40px;
  font-weight: bold;
  text-align: center;
  margin-bottom: 20px;
  color: #ddd;
}

.textarea {
  width: 100%;
  padding: 10px;
  font-size: 14px;
  border: 1px solid #ccc;
  border-radius: 5px;
  background-color: #444;
  color: #ccc;
}
</style>
