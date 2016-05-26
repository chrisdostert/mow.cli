package cli

import (
	"flag"
	"fmt"
	"strings"
)

// BoolOpt describes a boolean option
type BoolOpt struct {
	*boolParam

	// A space separated list of the option names *WITHOUT* the dashes, e.g. `f force` and *NOT* `-f --force`.
	// The one letter names will then be called with a single dash (short option), the others with two (long options).
	Name string
	// The option description as will be shown in help messages
	Desc string
	// A space separated list of environment variables names to be used to initialize this option
	EnvVar string
	// The option's inital value
	Value bool
	// A boolean to display or not the current value of the option in the help message
	HideValue bool
}

// StringOpt describes a string option
type StringOpt struct {
	*stringParam

	// A space separated list of the option names *WITHOUT* the dashes, e.g. `f force` and *NOT* `-f --force`.
	// The one letter names will then be called with a single dash (short option), the others with two (long options).
	Name string
	// The option description as will be shown in help messages
	Desc string
	// A space separated list of environment variables names to be used to initialize this option
	EnvVar string
	// The option's inital value
	Value string
	// A boolean to display or not the current value of the option in the help message
	HideValue bool
}

// IntOpt describes an int option
type IntOpt struct {
	*intParam

	// A space separated list of the option names *WITHOUT* the dashes, e.g. `f force` and *NOT* `-f --force`.
	// The one letter names will then be called with a single dash (short option), the others with two (long options).
	Name string
	// The option description as will be shown in help messages
	Desc string
	// A space separated list of environment variables names to be used to initialize this option
	EnvVar string
	// The option's inital value
	Value int
	// A boolean to display or not the current value of the option in the help message
	HideValue bool
}

// StringsOpt describes a string slice option
type StringsOpt struct {
	*stringsParam

	// A space separated list of the option names *WITHOUT* the dashes, e.g. `f force` and *NOT* `-f --force`.
	// The one letter names will then be called with a single dash (short option), the others with two (long options).
	Name string
	// The option description as will be shown in help messages
	Desc string
	// A space separated list of environment variables names to be used to initialize this option.
	// The env variable should contain a comma separated list of values
	EnvVar string
	// The option's inital value
	Value []string
	// A boolean to display or not the current value of the option in the help message
	HideValue bool
}

// IntsOpt describes an int slice option
type IntsOpt struct {
	*intsParam

	// A space separated list of the option names *WITHOUT* the dashes, e.g. `f force` and *NOT* `-f --force`.
	// The one letter names will then be called with a single dash (short option), the others with two (long options).
	Name string
	// The option description as will be shown in help messages
	Desc string
	// A space separated list of environment variables names to be used to initialize this option.
	// The env variable should contain a comma separated list of values
	EnvVar string
	// The option's inital value
	Value []int
	// A boolean to display or not the current value of the option in the help message
	HideValue bool
}

// VarOpt describes a user-settable option
type VarOpt struct {
	VarParam

	// A space separated list of the option names *WITHOUT* the dashes, e.g. `f force` and *NOT* `-f --force`.
	// The one letter names will then be called with a single dash (short option), the others with two (long options).
	Name string
	// The option description as will be shown in help messages
	Desc string
	// A space separated list of environment variables names to be used to initialize this option
	EnvVar string

	Value flag.Value
	// A boolean to display or not the current value of the option in the help message
	HideValue bool
}

func (vo *VarOpt) Set(s string) error {
	return vo.Value.Set(s)
}

func (vo *VarOpt) String() string {
	return vo.Value.String()
}

func (vo *VarOpt) IsBoolFlag() bool {
	if bf, ok := vo.Value.(boolValued); ok {
		return bf.IsBoolFlag()
	}

	return false
}

func (vo *VarOpt) IsMultiValued() bool {
	if bf, ok := vo.Value.(multiValued); ok {
		return bf.IsMultiValued()
	}

	return false
}

