package bolt

func Group(children ...Element) *SvgGroup {
	root := NewDefaultElement("g")
	root.Children(children...)
	return &SvgGroup{
		DefaultElement: root,
	}
}

type SvgGroup struct {
	DefaultElement
}
