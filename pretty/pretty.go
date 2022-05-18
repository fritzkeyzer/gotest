// Package pretty prints values.
package pretty

import (
	"fmt"
	"io"
	"reflect"
	"strings"
)

var Indent = "   "

type Flag int

const (
	FlagPointerAddress Flag = -iota
	FlagSingleLine          = -iota
	FlagEscapeString        = -iota
)

func FlagMaxDepth(depth int) Flag {
	return Flag(depth)
}

func String(value interface{}, flags ...Flag) string {
	w := &strings.Builder{}
	Write(w, "", "", value, flags...)
	return w.String()
}

func Print(value interface{}, flags ...Flag) {
	fmt.Println(String(value, flags...))
}

func Write(w io.Writer, prefix, indent string, value interface{}, flags ...Flag) {
	var pointerAddress, singleLine, escapeString bool
	var maxDepth int
	for _, flag := range flags {
		if flag >= 0 {
			maxDepth = int(flag)
			continue
		}
		switch flag {
		case FlagPointerAddress:
			pointerAddress = true
		case FlagSingleLine:
			singleLine = true
		case FlagEscapeString:
			escapeString = true

		}
	}
	write(w, "", reflect.ValueOf(value), 0, maxDepth, pointerAddress, singleLine, escapeString)
}

func write(w io.Writer, prefix string, v reflect.Value, depth, maxDepth int, pointerAddress, singleLine, escapeString bool) {

	if maxDepth > 0 && depth > maxDepth {
		fmt.Fprintf(w, "...")
		return
	}

	if v.CanInterface() {
		value := v.Interface()
		if stringer, ok := value.(fmt.Stringer); ok {
			if v.Kind() == reflect.Ptr && v.IsNil() {
				// do nothing, let nil pointer below handle code
			} else {
				fmt.Fprintf(w, stringer.String())
				return
			}
		}
	}

	switch v.Kind() {
	case reflect.Bool:
		fmt.Fprintf(w, "%t", v.Bool())

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Fprintf(w, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		fmt.Fprintf(w, "%d", v.Uint())

	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(w, "%f", v.Float())

	case reflect.String:
		if escapeString {
			fmt.Fprintf(w, "%q", v.String())
		} else {
			fmt.Fprintf(w, "\"%s\"", v.String())
		}

	case reflect.Ptr:
		if v.IsNil() {
			fmt.Fprint(w, "nil")
			return
		}
		if pointerAddress {
			fmt.Fprintf(w, "%p -> ", v.Interface())
		}
		write(w, prefix, v.Elem(), depth+1, maxDepth, pointerAddress, singleLine, escapeString)

	case reflect.Interface:
		if v.IsNil() {
			fmt.Fprint(w, "nil")
			return
		}
		write(w, prefix, v.Elem(), depth+1, maxDepth, pointerAddress, singleLine, escapeString)

	case reflect.Slice:
		if v.IsNil() {
			fmt.Fprint(w, "nil")
			return
		}
		fallthrough

	case reflect.Array:
		if v.Type().Elem().Kind() == reflect.Uint8 {
			fmt.Fprintf(w, "%d bytes", v.Len())
			return
		}

		fmt.Fprintf(w, "[")
		for i := 0; i < v.Len(); i += 1 {
			if !singleLine {
				fmt.Fprintf(w, "\n%s", prefix+Indent)
			}
			write(w, prefix+Indent, v.Index(i), depth+1, maxDepth, pointerAddress, singleLine, escapeString)
			fmt.Fprintf(w, ", ")
		}
		if !singleLine {
			fmt.Fprintf(w, "\n%s", prefix)
		}
		fmt.Fprintf(w, "]")

	case reflect.Map:
		fmt.Fprintf(w, "{")
		for _, key := range v.MapKeys() {
			if !singleLine {
				fmt.Fprintf(w, "\n%s", prefix+Indent)
			}
			write(w, prefix+Indent, key, depth+1, maxDepth, pointerAddress, singleLine, escapeString)
			fmt.Fprintf(w, ": ")
			write(w, prefix+Indent, v.MapIndex(key), depth+1, maxDepth, pointerAddress, singleLine, escapeString)
			fmt.Fprintf(w, ", ")
		}
		if !singleLine {
			fmt.Fprintf(w, "\n%s", prefix)
		}
		fmt.Fprintf(w, "}")

	case reflect.Struct:
		fmt.Fprintf(w, "{")
		for i := 0; i < v.NumField(); i += 1 {
			if !singleLine {
				fmt.Fprintf(w, "\n%s", prefix+Indent)
			}
			fmt.Fprintf(w, "%s: ", v.Type().Field(i).Name)
			write(w, prefix+Indent, v.Field(i), depth+1, maxDepth, pointerAddress, singleLine, escapeString)
			fmt.Fprintf(w, " ")
		}
		if !singleLine {
			fmt.Fprintf(w, "\n%s", prefix)
		}
		fmt.Fprintf(w, "}")

	default:
		fmt.Fprintf(w, "%s (no pretty printer)", v.Type().String())
	}
}
