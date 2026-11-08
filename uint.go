package cli

import (
	"strconv"
)

func (set *FlagSet) UintVar(name string, variable *uint64, defaultValue uint64, usage string) *Flag {
	return uintVarInternal(set, name, variable, defaultValue, usage)
}

func (set *FlagSet) UintMultiple(name string, collection *[]uint64, usage string) *Flag {
	return uintMultipleInternal(set, name, collection, usage)
}

func uintVarInternal(flagSet *FlagSet, name string, variable *uint64, defaultValue uint64, usage string) *Flag {
	return anyVarInternal(flagSet, name, variable, defaultValue, usage, func(s string) (uint64, error) {
		return strconv.ParseUint(s, 10, 64)
	})
}

func uintMultipleInternal(flagSet *FlagSet, name string, collection *[]uint64, usage string) *Flag {
	return anyMultipleInternal(flagSet, name, collection, usage, func(s string) (uint64, error) {
		return strconv.ParseUint(s, 10, 64)
	})
}
