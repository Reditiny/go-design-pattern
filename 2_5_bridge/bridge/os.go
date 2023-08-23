package bridge

// OperationSystem 抽象化角色 后续扩展操作系统平台 实现这个接口即可
type OperationSystem interface {
	PlayVideo(fileName string)
}
