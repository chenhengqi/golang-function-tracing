package codegen

import "testing"

func TestPreamble(t *testing.T) {
	t.Log(preamble())
}

func TestCode(t *testing.T) {
	fields := []string{
		"u64 x1;",
		"char x2[128];",
	}
	snippets := []string{
		`
    u64* sp = (u64*)ctx->sp;
    // string length
    bpf_probe_read(&ev.len, sizeof(ev.len), sp + 2);
    // string content
    u64 addr = 0;
    bpf_probe_read(&addr, sizeof(addr), sp + 1);
    bpf_probe_read(ev.data, sizeof(ev.data), (u64*)addr);
        `,
	}
	t.Log(code(fields, snippets))
}
