package globals

type ErrCode struct {
	Code    int
	Message string
}

var RespOK = ErrCode{
	Code:    0,
	Message: "ok",
}

//ErrParamValidException
//基础错误类型定义
var (
	ErrParamValidException = ErrCode{Code: 101, Message: "param is not valid"}
	ErrUnknownException    = ErrCode{Code: 102, Message: "unknown exception"}
	ErrSystemException     = ErrCode{Code: 103, Message: "system exception"}
	ErrDbAddException      = ErrCode{Code: 104, Message: "add data to mysql failed"}
	ErrDbUpdateException   = ErrCode{Code: 105, Message: "update data to mysql failed"}
	ErrDbDeleteException   = ErrCode{Code: 106, Message: "delete data from mysql failed"}
	ErrDbQueryException    = ErrCode{Code: 107, Message: "query data from mysql failed"}
	ErrDuplicateException  = ErrCode{Code: 108, Message: "data duplicate exception"}
)

var (
	ErrAuthUserUnauthorized = ErrCode{Code: 1010401, Message: "User Unauthorized"}
	ErrAuthUserForbidden    = ErrCode{Code: 1010403, Message: "User Forbidden"}
	ErrRedisConnectFailed   = ErrCode{Code: 1010404, Message: "Redis Connect Failed"}
	ErrGoogleReqFailed      = ErrCode{Code: 1010405, Message: "Send Google Req Failed"}
	ErrUnmarshalFailed      = ErrCode{Code: 1010405, Message: "Unmarshal failed"}
	ErrDBConnectFailed      = ErrCode{Code: 1010406, Message: "DB connect failed"}
)

var (
	ErrReportFormatInvalid       = ErrCode{Code: 1011001, Message: "Report Format Invalid"}
	ErrReportAppInvalid          = ErrCode{Code: 1011002, Message: "Report App Invalid"}
	ErrReportFeedbackTypeInvalid = ErrCode{Code: 1011003, Message: "Report FeedbackType Invalid"}
	ErrReportCreateFailed        = ErrCode{Code: 1011004, Message: "Report Create Failed"}
)

var (
	ErrModifyFormatInvalid        = ErrCode{Code: 1012001, Message: "Modify Format Invalid"}
	ErrModifyStatusInvalid        = ErrCode{Code: 1012002, Message: "Modify Status Invalid"}
	ErrModifyLevelInvalid         = ErrCode{Code: 1012003, Message: "Modify Level Invalid"}
	ErrModifyQuestionEpochInvalid = ErrCode{Code: 1012004, Message: "Modify QuestionEpoch Invalid"}
	ErrModifyCategoryInvalid      = ErrCode{Code: 1012005, Message: "Modify Category Invalid"}
	ErrModifyProcessorInvalid     = ErrCode{Code: 1012006, Message: "Modify Processor Invalid"}
	ErrModifyModifyFailed         = ErrCode{Code: 1012007, Message: "Modify Failed"}
)

var (
	ErrQueryFormatInvalid        = ErrCode{Code: 1013001, Message: "Query Format Invalid"}
	ErrQueryStatusInvalid        = ErrCode{Code: 1013002, Message: "Query Status Invalid"}
	ErrQueryLevelInvalid         = ErrCode{Code: 1013003, Message: "Query Level Invalid"}
	ErrQueryQuestionEpochInvalid = ErrCode{Code: 1013004, Message: "Query QuestionEpoch Invalid"}
	ErrQueryCategoryInvalid      = ErrCode{Code: 1013005, Message: "Query Category Invalid"}
	ErrQueryProcessorInvalid     = ErrCode{Code: 1013006, Message: "Query Processor Invalid"}
	ErrQueryFailed               = ErrCode{Code: 1013007, Message: "Query Failed"}
)

var (
	ErrDetailFormatInvalid = ErrCode{Code: 1014001, Message: "Detail Format Invalid"}
	ErrDetailFailed        = ErrCode{Code: 1014002, Message: "Detail Failed"}
)

