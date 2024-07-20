package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Comando que se ejecutará
	cmd := exec.Command("ls", "-la")

	// Establecer el comando para usar la entrada y salida estándar
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Ejecutar el comando
	if err := cmd.Run(); err != nil {
		fmt.Println("Error:", err)
	}
}
