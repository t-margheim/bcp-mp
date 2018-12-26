package canticles

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"

	"github.com/t-margheim/bcp-mp/pkg/calendar"
)

var canticles []Canticle

func init() {
	contents, err := ioutil.ReadFile(fmt.Sprintf("%s/src/github.com/t-margheim/bcp-mp/pkg/canticles/data.json", os.Getenv("GOPATH")))
	if err != nil {
		log.Fatal("failed to read file", err)
	}
	err = json.Unmarshal(contents, &canticles)
	if err != nil {
		log.Fatal("failed to parse json:", err)
	}
}

func Get(keys calendar.KeyChain) []Canticle {

	return []Canticle{canticles[0]}
}

type Canticle struct {
	EnglishTitle string
	LatinTitle   string
	Content      template.HTML
}
