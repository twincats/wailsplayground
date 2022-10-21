import { createRouter, createWebHashHistory } from 'vue-router'

export default createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: '/home',
      name: 'home',
      component: () => import('../views/Home.vue'),
    },
    {
      path: '/',
      name: 'mangahome',
      component: () => import('../views/MangaHome.vue'),
    },
    {
      path: '/chapter/:mid',
      name: 'chapter',
      component: () => import('../views/Chapter.vue'),
    },
    {
      path: '/page/:cid',
      name: 'page',
      component: () => import('../views/Page.vue'),
    },
    {
      path: '/convert',
      name: 'convert',
      component: () => import('../views/Convert.vue'),
    },
    {
      path: '/about',
      name: 'about',
      component: () => import('../views/About.vue'),
    },
    {
      path: '/test',
      name: 'test',
      component: () => import('../views/Test.vue'),
    },
  ],
})
