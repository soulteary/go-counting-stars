package badge

import (
	"bytes"
	"html/template"
	"io"
	"strings"
	"sync"
	"testing"
)

func TestBadgeDrawerRender(t *testing.T) {
	const tplName = "flat"
	mockTemplate := strings.TrimSpace(`
	{{.Subject}},{{.Status}},{{.Color}},{{with .Bounds}}{{.SubjectX}},{{.SubjectDx}},{{.StatusX}},{{.StatusDx}},{{.Dx}}{{end}}
	`)
	mockFontSize := 11.0
	mockDPI := 72.0

	d := &badgeDrawer{
		fd:    mustNewFontDrawer(mockFontSize, mockDPI),
		tmpl:  template.Must(template.New("mock-template").Parse(mockTemplate)),
		mutex: &sync.Mutex{},
	}

	output := "XXX,YYY,#c0c0c0,18,34,50,34,68"

	var buf bytes.Buffer
	err := d.Render(tplName, "XXX", "YYY", "#c0c0c0", &buf)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	result := buf.String()
	if result != output {
		t.Errorf("expect %q got %q", output, result)
	}
}

func TestBadgeDrawerRenderBytes(t *testing.T) {
	const tplName = "flat"
	mockTemplate := strings.TrimSpace(`
	{{.Subject}},{{.Status}},{{.Color}},{{with .Bounds}}{{.SubjectX}},{{.SubjectDx}},{{.StatusX}},{{.StatusDx}},{{.Dx}}{{end}}
	`)
	mockFontSize := 11.0
	mockDPI := 72.0

	d := &badgeDrawer{
		fd:    mustNewFontDrawer(mockFontSize, mockDPI),
		tmpl:  template.Must(template.New("mock-template").Parse(mockTemplate)),
		mutex: &sync.Mutex{},
	}

	output := "XXX,YYY,#c0c0c0,18,34,50,34,68"

	bytes, err := d.RenderBytes(tplName, "XXX", "YYY", "#c0c0c0")
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if string(bytes) != output {
		t.Errorf("expect %q got %q", output, string(bytes))
	}
}

func BenchmarkRender(b *testing.B) {
	const tplName = "flat"
	// warm up
	Render(tplName, "XXX", "YYY", ColorBlue, io.Discard)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := Render(tplName, "XXX", "YYY", ColorBlue, io.Discard)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkRenderParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		const tplName = "flat"
		for pb.Next() {
			err := Render(tplName, "XXX", "YYY", ColorBlue, io.Discard)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}
