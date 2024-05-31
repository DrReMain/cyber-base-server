namespace go common.base

include "code.thrift"

struct Base {
    1: i64 t
    2: bool success
    3: code.Code code
    4: string msg
}
