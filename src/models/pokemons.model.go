package models
import ()

// Pokemon représente un Pokémon avec ses attributs de base.
type Pokemon struct {
	ID    int
	Name  string
	Type  string
	Level int
}