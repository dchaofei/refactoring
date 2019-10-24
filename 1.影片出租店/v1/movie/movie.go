package movie

const (
	// 普通片
	Regular = iota
	// 新片
	NewRelease
	// 儿童片
	ChildRens
)

type Movie struct {
	title     string
	priceCode int // 影片类型
}

func NewMovie(title string, priceCode int) *Movie {
	return &Movie{
		title:     title,
		priceCode: priceCode,
	}
}
