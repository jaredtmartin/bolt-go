package bolt

import "github.com/jaredtmartin/bolt-go"

func Group(children ...bolt.Element) *SvgGroup {
	root := bolt.NewDefaultElement("g")
	root.Children(children...)
	return &SvgGroup{
		DefaultElement: root,
	}
}

type SvgGroup struct {
	bolt.DefaultElement
}
