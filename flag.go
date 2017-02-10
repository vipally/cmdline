//2016-08-16
//Ally(vipally@gmail.com) modify from std.flag version 1.7
//change list see cmdline.go

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
	Package flag implements command-line flag parsing.

	Usage:

	Define flags using flag.String(), Bool(), Int(), etc.

	This declares an integer flag, -flagname, stored in the pointer ip, with type *int.
		import "flag"
		var ip = flag.Int("flagname", 1234, "help message for flagname")
	If you like, you can bind the flag to a variable using the Var() functions.
		var flagvar int
		func init() {
			flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
		}
	Or you can create custom flags that satisfy the Value interface (with
	pointer receivers) and couple them to flag parsing by
		flag.Var(&flagVal, "name", "help message for flagname")
	For such flags, the default value is just the initial value of the variable.

	After all flags are defined, call
		flag.Parse()
	to parse the command line into the defined flags.

	Flags may then be used directly. If you're using the flags themselves,
	they are all pointers; if you bind to variables, they're values.
		fmt.Println("ip has value ", *ip)
		fmt.Println("flagvar has value ", flagvar)

	After parsing, the arguments following the flags are available as the
	slice flag.Args() or individually as flag.Arg(i).
	The arguments are indexed from 0 through flag.NArg()-1.

	Command line flag syntax:
		-flag
		-flag=x
		-flag x  // non-boolean flags only
	One or two minus signs may be used; they are equivalent.
	The last form is not permitted for boolean flags because the
	meaning of the command
		cmd -x *
	will change if there is a file called 0, false, etc.  You must
	use the -flag=false form to turn off a boolean flag.

	Flag parsing stops just before the first non-flag argument
	("-" is a non-flag argument) or after the terminator "--".

	Integer flags accept 1234, 0664, 0x1234 and may be negative.
	Boolean flags may be:
		1, 0, t, f, T, F, true, false, TRUE, FALSE, True, False
	Duration flags accept any input valid for time.ParseDuration.

	The default set of command-line flags is controlled by
	top-level functions.  The FlagSet type allows one to define
	independent sets of flags, such as to implement subcommands
	in a command-line interface. The methods of FlagSet are
	analogous to the top-level functions for the command-line
	flag set.
*/
//package cmdline support a friendly command line interface based on flag

package cmdline

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	gNoNamePrefix = "{noname#"
)

// ErrHelp is the error returned if the -help or -h flag is invoked
// but no such flag is defined.
var ErrHelp = errors.New("flag: help requested")

// -- bool Value
type boolValue bool

func newBoolValue(val bool, p *bool) *boolValue {
	*p = val
	return (*boolValue)(p)
}

func (b *boolValue) Set(s string) error {
	v, err := strconv.ParseBool(s)
	*b = boolValue(v)
	return err
}

func (b *boolValue) Get() interface{} { return bool(*b) }

func (b *boolValue) String() string { return fmt.Sprintf("%v", *b) }

func (b *boolValue) IsBoolFlag() bool { return true }

// optional interface to indicate boolean flags that can be
// supplied without "=value" text
type boolFlag interface {
	Value
	IsBoolFlag() bool
}

// -- int Value
type intValue int

func newIntValue(val int, p *int) *intValue {
	*p = val
	return (*intValue)(p)
}

func (i *intValue) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	*i = intValue(v)
	return err
}

func (i *intValue) Get() interface{} { return int(*i) }

func (i *intValue) String() string { return fmt.Sprintf("%v", *i) }

// -- int64 Value
type int64Value int64

func newInt64Value(val int64, p *int64) *int64Value {
	*p = val
	return (*int64Value)(p)
}

func (i *int64Value) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	*i = int64Value(v)
	return err
}

func (i *int64Value) Get() interface{} { return int64(*i) }

func (i *int64Value) String() string { return fmt.Sprintf("%v", *i) }

// -- uint Value
type uintValue uint

func newUintValue(val uint, p *uint) *uintValue {
	*p = val
	return (*uintValue)(p)
}

func (i *uintValue) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, 64)
	*i = uintValue(v)
	return err
}

func (i *uintValue) Get() interface{} { return uint(*i) }

func (i *uintValue) String() string { return fmt.Sprintf("%v", *i) }

// -- uint64 Value
type uint64Value uint64

func newUint64Value(val uint64, p *uint64) *uint64Value {
	*p = val
	return (*uint64Value)(p)
}

func (i *uint64Value) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, 64)
	*i = uint64Value(v)
	return err
}

func (i *uint64Value) Get() interface{} { return uint64(*i) }

func (i *uint64Value) String() string { return fmt.Sprintf("%v", *i) }

// -- string Value
type stringValue string

func newStringValue(val string, p *string) *stringValue {
	*p = val
	return (*stringValue)(p)
}

func (s *stringValue) Set(val string) error {
	*s = stringValue(val)
	return nil
}

func (s *stringValue) Get() interface{} { return string(*s) }

