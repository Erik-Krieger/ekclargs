package cli

func StringVar(name string, variable *string, defaultValue string, usage string) *Flag {
	return stringVarInternal(&mainFlagSet, name, variable, defaultValue, usage)
}

func StringMultiple(name string, collection *[]string, usage string) *Flag {
	return stringMultipleInternal(&mainFlagSet, name, collection, usage)
}

func (set *FlagSet) StringVar(name string, variable *string, defaultValue string, usage string) *Flag {
	return stringVarInternal(set, name, variable, defaultValue, usage)
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

func stringMultipleInternal(flagSet *FlagSet, name string, collection *[]string, usage string) *Flag {
	flg := anyMultipleInternal(flagSet, name, collection, usage, func(s string) (string, error) {
		return s, nil
	})

	return flg
}
