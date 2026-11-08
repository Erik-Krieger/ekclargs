package cli

import (
	"fmt"
	"os"
	"strings"
)

func Parse() error {
	remainingCli := os.Args[1:]
	return parseInner(&mainFlagSet, remainingCli)
}

func parseInner(flagSet *FlagSet, cli []string) error {
	if len(cli) < 1 {
		return nil
	}

	if !strings.HasPrefix(cli[0], "-") {
		for _, subCmd := range flagSet.subCommands {
			cli = cli[1:]
			return parseInner(subCmd, cli)
		}

		return fmt.Errorf("invalid param")
	}

	for _, flg := range flagSet.flags {
		err := flg.fn(cli, flg)
		if err != nil {
			return err
		}
	}

	return nil
}
