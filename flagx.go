//2016-08-16
//Ally(vipally@gmail.com) modify from std.flag version 1.11
//change list see cmdline.go

//package cmdline support a friendly command line interface based on flag
package cmdline

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"
)

const (
	gNoNamePrefix = "{noname#"
)

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

func (f *FlagSet) PrintDefaults() {
	fmt.Fprint(f.Output(), f.GetUsage())
}

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
		fmt.Fprintln(f.Output(), msg)
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

func isFlagLeadByte(c byte) bool {
	return c == '-' || c == '/'
}

func isFlagLead(s string) bool {
	return s == "-" || s == "--" || s == "/"
}

func isHelpFlag(s string) bool {
	return s == "help" || s == "h" || s == "?"
}

//GetUsage returns the usage string
func GetUsage() string {
	return CommandLine.GetUsage()
}

//GetUsage returns the usage string
func (f *FlagSet) GetUsage() string {
	buf := bytes.NewBufferString("")
	buf.WriteString(fmt.Sprintf("Usage of ([%s] Build %s):\n", thisCmd, BuildTime()))
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

func (f *FlagSet) handleError(err error) (bool, error) {
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
func (f *FlagSet) checkRequiredFlag() error {
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

//Alias add a synonym flag newname for old
func Alias(newname, old string) (ok bool) {
	return CommandLine.Alias(newname, old)
}

//Alias add a synonym flag newname for old
func (f *FlagSet) Alias(newname, old string) (ok bool) {
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
		fmt.Fprintln(f.Output(), err)
		panic(err)
	}
	return
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
		if ok, err := f.handleError(err); ok {
			return err
		}
	}
	if err := f.checkRequiredFlag(); err != nil {
		if ok, err_r := f.handleError(err); ok {
			return err_r
		}
	}
	return nil
}