import request from '@/utils/request'

// 登录
export function login(data) {
  return request({
    url: '/auth/login', // 请求地址
    method: 'post', // 请求类型
    data: data // 请求数据
  })
}
export function getInfo(data) {
  return request({
    url: '/user/info', // 请求地址
    method: 'post', // 请求类型
    data: data // 请求数据
  })
}
