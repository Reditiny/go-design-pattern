package bridge

// Windows 扩展抽象化角色
type Windows struct {
	Video
}

func (w *Windows) PlayVideo(fileName string) {
	w.Video.DecodeVideo(fileName)
}
