package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"proy1/proctree"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	//Crear nuestra aplicación de Fiber
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Content-Type, Authorization",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/api/ram", func(c *fiber.Ctx) error {
		type data struct {
			RamFree float64 `json:"RamFree"`
			CpuFree float64 `json:"CpuFree"`
		}

		ram, cpu := getData()
		return c.JSON(data{ram, cpu})
	})

	app.Get("/api/proc/:pid", func(c *fiber.Ctx) error {
		parametro1 := c.Params("pid")
		parametro1Int, err := strconv.Atoi(parametro1)
		if err != nil {
			log.Fatal(err)
		}

		tree := proctree.GetProc(parametro1Int)
		return c.SendString(tree)
	})

	app.Get("/api/proclist", func(c *fiber.Ctx) error {

		return c.SendString(proctree.GetProcList())
	})

	log.Fatal(app.Listen(":3000"))

}

func getData() (float64, float64) {
	//leer ram
	data, err := os.ReadFile("/proc/ram_202004796")
	if err != nil {
		log.Fatal(err)
	}
	ram := strings.Split(string(data), "  ")
	ramValue1, _ := strconv.ParseFloat(ram[0], 64)
	ramValue2, _ := strconv.ParseFloat(ram[1], 64)
	freeRam := ramValue1 * 100 / ramValue2
	fmt.Println("free RAM", freeRam, "%")

	//leer cpu
	// Ejecutar mpstat
	cmd := exec.Command("mpstat")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error al ejecutar mpstat:", err)
	}

	// Convertir la salida a una cadena
	outputStr := string(output)

	// Dividir la salida en líneas
	lines := strings.Split(outputStr, "\n")
	fields := strings.Fields(lines[len(lines)-2])
	cpu := fields[len(fields)-1]
	freeCpu, _ := strconv.ParseFloat(cpu, 64)

	fmt.Println("free CPU", freeCpu, "%")
	return freeRam, freeCpu
}
