package bridge

// Mac 扩展抽象化角色
type Mac struct {
	Video
}

func (w *Mac) PlayVideo(fileName string) {
	w.Video.DecodeVideo(fileName)
}
