package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

/*
order:
- DONE: ingest
- DONE: determine movie or pic
- reorg /mov|pic/yyyy/mm/dd/<filename> -> avoid collide, copy instead of move?
- skip ignores
*/

func main() {
	var loadDatabase string

	flag.StringVar(&loadDatabase, "db", "./pic-man.db", "The full path for pic-man database")

	if _, err := os.Stat(loadDatabase); err == nil {
		fmt.Printf("Will load database from %v\n", loadDatabase)
	} else {
		panic(fmt.Errorf("missing %v - pic manager database", loadDatabase))
	}

	data, err := ioutil.ReadFile(loadDatabase)
	if err != nil {
		panic(fmt.Errorf("could not read %v: %w", loadDatabase, err))
	}

	allImages := make(map[string]ImageMeta)

	err = yaml.Unmarshal([]byte(data), &allImages)
	if err != nil {
		panic(fmt.Errorf("failed to load %v: %w", loadDatabase, err))
	}
	fmt.Printf("Loaded %v meta records", len(allImages))

	supportedTypes := getTypeMapping()

	// make sure we map the type
	for _, val := range allImages {
		if _, ok := supportedTypes[strings.ToLower(val.Extensions[0])]; !ok {
			panic(fmt.Sprintf("could not map type %v", strings.ToLower(val.Extensions[0])))
		}
	}
}

func getTypeMapping() map[string]string {
	supportedTypes := make(map[string]string)
	supportedTypes["m4v"] = "mov"
	supportedTypes["mp4"] = "mov"
	supportedTypes["png"] = "pic"
	supportedTypes["gif"] = "pic"
	supportedTypes["bmp"] = "pic"
	supportedTypes["jpeg"] = "pic"
	supportedTypes["jpg"] = "pic"
	supportedTypes["mov"] = "mov"
	supportedTypes["cr2"] = "mov"
	supportedTypes["avi"] = "mov"
	supportedTypes["mpg"] = "mov"
	return supportedTypes
}

// ImageMeta is a single entry in our pic-man.db
type ImageMeta struct {
	Sha        string   `yaml:"sha256"`
	Extensions []string `yaml:"extensions"`
	Paths      []string `yaml:"paths"`
	Date       uint64   `yaml:"earliestDate"`
	Review     bool     `yaml:"reviewDone"`
	Ignore     bool     `yaml:"ignore"`
}

func (meta ImageMeta) String() string {
	return fmt.Sprintf("Meta{sha: %v, pathCount: %v, date: %v, reviewDone: %v, Ignore: %v}")
}
