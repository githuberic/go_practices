package model

type Article struct {
	ArticleId int64  `gorm:"column:articleId",json:"articleId"`
	Subject   string `gorm:"column:subject",json:"title"`
	Url       string `gorm:"column:url",json:"url"`
	ImgUrl    string `json:"imgurl"`
	HeadUrl   string `json:"headurl"`
}

func (Article) TableName() string {
	return "article"
}
