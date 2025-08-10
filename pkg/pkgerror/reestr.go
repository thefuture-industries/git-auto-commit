package pkgerror

import "errors"

var (
	Err_FileNotFound            = errors.New("the file was not found")
	Err_DirectoryNotFound       = errors.New("the directory was not found")
	Err_Timeout                 = errors.New("connection exceeded the waiting time")
	Err_NetworkTemporary        = errors.New("there is a temporary network problem, please try again")
	Err_Network                 = errors.New("network error when connecting")
	Err_ContextCanceled         = errors.New("the operation was cancelled")
	Err_ContextDeadlineExceeded = errors.New("waiting for a long time, please try again later")
)
