// Package main checks YAML files for a given JSON schema.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"

	validator "github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v3"
)

func main() {
	var (
		schemaFile string
		files      []string
	)
	flag.StringVar(&schemaFile, "s", "", "schema file path")
	flag.Parse()
	files = flag.Args()

	if len(files) == 0 {
		flag.Usage()
		log.Fatal("no YAML files to check")
	}

	p, err := filepath.Abs(schemaFile)
	if err != nil {
		log.Panic(err)
	}

	refURL := &url.URL{Scheme: "file", Path: p}
	schemaLoader := validator.NewReferenceLoader(refURL.String())

	for _, file := range files {
		log.Println("checking", file)

		d, err := os.ReadFile(file)
		if err != nil {
			log.Panic(err)
		}

		m := make(map[string]interface{})
		if err := yaml.Unmarshal(d, &m); err != nil {
			log.Panic(err)
		}

		ret, err := validator.Validate(schemaLoader, validator.NewGoLoader(m))
		if err != nil {
			log.Panic(err)
		}
		if ret.Valid() {
			continue
		}
		for _, err := range ret.Errors() {
			log.Printf("%s\n", err)
		}
		log.Panic(fmt.Sprintf("invalid file: %s", file))
	}
}
