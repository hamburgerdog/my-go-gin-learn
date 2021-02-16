package models

//	use orm to get data
type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

/*
func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}
*/

func GetTags(pageNum int, pageSize int, maps interface{}) ([]Tag, error) {
	//	db from models/models.go [var db *gorm.DB]
	var tags []Tag
	if err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func GetTagTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Tag{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func AddTag(name string, state int, createdBy string) error {
	if err := db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTag(id int) error {
	if err := db.Where("id=?", id).Delete(&Tag{}).Error; err != nil {
		return err
	}
	return nil
}

func EditTag(id int, data interface{}) error {
	if err := db.Model(&Tag{}).Where("id = ? AND deleted_on = ?", id, 0).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func ExistTagByName(name string) (bool, error) {
	var tag Tag
	if err := db.Select("id").Where("name = ? AND deleted_on = ?", name, 0).First(&tag).Error; err != nil {
		return false, err
	}
	if tag.ID > 0 {
		return true, nil
	}
	return false, nil
}

func ExistTagById(id int) (bool, error) {
	var tag Tag
	if err := db.Select("id").Where("id = ? AND deleted_on = ?", id, 0).First(&tag).Error; err != nil {
		return false, err
	}
	if tag.ID > 0 {
		return true, nil
	}
	return false, nil
}
