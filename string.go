package cli

import (
	"fmt"
	"slices"
)

func (set *FlagSet) StringVar(name string, variable *string, defaultValue string, usage string) *Flag {
	return stringVarInternal(set, name, variable, defaultValue, usage)
}

func (set *FlagSet) StringVarValidSet(name string, variable *string, defaultValue string, usage string, validSet []string) *Flag {
	return stringVarValidSetInternal(set, name, variable, defaultValue, usage, validSet)
}

func (set *FlagSet) StringMultiple(name string, collection *[]string, usage string) *Flag {
	return stringMultipleInternal(set, name, collection, usage)
}

func stringVarInternal(flagSet *FlagSet, name string, variable *string, defaultValue string, usage string) *Flag {
	flg := anyVarInternal(flagSet, name, variable, defaultValue, usage, func(s string) (string, error) {
		return s, nil
	})

	return flg
}

func stringVarValidSetInternal(flagSet *FlagSet, name string, variable *string, defaultValue string, usage string, validValues []string) *Flag {
	flg := anyVarInternal(flagSet, name, variable, defaultValue, usage, func(s string) (string, error) {
		if !slices.Contains(validValues, s) {
			return "", fmt.Errorf("Provided value: \"%s\" not part of the set of valid values", s)
		}

		return s, nil
	})

	return flg
}

func stringMultipleInternal(flagSet *FlagSet, name string, collection *[]string, usage string) *Flag {
	flg := anyMultipleInternal(flagSet, name, collection, usage, func(s string) (string, error) {
		return s, nil
	})

	return flg
}
