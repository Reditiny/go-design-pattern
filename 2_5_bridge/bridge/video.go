package bridge

// Video 实现化接口 定义视频播放器的解码标准 后续扩展文件类型 实现这个接口即可
type Video interface {
	DecodeVideo(fileName string)
}
