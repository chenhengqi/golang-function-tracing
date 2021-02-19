package codegen

import (
	"bytes"
	"text/template"
)

type preambleDef struct {
	Includes []string
	Defs     []string
	Funcs    []string
}

type codeDef struct {
	Preamble string
	Fields   []string
	Snippets []string
}

const preambleTmpl = `
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

const codeTmpl = `
{{.Preamble}}

BPF_PERF_OUTPUT(output);

struct event {
{{range $field := .Fields}}
    {{$field}}
{{end}}
};

int golang_func_trace(struct pt_regs *ctx) {
    struct event ev = {0};

{{range $snippet := .Snippets}}
    {{$snippet}}
{{end}}

    output.perf_submit(ctx, &ev, sizeof(ev));
    return 0;
}
`

func preamble() string {
	c := preambleDef{
		Includes: []string{ptraceHeader},
		Defs:     []string{stringDef},
		Funcs:    []string{primitveInspectFunc, ptrInspectFunc, stringInspectFunc},
	}
	t := template.Must(template.New("preamble").Parse(preambleTmpl))
	buf := bytes.Buffer{}
	err := t.Execute(&buf, &c)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func code(fileds, snippets []string) string {
	c := codeDef{
		Preamble: preamble(),
		Fields:   fileds,
		Snippets: snippets,
	}
	t := template.Must(template.New("code").Parse(codeTmpl))
	buf := bytes.Buffer{}
	err := t.Execute(&buf, &c)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
