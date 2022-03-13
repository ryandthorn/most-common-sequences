package main

type file struct {
	name     string
	contents []byte
	top100   []wordGroup
}

func (f *file) Name() string {
	return f.name
}

func (f *file) Contents() []byte {
	return f.contents
}

func (f *file) ContentsToString() string {
	return string(f.contents)
}

func (f *file) SetTop100(wgs []wordGroup) {
	if len(wgs) <= 100 {
		f.top100 = wgs
	} else {
		f.top100 = wgs[:100]
	}
}

func (f *file) GetTop100() []wordGroup {
	return f.top100
}
