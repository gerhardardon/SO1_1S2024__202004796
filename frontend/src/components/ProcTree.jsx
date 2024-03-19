import React, { useRef, useEffect } from "react";
import Viz from "viz.js";
import { Module, render } from "viz.js/full.render.js"; // Importa el renderizador completo

function GraphvizComponent({ dot }) {
  const graphRef = useRef(null);

  useEffect(() => {
    if (!graphRef.current) return;

    // Renderiza el gráfico Graphviz
    const viz = new Viz({ Module, render });
    viz
      .renderSVGElement(dot)
      .then((element) => {
        // Limpia cualquier contenido previo
        graphRef.current.innerHTML = "";
        // Agrega el gráfico al elemento ref
        graphRef.current.appendChild(element);
      })
      .catch((error) => {
        console.error("Error al renderizar el gráfico:", error);
      });
  }, [dot]);

  return (
    <>
      
      <div ref={graphRef} />
    </>
  );
}

export default GraphvizComponent;
