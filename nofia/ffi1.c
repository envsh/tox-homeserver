#include <ffi.h>

int ffi_get_default_abi() { return FFI_DEFAULT_ABI; }
int ffi_type_size() { return sizeof(ffi_type); }
int ffi_cif_size() { return sizeof(ffi_cif); }

void dump_pointer_array(int n, void** ptr) {
    for (int i = 0;i < n; i ++) {
        printf("%p %d, = %p\n", ptr, i, ptr[i]);
    }
}