var (
	ErrConfigListTypeInvalid = ErrCode{Code: 1015001, Message: "ConfigList Type Invalid"}
	ErrConfigListFailed      = ErrCode{Code: 1015002, Message: "ConfigList Failed"}
)

var (
	ErrConfigDetailTypeInvalid = ErrCode{Code: 1016001, Message: "ConfigDetail Type Invalid"}
	ErrConfigDetailIDInvalid   = ErrCode{Code: 1016002, Message: "ConfigDetail ID Invalid"}
	ErrConfigDetailFailed      = ErrCode{Code: 1016003, Message: "ConfigDetail Failed"}
)

var (
	ErrConfigCreateTypeInvalid   = ErrCode{Code: 1017001, Message: "ConfigCreate Type Invalid"}
	ErrConfigCreateFormatInvalid = ErrCode{Code: 1017002, Message: "ConfigCreate Format Invalid"}
	ErrConfigCreateFailed        = ErrCode{Code: 1017003, Message: "ConfigCreate Failed"}
)

var (
	ErrConfigModifyTypeInvalid   = ErrCode{Code: 1018001, Message: "ConfigModify Type Invalid"}
	ErrConfigModifyIDInvalid     = ErrCode{Code: 1018002, Message: "ConfigModify ID Invalid"}
	ErrConfigModifyFormatInvalid = ErrCode{Code: 1018003, Message: "ConfigModify Format Invalid"}
	ErrConfigModifyFailed        = ErrCode{Code: 1018004, Message: "ConfigModify Failed"}
)

var (
	ErrConfigDeleteTypeInvalid = ErrCode{Code: 1019001, Message: "ConfigDelete Type Invalid"}
	ErrConfigDeleteIDInvalid   = ErrCode{Code: 1019002, Message: "ConfigDelete ID Invalid"}
	ErrConfigDeleteFailed      = ErrCode{Code: 1019003, Message: "ConfigDelete Failed"}
)

var (
	ErrConfigFetchTokenInvalid  = ErrCode{Code: 1020001, Message: "ConfigFetch Token Invalid"}
	ErrConfigFetchFailed        = ErrCode{Code: 1020002, Message: "ConfigFetch Failed"}
	ErrConfigFetchEncodeFailed  = ErrCode{Code: 1020003, Message: "ConfigFetch Encode Failed"}
	ErrConfigFetchMarshalFailed = ErrCode{Code: 1020004, Message: "ConfigFetch Marshal Failed"}
	ErrConfigFetchHashFailed    = ErrCode{Code: 1020005, Message: "ConfigFetch Hash Failed"}
)

var (
	ErrUploadPresignKeyInvalid = ErrCode{Code: 1021001, Message: "UploadPresign Key Invalid"}
	ErrUploadPresignFailed     = ErrCode{Code: 1021002, Message: "UploadPresign Failed"}
)

var (
	ErrInsertMarkModelFailed       = ErrCode{Code: 1022001, Message: "Insert Mark Model Failed"}
	ErrUpdateMarkModelFailed       = ErrCode{Code: 1022002, Message: "Update Mark Model Failed"}
	ErrDeleteMarkModelFailed       = ErrCode{Code: 1022003, Message: "Delete Mark Model Failed"}
	ErrMarkModelRepeat             = ErrCode{Code: 1022004, Message: "Mark Model Already Exist"}
	ErrCancelDeleteMarkModelFailed = ErrCode{Code: 1022005, Message: "Cancel Delete Mark Model Failed"}
)

var (
	ErrInsertReplyModelFailed = ErrCode{Code: 1023001, Message: "Insert Reply Model Failed"}
	ErrUpdateReplyModelFailed = ErrCode{Code: 1023002, Message: "Update Reply Model Failed"}
	ErrDeleteReplyModelFailed = ErrCode{Code: 1023003, Message: "Delete Reply Model Failed"}
)

var (
	ErrGoogleReplyFailed = ErrCode{Code: 1024001, Message: "Google Reply failed"}
)