func (s *stringValue) String() string { return fmt.Sprintf("%s", *s) }

// -- float64 Value
type float64Value float64

func newFloat64Value(val float64, p *float64) *float64Value {
	*p = val
	return (*float64Value)(p)
}

func (f *float64Value) Set(s string) error {
	v, err := strconv.ParseFloat(s, 64)
	*f = float64Value(v)
	return err
}

func (f *float64Value) Get() interface{} { return float64(*f) }

func (f *float64Value) String() string { return fmt.Sprintf("%v", *f) }

// -- time.Duration Value
type durationValue time.Duration

func newDurationValue(val time.Duration, p *time.Duration) *durationValue {
	*p = val
	return (*durationValue)(p)
}

func (d *durationValue) Set(s string) error {
	v, err := time.ParseDuration(s)
	*d = durationValue(v)
	return err
}

func (d *durationValue) Get() interface{} { return time.Duration(*d) }

func (d *durationValue) String() string { return (*time.Duration)(d).String() }

// Value is the interface to the dynamic value stored in a flag.
// (The default value is represented as a string.)
//
// If a Value has an IsBoolFlag() bool method returning true,
// the command-line parser makes -name equivalent to -name=true
// rather than using the next command-line argument.
//
// Set is called once, in command line order, for each flag present.
type Value interface {
	String() string
	Set(string) error
}

// Getter is an interface that allows the contents of a Value to be retrieved.
// It wraps the Value interface, rather than being part of it, because it
// appeared after Go 1 and its compatibility rules. All Value types provided
// by this package satisfy the Getter interface.
type Getter interface {
	Value
	Get() interface{}
}

// ErrorHandling defines how FlagSet.Parse behaves if the parse fails.
type ErrorHandling int

// These constants cause FlagSet.Parse to behave as described if the parse fails.
const (
	ContinueOnError ErrorHandling = iota // Return a descriptive error.
	ExitOnError                          // Call os.Exit(2).
	PanicOnError                         // Call panic with a descriptive error.
)

// A FlagSet represents a set of defined flags. The zero value of a FlagSet
// has no name and has ContinueOnError error handling.
type FlagSet struct {
	// Usage is the function called when an error occurs while parsing flags.
	// The field is a function (not a method) that may be changed to point to
	// a custom error handler.
	Usage func()

	name          string
	parsed        bool
	actual        map[string]*Flag
	formal        map[string]*Flag
	args          []string // arguments after flags
	errorHandling ErrorHandling
	output        io.Writer // nil means stderr; use out() accessor

	summary   string //summary of this application
	copyright string //copyright of this application
	details   string //detail use of this application
	auto_id   int    //no-name flag uses auto_id++ as auto-name suffix
}

// A Flag represents the state of a flag.
type Flag struct {
	Name     string // name as it appears on command line(instead by Synonyms )
	Usage    string // help message
	Value    Value  // value as set
	DefValue string // default value (as text); for usage message

	LogicName string   //logic name of this flag
	Required  bool     //if this flag is force required
	Synonyms  []string //different flags(eg:-f/-flag) maybe the same ones, they are synonyms
	Visitor   string   //name of what synonym is visiting this flag
}

// GetShowName return name that will show in usage page. with "-f|flag=" format
//
// no-name ones returns empty and others synonyms
func (f *Flag) GetShowName() (r string) {
	if !strings.HasPrefix(f.Name, gNoNamePrefix) {
		r = fmt.Sprintf("-%s=", f.GetSynonyms())
	}
	return
}

//GetSynonyms return synonyms of this flag, as "f|flag" format
func (f *Flag) GetSynonyms() string {
	var b bytes.Buffer
	for _, v := range f.Synonyms {
		b.WriteString(v)
		b.WriteByte('|')
	}
	if len(f.Synonyms) > 0 {
		b.Truncate(b.Len() - 1) //remove last '|'
	}

	return b.String()
}

// sortFlags returns the flags as a slice in lexicographical sorted order.
func sortFlags(flags map[string]*Flag) ([]*Flag, []string) {
	list := make(sort.StringSlice, len(flags))
	i := 0
	for name, _ := range flags {
		list[i] = name
		i++
	}
	list.Sort()
	result := make([]*Flag, len(list))
	for i, name := range list {
		result[i] = flags[name]
	}
	return result, list
}

func (f *FlagSet) out() io.Writer {
	if f.output == nil {
		return os.Stderr
	}
	return f.output
}

// SetOutput sets the destination for usage and error messages.
// If output is nil, os.Stderr is used.
func (f *FlagSet) SetOutput(output io.Writer) {
	f.output = output
}

// VisitAll visits the flags in lexicographical order, calling fn for each.
// It visits all flags, even those not set.
func (f *FlagSet) VisitAll(fn func(*Flag)) {
	list, names := sortFlags(f.formal)
	for i, flag := range list {
		flag.Visitor = names[i]
		fn(flag)
	}
}

