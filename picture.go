package bolt

import "log"

type PictureSource struct {
	Media string
	Src   string
}

func Picture(alt string, sources ...PictureSource) Element {
	pic := NewElement("picture")
	for i, v := range sources {
		log.Println(`i: `, i)
		log.Println(`len(sources): `, len(sources))
		if i < len(sources)-1 {
			pic.Add(NewElement("source").Attr("srcset", v.Src).Attr("media", v.Media))
		} else {
			pic.Add(NewElement("img").Attr("src", v.Src).Attr("alt", alt))
		}
	}
	return pic
}
