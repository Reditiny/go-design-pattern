package bridge

import "fmt"

// Avi 播放器的具体实现化
type Avi struct {
}

func (r *Avi) DecodeVideo(fileName string) {
	fmt.Printf("Avi Decoding %s\n", fileName)
}
