package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/Akshit8/tdm/internal/rest"
	"github.com/ghodss/yaml"
)

func main() {
	var output string

	flag.StringVar(&output, "path", "", "path to generate OpenAPI3 files")
	flag.Parse()

	if output == "" {
		log.Fatalln("path is required")
	}

	swagger := rest.NewOpenAPI3()

	// openapi3.json
	data, err := json.Marshal(&swagger)
	if err != nil {
		log.Fatalln("Couldn't marshal json: %w", err)
	}

	err = os.WriteFile(path.Join(output, "openapi3.json"), data, 0644)
	if err != nil {
		log.Fatalf("Couldn't write json: %s", err)
	}

	// openapi3.yaml
	data, err = yaml.Marshal(&swagger)
	if err != nil {
		log.Fatalln("Couldn't marshal json: %w", err)
	}

	err = os.WriteFile(path.Join(output, "openapi3.yaml"), data, 0644)
	if err != nil {
		log.Fatalln("Couldn't write json: %w", err)
	}

	fmt.Println("open api files generated")
}
