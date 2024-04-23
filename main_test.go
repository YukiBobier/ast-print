package main

import (
	"bytes"
	"embed"
	"testing"
)

//go:embed testdata
var testdata embed.FS

var helloWorldGo []byte
var helloWorldWantTxt []byte

func init() {
	var err error
	helloWorldGo, err = testdata.ReadFile("testdata/helloworld.go")
	if err != nil {
		panic(err)
	}
	helloWorldWantTxt, err = testdata.ReadFile("testdata/helloworld.want.txt")
	if err != nil {
		panic(err)
	}
}

func TestRun(t *testing.T) {
	in := bytes.NewReader(helloWorldGo)
	out := new(bytes.Buffer)
	run(in, out)

	got := out.Bytes()
	want := helloWorldWantTxt
	if !bytes.Equal(got, want) {
		t.Errorf("got %s, want %s", got, want)
	}
}