// VisitAll visits the command-line flags in lexicographical order, calling
// fn for each. It visits all flags, even those not set.
func VisitAll(fn func(*Flag)) {
	CommandLine.VisitAll(fn)
}

// Visit visits the flags in lexicographical order, calling fn for each.
// It visits only those flags that have been set.
func (f *FlagSet) Visit(fn func(*Flag)) {
	list, names := sortFlags(f.actual)
	for i, flag := range list {
		flag.Visitor = names[i]
		fn(flag)
	}
}

// Visit visits the command-line flags in lexicographical order, calling fn
// for each. It visits only those flags that have been set.
func Visit(fn func(*Flag)) {
	CommandLine.Visit(fn)
}

// Lookup returns the Flag structure of the named flag, returning nil if none exists.
func (f *FlagSet) Lookup(name string) *Flag {
	return f.formal[name]
}

// Lookup returns the Flag structure of the named command-line flag,
// returning nil if none exists.
func Lookup(name string) *Flag {
	return CommandLine.formal[name]
}

// Set sets the value of the named flag.
func (f *FlagSet) Set(name, value string) error {
	flag, ok := f.formal[name]
	if !ok {
		return fmt.Errorf("no such flag -%v", name)
	}
	err := flag.Value.Set(value)
	if err != nil {
		return err
	}
	if f.actual == nil {
		f.actual = make(map[string]*Flag)
	}
	f.actual[name] = flag
	return nil
}

// Set sets the value of the named command-line flag.
func Set(name, value string) error {
	return CommandLine.Set(name, value)
}

// isZeroValue guesses whether the string represents the zero
// value for a flag. It is not accurate but in practice works OK.
func isZeroValue(flag *Flag, value string) bool {
	// Build a zero value of the flag's Value type, and see if the
	// result of calling its String method equals the value passed in.
	// This works unless the Value type is itself an interface type.
	typ := reflect.TypeOf(flag.Value)
	var z reflect.Value
	if typ.Kind() == reflect.Ptr {
		z = reflect.New(typ.Elem())
	} else {
		z = reflect.Zero(typ)
	}
	if value == z.Interface().(Value).String() {
		return true
	}

	switch value {
	case "false":
		return true
	case "":
		return true
	case "0":
		return true
	}
	return false
}

// UnquoteUsage extracts a back-quoted name from the usage
// string for a flag and returns it and the un-quoted usage.
// Given "a `name` to show" it returns ("name", "a name to show").
// If there are no back quotes, the name is an educated guess of the
// type of the flag's value, or the empty string if the flag is boolean.
func UnquoteUsage(flag *Flag) (name string, usage string) {
	// Look for a back-quoted name, but avoid the strings package.
	usage = flag.Usage
	for i := 0; i < len(usage); i++ {
		if usage[i] == '`' {
			for j := i + 1; j < len(usage); j++ {
				if usage[j] == '`' {
					name = usage[i+1 : j]
					usage = usage[:i] + name + usage[j+1:]
					return name, usage
				}
			}
			break // Only one back quote; use type name.
		}
	}
	// No explicit name, so use type if we can find one.
	name = "value"
	switch flag.Value.(type) {
	case boolFlag:
		name = ""
	case *durationValue:
		name = "duration"
	case *float64Value:
		name = "float"
	case *intValue, *int64Value:
		name = "int"
	case *stringValue:
		name = "string"
	case *uintValue, *uint64Value:
		name = "uint"
	}
	return
}

// PrintDefaults prints to standard error the default values of all
// defined command-line flags in the set. See the documentation for
// the global function PrintDefaults for more information.
func (f *FlagSet) PrintDefaults() {
	fmt.Fprintf(f.out(), f.GetUsage())
}

// PrintDefaults prints, to standard error unless configured otherwise,
// a usage message showing the default settings of all defined
// command-line flags.
// For an integer valued flag x, the default output has the form
//	-x int
//		usage-message-for-x (default 7)
// The usage message will appear on a separate line for anything but
// a bool flag with a one-byte name. For bool flags, the type is
// omitted and if the flag name is one byte the usage message appears
// on the same line. The parenthetical default is omitted if the
// default is the zero value for the type. The listed type, here int,
// can be changed by placing a back-quoted name in the flag's usage
// string; the first such item in the message is taken to be a parameter
// name to show in the message and the back quotes are stripped from
// the message when displayed. For instance, given
//	flag.String("I", "", "search `directory` for include files")
// the output will be
//	-I directory
//		search directory for include files.
func PrintDefaults() {
	CommandLine.PrintDefaults()
}

// defaultUsage is the default function to print a usage message.
func defaultUsage(f *FlagSet) {
	if f.name == "" {
		fmt.Fprintf(f.out(), "Usage:\n")
	} else {
		fmt.Fprintf(f.out(), "Usage of %s:\n", f.name)
	}
	f.PrintDefaults()
}

// NOTE: Usage is not just defaultUsage(CommandLine)
// because it serves (via godoc flag Usage) as the example
// for how to write your own usage function.

