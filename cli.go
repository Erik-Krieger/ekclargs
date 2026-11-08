package cli

import (
	"fmt"
	"strings"
)

func New() *FlagSet {
	return &FlagSet{
		name:        "",
		subCommands: make([]*FlagSet, 0),
		flags:       make([]*Flag, 0),
	}
}

func (set *FlagSet) Parse(args []string) (*FlagSet, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("The provided args should be a non nil non empty value")
	}

	remainingCli := args[1:]
	return parseInner(set, remainingCli)
}

func parseInner(flagSet *FlagSet, cli []string) (*FlagSet, error) {
	if len(cli) < 1 {
		return flagSet, nil
	}

	if !strings.HasPrefix(cli[0], "-") {
		for _, subCmd := range flagSet.subCommands {
			if subCmd.name == cli[0] {
				cli = cli[1:]
				return parseInner(subCmd, cli)
			}
		}

		return flagSet, fmt.Errorf("invalid param")
	}

	for _, flg := range flagSet.flags {
		err := flg.fn(cli, flg)
		if err != nil {
			return flagSet, err
		}
	}

	return flagSet, nil
}
