package cartovl

import "github.com/gopherjs/gopherjs/js"

type MapOptions struct {
	Container       string
	Style           string
	Center          []interface{}
	Zoom            int
	ScrollZoom      bool
	DragRotate      bool
	TouchZoomRotate bool
}

type Map struct {
	*js.Object
}

// NewMap creates a new map
func NewMap(mo *MapOptions) *Map {
	options := js.M{
		"container":       mo.Container,
		"style":           mo.Style,
		"center":          mo.Center,
		"zoom":            mo.Zoom,
		"scrollZoom":      mo.ScrollZoom,
		"dragRotate":      mo.DragRotate,
		"touchZoomRotate": mo.TouchZoomRotate,
	}
	return &Map{Object: js.Global.Get("mapboxgl").Get("Map").New(options)}
}

type Dataset struct {
	*js.Object
}

func NewDataset(name string) *Dataset {
	return &Dataset{Object: js.Global.Get("carto").Get("source").Get("Dataset").New(name)}
}

type Viz struct {
	*js.Object
}

func NewDefaultViz() *Viz {
	return &Viz{Object: js.Global.Get("carto").Get("Viz").New()}
}

func NewViz(style string) *Viz {
	return &Viz{Object: js.Global.Get("carto").Get("Viz").New(style)}
}

func SetDefaultAuth(user, apiKey string) {
	carto := js.Global.Get("carto")
	carto.Call("setDefaultAuth", js.M{"user": user, "apiKey": apiKey})
}

type Layer struct {
	*js.Object
}

func (l *Layer) AddTo(m *Map, name string) {
	l.Call("addTo", m, name)
}

func NewLayer(name string, source *Dataset, viz *Viz) *Layer {
	return &Layer{Object: js.Global.Get("carto").Get("Layer").New(name, source, viz)}
}
