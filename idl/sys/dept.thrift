namespace go sys

include "../base.thrift"
include "../pagination.thrift"

struct Dept {
    1: required string   id
    2: required string   dept_name
    3: required string   remark
}
struct PResult {
    1: required pagination.P    p
    2: required list<Dept>      list
}

struct CreateDeptReq {
    1: optional string dept_name    (api.body="dept_name", api.vd="(len($)>0 && len($)<=100);msg:'部门名称不能为空也不能多于100个字符'")
    2: optional string remark       (api.body="remark", api.vd="len($)<=500")
}
struct CreateDeptRes {
    1: required base.Base   base
    2: optional bool        result
}

struct UpdateDeptReq {
    1: required string id           (api.path="id")
    2: optional string dept_name    (api.body="dept_name", api.vd="len($)<=100")
    3: optional string remark       (api.body="remark", api.vd="len($)<=500")
}
struct UpdateDeptRes {
    1: required base.Base   base
    2: optional bool        result
}

struct DeleteDeptReq {
    1: required string id           (api.path="id")
}
struct DeleteDeptRes {
    1: required base.Base   base
    2: optional bool        result
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
    1: required string id           (api.path="id")
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
