import React from "react";

import { useEffect, useState } from "react";
import CreateDoughnutData from "./components/Graph.jsx";
import GraphvizComponent from "./components/ProcTree.jsx";
import "./App.css";
import FloatingLabel from "react-bootstrap/FloatingLabel";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";

function App() {
  const [freeRam, setFreeRam] = useState(0);
  const [freeCpu, setFreeCpu] = useState(0);
  const [dot, setDot] = useState(`digraph G {Procesos}`);
  const [pid, setPid] = useState("");

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch("apigo/api/ram");
        if (!response.ok) {
          throw new Error("Error");
        }
        let data = await response.json();
        const { RamFree, CpuFree } = data;
        setFreeRam(RamFree);
        setFreeCpu(CpuFree);
      } catch (error) {
        console.error("Error al obtener los datos:", error);
      }
    };

    // Llama a la función fetchData inicialmente
    fetchData();

    // Establece un intervalo para llamar a fetchData cada 2 segundos
    const intervalId = setInterval(fetchData, 2000);

    // Limpia el intervalo cuando el componente se desmonta
    return () => clearInterval(intervalId);
  }, []); // La dependencia está vacía, lo que significa que el efecto se ejecuta solo una vez al montar el componente

  const handleChange = (event) => {
    setPid(event.target.value);
  };

  const handleClick = () => {
    console.log("Contenido del FloatingLabel:", pid);
    fetch("apigo/api/proc/" + pid)
      .then((response) => {
        if (!response.ok) {
          throw new Error("Error al obtener los datos");
        }
        // Si la respuesta es exitosa, retorna los datos en formato JSON
        return response.text();
      })
      .then((data) => {
        setDot(data);
        console.log(data);
      })
      .catch((error) => {
        // Maneja los errores de la solicitud
        console.error("Error al obtener los datos:", error);
      });
  };

  return (
    <div className="todo">
      <h1 className="title">SO1_202004796</h1>

      <div className="graficas">
        <CreateDoughnutData name="RAM Usage:" free={freeRam} />
        <CreateDoughnutData name="CPU Usage:" free={freeCpu} />
      </div>

      <h2>Arbol de Procesos</h2>
      <div className="proctree">
        <>
          <FloatingLabel
            controlId="floatingTextarea"
            label="Ingresa un PID"
            className="mb-3"
          >
            <Form.Control
              as="textarea"
              placeholder="Ingresa un PID"
              onChange={handleChange}
            />
          </FloatingLabel>
          <Button variant="primary" onClick={handleClick}>
            Buscar
          </Button>{" "}
        </>
        <GraphvizComponent dot={dot} />
      </div>
    </div>
  );
}

export default App;