// Usage prints to standard error a usage message documenting all defined command-line flags.
// It is called when an error occurs while parsing flags.
// The function is a variable that may be changed to point to a custom function.
// By default it prints a simple header and calls PrintDefaults; for details about the
// format of the output and how to control it, see the documentation for PrintDefaults.
var Usage = func() {
	PrintDefaults()
}

// NFlag returns the number of flags that have been set.
func (f *FlagSet) NFlag() int { return len(f.actual) }

// NFlag returns the number of command-line flags that have been set.
func NFlag() int { return len(CommandLine.actual) }

// Arg returns the i'th argument. Arg(0) is the first remaining argument
// after flags have been processed. Arg returns an empty string if the
// requested element does not exist.
func (f *FlagSet) Arg(i int) string {
	if i < 0 || i >= len(f.args) {
		return ""
	}
	return f.args[i]
}

// Arg returns the i'th command-line argument. Arg(0) is the first remaining argument
// after flags have been processed. Arg returns an empty string if the
// requested element does not exist.
func Arg(i int) string {
	return CommandLine.Arg(i)
}

// NArg is the number of arguments remaining after flags have been processed.
func (f *FlagSet) NArg() int { return len(f.args) }

// NArg is the number of arguments remaining after flags have been processed.
func NArg() int { return len(CommandLine.args) }

// Args returns the non-flag arguments.
func (f *FlagSet) Args() []string { return f.args }

// Args returns the non-flag command-line arguments.
func Args() []string { return CommandLine.args }

// BoolVar defines a bool flag with specified name, default value, and usage string.
// The argument p points to a bool variable in which to store the value of the flag.
func (f *FlagSet) BoolVar(p *bool, name string, logic_name string, value bool, required bool, usage string) {
	f.Var(newBoolValue(value, p), name, logic_name, required, usage)
}

// BoolVar defines a bool flag with specified name, default value, and usage string.
// The argument p points to a bool variable in which to store the value of the flag.
func BoolVar(p *bool, name string, logic_name string, value bool, required bool, usage string) {
	CommandLine.Var(newBoolValue(value, p), name, logic_name, required, usage)
}

// Bool defines a bool flag with specified name, default value, and usage string.
// The return value is the address of a bool variable that stores the value of the flag.
func (f *FlagSet) Bool(name string, logic_name string, value bool, required bool, usage string) *bool {
	p := new(bool)
	f.BoolVar(p, name, logic_name, value, required, usage)
	return p
}

// Bool defines a bool flag with specified name, default value, and usage string.
// The return value is the address of a bool variable that stores the value of the flag.
func Bool(name string, logic_name string, value bool, required bool, usage string) *bool {
	return CommandLine.Bool(name, logic_name, value, required, usage)
}

// IntVar defines an int flag with specified name, default value, and usage string.
// The argument p points to an int variable in which to store the value of the flag.
func (f *FlagSet) IntVar(p *int, name string, logic_name string, value int, required bool, usage string) {
	f.Var(newIntValue(value, p), name, logic_name, required, usage)
}

// IntVar defines an int flag with specified name, default value, and usage string.
// The argument p points to an int variable in which to store the value of the flag.
func IntVar(p *int, name string, logic_name string, value int, required bool, usage string) {
	CommandLine.Var(newIntValue(value, p), name, logic_name, required, usage)
}

// Int defines an int flag with specified name, default value, and usage string.
// The return value is the address of an int variable that stores the value of the flag.
func (f *FlagSet) Int(name string, logic_name string, value int, required bool, usage string) *int {
	p := new(int)
	f.IntVar(p, name, logic_name, value, required, usage)
	return p
}

// Int defines an int flag with specified name, default value, and usage string.
// The return value is the address of an int variable that stores the value of the flag.
func Int(name string, logic_name string, value int, required bool, usage string) *int {
	return CommandLine.Int(name, logic_name, value, required, usage)
}

// Int64Var defines an int64 flag with specified name, default value, and usage string.
// The argument p points to an int64 variable in which to store the value of the flag.
func (f *FlagSet) Int64Var(p *int64, name string, logic_name string, value int64, required bool, usage string) {
	f.Var(newInt64Value(value, p), name, logic_name, required, usage)
}

// Int64Var defines an int64 flag with specified name, default value, and usage string.
// The argument p points to an int64 variable in which to store the value of the flag.
func Int64Var(p *int64, name string, logic_name string, value int64, required bool, usage string) {
	CommandLine.Var(newInt64Value(value, p), name, logic_name, required, usage)
}

// Int64 defines an int64 flag with specified name, default value, and usage string.
// The return value is the address of an int64 variable that stores the value of the flag.
func (f *FlagSet) Int64(name string, logic_name string, value int64, required bool, usage string) *int64 {
	p := new(int64)
	f.Int64Var(p, name, logic_name, value, required, usage)
	return p
}

