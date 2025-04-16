package internals

type YoMamaJoke struct {
	Joke     string `json:"joke,omitzero,omitempty"`
	Category string `json:"category,omitzero,omitempty"`
}
