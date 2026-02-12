package models

const ShortUrlPrefix = "http://wu.x/"

type ConvertedUrl struct {
	Long_url string `gorm:"primaryKey"`
	Short_url string
}

func NewConvertedUrl (long_url, short_url string) *ConvertedUrl {
	return &ConvertedUrl{
		Long_url: long_url,
		Short_url: short_url,
	}
}