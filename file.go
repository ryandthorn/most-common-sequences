package main

import (
	"fmt"
	"strings"
)

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

func (f *file) Top100ToString() string {
	b := strings.Builder{}
	b.WriteString(fmt.Sprintf("\nResults for %s\n", f.Name()))
	b.WriteString("---------------------------------\n")
	for _, wg := range f.GetTop100() {
		b.WriteString(fmt.Sprintf("%s: %d\n", wg.Text(), wg.Count()))
	}
	return b.String()
}
