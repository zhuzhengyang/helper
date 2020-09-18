package helper

type Foo struct {
}

func (f *Foo) Init() {
	f.bar()
	f.Public()
}

//go:noinline
func (f *Foo) bar() string {
	return "o bar"
}

//go:noinline
func (f *Foo) Public() string {
	return "o public"
}