// Int64 defines an int64 flag with specified name, default value, and usage string.
// The return value is the address of an int64 variable that stores the value of the flag.
func Int64(name string, logic_name string, value int64, required bool, usage string) *int64 {
	return CommandLine.Int64(name, logic_name, value, required, usage)
}

// UintVar defines a uint flag with specified name, default value, and usage string.
// The argument p points to a uint variable in which to store the value of the flag.
func (f *FlagSet) UintVar(p *uint, name string, logic_name string, value uint, required bool, usage string) {
	f.Var(newUintValue(value, p), name, logic_name, required, usage)
}

// UintVar defines a uint flag with specified name, default value, and usage string.
// The argument p points to a uint  variable in which to store the value of the flag.
func UintVar(p *uint, name string, logic_name string, value uint, required bool, usage string) {
	CommandLine.Var(newUintValue(value, p), name, logic_name, required, usage)
}

// Uint defines a uint flag with specified name, default value, and usage string.
// The return value is the address of a uint  variable that stores the value of the flag.
func (f *FlagSet) Uint(name string, logic_name string, value uint, required bool, usage string) *uint {
	p := new(uint)
	f.UintVar(p, name, logic_name, value, required, usage)
	return p
}

// Uint defines a uint flag with specified name, default value, and usage string.
// The return value is the address of a uint  variable that stores the value of the flag.
func Uint(name string, logic_name string, value uint, required bool, usage string) *uint {
	return CommandLine.Uint(name, logic_name, value, required, usage)
}

// Uint64Var defines a uint64 flag with specified name, default value, and usage string.
// The argument p points to a uint64 variable in which to store the value of the flag.
func (f *FlagSet) Uint64Var(p *uint64, name string, logic_name string, value uint64, required bool, usage string) {
	f.Var(newUint64Value(value, p), name, logic_name, required, usage)
}

// Uint64Var defines a uint64 flag with specified name, default value, and usage string.
// The argument p points to a uint64 variable in which to store the value of the flag.
func Uint64Var(p *uint64, name string, logic_name string, value uint64, required bool, usage string) {
	CommandLine.Var(newUint64Value(value, p), name, logic_name, required, usage)
}

// Uint64 defines a uint64 flag with specified name, default value, and usage string.
// The return value is the address of a uint64 variable that stores the value of the flag.
func (f *FlagSet) Uint64(name string, logic_name string, value uint64, required bool, usage string) *uint64 {
	p := new(uint64)
	f.Uint64Var(p, name, logic_name, value, required, usage)
	return p
}

// Uint64 defines a uint64 flag with specified name, default value, and usage string.
// The return value is the address of a uint64 variable that stores the value of the flag.
func Uint64(name string, logic_name string, value uint64, required bool, usage string) *uint64 {
	return CommandLine.Uint64(name, logic_name, value, required, usage)
}

// StringVar defines a string flag with specified name, default value, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
func (f *FlagSet) StringVar(p *string, name string, logic_name string, value string, required bool, usage string) {
	f.Var(newStringValue(value, p), name, logic_name, required, usage)
}

// StringVar defines a string flag with specified name, default value, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
func StringVar(p *string, name string, logic_name string, value string, required bool, usage string) {
	CommandLine.Var(newStringValue(value, p), name, logic_name, required, usage)
}

// String defines a string flag with specified name, default value, and usage string.
// The return value is the address of a string variable that stores the value of the flag.
func (f *FlagSet) String(name string, logic_name string, value string, required bool, usage string) *string {
	p := new(string)
	f.StringVar(p, name, logic_name, value, required, usage)
	return p
}

// String defines a string flag with specified name, default value, and usage string.
// The return value is the address of a string variable that stores the value of the flag.
func String(name string, logic_name string, value string, required bool, usage string) *string {
	return CommandLine.String(name, logic_name, value, required, usage)
}

// Float64Var defines a float64 flag with specified name, default value, and usage string.
// The argument p points to a float64 variable in which to store the value of the flag.
func (f *FlagSet) Float64Var(p *float64, name string, logic_name string, value float64, required bool, usage string) {
	f.Var(newFloat64Value(value, p), name, logic_name, required, usage)
}

// Float64Var defines a float64 flag with specified name, default value, and usage string.
// The argument p points to a float64 variable in which to store the value of the flag.
func Float64Var(p *float64, name string, logic_name string, value float64, required bool, usage string) {
	CommandLine.Var(newFloat64Value(value, p), name, logic_name, required, usage)
}

// Float64 defines a float64 flag with specified name, default value, and usage string.
// The return value is the address of a float64 variable that stores the value of the flag.
func (f *FlagSet) Float64(name string, logic_name string, value float64, required bool, usage string) *float64 {
	p := new(float64)
	f.Float64Var(p, name, logic_name, value, required, usage)
	return p
}

// Float64 defines a float64 flag with specified name, default value, and usage string.
// The return value is the address of a float64 variable that stores the value of the flag.
func Float64(name string, logic_name string, value float64, required bool, usage string) *float64 {
	return CommandLine.Float64(name, logic_name, value, required, usage)
}

