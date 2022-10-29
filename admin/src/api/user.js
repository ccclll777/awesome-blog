import request from '@/utils/request'

export function getInfo(data) {
  return request({
    url: '/user/info', // 请求地址
    method: 'post', // 请求类型
    data: data // 请求数据
  })
}
