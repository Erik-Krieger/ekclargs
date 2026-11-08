package cli

import (
	"fmt"
	"strings"
)

type FlagSet struct {
	name        string
	subCommands []*FlagSet
	flags       []*Flag
	usage       string
}

func (set *FlagSet) Usage() {
	if set == nil {
		return
	}

	fmt.Println("Subcommands:")
	for _, subCmd := range set.subCommands {
		fmt.Printf("  %s\n", subCmd.name)
		if subCmd.usage != "" {
			fmt.Printf("    %s", subCmd.usage)
		}
	}

	fmt.Println("\n\nFlags:")
	for _, flag := range set.flags {
		names := strings.Join(flag.names, ", -")
		fmt.Printf("  -%s\n", names)
		if flag.usage != "" {
			fmt.Printf("      %s\n", flag.usage)
		}
	}
}
