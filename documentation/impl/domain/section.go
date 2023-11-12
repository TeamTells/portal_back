package domain

type Section struct {
	Id           int
	Title        string
	ThumbnailUrl string
	IsFavorite   bool
}

const NO_ID = -1
