package main

type file struct {
	name     string
	contents []byte
}

func (f *file) Name() string {
	return f.name
}

func (f *file) Contents() []byte {
	return f.contents
}
