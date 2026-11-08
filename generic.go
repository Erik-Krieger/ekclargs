package cli

import (
	"fmt"
	"slices"
	"strings"
)

func anyVarInternal[T any](flagSet *FlagSet, name string, variable *T, defaultValue T, usage string, convFunc func(s string) (T, error)) *Flag {
	fn := func(cli []string, flg *Flag) error {
		_, err := checkForMatchWithValue(cli, flg.names, variable, defaultValue, convFunc)
		return err
	}

	// It doesn't make sense to add an empty flag.
	if name == "" {
		return nil
	}

	newFlag := &Flag{
		names: []string{name},
		fn:    fn,
		usage: usage,
	}

	flagSet.flags = append(flagSet.flags, newFlag)
	return newFlag
}

func anyMultipleInternal[T any](flagSet *FlagSet, name string, collection *[]T, usage string, convFunc func(s string) (T, error)) *Flag {
	// It doesn't make sense to add an empty flag.
	if name == "" {
		return nil
	}

	fn := func(cli []string, flg *Flag) error {
		if collection == nil {
			return fmt.Errorf("The passed in collection was nil")
		}

		if convFunc == nil {
			return fmt.Errorf("The passed in convFunc was nil")
		}

		skip := false

		for idx, entry := range cli {
			if skip {
				skip = false
				continue
			}

			entry = strings.TrimLeft(entry, "-")

			if !slices.Contains(flg.names, entry) {
				continue
			}

			if idx+1 >= len(cli) {
				return fmt.Errorf("Missing parameter for flag: \"%s\"", entry)
			}

			val, err := convFunc(cli[idx+1])
			if err != nil {
				return err
			}

			*collection = append(*collection, val)
		}

		return nil
	}

	newFlag := &Flag{
		names: []string{name},
		fn:    fn,
		usage: usage,
	}

	flagSet.flags = append(flagSet.flags, newFlag)
	return newFlag
}

func checkForMatchWithValue[T any](cli []string, names []string, variable *T, defaultValue T, convFunc func(s string) (T, error)) (bool, error) {
	for idx, entry := range cli {
		entry = strings.TrimLeft(entry, "-")

		if slices.Contains(names, entry) {
			// handle the case for presence flags
			if convFunc == nil {
				return true, nil
			}

			if idx >= len(cli)-1 {
				break
			}

			val, err := convFunc(cli[idx+1])
			if err != nil {
				return false, err
			}

			*variable = val
			return true, nil
		}
	}

	*variable = defaultValue
	return false, nil
}
