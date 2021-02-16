package models

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	Model

	TagID         int    `json:"tag_id" gorm:"index"` //	外键
	Tag           Tag    `json:"tag"`                 //	Related进行关联
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	CoverImageUrl string `json:"cover_image_url"`
	Content       string `json:"content"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         int    `json:"state"`
}

/*
func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}
*/
func ExistArticleByID(id int) (bool, error) {
	var article Article

	err := db.Select("id").Where("id=? AND deleted_on = ?", id, 0).First(&article).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if article.ID > 0 {
		return true, nil
	}

	return false, nil
}

func GetARticleTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Article{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func GetArticles(pageNum, pageSize int, maps interface{}) ([]*Article, error) {
	/*
		Preload就是一个预加载器，它会执行两条 SQL，分别是SELECT * FROM blog_articles;和SELECT * FROM blog_tag WHERE id IN (1,2,3,4);，那么在查询出结构后，gorm内部处理对应的映射逻辑，将其填充到Article的Tag中，会特别方便，并且避免了循环查询
	*/
	var article []*Article
	if err := db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&article).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return article, nil
}

func GetArticle(id int) (*Article, error) {
	var article Article
	if err := db.Where("id = ? ", id).First(&article).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	//	关联查询！！
	if err := db.Model(&article).Related(&article.Tag).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &article, nil
}

func EditArticle(id int, data interface{}) error {
	if err := db.Model(&Article{}).Where("id=? AND deleted_on = ?", id, 0).Update(data).Error; err != nil {
		return err
	}

	return nil
}

func AddArticle(data map[string]interface{}) error {
	if err := db.Create(&Article{
		//	v表示一个接口值，I表示接口类型。这个实际就是 Golang 中的类型断言，用于判断一个接口值的实际类型是否为某个类型，或一个非接口值的类型是否实现了某个接口类型
		TagID:         data["tag_id"].(int),
		Title:         data["title"].(string),
		Desc:          data["desc"].(string),
		Content:       data["content"].(string),
		CreatedBy:     data["created_by"].(string),
		State:         data["state"].(int),
		CoverImageUrl: data["cover_image_url"].(string),
	}).Error; err != nil {
		return err
	}

	return nil
}

func DeleteArticle(id int) error {
	if err := db.Where("id = ? ", id).Delete(Article{}).Error; err != nil {
		return err
	}

	return nil
}
