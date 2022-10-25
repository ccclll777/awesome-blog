package service

import (
	"awesome-blog/models"
	"sort"
	"sync"
)

type CategoryService struct {
}

var indexPostService IndexService

type CategoryPosts []models.CategoryJson

func (c *CategoryService) GetGroupByCategory(postQuantity int) CategoryPosts {
	categories, _ := categoryData.GetAllCategory()
	//fmt.Println("categories", categories)

	var categoryList CategoryPosts
	wg := sync.WaitGroup{}
	for i := 0; i < len(categories); i++ {
		category := categories[i]
		//fmt.Println("category", category)
		wg.Add(1)
		go func() {
			posts, _ := postData.GetPostByCategoryId(category.Id, postQuantity)
			var postJsonList []models.PostJson
			for _, post := range posts {
				postJsonList = append(postJsonList, indexPostService.GetPostJson(&post))
			}
			categoryJson := models.CategoryJson{category.Name, len(posts), postJsonList}
			categoryList = append(categoryList, categoryJson)
			//fmt.Println("categoryJson", categoryJson)
			//categoryList = append(categoryList, )
			wg.Done()
		}()
	}
	wg.Wait()
	//fmt.Println(postJsons)
	sort.Sort(categoryList)
	return categoryList
}
func (c CategoryPosts) Len() int { return len(c) }

func (c CategoryPosts) Less(i, j int) bool { return c[i].Quantity > c[j].Quantity }

func (c CategoryPosts) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
