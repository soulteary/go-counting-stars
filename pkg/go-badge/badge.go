package badge

import (
	"bytes"
	"html/template"
	"io"
	"sync"

	"github.com/golang/freetype/truetype"
	fonts "github.com/soulteary/go-counting-stars/pkg/go-badge-fonts"
	badgeTemplate "github.com/soulteary/go-counting-stars/pkg/go-badge-template"
	"golang.org/x/image/font"
)

type badge struct {
	Subject string
	Status  string
	Color   Color
	Bounds  bounds
}

type bounds struct {
	// SubjectDx is the width of subject string of the badge.
	SubjectDx float64
	SubjectX  float64
	// StatusDx is the width of status string of the badge.
	StatusDx float64
	StatusX  float64
}

func (b bounds) Dx() float64 {
	return b.SubjectDx + b.StatusDx
}

type badgeDrawer struct {
	fd    *font.Drawer
	tmpl  *template.Template
	mutex *sync.Mutex
	name  string
}

func (d *badgeDrawer) Render(tplName, subject, status string, color Color, w io.Writer) error {
	d.mutex.Lock()
	subjectDx := d.measureString(subject)
	statusDx := d.measureString(status)
	d.mutex.Unlock()

	if tplName == "flat-square" {
		const iconPadding = 18
		return d.tmpl.Execute(w, badge{
			Subject: subject,
			Status:  status,
			Color:   color,
			Bounds: bounds{
				SubjectDx: subjectDx + iconPadding,
				SubjectX:  subjectDx/2.0 + 1 + iconPadding,
				StatusDx:  statusDx,
				StatusX:   subjectDx + statusDx/2.0 - 1 + iconPadding,
			},
		})
	}

	if tplName == "plastic" {
		const iconPadding = 18
		return d.tmpl.Execute(w, badge{
			Subject: subject,
			Status:  status,
			Color:   color,
			Bounds: bounds{
				SubjectDx: subjectDx + iconPadding,
				SubjectX:  subjectDx/2.0 + 1 + iconPadding,
				StatusDx:  statusDx,
				StatusX:   subjectDx + statusDx/2.0 - 1 + iconPadding,
			},
		})
	}

	if tplName == "for-the-badge" {
		const iconPadding = 18
		return d.tmpl.Execute(w, badge{
			Subject: subject,
			Status:  status,
			Color:   color,
			Bounds: bounds{
				SubjectDx: subjectDx + iconPadding,
				SubjectX:  subjectDx/2.0 + 1 + iconPadding,
				StatusDx:  statusDx,
				StatusX:   subjectDx + statusDx/2.0 - 1 + iconPadding,
			},
		})
	}

	// flat
	return d.tmpl.Execute(w, badge{
		Subject: subject,
		Status:  status,
		Color:   color,
		Bounds: bounds{
			SubjectDx: subjectDx,
			SubjectX:  subjectDx/2.0 + 1,
			StatusDx:  statusDx,
			StatusX:   subjectDx + statusDx/2.0 - 1,
		},
	})
}

func (d *badgeDrawer) RenderBytes(tplName, subject, status string, color Color) ([]byte, error) {
	buf := &bytes.Buffer{}
	err := d.Render(tplName, subject, status, color, buf)
	return buf.Bytes(), err
}

// shield.io uses Verdana.ttf to measure text width with an extra 10px.
// As we use Vera.ttf, we have to tune this value a little.
const extraDx = 13

func (d *badgeDrawer) measureString(s string) float64 {
	return float64(d.fd.MeasureString(s)>>6) + extraDx
}

// Render renders a badge of the given color, with given subject and status to w.
func Render(tplName, subject, status string, color Color, w io.Writer) error {
	return drawer.Render(tplName, subject, status, color, w)
}

// RenderBytes renders a badge of the given color, with given subject and status to bytes.
func RenderBytes(tplName, subject, status string, color Color) ([]byte, error) {
	return drawer.RenderBytes(tplName, subject, status, color)
}

const (
	dpi      = 72
	fontsize = 11
)

var drawer *badgeDrawer

// TODO change theme via parameters
func init() {
	// drawer = &badgeDrawer{
	// 	fd:    mustNewFontDrawer(fontsize, dpi),
	// 	tmpl:  template.Must(template.New("flat-template").Parse(stripXmlWhitespace(badgeTemplate.FlatTemplate))),
	// 	mutex: &sync.Mutex{},
	// 	name:  "flat-template",
	// }

	// drawer = &badgeDrawer{
	// 	fd:    mustNewFontDrawer(fontsize, dpi),
	// 	tmpl:  template.Must(template.New("flat-square-template").Parse(stripXmlWhitespace(badgeTemplate.FlatSquareTemplate))),
	// 	mutex: &sync.Mutex{},
	// 	name:  "flat-square",
	// }

	// drawer = &badgeDrawer{
	// 	fd:    mustNewFontDrawer(fontsize, dpi),
	// 	tmpl:  template.Must(template.New("plastic-template").Parse(stripXmlWhitespace(badgeTemplate.PlasticTemplate))),
	// 	mutex: &sync.Mutex{},
	// 	name:  "plastic",
	// }

	drawer = &badgeDrawer{
		fd:    mustNewFontDrawer(fontsize, dpi),
		tmpl:  template.Must(template.New("for-the-badge-template").Parse(stripXmlWhitespace(badgeTemplate.ForTheBadgeTemplate))),
		mutex: &sync.Mutex{},
		name:  "for-the-badge",
	}

}

func mustNewFontDrawer(size, dpi float64) *font.Drawer {
	ttf, err := truetype.Parse(fonts.VeraSans)
	if err != nil {
		panic(err)
	}
	return &font.Drawer{
		Face: truetype.NewFace(ttf, &truetype.Options{
			Size:    size,
			DPI:     dpi,
			Hinting: font.HintingFull,
		}),
	}
}
