package container

// Stat container
type Stat struct {
	Mutants int     `json:"count_mutant_dna"`
	Humans  int     `json:"count_human_dna"`
	Ratio   float64 `json:"ratio"`
}
