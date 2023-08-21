package phone

type Builder struct {
	Phone
}

func (b *Builder) Cpu(c string) *Builder {
	b.cpu = c
	return b
}

func (b *Builder) Screen(s string) *Builder {
	b.screen = s
	return b
}

func (b *Builder) Memory(m string) *Builder {
	b.memory = m
	return b
}

func (b *Builder) Build() *Phone {
	return &Phone{cpu: b.cpu, memory: b.memory, screen: b.screen}
}
