import { createApp } from "vue"
import { createRouter, createWebHistory } from 'vue-router'
import ElementPlus from "element-plus"

const routes = [
  { path: '/', name: 'home', component: () => import('./views/home.js') },
  { path: '/about', name: 'about', component: () => import('./views/about.js') },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

const app = createApp({
  template: `<el-container>
    <el-header>
      <el-menu mode="horizontal" :default-active="$route.name" :router="true">
        <el-menu-item v-for="route in topRoutes" :key="route.name" :index="route.name" :route="{name: route.name}">{{ route.title }}</el-menu-item>
      </el-menu>
    </el-header>
    <el-main>
      <router-view></router-view>
    </el-main>
  </el-container>`,
  computed: {
    topRoutes() {
      // 顶级的 route 显示在顶栏上
      return [
        {name: 'home', title: 'Home'},
        {name: 'about', title: 'About'},
      ]
    }
  }
})
app.use(router)
app.use(ElementPlus)
app.mount('#root')
