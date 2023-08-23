package main

import "design_pattern/2_1_adapter/adapter"

func main() {
	TF1 := adapter.SpecificTF{}
	tfAdapterSD := &adapter.TFAdapterSD{SpecificTF: TF1}
	computer1 := &adapter.Computer{}
	// 通过适配器可以让原本只能读取SD卡的电脑也能读取TF卡
	computer1.ReadComputerSD(tfAdapterSD)
	computer1.WriteComputerSD(tfAdapterSD, "new data")
}
