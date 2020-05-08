package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {
	// os.Args[1] is already used as config name
	manifestFile := os.Args[2]

	manifest, err := ioutil.ReadFile(manifestFile)
	if err != nil {
		panic(err)
	}

	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(manifest, &m)
	if err != nil {
		panic(err)
	}

	// Overwrite metadata.name to metadata.generateName
	n := m["metadata"].(map[interface{}]interface{})
	n["generateName"] = m["metadata"].(map[interface{}]interface{})["name"]

	// Delete metadata.name key
	delete(n, "name")

	generatedManifest, err := yaml.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", generatedManifest)
}
