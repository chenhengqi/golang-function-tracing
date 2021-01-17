package codegen

const ptrInspectFunc = `
// str should be zero-initialized
void ptr_inspect(u64 *base_addr, u64 *ptr_val) {
    bpf_probe_read(ptr_val, sizeof(u64), base_addr);
}
`
