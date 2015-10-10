package godot

import "strings"

func (o Options) export(indent, separator string) string {
	ret := make([]string, len(o))
	i := 0
	for k, v := range o {
		ret[i] = indent + k + "=" + escape(v)
		i++
	}
	return strings.Join(ret, separator)
}

// AsBlock exports options as a block of text.
func (o Options) AsBlock(indent string) string {
	if o == nil || len(o) == 0 {
		return ""
	}
	return o.export(indent, ";\n") + ";"
}

// AsLine exports options as single line text.
func (o Options) AsLine() string {
	if o == nil || len(o) == 0 {
		return ""
	}
	return "[" + o.export("", ", ") + "]"
}
