import request from '@/utils/request'

export function fetchTagCount() {
  return request({
    url: '/tag/count',
    method: 'get'
  })
}
export function fetchTagList(query) {
  return request({
    url: '/tag/list',
    method: 'get',
    params: query
  })
}
export function fetchAllTag() {
  return request({
    url: '/tag/all',
    method: 'get'
  })
}

export function updateTag(data) {
  return request({
    url: '/tag/edit', // 请求地址
    method: 'post', // 请求类型
    data: data // 请求数据
  })
}
// 删除标签
export function deleteTag(id) {
  return request({
    url: `/tag/${id}`,
    method: 'delete'
  })
}
export function addTag(data) {
  return request({
    url: '/tag/add', // 请求地址
    method: 'post', // 请求类型
    data: data // 请求数据
  })
}
// 批量删除分类
export function multiDelTags(ids) {
  return request({
    url: `/tag/multiDelete/${ids}`,
    method: 'delete'
  })
}

