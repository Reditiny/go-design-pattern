package bridge

import "fmt"

// Rmv 播放器的具体实现化
type Rmv struct {
}

func (r *Rmv) DecodeVideo(fileName string) {
	fmt.Printf("Rmv Decoding %s\n", fileName)
}
