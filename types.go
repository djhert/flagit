package flagit

import (
	"fmt"
)

type flagBase struct {
	flags []string
	usage string
}

func (f flagBase) checkFlag(s string) bool {
	for i := range f.flags {
		if s == f.flags[i] {
			return true
		}
	}
	return false
}

func (f flagBase) GetFlag() string {
	if len(f.flags) < 1 {
		return "Missing flag, something major went wrong\nPlease file a bug report"
	}
	s := fmt.Sprintf("  %s", f.flags[0])
	for i := 1; i < len(f.flags); i++ {
		s = fmt.Sprintf("%s, %s", s, f.flags[i])
	}
	return s
}

func (f flagBase) GetFlagAndUsage() string {
	return fmt.Sprintf("%-25s\t%-s", f.GetFlag(), f.usage)
}

type boolFlag struct {
	flagBase
	value *bool
}

func createBoolFlag(flags []string, usage string, v *bool) boolFlag {
	b := boolFlag{}
	b.flags = flags
	b.usage = usage
	b.value = v
	return b
}

type intFlag struct {
	flagBase
	value *int
}

func createIntFlag(flags []string, usage string, v *int) intFlag {
	i := intFlag{}
	i.flags = flags
	i.usage = usage
	i.value = v
	return i
}

type stringFlag struct {
	flagBase
	value *string
}

func createStringFlag(flags []string, usage string, v *string) stringFlag {
	s := stringFlag{}
	s.flags = flags
	s.usage = usage
	s.value = v
	return s
}
