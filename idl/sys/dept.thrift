namespace go sys

include "../base.thrift"
include "../pagination.thrift"

struct Null {}
struct Dept {
    1: i64      id
    2: string   uuid
    3: string   dept_name
    4: string   remark
}
struct PResult {
    1: required pagination.P    p
    2: required list<Dept>      list
}

struct CreateDeptReq {
    1: optional string dept_name    (api.body="dept_name", api.vd="(len($)>0 && len($)<100);msg:'参数错误'")
    2: optional string remark       (api.body="remark", api.vd="len($) < 500;msg:'参数错误'")
}
struct CreateDeptRes {
    1: required base.Base   base
    2: required Null        result
}

struct UpdateDeptReq {
    1: required i64 id              (api.path="id", api.vd="$>0;msg:'参数错误'")
    2: optional string dept_name    (api.body="dept_name", api.vd="len($) < 100;msg:'参数错误'")
    3: optional string remark       (api.body="remark", api.vd="len($) < 500;msg:'参数错误'")
}
struct UpdateDeptRes {
    1: required base.Base   base
    2: required Null        result
}

struct DeleteDeptReq {
    1: required i64 id              (api.path="id", api.vd="$>0;msg:'参数错误'")
}
struct DeleteDeptRes {
    1: required base.Base   base
    2: required Null        result
}

struct QueryAllDeptReq {
    1: optional string dept_name    (api.query="dept_name")
}
struct QueryAllDeptRes {
    1: required base.Base   base
    2: required list<Dept>  result
}

struct QueryListDeptReq {
    1: optional i32     page_size   (api.query="page_size")
    2: optional i32     page_num    (api.query="page_num")
    3: optional string  dept_name   (api.query="dept_name")
}
struct QueryListDeptRes {
    1: required base.Base   base
    2: required PResult     result
}

struct QueryItemDeptReq {
    1: required i64 id              (api.path="id", api.vd="$>0;msg:'参数错误'")
}
struct QueryItemDeptRes {
    1: required base.Base   base
    2: required Dept        result
}

service DeptService {
    CreateDeptRes       CreateDept(1:       CreateDeptReq       req) (api.post=     "/v1/sys/dept/create")
    UpdateDeptRes       UpdateDept(1:       UpdateDeptReq       req) (api.put=      "/v1/sys/dept/:id")
    DeleteDeptRes       DeleteDept(1:       DeleteDeptReq       req) (api.delete=   "/v1/sys/dept/:id")
    QueryAllDeptRes     QueryAllDept(1:     QueryAllDeptReq     req) (api.get=      "/v1/sys/dept/list_all")
    QueryListDeptRes    QueryListDept(1:    QueryListDeptReq    req) (api.get=      "/v1/sys/dept/list")
    QueryItemDeptRes    QueryItemDept(1:    QueryItemDeptReq    req) (api.get=      "/v1/sys/dept/:id")
}
