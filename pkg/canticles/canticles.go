package canticles

import (
	"html/template"
)

func Get(iterator int) []Canticle {
	halfLength := len(canticles) / 2
	index1 := iterator % (halfLength)
	index2 := index1 + halfLength
	return []Canticle{canticles[index1], canticles[index2]}
}

type Canticle struct {
	EnglishTitle string
	LatinTitle   string
	Content      template.HTML
}
