import request from '@/utils/request'

export function fetchCategoryCount() {
  return request({
    url: '/category/count',
    method: 'get'
  })
}
export function fetchCategoryList(query) {
  return request({
    url: '/category/list',
    method: 'get',
    params: query
  })
}

export function updateCategory(data) {
  return request({
    url: '/category/edit', // 请求地址
    method: 'post', // 请求类型
    data: data // 请求数据
  })
}
// 删除分类
export function deleteCategory(id) {
  return request({
    url: `/category/${id}`,
    method: 'delete'
  })
}
export function addCategory(data) {
  return request({
    url: '/category/add', // 请求地址
    method: 'post', // 请求类型
    data: data // 请求数据
  })
}
// 批量删除分类
export function multiDelCategories(ids) {
  return request({
    url: `/category/multiDelete/${ids}`,
    method: 'delete'
  })
}

