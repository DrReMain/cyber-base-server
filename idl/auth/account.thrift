namespace go auth.account

include "../base.thrift"

struct RegisterAccountReq {
    1: required string      email       (api.body="email", api.vd="regexp('^[a-zA-Z0-9_\\.]+@[a-zA-Z0-9-]+[\\.a-zA-Z]+$');msg:'Email格式错误'")
    2: required string      password    (api.body="password", api.vd="len($)>0;msg:'密码不能为空'")
    3: required string      confirm     (api.body="confirm", api.vd="len($)>0;msg:'确认密码不能为空'")
}
struct RegisterAccountRes {
    1: required base.Base   base
    2: optional bool        result
}

struct LoginAccountReq {
    1: required string email            (api.body="email", api.vd="regexp('^[a-zA-Z0-9_\\.]+@[a-zA-Z0-9-]+[\\.a-zA-Z]+$');msg:'Email格式错误'")
    2: required string password         (api.body="password", api.vd="len($)>0;msg:'密码不能为空'")
}
struct LoginAccountResult {
    1: required string access_token
}
struct LoginAccountRes {
    1: required base.Base           base
    2: optional LoginAccountResult  result
}

service AccountHandler {
    RegisterAccountRes Register(1: RegisterAccountReq req) (api.post="/v1/auth/register")
    LoginAccountRes    Login(1:    LoginAccountReq    req) (api.post="/v1/auth/login")
}
