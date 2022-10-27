import router from './router'
import store from './store'
import { Message } from 'element-ui'
import NProgress from 'nprogress' // progress bar
import 'nprogress/nprogress.css' // progress bar style
import { getToken } from '@/utils/auth' // get token from cookie
import getPageTitle from '@/utils/get-page-title'

NProgress.configure({ showSpinner: false }) // NProgress Configuration
// 白名单
const whiteList = ['/login', '/auth-redirect'] // no redirect whitelist
// 只有管理员才有权限访问的页面
// const adminList = /^\/admin/
router.beforeEach(async(to, from, next) => {
  NProgress.start()
  document.title = getPageTitle(to.meta.title)

  // 获取用户token
  const hasToken = getToken()
  // 有token的处理情况
  if (hasToken) {
    if (to.path === '/login') {
      // 不能重复登录
      next({ path: '/' })
      NProgress.done()
    } else {
      const userRole = store.getters.role
      if (userRole === 1) {
        next()
      } else {
        try {
          await store.dispatch('user/getInfo')
          // 判断用户权限
          // const userRole = store.getters.role
          // 管理员可以访问所有页面
          next()
        } catch (error) {
          // remove token and go to login page to re-login
          await store.dispatch('user/resetToken')
          Message.error(error || 'Has Error')
          next(`/login?redirect=${to.path}`)
          NProgress.done()
        }
      }
    }
  } else {
    // 没有token 跳转登录或者白名单

    if (whiteList.indexOf(to.path) !== -1) {
      next()
    } else {
      next(`/login?redirect=${to.path}`)
      NProgress.done()
    }
  }
})

router.afterEach(() => {
  NProgress.done()
})
