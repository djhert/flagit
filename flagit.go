package flagit

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	ErrNoFlags = errors.New("No flags passed")
)

type Flag struct {
	boolFlags      []boolFlag
	intFlags       []intFlag
	stringFlags    []stringFlag
	availableFlags []flagBase
}

func NewFlag() *Flag {
	f := new(Flag)
	f.boolFlags = make([]boolFlag, 0)
	f.intFlags = make([]intFlag, 0)
	f.stringFlags = make([]stringFlag, 0)
	return f
}

func (f *Flag) Bool(v *bool, flags []string, usage string) {
	f.boolFlags = append(f.boolFlags, createBoolFlag(flags, usage, v))
	f.addAvailable(flags, usage)
}

func (f *Flag) Int(v *int, flags []string, usage string) {
	f.intFlags = append(f.intFlags, createIntFlag(flags, usage, v))
	f.addAvailable(flags, usage)
}

func (f *Flag) String(v *string, flags []string, usage string) {
	f.stringFlags = append(f.stringFlags, createStringFlag(flags, usage, v))
	f.addAvailable(flags, usage)
}

func (f *Flag) addAvailable(flags []string, usage string) {
	f.availableFlags = append(f.availableFlags, flagBase{flags, usage})
}

func (f Flag) PrintUsage() {
	fmt.Println("Flags: ")
	for i := range f.availableFlags {
		fmt.Println(f.availableFlags[i].GetFlagAndUsage())
	}
}

func (f Flag) PrintUsageOf(s string) {
	fmt.Println("Usage: ")
	for i := range f.availableFlags {
		if f.availableFlags[i].checkFlag(s) {
			fmt.Println(f.availableFlags[i].GetFlagAndUsage())
			return
		}
	}
	fmt.Printf("  Error: %s is not a valid flag\n", s)
}

func sanitizeFlags(s []string) []string {
	a := make([]string, 0)
	for i := range s {
		if len(s[i]) > 0 {
			if s[i][:1] == "-" && s[i][:2] != "--" {
				if len(s[i]) == 2 {
					a = append(a, s[i])
				} else {
					a = append(a, s[i][:2])
					temp := s[i][2:]
					x := 0
					for y := x + 1; y <= len(temp); y++ {
						a = append(a, "-"+temp[x:y])
						x++
					}
				}
			} else {
				a = append(a, s[i])
			}
		}
	}
	return a
}

func (f Flag) Parse(flags []string) ([]string, error) {
	found := false
	var err error
	if len(flags) < 1 {
		return nil, ErrNoFlags
	}
	data := make([]string, 0)
	availFlags := sanitizeFlags(flags)
	for i := 0; i < len(availFlags); i++ {
		found = false
		if availFlags[i][:1] != "-" {
			data = append(data, availFlags[i])
			found = true
		}
		if !found {
			for j := range f.intFlags {
				if f.intFlags[j].checkFlag(availFlags[i]) {
					if (i + 1) < len(availFlags) {
						*f.intFlags[j].value, err = strconv.Atoi(availFlags[i+1])
						if err != nil {
							return nil, fmt.Errorf("%s is not satisfied by %s", availFlags[i], availFlags[i+1])
						}
						i++
						found = true
					} else {
						return nil, fmt.Errorf("No value passed into flag %s", availFlags[i])
					}
				}
				if found {
					break
				}
			}
		}
		if !found {
			for j := range f.stringFlags {
				if f.stringFlags[j].checkFlag(availFlags[i]) {
					if (i + 1) < len(availFlags) {
						*f.stringFlags[j].value = availFlags[i+1]
						found = true
						i++
					} else {
						return nil, fmt.Errorf("No string passed into flag %s", availFlags[i])
					}
				}
				if found {
					break
				}
			}
		}
		if !found {
			for j := range f.boolFlags {
				if f.boolFlags[j].checkFlag(availFlags[i]) {
					*f.boolFlags[j].value = true
					found = true
				}
				if found {
					break
				}
			}
		}
		if !found {
			return nil, fmt.Errorf("Invalid Flag: %s", availFlags[i])
		}
	}
	return data, nil
}
