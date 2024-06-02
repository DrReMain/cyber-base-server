namespace go template.greet

include "../base.thrift"

struct GreetReq {
    1: required string name  (
        api.path="name",
        api.vd="($ == 'you' || $ == 'me'); msg: '参数只能是you或者me'"
    )
}

struct Result {
    1: required string text
}
struct GreetRes {
    1: required base.Base base
    2: required Result result
}

service GreetHandler {
    GreetRes Greet(1: GreetReq req) (api.get="/v1/template/greet/:name_content")
}
