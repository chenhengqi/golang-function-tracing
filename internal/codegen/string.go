package codegen

const stringDef = `
struct string {
    u64 len;
    char data[24];
};
`

const stringInspect = `
// str should be initialized
void string_inspect(u64 *base_addr, struct string *str) {
    // string length
    bpf_probe_read(&str.len, sizeof(str.len), base_addr + 2);
    // string content
    u64 addr = 0;
    bpf_probe_read(&addr, sizeof(addr), base_addr + 1);
    bpf_probe_read(str.data, sizeof(str.data), (u64*)addr);
}
`
