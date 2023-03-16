package yasch

import (
	"os"

	"github.com/alecthomas/jsonschema"
)

// WriteFile Write the structure definition to a JSON schema file.
func WriteFile(data interface{}, path string) {
	d, err := jsonschema.Reflect(data).MarshalJSON()
	if err != nil {
		panic(err)
	}
	if err := os.WriteFile(path, d, 0644); err != nil {
		panic(err)
	}
}
