package flagit

import (
	"fmt"
	"strconv"
)

type boolFlag struct {
	value *bool
	flag  string
}

type intFlag struct {
	value *int
	flag  string
}

type stringFlag struct {
	value *string
	flag  string
}

type Flag struct {
	boolFlags   []boolFlag
	intFlags    []intFlag
	stringFlags []stringFlag
}

func CreateFlag() *Flag {
	f := new(Flag)
	f.boolFlags = make([]boolFlag, 0)
	f.intFlags = make([]intFlag, 0)
	f.stringFlags = make([]stringFlag, 0)
	return f
}

func (f *Flag) AddBoolFlag(v *bool, flags ...string) {
	for i := range flags {
		f.boolFlags = append(f.boolFlags, boolFlag{v, flags[i]})
	}
}

func (f *Flag) AddIntFlag(v *int, flags ...string) {
	for i := range flags {
		f.intFlags = append(f.intFlags, intFlag{v, flags[i]})
	}
}

func (f *Flag) AddStringFlag(v *string, flags ...string) {
	for i := range flags {
		f.stringFlags = append(f.stringFlags, stringFlag{v, flags[i]})
	}
}

func sanitizeFlags(s []string) []string {
	a := make([]string, 0)
	for i := range s {
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

	return a
}

func (f Flag) ParseFlags(flags []string) error {
	found := false
	var err error
	availFlags := sanitizeFlags(flags)
	for i := range availFlags {
		found = false
		if !found {
			for j := range f.intFlags {
				if availFlags[i] == f.intFlags[j].flag {
					if (i + 1) < len(availFlags) {
						*f.intFlags[j].value, err = strconv.Atoi(availFlags[i+1])
						if err != nil {
							return fmt.Errorf("%s is not satisfied by %s", f.intFlags[j].flag, availFlags[i+1])
						}
						i += 2
						found = true
					} else {
						return fmt.Errorf("No value passed into flag %s", f.intFlags[j].flag)
					}
				}
				if found {
					break
				}
			}
		}
		if !found {
			for j := range f.stringFlags {
				if availFlags[i] == f.stringFlags[j].flag {
					if (i + 1) < len(availFlags) {
						*f.stringFlags[j].value = availFlags[i+1]
						found = true
						i += 2
					} else {
						return fmt.Errorf("No string passed into flag %s", f.stringFlags[j].flag)
					}
				}

				if found {
					break
				}
			}
		}
		if !found {
			for j := range f.boolFlags {
				if availFlags[i] == f.boolFlags[j].flag {
					*f.boolFlags[j].value = true
					found = true
				}
				if found {
					break
				}
			}
		}
	}
	return nil
}
