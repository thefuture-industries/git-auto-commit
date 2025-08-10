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
	Err_GitNotInstalled         = errors.New("git is not installed or not found in PATH")
	Err_GitNotRepository        = errors.New("current directory is not a git repository")
	Err_NoStagedFiles           = errors.New("no staged files to commit")
	Err_FailedToGetDiff         = errors.New("failed to get file diff")
	Err_FailedToCreateCommit    = errors.New("failed to create commit")
	Err_FailedToReadFile        = errors.New("error reading file")
	Err_EmptyCommitMessage      = errors.New("commit message is empty")
	Err_InvalidConfig           = errors.New("invalid configuration file")
	Err_RemoteUnavailable       = errors.New("remote repository is unavailable")
	Err_BranchNotFound          = errors.New("current branch not found")
)
