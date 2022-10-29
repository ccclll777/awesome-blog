import request from '@/utils/request'

// 登录
export function login(data) {
  return request({
    url: '/auth/login', // 请求地址
    method: 'post', // 请求类型
    data: data // 请求数据
  })
}
// 退出登录
export function logout() {
  return request({
    url: '/auth/logout',
    method: 'post'
  })
}
