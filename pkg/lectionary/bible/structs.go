package bible

type resp struct {
	Canonical string   `json:"canonical"`
	Passages  []string `json:"passages"`
}
