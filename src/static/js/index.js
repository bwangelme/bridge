import { createApp } from "vue"
import { createRouter, createWebHistory } from 'vue-router'
import ElementPlus from "element-plus"

const routes = [
  { path: '/', name: 'app', component: () => import('./views/app.js') },
  { path: '/about', name: 'about', component: () => import('./views/about.js') },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

const app = createApp({
  template: `<el-container>
    <el-header>
      <el-menu :default-active="$route.name" :router="true" mode="horizontal">
        <el-menu-item v-for="route in topRoutes" :key="route.name" :index="route.name" :route="{name: route.name}">{{ route.title }}</el-menu-item>
      </el-menu>
    </el-header>
    <el-container>
      <el-main>
        <router-view></router-view>
      </el-main>
    </el-container>
  </el-container>`,
  computed: {
    topRoutes() {
      // 顶级的 route 显示在顶栏上
      return [
        {name: 'app', title: 'Apps'},
        {name: 'about', title: 'About'},
      ]
    }
  }
})
app.use(router)
app.use(ElementPlus)
app.mount('#root')
