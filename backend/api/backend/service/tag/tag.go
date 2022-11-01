package tag

import (
	"awesome-blog/models"
	"strconv"
	"strings"
)

type TagService struct {
	tagData     models.Tag
	postData    models.Post
	postTagData models.PostTag
}

func (e *TagService) GetTagCount() (int, error) {
	total, err := e.tagData.GetTagCount()
	return total, err
}
func (e *TagService) GetTagList(page, pageSize int) ([]models.Tag, error) {
	tagList, err := e.tagData.GetTagList(page, pageSize)
	return tagList, err
}

func (e *TagService) GetAllTag() ([]models.Tag, error) {
	tagList, err := e.tagData.GetAllTag()
	return tagList, err
}

// 删除标签
func (e *TagService) DeleteTagById(id int) error {
	count, _ := e.postTagData.GetPostCountByTagId(id)
	if count == 0 {
		err := e.tagData.DeleteTagById(id)
		return err
	} else {
		//需要先在post tag中删除 对应的post tag
		_ = e.postTagData.DeleteByTagId(id)
		err := e.tagData.DeleteTagById(id) //然后在删除tag
		return err
	}
}
func (e *TagService) GetTagById(id int) (models.Tag, error) {
	tag, err := e.tagData.GetTagById(id)
	return tag, err
}
func (e *TagService) UpdateTag(tag models.Tag) (models.Tag, error) {
	tag, err := e.tagData.UpdateTag(tag)
	return tag, err
}
func (e *TagService) AddTag(name string, description string) (models.Tag, error) {
	tag := models.Tag{Name: name, Description: description}
	tag, err := e.tagData.AddTag(tag)
	return tag, err
}

/*
批量删除标签
*/
func (e *TagService) MultiDeleteByIds(ids string) error {
	idList := strings.Split(ids, ",") // 根据 , 分割成字符串数组
	for i := 0; i < len(idList); i++ {
		id, _ := strconv.Atoi(idList[i])
		_ = e.DeleteTagById(id)
	}
	return nil
}
