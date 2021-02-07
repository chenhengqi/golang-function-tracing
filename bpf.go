package main

const (
	bpfTableName    = "output"
	golangFuncProbe = "golang_func_trace"
)

const bpfProgram0 = `
#include <uapi/linux/ptrace.h>

struct event {
    u64 len;
    char data[24];
};

BPF_PERF_OUTPUT(output);

int golang_func_trace(struct pt_regs *ctx) {
    struct event ev = {0};
    u64* sp = (u64*)ctx->sp;
    // string length
    bpf_probe_read(&ev.len, sizeof(ev.len), sp + 2);
    // string content
    u64 addr = 0;
    bpf_probe_read(&addr, sizeof(addr), sp + 1);
    bpf_probe_read(ev.data, sizeof(ev.data), (u64*)addr);
    output.perf_submit(ctx, &ev, sizeof(ev));
    return 0;
}
`
