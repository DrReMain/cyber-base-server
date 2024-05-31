namespace go template.greet

include "../base.thrift"

struct GreetReq {
    1: string name_content  (
        api.path="name_content",
        api.vd="($ == 'you' || $ == 'me'); msg: '参数只能是you或者me'"
    )
}

struct Result {
    1: string text_content
}
struct GreetRes {
    1: required base.Base base
    2: required Result result
}

service GreetService {
    GreetRes Greet(1: GreetReq req) (api.get="/v1/template/greet/:name_content")
}
