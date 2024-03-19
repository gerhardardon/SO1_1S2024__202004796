import React from "react";
import { Chart, ArcElement, Tooltip, Legend, Title } from "chart.js";
import { Doughnut } from "react-chartjs-2";
import Card from "react-bootstrap/Card";

Chart.register(ArcElement, Tooltip, Legend, Title);
Chart.defaults.plugins.tooltip.backgroundColor = "rgb(0, 0, 156)";
Chart.defaults.plugins.legend.position = "right";
Chart.defaults.plugins.legend.title.display = true;
Chart.defaults.plugins.legend.title.font = "system-ui";

const options = {
  responsive: true, // Hace que el gráfico sea responsive
  maintainAspectRatio: true, // Permite que el gráfico ajuste su tamaño automáticamente

  // Otras opciones personalizadas según tus necesidades
};

function CreateDoughnutData(props) {
  const data = {
    labels: ["free", "used"],
    datasets: [
      {
        data: [props.free, 100 - props.free],
        backgroundColor: ["rgb(116, 226, 192)", "rgb(255, 87, 51)"],

        borderWidth: 4,
        radius: "80%",
      },
    ],
  };

  return (
    <Card style={{ width: "500px",height: "500px"}}>
      <Card.Body>
        <Card.Title>{props.name}</Card.Title>
        
          <div style={{ width: "80%", height: "400px", margin: "auto" }}>
            <Doughnut data={data} options={options} />
            <h5>libre: {props.free.toFixed(2)} %</h5>
            <h5>usado: {(100 - props.free).toFixed(2)} %</h5>
          </div>
        
      </Card.Body>
    </Card>
  );
}

export default CreateDoughnutData;
