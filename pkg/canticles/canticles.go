package canticles

import (
	"html/template"

	"github.com/t-margheim/bcp-mp/pkg/calendar"
)

// func init() {
// 	path := os.Getenv("CANTICLE_PATH")
// 	if path == "" {
// 		path = fmt.Sprintf("%s/src/github.com/t-margheim/bcp-mp/pkg/canticles", os.Getenv("GOPATH"))
// 	}
// 	contents, err := ioutil.ReadFile(fmt.Sprintf("%s/data.json", path))
// 	if err != nil {
// 		log.Fatal("failed to read file", err)
// 	}
// 	err = json.Unmarshal(contents, &canticles)
// 	if err != nil {
// 		log.Fatal("failed to parse json:", err)
// 	}
// }

func Get(keys calendar.KeyChain) []Canticle {
	halfLength := len(canticles) / 2
	index1 := keys.Iterator % (halfLength)
	index2 := index1 + halfLength
	return []Canticle{canticles[index1], canticles[index2]}
}

type Canticle struct {
	EnglishTitle string
	LatinTitle   string
	Content      template.HTML
}
