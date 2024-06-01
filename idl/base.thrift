namespace go common.base

include "code.thrift"

struct Base {
    1: required i64         t
    2: required bool        success
    3: required code.Code   code
    4: required string      msg
}
