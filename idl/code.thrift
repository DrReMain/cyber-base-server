namespace go common.code

enum Code {
    Success                 = 0
    InternalErr             = 100000
    ParamsInvalidErr        = 100001
    AuthorizationFailedErr  = 100002
    AlreadyExistErr         = 100003
    NotExistErr             = 100004
}