// DurationVar defines a time.Duration flag with specified name, default value, and usage string.
// The argument p points to a time.Duration variable in which to store the value of the flag.
// The flag accepts a value acceptable to time.ParseDuration.
func (f *FlagSet) DurationVar(p *time.Duration, name string, logic_name string, value time.Duration, required bool, usage string) {
	f.Var(newDurationValue(value, p), name, logic_name, required, usage)
}

// DurationVar defines a time.Duration flag with specified name, default value, and usage string.
// The argument p points to a time.Duration variable in which to store the value of the flag.
// The flag accepts a value acceptable to time.ParseDuration.
func DurationVar(p *time.Duration, name string, logic_name string, value time.Duration, required bool, usage string) {
	CommandLine.Var(newDurationValue(value, p), name, logic_name, required, usage)
}

// Duration defines a time.Duration flag with specified name, default value, and usage string.
// The return value is the address of a time.Duration variable that stores the value of the flag.
// The flag accepts a value acceptable to time.ParseDuration.
func (f *FlagSet) Duration(name string, logic_name string, value time.Duration, required bool, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationVar(p, name, logic_name, value, required, usage)
	return p
}

// Duration defines a time.Duration flag with specified name, default value, and usage string.
// The return value is the address of a time.Duration variable that stores the value of the flag.
// The flag accepts a value acceptable to time.ParseDuration.
func Duration(name string, logic_name string, value time.Duration, required bool, usage string) *time.Duration {
	return CommandLine.Duration(name, logic_name, value, required, usage)
}

func getValuePtr(value Value) (r uintptr) {
	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Ptr {
		r = v.Pointer()
	}
	return
}

// Var defines a flag with the specified name and usage string. The type and
// value of the flag are represented by the first argument, of type Value, which
// typically holds a user-defined implementation of Value. For instance, the
// caller could create a flag that turns a comma-separated string into a slice
// of strings by giving the slice the methods of Value; in particular, Set would
// decompose the comma-separated string into the slice.
func (f *FlagSet) Var(value Value, name string, logic_name string, required bool, usage string) {
	// Remember the default value as a string; it won't change.
	name = f.getAutoName(name) //auto generate a name if not assigned a flag name

	_, alreadythere := f.formal[name]
	if alreadythere {
		var msg string
		if f.name == "" {
			msg = fmt.Sprintf("flag redefined: %s", name)
		} else {
			msg = fmt.Sprintf("%s flag redefined: %s", f.name, name)
		}
		fmt.Fprintln(f.out(), msg)
		panic(msg) // Happens only if flags are declared with identical names
	}
	var flag *Flag
	if !strings.HasPrefix(name, gNoNamePrefix) { //no-name flags do not support Synonyms
		for _, f := range f.formal { //find if there is a Synonyms flag
			if getValuePtr(value) == getValuePtr(f.Value) { //the same one
				flag = f
				f.Synonyms = append(f.Synonyms, name) //Synonyms
				break
			}
		}
	}
	if nil == flag {
		flag = &Flag{name, usage, value, value.String(), logic_name, required, []string{name}, ""}
	}
	if f.formal == nil {
		f.formal = make(map[string]*Flag)
	}
	f.formal[name] = flag
}

// Var defines a flag with the specified name and usage string. The type and
// value of the flag are represented by the first argument, of type Value, which
// typically holds a user-defined implementation of Value. For instance, the
// caller could create a flag that turns a comma-separated string into a slice
// of strings by giving the slice the methods of Value; in particular, Set would
// decompose the comma-separated string into the slice.
func Var(value Value, name string, logic_name string, required bool, usage string) {
	CommandLine.Var(value, name, logic_name, required, usage)
}

// failf prints to standard error a formatted error and usage message and
// returns the error.
func (f *FlagSet) failf(format string, a ...interface{}) error {
	err := fmt.Errorf(format, a...)
	fmt.Fprintln(f.out(), err)
	f.usage()
	return err
}

// usage calls the Usage method for the flag set if one is specified,
// or the appropriate default usage function otherwise.
func (f *FlagSet) usage() {
	if f.Usage == nil {
		if f == CommandLine {
			Usage()
		} else {
			defaultUsage(f)
		}
	} else {
		f.Usage()
	}
}

func isFlagPrefix(c byte) bool {
	return c == '-' || c == '/'
}

