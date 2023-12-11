package definition_collector

import "fmt"

type Definition struct {
	Definition string
	Word       string
	Strength   int
}

func (s Definition) String() string {
	return fmt.Sprintf("{\n\tdefinition: %v,\n\tword: %v,\n\tstrength: %v\n}", s.Definition, s.Word, s.Strength)
}
