import request from '@/utils/request'

export function fetchPostCount() {
  return request({
    url: '/post/count',
    method: 'get'
  })
}
export function fetchPostList(query) {
  return request({
    url: '/post/list',
    method: 'get',
    params: query
  })
}
export function fetchPostById(query) {
  return request({
    url: '/post/fetchPostById',
    method: 'get',
    params: query
  })
}

export function addPost(data) {
  return request({
    url: '/post/add',
    method: 'post',
    data
  })
}
export function updatePost(data) {
  return request({
    url: '/post/update',
    method: 'post',
    data
  })
}
export function publishPost(data) {
  return request({
    url: '/post/publish',
    method: 'post',
    data
  })
}
export function stopPublishPost(data) {
  return request({
    url: '/post/stopPublish',
    method: 'post',
    data
  })
}
export function deletePost(data) {
  return request({
    url: '/post/delete',
    method: 'post',
    data
  })
}
export function undeletePost(data) {
  return request({
    url: '/post/undelete',
    method: 'post',
    data
  })
}

