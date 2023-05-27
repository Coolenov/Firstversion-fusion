package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {

	pythonScript := "./cmd/serviceUaNews/api.py"
	cmd := exec.Command("python", pythonScript)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Файл Python успешно выполнен.")
}
