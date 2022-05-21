package e

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "invalid params",

	ERROR_NOT_EXIST_JOB: "job does not exist",
	ERROR_ADD_JOB_FAIL:  "failed to add the job",
	ERROR_GET_JOB_FAIL:  "failed to get the job",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
