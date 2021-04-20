package models

type CorpseStage uint8

// Corpse bodypart enum
const (
	head CorpseStage = 1
	body CorpseStage = 2
	legs CorpseStage = 3
)

type Corpse struct {
	Name  string
	Stage CorpseStage
}

type User struct {
	userName string
}

type ExqCorpModeler interface {
	GetCorpse(name string) (*Corpse, error)
	Close() error
}