// parseOne parses one flag. It reports whether a flag was seen.
func (f *FlagSet) parseOne() (bool, error) {
	if len(f.args) == 0 {
		return false, nil
	}
	s := f.args[0]
	if s == "" || s == "-" || s == "--" || s == "/" {
		return false, nil
	}

	name, value := "", ""
	if isFlagPrefix(s[0]) {
		numMinuses := 1
		if s[1] == '-' {
			numMinuses++
		}

		name, value = s[numMinuses:], ""
		for i := 0; i < len(name); i++ { //split at first '='
			if name[i] == '=' {
				value = name[i+1:]
				name = name[:i]
				break
			}
		}

	} else {
		name = f.getAutoName("") //auto generate a name if not assigned a flag name
		value = s
	}

	if len(name) == 0 || isFlagPrefix(name[0]) || name[0] == '=' {
		return false, f.failf("[error] bad flag syntax: %s", s)
	}

	// it's a flag. does it have an argument?
	f.args = f.args[1:]

	m := f.formal
	flag, alreadythere := m[name] // BUG
	if !alreadythere {
		if name == "help" || name == "h" || name == "?" { // special case for nice help message.
			f.usage()
			return false, ErrHelp
		}
		return false, f.failf("[error] flag provided but not defined: -%s", name)
	}

	if fv, ok := flag.Value.(boolFlag); ok && fv.IsBoolFlag() { // special case: doesn't need an arg
		if value != "" {
			if err := fv.Set(value); err != nil {
				return false, f.failf("[error] invalid boolean value %q for -%s: %v", value, name, err)
			}
		} else {
			if err := fv.Set("true"); err != nil {
				return false, f.failf("[error] invalid boolean flag %s: %v", name, err)
			}
		}
	} else {
		// It must have a value, which might be the next argument.
		for len(f.args) > 0 && (value == "" || value == "=") { //consider "-f = 1" or "-f= 1" or "-f =1" cases
			value, f.args = f.args[0], f.args[1:]
			if value[0] == '=' {
				value = value[1:]
			}
		}

		if value == "" {
			return false, f.failf("[error] flag needs an argument: -%s", name)
		}
		if err := flag.Value.Set(value); err != nil {
			return false, f.failf("[error] invalid value %q for flag -%s: %v", value, name, err)
		}
	}
	if f.actual == nil {
		f.actual = make(map[string]*Flag)
	}
	f.actual[name] = flag
	return true, nil
}

// Parse parses flag definitions from the argument list, which should not
// include the command name. Must be called after all flags in the FlagSet
// are defined and before flags are accessed by the program.
// The return value will be ErrHelp if -help or -h were set but not defined.
func (f *FlagSet) Parse(arguments []string) error {
	f.parsed = true
	f.args = arguments
	f.auto_id = 0 //reset auto_id for parse logic to generate noname flags
	for {
		seen, err := f.parseOne()
		if seen {
			continue
		}
		if err == nil {
			break
		}
		if ok, err_r := f.handle_error(err); ok {
			return err_r
		}
	}
	if err := f.check_require(); err != nil {
		if ok, err_r := f.handle_error(err); ok {
			return err_r
		}
	}
	return nil
}

// Parse parses the command-line flags from os.Args[1:].  Must be called
// after all flags are defined and before flags are accessed by the program.
func Parse() {
	// Ignore errors; CommandLine is set for ExitOnError.
	CommandLine.Parse(os.Args[1:])
}

// Parsed reports whether f.Parse has been called.
func (f *FlagSet) Parsed() bool {
	return f.parsed
}

// Parsed reports whether the command-line flags have been parsed.
func Parsed() bool {
	return CommandLine.Parsed()
}

//GetUsage returns the usage string
func GetUsage() string {
	return CommandLine.GetUsage()
}

//GetUsage returns the usage string
func (f *FlagSet) GetUsage() string {
	buf := bytes.NewBufferString("")
	buf.WriteString(fmt.Sprintf("Usage of ([%s] Build %s):\n", thisCmd, "" /*BuildTime()*/))
	if f.summary != "" {
		buf.WriteString(fmt.Sprintf("  Summary:\n%s\n\n", FormatLineHead(f.summary, "    ")))
	}

	buf.WriteString(fmt.Sprintf("  Usage:\n    %s", thisCmd))
	f.VisitAll(func(flag *Flag) {
		if flag.Visitor != flag.Name { //Synonyms show at the first one only
			return
		}

		_fmt := ""
		if flag.Required {
			_fmt = " %s<%s>"
		} else {
			_fmt = " [%s<%s>]"
		}
		s := fmt.Sprintf(_fmt, flag.GetShowName(), flag.LogicName)
		buf.WriteString(s)
	})
	buf.WriteString("\n")

	f.VisitAll(func(flag *Flag) {
		if flag.Visitor != flag.Name { //Synonyms show at the first one only
			return
		}

		s := fmt.Sprintf("  %s<%s>", flag.GetShowName(), flag.LogicName) // Two spaces before -; see next two comments.
		buf.WriteString(s)
		if flag.Required {
			buf.WriteString("  required")
		}
		name, usage := UnquoteUsage(flag)
		if len(name) > 0 {
			buf.WriteString("  ")
			buf.WriteString(name)
		}
		if !isZeroValue(flag, flag.DefValue) {
			if _, ok := flag.Value.(*stringValue); ok {
				// put quotes on the value
				buf.WriteString(fmt.Sprintf(" (default %q)", flag.DefValue))
			} else {
				buf.WriteString(fmt.Sprintf(" (default %v)", flag.DefValue))
			}
		}
		// Boolean flags of one ASCII letter are so common we
		// treat them specially, putting their usage on the same line.
		if len(s) <= 4 { // space, space, '-', 'x'.
			buf.WriteString("\t")
		} else {
			// Four spaces before the tab triggers good alignment
			// for both 4- and 8-space tab stops.
			buf.WriteString("\n")
		}
		buf.WriteString(FormatLineHead(usage, "    "))
		buf.WriteString("\n")
	})

	if f.copyright != "" {
		buf.WriteString(fmt.Sprintf("\n  CopyRight:\n%s", FormatLineHead(f.copyright, "    ")))
		if f.copyright[len(f.copyright)-1] != '\n' {
			buf.WriteRune('\n')
		}
	}

	if f.details != "" {
		buf.WriteString(fmt.Sprintf("\n  Details:\n%s\n", FormatLineHead(f.details, "    ")))
	}

	return buf.String()
}

