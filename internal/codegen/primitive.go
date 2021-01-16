package codegen

const primitveInspectFunc = `
// for bool, int8, uint8, byte
void primitive8_inspect(u64 *base_addr, u8 *dest) {
    bpf_probe_read(dest, sizeof(u8), base_addr);
}

// for int16, uint16
void primitive16_inspect(u64 *base_addr, u16 *dest) {
    bpf_probe_read(dest, sizeof(u16), base_addr);
}

// for int32, uint32, rune, float32
void primitive32_inspect(u64 *base_addr, u32 *dest) {
    bpf_probe_read(dest, sizeof(u32), base_addr);
}

// for int64, uint64, int, uint, uintptr, float64
// for complex64, complex128
void primitive64_inspect(u64 *base_addr, u64 *dest) {
    bpf_probe_read(dest, sizeof(u64), base_addr);
}
`
