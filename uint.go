package cli

import (
	"strconv"
)

func UintVar(name string, variable *uint64, defaultValue uint64, usage string) *Flag {
	return uintVarInternal(&mainFlagSet, name, variable, defaultValue, usage)
}

func UintMultiple(name string, collection *[]uint64, usage string) *Flag {
	return uintMultipleInternal(&mainFlagSet, name, collection, usage)
}

func (set *FlagSet) UintVar(name string, variable *uint64, defaultValue uint64, usage string) *Flag {
	return uintVarInternal(set, name, variable, defaultValue, usage)
}

func (set *FlagSet) UintMultiple(name string, collection *[]uint64, usage string) *Flag {
	return uintMultipleInternal(set, name, collection, usage)
}

func uintVarInternal(flagSet *FlagSet, name string, variable *uint64, defaultValue uint64, usage string) *Flag {
	flg := anyVarInternal(flagSet, name, variable, defaultValue, usage, func(s string) (uint64, error) {
		return strconv.ParseUint(s, 10, 64)
	})

	return flg
}

func uintMultipleInternal(flagSet *FlagSet, name string, collection *[]uint64, usage string) *Flag {
	flg := anyMultipleInternal(flagSet, name, collection, usage, func(s string) (uint64, error) {
		return strconv.ParseUint(s, 10, 64)
	})

	return flg
}
