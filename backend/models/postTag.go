package models

import "awesome-blog/common/global"

type PostTag struct {
	Id     int `gorm:"id" `
	PostId int `gorm:"post_id" `
	TagId  int `gorm:"tag_id" `
}

// type Categories []Category
func (PostTag) TableName() string {
	return "post_tag"
}
func (e *PostTag) DeleteByTagId(tag_id int) error {
	var postTag []PostTag
	if err := global.Db.Table(e.TableName()).Where("tag_id = ?", tag_id).Delete(&postTag).Error; err != nil {
		return err
	}
	return nil
}
func (e *PostTag) GetPostCountByTagId(tag_id int) (int, error) {
	var total int64
	if err := global.Db.Table(e.TableName()).Where("tag_id =?", tag_id).Count(&total).Error; err != nil {
		return 0, err
	}
	return int(total), nil
}

func (e *PostTag) AddPostTag(postTag PostTag) (PostTag, error) {
	if err := global.Db.Table(e.TableName()).Create(&postTag).Error; err != nil {
		return postTag, err
	}
	return postTag, nil
}
func (e *PostTag) Count(post_id, tag_id int) (int, error) {
	var total int64
	if err := global.Db.Table(e.TableName()).Where("post_id = ? and tag_id =? ", post_id, tag_id).Count(&total).Error; err != nil {
		return 0, err
	}
	return int(total), nil
}
