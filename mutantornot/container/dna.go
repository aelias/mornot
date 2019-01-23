package container

// DnaMatrix represent a human or mutan DNA
type DnaMatrix struct {
	Dna []string `json:"dna" binding:"required"`
}