func (f *FlagSet) handle_error(err error) (bool, error) {
	if err != nil {
		switch f.errorHandling {
		case ContinueOnError:
			return true, err
		case ExitOnError:
			os.Exit(2)
		case PanicOnError:
			panic(err)
		}
	}
	return false, err
}

//check if there is a required flag and do not set it
func (f *FlagSet) check_require() error {
	for name, flg := range f.formal {
		if flg.Required {
			if _, ok := f.actual[name]; !ok {
				return f.failf("[error] require but lack of flag %s<%s>", flg.GetShowName(), flg.LogicName)
			}
		}
	}
	return nil
}

//Summary set the summary info of the command, this will show in usage page
func Summary(summary string) (old string) {
	return CommandLine.Summary(summary)
}

//Details set the detail info of the command, this will show in usage page
func Details(details string) (old string) {
	return CommandLine.Details(details)
}

//CopyRight set the copyright info of the command, this will show in usage page
func CopyRight(copyright string) (old string) {
	return CommandLine.CopyRight(copyright)
}

//Summary set the summary info of the command, this will show in usage page
func (f *FlagSet) Summary(summary string) (old string) {
	old, f.summary = f.summary, ReplaceTags(summary)
	return
}

//Details set the detail info of the command, this will show in usage page
func (f *FlagSet) Details(details string) (old string) {
	old, f.details = f.details, ReplaceTags(details)
	return
}

//CopyRight set the copyright info of the command, this will show in usage page
func (f *FlagSet) CopyRight(copyright string) (old string) {
	old, f.copyright = f.copyright, ReplaceTags(copyright)
	return
}

//auto genterate a name if name not assigned
func (f *FlagSet) getAutoName(name string) string {
	if name == "" || strings.HasPrefix(name, gNoNamePrefix) {
		f.auto_id++
		name = fmt.Sprintf("%s%d}", gNoNamePrefix, f.auto_id)
	}
	return name
}

//AnotherName add a synonym flag newname for old
func AnotherName(newname, old string) (ok bool) {
	return CommandLine.AnotherName(newname, old)
}

//AnotherName add a synonym flag newname for old
func (f *FlagSet) AnotherName(newname, old string) (ok bool) {
	var msg string
	ok = true
	if ok && strings.HasPrefix(newname, gNoNamePrefix) {
		msg = fmt.Sprintf("RepeatFlag: %s forbid newname", newname)
		ok = false
	}
	if _, _ok := f.formal[newname]; ok && _ok {
		msg = fmt.Sprintf("RepeatFlag: %s redefined", newname)
		ok = false
	}
	if ok {
		if flag, _ok := f.formal[old]; _ok {
			flag.Synonyms = append(flag.Synonyms, newname)
			f.formal[newname] = flag
		} else {
			msg = fmt.Sprintf("RepeatFlag: old %s not exists", old)
			ok = false
		}
	}

	if !ok {
		var err string
		if f.name == "" {
			err = msg
		} else {
			err = fmt.Sprintf("%s %s", f.name, msg)
		}
		fmt.Fprintln(f.out(), err)
		panic(err)
	}
	return
}

// CommandLine is the default set of command-line flags, parsed from os.Args.
// The top-level functions such as BoolVar, Arg, and so on are wrappers for the
// methods of CommandLine.
var CommandLine = NewFlagSet(get_cmd(os.Args[0]), ExitOnError)

// NewFlagSet returns a new, empty flag set with the specified name and
// error handling property.
func NewFlagSet(name string, errorHandling ErrorHandling) *FlagSet {
	f := &FlagSet{
		name:          name,
		errorHandling: errorHandling,
	}
	return f
}

// Init sets the name and error handling property for a flag set.
// By default, the zero FlagSet uses an empty name and the
// ContinueOnError error handling policy.
func (f *FlagSet) Init(name string, errorHandling ErrorHandling) {
	f.name = name
	f.errorHandling = errorHandling
}
