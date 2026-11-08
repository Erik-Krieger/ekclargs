package cli

import (
	"fmt"
	"slices"
)

type Flag struct {
	names []string
	fn    func([]string, *Flag) error
	usage string
}

func (f *Flag) AddAlias(name string) error {
	if slices.Contains(f.names, name) {
		return fmt.Errorf("The flag alias \"%s\" is already defined.", name)
	}

	f.names = append(f.names, name)
	return nil
}
