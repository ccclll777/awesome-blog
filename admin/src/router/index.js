import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/layout'
import PostCategory from '@/views/category/index'
import Dashboard from '@/views/dashboard/index'
import Tag from '@/views/tag/index'

import AddPost from '@/views/addPost/index'
import Post from '@/views/post/index'
import Login from '@/views/login/index'
import Error404 from '@/views/errorPage/404'
export const constantRoutes = [
  {
    path: '/login',
    component: Login,
    hidden: true
  },

  {
    path: '/404',
    component: Error404,
    hidden: true
  },
  {
    path: '/',
    component: Layout,
    redirect: '/dashboard',
    children: [{
      path: 'dashboard',
      name: 'dashboard',
      component: Dashboard,
      meta: { title: '首页', icon: 'dashboard' }
    }]
  },
  {
    path: '/post',
    component: Layout,
    redirect: '/dashboard',
    name: 'post',
    meta: { title: '博文管理', icon: 'documentation' },
    children: [
      {
        path: 'category',
        component: PostCategory,
        name: 'category',
        meta: { title: '文章分类', icon: 'list' }
      },
      {
        path: 'tag',
        component: Tag,
        name: 'tag',
        meta: { title: '文章标签', icon: 'component' }
      },
      {
        path: 'postList',
        component: Post,
        name: 'postList',
        meta: { title: '文章列表', icon: 'el-icon-document' }
      },
      {
        path: 'addPost',
        component: AddPost,
        name: 'addPost',
        meta: { title: '写文章', icon: 'el-icon-edit' }
      }
    ]
  },

  // 404 page must be placed at the end !!!
  { path: '*', redirect: '/404', hidden: true }
]

const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})

const router = createRouter()

export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
