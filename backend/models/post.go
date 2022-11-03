package models

import (
	"awesome-blog/common/global"
	"strconv"
	"strings"
	"time"
)

// 数据库，博客数据映射模型
type Post struct {
	ID          int    `gorm:"primaryKey"` // ID, store自行控制
	Title       string `gorm:"title" `
	Description string `gorm:"description"` //摘要/描述
	Keywords    string `gorm:"keywords"`    //关键词
	Count       int    `gorm:"count"`       // 评论数量
	Content     string `gorm:"content" `    //HTML内容
	MDContent   string `gorm:"md_content" ` //markdown 内容
	AuthorId    int    `gorm:"author_id"`   //作者
	//IsDraft     bool      `gorm:"is_draft" `   // 是否是草稿
	DeletedAt   time.Time `gorm:"deleted_at"` // 删除时间
	UpdatedAt   time.Time `gorm:"updated_at"` // 更新时间
	CreatedAt   time.Time `gorm:"created_at"` // 创建时间
	IsDelete    int       `gorm:"is_delete"`  // 是否删除
	MusicId     string    `gorm:"music_id"`
	Url         string    `gorm:"url"`
	IsPublished int       `gorm:"is_published"` // 是否已经发布 还是只是草稿
	CategoryId  int       `gorm:"category_id"`  //所属分类
	Category    string
	//TagList     []Tag
}
type PostJson struct {
	Id          int
	Title       string    `json:"title"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Tags        []string  `json:"tags"`
	Author      string    `json:"author"`
	MusicId     string    `json:"musicId"`
	Category    string
}
type PostDatail struct {
	Post
	Tags     []string
	Category string
	Body     string
}

var postTagData PostTag

func (Post) TableName() string {
	return "post"
}
func (e *Post) GetPostById(post_id int) (Post, error) {
	var post Post
	if err := global.Db.Table(e.TableName()).Where("id =?", post_id).Find(&post).Error; err != nil {
		return post, err
	}
	return post, nil
}

func (e *Post) GetPostList(page, pageSize int) ([]Post, error) {
	offset := (page - 1) * pageSize
	queryBuider := global.Db.Limit(pageSize).Offset(offset).Order("created_at asc").Table(e.TableName())
	var postList []Post
	if err := queryBuider.Find(&postList).Error; err != nil {
		return nil, err
	}
	return postList, nil
}
func (e *Post) GetPostCount() (int, error) {
	var total int64
	if err := global.Db.Table(e.TableName()).Count(&total).Error; err != nil {
		return 0, err
	}
	return int(total), nil
}
func (e *Post) GetPostCountByTag(tag_id int) (int, error) {
	var total int64
	if err := global.Db.Table("post_tag").Where("tag_id =?", tag_id).Count(&total).Error; err != nil {
		return 0, err
	}
	return int(total), nil
}
func (e *Post) GetPostCountByCategory(category_id int) (int, error) {
	var total int64
	if err := global.Db.Table(e.TableName()).Where("category_id =?", category_id).Count(&total).Error; err != nil {
		return 0, err
	}
	return int(total), nil
}
func (e *Post) GetPostsByCategoryId(category_id int, page, postQuantity int) ([]Post, error) {
	offset := (page - 1) * postQuantity
	var postList []Post
	if err := global.Db.Table(e.TableName()).Where("category_id =?", category_id).Limit(postQuantity).Offset(offset).Order("created_at asc").Find(&postList).Error; err != nil {
		return nil, err
	}
	return postList, nil
}
func (e *Post) GetPostsByTagId(tag_id int, page, postQuantity int) ([]Post, error) {
	offset := (page - 1) * postQuantity
	var posts []Post
	if err := global.Db.Table(e.TableName()).
		Where("id in ( select post_id as id from post_tag where tag_id = ?)", tag_id).Limit(postQuantity).Offset(offset).Order("created_at asc").Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (e *Post) GetPostCountByCategoryId(category_id int) (int, error) {
	var total int64
	if err := global.Db.Table(e.TableName()).Where("category_id =?", category_id).Count(&total).Error; err != nil {
		return 0, err
	}
	return int(total), nil
}
func (e *Post) AddPost(post Post, tagIds string) error {
	// 开始事务  保证添加文章的操作成功
	tx := global.Db.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	// 添加文章
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()
	err := tx.Exec("INSERT INTO `post` ( `created_at`,`updated_at`,`title`, `description`, `keywords`, `content`,"+
		" `md_content`, `author_id`, `category_id`, `is_published`)"+
		" VALUES( ?,?,?, ?, ?, ?, ?, ?, ?, ?)", post.CreatedAt, post.UpdatedAt,
		post.Title, post.Description, post.Keywords, post.Content,
		post.MDContent, post.AuthorId, post.CategoryId, post.IsPublished).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 根据标题获取文章
	err = tx.Table(e.TableName()).Where("title = ?", post.Title).First(&post).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 创建文章和标签关联，更新标签对应文章数量
	if tagIds != "" {
		tagIdList := strings.Split(tagIds, ",") // 根据 , 分割成字符串数组
		for _, tagId := range tagIdList {
			id, _ := strconv.Atoi(tagId)
			err = tx.Exec("insert into `post_tag` (`post_id`,`tag_id`) values (?,?)",
				post.ID, id).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	return tx.Commit().Error
}
func (e *Post) GetPostCountByTitle(title string) (int, error) {
	var total int64
	if err := global.Db.Table(e.TableName()).Where("title = ?", title).Count(&total).Error; err != nil {
		return 0, err
	}
	return int(total), nil
}

func (e *Post) PublishPost(postId int) error {
	if err := global.Db.Table(e.TableName()).Where("id = ?", postId).Update("is_published", 1).Error; err != nil {
		return err
	}
	return nil
}
func (e *Post) StopPublishPost(postId int) error {
	if err := global.Db.Table(e.TableName()).Where("id = ?", postId).Update("is_published", 0).Error; err != nil {
		return err
	}
	return nil
}

func (e *Post) DeletePost(postId int) error {
	if err := global.Db.Table(e.TableName()).Where("id = ?", postId).Update("is_delete", 1).Error; err != nil {
		return err
	}
	return nil
}

func (e *Post) UnDeletePost(postId int) error {
	if err := global.Db.Table(e.TableName()).Where("id = ?", postId).Update("is_delete", 0).Error; err != nil {
		return err
	}
	return nil
}

func (e *Post) UpdatePost(post Post, tagIds string) error {
	//post.UpdatedAt = time.Now()
	//if err := global.Db.Table(e.TableName()).Where("id = ?", post.ID).Updates(&post).Error; err != nil {
	//	return err
	//}
	//return nil
	// 开始事务  保证添加文章的操作成功
	tx := global.Db.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	post.UpdatedAt = time.Now()
	err := tx.Exec("INSERT INTO `post` ( `created_at`,`updated_at`,`title`, `description`, `keywords`, `content`,"+
		" `md_content`, `author_id`, `category_id`, `is_published`)"+
		" VALUES( ?,?,?, ?, ?, ?, ?, ?, ?, ?)", post.CreatedAt, post.UpdatedAt,
		post.Title, post.Description, post.Keywords, post.Content,
		post.MDContent, post.AuthorId, post.CategoryId, post.IsPublished).Error
	//title, description, keywords, content, mdContent, TagIds string, post_id, author_id, is_published, category_id int
	err = tx.Model(&Post{}).Where("`id` = ?", post.ID).
		Updates(map[string]interface{}{
			"title":        post.Title,
			"description":  post.Description,
			"keywords":     post.Keywords,
			"content":      post.Content,
			"md_content":   post.MDContent,
			"author_id":    post.AuthorId,
			"category_id":  post.CategoryId,
			"is_published": post.IsPublished,
			"updated_at":   post.UpdatedAt}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	// 删除标签文章表中关联记录
	err = tx.Exec("delete from `post_tag` where `post_id` = ?", post.ID).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	// 创建文章和标签关联，更新标签对应文章数量
	if tagIds != "" {
		tagIdList := strings.Split(tagIds, ",") // 根据 , 分割成字符串数组
		for _, tagId := range tagIdList {
			id, _ := strconv.Atoi(tagId)
			//fmt.Println("创建文章和标签关联，更新标签对应文章数量", id)
			err = tx.Exec("insert into `post_tag` (`post_id`,`tag_id`) values (?,?)",
				post.ID, id).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	return tx.Commit().Error
}
