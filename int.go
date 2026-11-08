package cli

import "strconv"

func IntVar(name string, variable *int64, defaultValue int64, usage string) *Flag {
	return intVarInternal(&mainFlagSet, name, variable, defaultValue, usage)
}

func IntMultiple(name string, collection *[]int64, defaultValue int64, usage string) *Flag {
	return intMultipleInternal(&mainFlagSet, name, collection, usage)
}

func (set *FlagSet) IntVar(name string, variable *int64, defaultValue int64, usage string) *Flag {
	return intVarInternal(set, name, variable, defaultValue, usage)
}

func (set *FlagSet) IntMultiple(name string, collection *[]int64, defaultValue int64, usage string) *Flag {
	return intMultipleInternal(set, name, collection, usage)
}

func intVarInternal(flagSet *FlagSet, name string, variable *int64, defaultValue int64, usage string) *Flag {
	flg := anyVarInternal(flagSet, name, variable, defaultValue, usage, func(s string) (int64, error) {
		return strconv.ParseInt(s, 10, 64)
	})

	return flg
}

func intMultipleInternal(flagSet *FlagSet, name string, collection *[]int64, usage string) *Flag {
	flg := anyMultipleInternal(flagSet, name, collection, usage, func(s string) (int64, error) {
		return strconv.ParseInt(s, 10, 64)
	})

	return flg
}
