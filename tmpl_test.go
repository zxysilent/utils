package utils

import (
	"bytes"
	"testing"
)

func TestLoadTmpl(t *testing.T) {
	tmpl := LoadTmpl("./testdata/views",nil)
	if tmpl == nil {
		t.Error("error")
	}
	tmpls := tmpl.Templates()
	for i := 0; i < len(tmpls); i++ {
		t.Log(tmpls[i].Name())
	}
}

func TestLoadTmplExec(t *testing.T) {
	mod := struct {
		Name string
		Arr  []int
	}{
		Name: "testName",
		Arr:  []int{1, 3, 5, 7},
	}
	tmpl := LoadTmpl("./testdata/views",nil)
	if tmpl == nil {
		t.Error("error")
	}
	w := bytes.NewBuffer(nil)
	tmpl.ExecuteTemplate(w, "subtmpl/subtmpl.html", mod)
	t.Log(w.String())
}
