package dao

import (
	"fmt"
	"go_practices/composited_exec/gin_gorm_viper/global"
	"go_practices/composited_exec/gin_gorm_viper/model"
)

//select一条记录
func SelectOneArticle(articleId uint64) (*model.Article, error) {
	fields := []string{"articleId", "subject", "url"}
	articleOne := &model.Article{}
	err := global.DBLink.Select(fields).Where("articleId=?", articleId).First(&articleOne).Error
	if err != nil {
		return nil, err
	} else {
		return articleOne, nil
	}
}

//select总数
func SelectcountAll() (int, error) {
	var count int
	err := global.DBLink.Table(model.Article{}.TableName()).Where("isPublish=?", 1).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

//select全部文章
func SelectAllArticle(pageOffset int, pageSize int) ([]*model.Article, error) {
	fields := []string{"articleId", "subject", "url"}
	rows, err := global.DBLink.Select(fields).Table(model.Article{}.TableName()).Where("isPublish=?", 1).Offset(pageOffset).Limit(pageSize).Rows()

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var articles []*model.Article
	for rows.Next() {
		fmt.Println("rows.next:")
		r := &model.Article{}
		if err := rows.Scan(&r.ArticleId, &r.Subject, &r.Url); err != nil {
			fmt.Println("rows.next:")
			fmt.Println(err)
			return nil, err
		}
		articles = append(articles, r)
	}
	return articles, nil
}
