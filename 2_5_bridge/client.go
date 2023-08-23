package main

import "design_pattern/2_5_bridge/bridge"

func main() {
	// video 和 os 的具体实现是独立的
	var video bridge.Video
	var os bridge.OperationSystem
	// mac 上播放 rmv 文件
	video = new(bridge.Rmv)
	os = &bridge.Mac{Video: video}
	os.PlayVideo("test1.mp4")
	// windows 上播放 avi 文件
	video = new(bridge.Avi)
	os = &bridge.Windows{Video: video}
	os.PlayVideo("test2.mp4")

}
