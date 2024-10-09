package router

import (
	"bytes"
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func reloadHtml(filename string, data *gin.H) (string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	// Parse the template
	dotSplit := strings.Split(filename, ".")
	parts := strings.Split(dotSplit[0], "/")
	lastPart := parts[len(parts)-1]
	log.Println("Read file:", lastPart)
	tmpl, err := template.New(lastPart).Parse(string(file))
	if err != nil {
		return "", err
	}

	// Create a buffer to store the rendered HTML
	var buf bytes.Buffer

	// Execute the template with the data
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return "", err
	}
	bufString := buf.String()
	log.Println("Read file:", bufString)
	return bufString, nil
}
