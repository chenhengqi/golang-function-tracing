package codegen

import (
	"bytes"
	"text/template"
)

type code struct {
	Includes []string
	Defs     []string
	Funcs    []string
}

const codeTmpl = `
{{range $inc := .Includes}}
{{$inc}}
{{end}}

{{range $def := .Defs}}
{{$def}}
{{end}}

{{range $fn := .Funcs}}
{{$fn}}
{{end}}
`

func preamble() string {
	c := code{
		Includes: []string{ptraceHeader},
		Defs:     []string{stringDef},
		Funcs:    []string{primitveInspectFunc, ptrInspectFunc, stringInspectFunc},
	}
	t := template.Must(template.New("code").Parse(codeTmpl))
	buf := bytes.Buffer{}
	err := t.Execute(&buf, &c)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
