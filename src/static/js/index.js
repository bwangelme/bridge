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
    <el-header style="justify-content: center; display: flex">
        <h1>Bridge For QAE</h1>
    </el-header>
    <el-container>
      <el-aside width="200px">
        <el-menu mode="vertical" :default-active="$route.name" :router="true">
          <el-menu-item v-for="route in topRoutes" :key="route.name" :index="route.name" :route="{name: route.name}">{{ route.title }}</el-menu-item>
        </el-menu>
      </el-aside>
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
