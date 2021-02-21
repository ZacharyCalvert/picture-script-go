package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	yaml "gopkg.in/yaml.v2"
)

/*
order:
- DONE: ingest
- DONE: determine movie or pic
- reorg /mov|pic/yyyy/mm/dd/<filename> -> avoid collide, copy instead of move?
- skip ignores
- find folder has more than 5 images in single hour, with not yet reviewed
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

	supportedTypes := GetTypeMapping()
	validateAllTypesKnown(supportedTypes, allImages)

	for _, meta := range allImages {
		if meta.Ignore {
			continue
		}
		t, _ := supportedTypes.GetType(meta.Extensions[0])
		earliestDate := time.Unix(0, meta.Date*int64(time.Millisecond))
		target := fmt.Sprintf("%v/%v/%v/%v/%v", t, earliestDate.Year(), earliestDate.Month(), earliestDate.Day(), pathToFname(meta.Paths[0]))
		fmt.Printf("Copying %v to %v\n", meta.StoredAt, target)
	}
}

func validateAllTypesKnown(types *TypeMap, allImages map[string]ImageMeta) {
	// make sure we map the type of file
	for _, val := range allImages {
		if _, ok := types.GetType(val.Extensions[0]); !ok {
			panic(fmt.Sprintf("could not map type %v", strings.ToLower(val.Extensions[0])))
		}
	}
}

func pathToFname(path string) string {
	parts := strings.Split(path, "/")
	return parts[len(parts)-1]
}

// ImageMeta is a single entry in our pic-man.db
type ImageMeta struct {
	Sha        string   `yaml:"sha256"`
	Extensions []string `yaml:"extensions"`
	Paths      []string `yaml:"paths"`
	Date       int64    `yaml:"earliestDate"`
	ReviewDone bool     `yaml:"reviewDone"`
	Ignore     bool     `yaml:"ignore"`
	StoredAt   string   `yaml:"storedAt"`
	Tags       []string `yaml:"tags"`
}

func (meta *ImageMeta) String() string {
	t := time.Unix(0, meta.Date*int64(time.Millisecond))
	return fmt.Sprintf("Meta{sha: %v, pathCount: %v, date: %v, reviewDone: %v, Ignore: %v}", meta.Sha, len(meta.Paths), t, meta.ReviewDone, meta.Ignore)
}
