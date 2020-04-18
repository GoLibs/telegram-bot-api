package structs

type StickerSet struct {
	Name          string    `json:"name"`
	Title         string    `json:"title"`
	IsAnimated    bool      `json:"is_animated"`
	ContainsMasks bool      `json:"contains_masks"`
	Stickers      []Sticker `json:"stickers"`
}
