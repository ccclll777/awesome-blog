package service

import (
	"awesome-blog/models"
	"sort"
	"sync"
)

type TagService struct {
}

// var indexPostService IndexService
//
// type Categories []models.CategoryJson
type TagPosts []models.TagJson

func (c *TagService) GetGroupByTag(postQuantity int) TagPosts {
	tags, _ := tagData.GetAllTag()
	//fmt.Println("categories", categories)
	var tagPosts TagPosts
	wg := sync.WaitGroup{}
	for i := 0; i < len(tags); i++ {
		tag := tags[i]
		//fmt.Println("category", category)
		wg.Add(1)
		go func() {
			posts, _ := postData.GetPostListByTag(tag.Id, postQuantity)
			var postJsonList []models.PostJson
			for _, post := range posts {
				postJsonList = append(postJsonList, indexPostService.GetPostJson(post))
			}
			tagJson := models.TagJson{tag.Name, len(posts), postJsonList}
			tagPosts = append(tagPosts, tagJson)
			//fmt.Println("categoryJson", categoryJson)
			//categoryList = append(categoryList, )
			wg.Done()
		}()
	}
	wg.Wait()
	sort.Sort(tagPosts)
	return tagPosts
}
func (c TagPosts) Len() int { return len(c) }

func (c TagPosts) Less(i, j int) bool { return c[i].Quantity > c[j].Quantity }

func (c TagPosts) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
