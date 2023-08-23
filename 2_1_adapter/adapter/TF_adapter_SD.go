package adapter

// TFAdapterSD 适配器类 需要组合原接口并实现目标接口
type TFAdapterSD struct {
	SpecificTF // 此处包含具体实现或原接口均可
}

func (s *TFAdapterSD) ReadSD() string {
	return s.SpecificTF.ReadTF()
}

func (s *TFAdapterSD) WriteSD(content string) {
	s.SpecificTF.WriteTF(content)
}
