package cli

func BoolVar(name string, variable *bool, defaultValue bool, usage string) *Flag {
	return boolVarInternal(&mainFlagSet, name, variable, defaultValue, usage)
}

func (set *FlagSet) BoolVar(name string, variable *bool, defaultValue bool, usage string) *Flag {
	return boolVarInternal(set, name, variable, defaultValue, usage)
}

func boolVarInternal(flagSet *FlagSet, name string, variable *bool, defaultValue bool, usage string) *Flag {
	fn := func(cli []string, flg *Flag) error {
		found, err := checkForMatchWithValue(cli, flg.names, variable, defaultValue, nil)
		*variable = found
		return err
	}

	newFlag := &Flag{
		names: []string{name},
		fn:    fn,
		usage: usage,
	}

	flagSet.flags = append(flagSet.flags, newFlag)
	return newFlag
}