func (vo *VarOpt) SetMulti(vs []string) error {
	mv, ok := vo.Value.(multiValued)
	if !ok || !mv.IsMultiValued() {
		panic("Bug")
	}

	return mv.SetMulti(vs)
}

/*
BoolOpt defines a boolean option on the command c named `name`, with an initial value of `value` and a description of `desc` which will be used in help messages.

The name is a space separated list of the option names *WITHOUT* the dashes, e.g. `f force` and *NOT* `-f --force`.
The one letter names will then be called with a single dash (short option), the others with two (long options).


The result should be stored in a variable (a pointer to a bool) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) BoolOpt(name string, value bool, desc string) *bool {
	return c.Bool(BoolOpt{
		Name:  name,
		Value: value,
		Desc:  desc,
	})
}

/*
StringOpt defines a string option on the command c named `name`, with an initial value of `value` and a description of `desc` which will be used in help messages.

The name is a space separated list of the option names *WITHOUT* the dashes, e.g. `f force` and *NOT* `-f --force`.
The one letter names will then be called with a single dash (short option), the others with two (long options).


The result should be stored in a variable (a pointer to a string) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) StringOpt(name string, value string, desc string) *string {
	return c.String(StringOpt{
		Name:  name,
		Value: value,
		Desc:  desc,
	})
}

/*
IntOpt defines an int option on the command c named `name`, with an initial value of `value` and a description of `desc` which will be used in help messages.

The name is a space separated list of the option names *WITHOUT* the dashes, e.g. `f force` and *NOT* `-f --force`.
The one letter names will then be called with a single dash (short option), the others with two (long options).


The result should be stored in a variable (a pointer to an int) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) IntOpt(name string, value int, desc string) *int {
	return c.Int(IntOpt{
		Name:  name,
		Value: value,
		Desc:  desc,
	})
}

/*
StringsOpt defines a string slice option on the command c named `name`, with an initial value of `value` and a description of `desc` which will be used in help messages.

The name is a space separated list of the option names *WITHOUT* the dashes, e.g. `f force` and *NOT* `-f --force`.
The one letter names will then be called with a single dash (short option), the others with two (long options).


The result should be stored in a variable (a pointer to a string slice) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) StringsOpt(name string, value []string, desc string) *[]string {
	return c.Strings(StringsOpt{
		Name:  name,
		Value: value,
		Desc:  desc,
	})
}

/*
IntsOpt defines an int slice option on the command c named `name`, with an initial value of `value` and a description of `desc` which will be used in help messages.

The name is a space separated list of the option names *WITHOUT* the dashes, e.g. `f force` and *NOT* `-f --force`.
The one letter names will then be called with a single dash (short option), the others with two (long options).


The result should be stored in a variable (a pointer to an int slice) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) IntsOpt(name string, value []int, desc string) *[]int {
	return c.Ints(IntsOpt{
		Name:  name,
		Value: value,
		Desc:  desc,
	})
}

func (c *Cmd) VarOpt(name string, value flag.Value, desc string) {
	c.mkOpt(opt{name: name, desc: desc, value: value})
}

type opt struct {
	name      string
	desc      string
	envVar    string
	names     []string
	hideValue bool
	value     flag.Value
}

func (o *opt) isBool() bool {
	if bf, ok := o.value.(boolValued); ok {
		return bf.IsBoolFlag()
	}

	return false
}

func (o *opt) String() string {
	return fmt.Sprintf("Opt(%v)", o.names)
}

func mkOptStrs(optName string) []string {
	namesSl := strings.Split(optName, " ")
	for i, name := range namesSl {
		prefix := "-"
		if len(name) > 1 {
			prefix = "--"
		}
		namesSl[i] = prefix + name
	}
	return namesSl
}

func (c *Cmd) mkOpt(opt opt) {
	setFromEnv(opt.value, opt.envVar)

	opt.names = mkOptStrs(opt.name)

	c.options = append(c.options, &opt)
	for _, name := range opt.names {
		c.optionsIdx[name] = &opt
	}
}
