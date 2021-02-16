package codegen

const bpfProgramForNoArgs = `
#include <uapi/linux/ptrace.h>

BPF_PERF_OUTPUT(output);

int golang_func_trace(struct pt_regs *ctx) {
    u8 called =  1;
    output.perf_submit(ctx, &called, sizeof(called));
    return 0;
}
`
