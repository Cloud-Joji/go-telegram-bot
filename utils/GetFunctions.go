package utils

import (
	"fmt"
	"os"
	"log"
	"strings"
)

// GetFunctions obtiene una lista de nombres de archivos del directorio "functions"
// dentro del directorio de trabajo actual y devuelve estos nombres de archivo
// como una sola cadena, donde cada nombre de archivo está precedido por "/".
func GetFunctions() string {
	// Slice para almacenar los nombres de archivo
	var fileNames []string

	// Obtener el directorio de trabajo actual
	folder, err := os.Getwd()
	if err != nil {
		// Manejar el error si no se puede obtener el directorio de trabajo
		errorMessage := fmt.Sprintf("Error in folder: %s", err)
		log.Fatal(errorMessage)
	}

	// Agregar "/functions" al final de la ruta del directorio
	folder = folder + "/functions"

	// Leer los archivos en el directorio "functions"
	files, err := os.ReadDir(folder)
	if err != nil {
		// Manejar el error si no se pueden leer los archivos del directorio
		errorMessage := fmt.Sprintf("Error reading dir: %s", err)
		log.Fatal(errorMessage)
	}

	// Iterar sobre los archivos, reemplazar ".go" y agregar "/" al principio del nombre del archivo
	for _, file := range files {
		fileNames = append(fileNames, "/" + strings.Replace(file.Name(), ".go", "", -1))
	}

	// Unir los nombres de archivo en una sola cadena, separados por espacios
	commands := strings.Join(fileNames, "\n")

	// Devolver la cadena resultante con los nombres de archivo modificados
	return commands
}
