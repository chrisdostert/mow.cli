package cli

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

// BoolArg describes a boolean argument
type BoolArg struct {
	*boolParam

	BoolParam

	// The argument name as will be shown in help messages
	Name string
	// The argument description as will be shown in help messages
	Desc string
	// A space separated list of environment variables names to be used to initialize this argument
	EnvVar string
	// The argument's inital value
	Value bool
	// A boolean to display or not the current value of the argument in the help message
	HideValue bool
}

// StringArg describes a string argument
type StringArg struct {
	*stringParam

	// The argument name as will be shown in help messages
	Name string
	// The argument description as will be shown in help messages
	Desc string
	// A space separated list of environment variables names to be used to initialize this argument
	EnvVar string
	// The argument's inital value
	Value string
	// A boolean to display or not the current value of the argument in the help message
	HideValue bool
}

var (
	_ flag.Value = &StringOpt{}
)

func (sa *StringArg) Set(s string) error {
	*sa.into = s
	return nil
}

func (sa *StringArg) String() string {
	return fmt.Sprintf("%#v", *sa.into)
}

// IntArg describes an int argument
type IntArg struct {
	*intParam

	// The argument name as will be shown in help messages
	Name string
	// The argument description as will be shown in help messages
	Desc string
	// A space separated list of environment variables names to be used to initialize this argument
	EnvVar string
	// The argument's inital value
	Value int
	// A boolean to display or not the current value of the argument in the help message
	HideValue bool
}

var (
	_ flag.Value = &IntArg{}
)

func (ia *IntArg) Set(s string) error {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	*ia.into = int(i)
	return nil
}

func (ia *IntArg) String() string {
	return fmt.Sprintf("%v", *ia.into)
}

// StringsArg describes a string slice argument
type StringsArg struct {
	*stringsParam

	// The argument name as will be shown in help messages
	Name string
	// The argument description as will be shown in help messages
	Desc string
	// A space separated list of environment variables names to be used to initialize this argument.
	// The env variable should contain a comma separated list of values
	EnvVar string
	// The argument's inital value
	Value []string
	// A boolean to display or not the current value of the argument in the help message
	HideValue bool
}

var (
	_ flag.Value  = &StringsArg{}
	_ multiValued = &StringsArg{}
)

func (sa *StringsArg) Set(s string) error {
	*sa.into = append(*sa.into, s)
	return nil
}

func (sa *StringsArg) String() string {
	res := "["
	for idx, s := range *sa.into {
		if idx > 0 {
			res += ", "
		}
		res += fmt.Sprintf("%#v", s)
	}
	return res + "]"
}

func (sa *StringsArg) SetMulti(vs []string) error {
	newValue := make([]string, len(vs))
	for idx, v := range vs {
		v = strings.TrimSpace(v)
		newValue[idx] = v
	}
	sa.into = &newValue
	return nil
}

// IntsArg describes an int slice argument
type IntsArg struct {
	*intsParam

	// The argument name as will be shown in help messages
	Name string
	// The argument description as will be shown in help messages
	Desc string
	// A space separated list of environment variables names to be used to initialize this argument.
	// The env variable should contain a comma separated list of values
	EnvVar string
	// The argument's inital value
	Value []int
	// A boolean to display or not the current value of the argument in the help message
	HideValue bool
}

var (
	_ flag.Value  = &IntsArg{}
	_ multiValued = &IntsArg{}
)

func (ia *IntsArg) Set(s string) error {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	*ia.into = append(*ia.into, int(i))
	return nil
}

func (ia *IntsArg) String() string {
	res := "["
	for idx, s := range *ia.into {
		if idx > 0 {
			res += ", "
		}
		res += fmt.Sprintf("%v", s)
	}
	return res + "]"
}

func (ia *IntsArg) SetMulti(vs []string) error {
	newValue := []int{}
	for _, v := range vs {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return err
		}
		newValue = append(newValue, int(i))
	}
	ia.into = &newValue
	return nil
}

// VarOpt describes a user-settable option
type VarArg struct {
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

func (va *VarArg) Set(s string) error {
	return va.Value.Set(s)
}

func (va *VarArg) String() string {
	return va.Value.String()
}

func (va *VarArg) IsBoolFlag() bool {
	if bf, ok := va.Value.(boolValued); ok {
		return bf.IsBoolFlag()
	}

	return false
}

func (va *VarArg) IsMultiValued() bool {
	if bf, ok := va.Value.(multiValued); ok {
		return bf.IsMultiValued()
	}

	return false
}

func (va *VarArg) SetMulti(vs []string) error {
	mv, ok := va.Value.(multiValued)
	if !ok || !mv.IsMultiValued() {
		panic("Bug")
	}

	return mv.SetMulti(vs)
}

/*
BoolArg defines a boolean argument on the command c named `name`, with an initial value of `value` and a description of `desc` which will be used in help messages.

The result should be stored in a variable (a pointer to a bool) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) BoolArg(name string, value bool, desc string) *bool {
	return c.Bool(BoolArg{
		Name:  name,
		Value: value,
		Desc:  desc,
	})
}

/*
StringArg defines a string argument on the command c named `name`, with an initial value of `value` and a description of `desc` which will be used in help messages.

The result should be stored in a variable (a pointer to a string) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) StringArg(name string, value string, desc string) *string {
	return c.String(StringArg{
		Name:  name,
		Value: value,
		Desc:  desc,
	})
}

/*
IntArg defines an int argument on the command c named `name`, with an initial value of `value` and a description of `desc` which will be used in help messages.

The result should be stored in a variable (a pointer to an int) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) IntArg(name string, value int, desc string) *int {
	return c.Int(IntArg{
		Name:  name,
		Value: value,
		Desc:  desc,
	})
}

/*
StringsArg defines a string slice argument on the command c named `name`, with an initial value of `value` and a description of `desc` which will be used in help messages.

The result should be stored in a variable (a pointer to a string slice) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) StringsArg(name string, value []string, desc string) *[]string {
	return c.Strings(StringsArg{
		Name:  name,
		Value: value,
		Desc:  desc,
	})
}

/*
IntsArg defines an int slice argument on the command c named `name`, with an initial value of `value` and a description of `desc` which will be used in help messages.

The result should be stored in a variable (a pointer to an int slice) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) IntsArg(name string, value []int, desc string) *[]int {
	return c.Ints(IntsArg{
		Name:  name,
		Value: value,
		Desc:  desc,
	})
}

/*
IntsArg defines an int slice argument on the command c named `name`, with an initial value of `value` and a description of `desc` which will be used in help messages.

The result should be stored in a variable (a pointer to an int slice) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) VarArg(name string, value flag.Value, desc string) {
	c.mkArg(arg{name: name, desc: desc, value: value})
}

type arg struct {
	name          string
	desc          string
	envVar        string
	helpFormatter func(interface{}) string
	hideValue     bool
	value         flag.Value
}

func (a *arg) String() string {
	return fmt.Sprintf("ARG(%s)", a.name)
}

func (c *Cmd) mkArg(arg arg) {
	setFromEnv(arg.value, arg.envVar)

	c.args = append(c.args, &arg)
	c.argsIdx[arg.name] = &arg
}
