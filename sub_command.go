package cli

func AddSubCommand(name string) *FlagSet {
	return addSubCommand(&mainFlagSet, name)
}

func (set *FlagSet) AddSubCommand(name string) *FlagSet {
	return addSubCommand(set, name)
}

func addSubCommand(parentCommand *FlagSet, name string) *FlagSet {
	subCommand := &FlagSet{
		name:        name,
		subCommands: make([]*FlagSet, 0),
		flags:       make([]*Flag, 0),
	}

	parentCommand.subCommands = append(parentCommand.subCommands, subCommand)

	return subCommand
}
