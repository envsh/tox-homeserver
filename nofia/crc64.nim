
{.compile:"crc64.c.ngo".}

proc crc64*(crc: uint64, s : cstring, length:  uint64) : uint64 {.importc.}

