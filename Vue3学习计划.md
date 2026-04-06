# Vue.js 四周学习计划

> 目标明确、实践驱动、避免弯路。**边学边做，项目为王。**

---

## 📋 学习路线总览

| 阶段 | 时间 | 核心目标 |
|:---:|:---:|:---|
| 第一阶段 | 第 1-2 天 | 前置准备，环境就绪 |
| 第二阶段 | 第 3-10 天 | Vue 核心精要 + Todo List 项目 |
| 第三阶段 | 第 11-21 天 | 生态与项目驱动 |
| 第四阶段 | 第 22-28 天 | 深化与工具链 |

---

## 第一阶段：前置准备（第 1-2 天）

> **目标：** 确保基础无碍，环境就绪。

### 1.1 确认基础

如果以下内容你已经熟悉，可以直接跳过，进入环境搭建：

- ✅ HTML 基础
- ✅ CSS 基础
- ✅ JavaScript 基础
- ✅ ES6+ 语法（箭头函数、解构、模块）

**如果不熟悉？** 建议花 1-2 天快速过一遍：
- [MDN JavaScript 入门](https://developer.mozilla.org/zh-CN/docs/Learn/JavaScript/First_steps)
- [ES6 入门教程](https://es6.ruanyifeng.com/)

### 1.2 环境搭建

#### 1.2.1 安装 Node.js

访问 [Node.js 官网](https://nodejs.org/)，下载并安装最新版（LTS 版本）。

```bash
# 验证安装
node -v
npm -v
```

#### 1.2.2 安装 VS Code

下载安装 [Visual Studio Code](https://code.visualstudio.com/)，并安装以下插件：

| 插件名称 | 说明 |
|:---|:---|
| **Volar** | Vue 3 官方推荐的 VS Code 插件，提供语法高亮、类型提示、代码补全 |
| **Vue - Official** | Vue 官方维护的扩展包（新版） |
| **Auto Rename Tag** | 自动重命名配对的 HTML/XML 标签 |
| **Bracket Pair Colorizer** | 高亮显示匹配的括号 |
| **ESLint** | 代码检查 |
| **Prettier** | 代码格式化 |

#### 1.2.3 熟悉终端基本操作

```bash
# 常用命令
cd <目录>          # 进入目录
cd ..              # 返回上级目录
ls                 # 列出文件（Linux/Mac）
dir                # 列出文件（Windows）
mkdir <目录名>     # 创建目录
rm -rf <目录>      # 删除目录（慎用）
```

---

## 第二阶段：Vue 核心精要（第 3-10 天）

> **目标：** 掌握 Vue 3 的核心概念，并用最小项目巩固。

### 2.1 创建第一个应用

使用官方推荐的构建工具 `create-vue` 快速创建项目：

```bash
npm init vue@latest
```

运行命令后，会进入交互式引导：

```
Vue CLI v5.x
? Project name: ... my-vue-app
? Add TypeScript? ... No / Yes
? Add JSX Support? ... No / Yes
? Add Vue Router? ... No / Yes
? Add Pinia for state management? ... No / Yes
? Add Vitest for Unit Testing? ... No / Yes
? Add ESLint for code quality? ... No / Yes
```

或者使用 Vite 手动创建：

```bash
# 使用 Vite 创建 Vue 项目
npm create vite@latest my-vue-app -- --template vue
cd my-vue-app
npm install
npm run dev
```

### 2.2 核心概念学习与实践

#### 2.2.1 模板语法与响应式基础

**插值与指令：**

```vue
<script setup>
import { ref, reactive } from 'vue'

// ref 用于基本类型
const message = ref('Hello Vue!')
const count = ref(0)

// reactive 用于对象
const state = reactive({
  name: 'Vue',
  version: '3'
})

// 修改值
const updateMessage = () => {
  message.value = 'Hello World!'
  count.value++
}
</script>

<template>
  <h1>{{ message }}</h1>
  <p>计数: {{ count }}</p>
  
  <!-- 响应式绑定 -->
  <div :class="state.name">
    {{ state.name }} 版本: {{ state.version }}
  </div>
  
  <!-- 事件处理 -->
  <button @click="updateMessage">点击更新</button>
</template>
```

**常用指令：**

| 指令 | 用法 | 说明 |
|:---|:---|:---|
| `v-bind` | `:属性名="值"` | 绑定属性 |
| `v-on` | `@事件名="处理函数"` | 绑定事件 |
| `v-model` | `v-model="变量"` | 双向绑定 |
| `v-if` | `v-if="条件"` | 条件渲染 |
| `v-show` | `v-show="条件"` | 显示/隐藏 |
| `v-for` | `v-for="item in 数组"` | 列表渲染 |
| `v-once` | `v-once` | 只渲染一次 |
| `v-html` | `v-html="html字符串"` | 插入 HTML |

#### 2.2.2 计算属性与侦听器

**computed（计算属性）：**

```vue
<script setup>
import { ref, computed } from 'vue'

const firstName = ref('张')
const lastName = ref('三')

// 计算属性 - 基于响应式依赖缓存
const fullName = computed(() => {
  return lastName.value + firstName.value
})

// 可写的计算属性
const fullNameWritable = computed({
  get: () => `${lastName.value}${firstName.value}`,
  set: (val) => {
    lastName.value = val.slice(0, 1)
    firstName.value = val.slice(1)
  }
})
</script>

<template>
  <p>姓名: {{ fullName }}</p>
</template>
```

**watch 与 watchEffect（侦听器）：**

```vue
<script setup>
import { ref, watch, watchEffect } from 'vue'

const todo = ref('')
const todos = ref([])

// watch - 监听特定数据源
watch(todo, (newValue, oldValue) => {
  console.log(`从 "${oldValue}" 变为 "${newValue}"`)
  if (newValue) {
    todos.value.push({ text: newValue, done: false })
  }
})

// watchEffect - 立即执行，依赖追踪
watchEffect(() => {
  console.log('todos 数量:', todos.value.length)
})
</script>

<template>
  <input v-model="todo" placeholder="输入待办事项" />
  <ul>
    <li v-for="(item, index) in todos" :key="index">
      {{ item.text }}
    </li>
  </ul>
</template>
```

#### 2.2.3 组件基础

**创建子组件：**

```vue
<!-- src/components/TodoItem.vue -->
<script setup>
// 定义 props
defineProps({
  todo: {
    type: Object,
    required: true
  },
  index: {
    type: Number,
    required: true
  }
})

// 定义 emit
const emit = defineEmits(['toggle', 'remove'])
</script>

<template>
  <div class="todo-item" :class="{ done: todo.done }">
    <input 
      type="checkbox" 
      :checked="todo.done"
      @change="emit('toggle', index)"
    />
    <span>{{ todo.text }}</span>
    <button @click="emit('remove', index)">删除</button>
  </div>
</template>

<style scoped>
.done {
  text-decoration: line-through;
  color: #999;
}
</style>
```

**使用子组件：**

```vue
<!-- src/App.vue -->
<script setup>
import { ref } from 'vue'
import TodoItem from './components/TodoItem.vue'

const todos = ref([
  { text: '学习 Vue', done: false },
  { text: '完成 Todo List', done: false }
])

const toggleTodo = (index) => {
  todos.value[index].done = !todos.value[index].done
}

const removeTodo = (index) => {
  todos.value.splice(index, 1)
}
</script>

<template>
  <h1>我的待办</h1>
  <TodoItem 
    v-for="(todo, index) in todos"
    :key="index"
    :todo="todo"
    :index="index"
    @toggle="toggleTodo"
    @remove="removeTodo"
  />
</template>
```

#### 2.2.4 生命周期钩子

```vue
<script setup>
import { ref, onMounted, onUpdated, onUnmounted } from 'vue'

const count = ref(0)

// 组件挂载后
onMounted(() => {
  console.log('组件已挂载')
  // 适合：获取 DOM、发起初始请求、设置定时器
})

// 数据更新后
onUpdated(() => {
  console.log('组件已更新')
  // 适合：更新后的 DOM 操作
})

// 组件卸载前
onUnmounted(() => {
  console.log('组件已卸载')
  // 适合：清理定时器、取消订阅、解绑事件
})

// 其他常用钩子
import { onBeforeMount, onBeforeUpdate, onBeforeUnmount } from 'vue'

onBeforeMount(() => {
  console.log('即将挂载')
})

onBeforeUpdate(() => {
  console.log('即将更新')
})

onBeforeUnmount(() => {
  console.log('即将卸载')
})
</script>
```

#### 2.2.5 组合式 API（Composition API）

**`<script setup>` 语法：**

```vue
<script setup>
// 代码在这里自动作为组件的 setup 函数
import { ref, computed, onMounted } from 'vue'

// 响应式数据
const message = ref('Hello')

// 方法
const greet = () => {
  alert(message.value)
}

// 生命周期
onMounted(() => {
  console.log('Ready!')
})

// 暴露给模板的属性
defineExpose({
  message,
  greet
})
</script>

<template>
  <button @click="greet">{{ message }}</button>
</template>
```

### 2.3 第一个实践项目：Todo List

> **目标：** 实现一个增删改查、可标记完成的任务列表。

**项目要求：**

- [ ] 添加新任务
- [ ] 标记任务完成/未完成
- [ ] 删除任务
- [ ] 编辑任务内容
- [ ] 筛选任务（全部/已完成/未完成）
- [ ] 数据持久化（本地存储）

**项目结构：**

```
src/
├── App.vue
├── main.js
├── assets/
│   └── main.css
└── components/
    ├── TodoHeader.vue      # 输入框和添加按钮
    ├── TodoList.vue         # 任务列表
    ├── TodoItem.vue         # 单个任务项
    ├── TodoFooter.vue       # 筛选和统计
    └── TodoEditModal.vue    # 编辑弹窗
```

**核心实现：**

```vue
<!-- src/App.vue 完整示例 -->
<script setup>
import { ref, computed, watch } from 'vue'
import TodoHeader from './components/TodoHeader.vue'
import TodoList from './components/TodoList.vue'
import TodoFooter from './components/TodoFooter.vue'

// 任务列表
const todos = ref([])

// 筛选状态
const filter = ref('all')

// 过滤后的任务
const filteredTodos = computed(() => {
  switch (filter.value) {
    case 'active':
      return todos.value.filter(t => !t.done)
    case 'completed':
      return todos.value.filter(t => t.done)
    default:
      return todos.value
  }
})

// 添加任务
const addTodo = (text) => {
  todos.value.push({
    id: Date.now(),
    text,
    done: false
  })
}

// 切换完成状态
const toggleTodo = (id) => {
  const todo = todos.value.find(t => t.id === id)
  if (todo) todo.done = !todo.done
}

// 删除任务
const removeTodo = (id) => {
  todos.value = todos.value.filter(t => t.id !== id)
}

// 编辑任务
const editTodo = (id, newText) => {
  const todo = todos.value.find(t => t.id === id)
  if (todo) todo.text = newText
}

// 清除已完成
const clearCompleted = () => {
  todos.value = todos.value.filter(t => !t.done)
}

// 统计
const activeCount = computed(() => 
  todos.value.filter(t => !t.done).length
)

// 本地存储
watch(todos, (val) => {
  localStorage.setItem('todos', JSON.stringify(val))
}, { deep: true })

// 初始化
const saved = localStorage.getItem('todos')
if (saved) {
  todos.value = JSON.parse(saved)
}
</script>

<template>
  <div class="app">
    <h1>Vue Todo List</h1>
    <TodoHeader @add="addTodo" />
    <TodoList 
      :todos="filteredTodos"
      @toggle="toggleTodo"
      @remove="removeTodo"
      @edit="editTodo"
    />
    <TodoFooter 
      :filter="filter"
      :active-count="activeCount"
      :total-count="todos.length"
      @set-filter="filter = $event"
      @clear="clearCompleted"
    />
  </div>
</template>
```

---

## 第三阶段：生态与项目驱动（第 11-21 天）

> **目标：** 学习必备生态库，并通过真实项目整合知识。

### 3.1 学习 Vue Router

#### 3.1.1 安装与配置

```bash
npm install vue-router
```

```javascript
// src/router/index.js
import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: () => import('../views/Home.vue')
  },
  {
    path: '/todos',
    name: 'Todos',
    component: () => import('../views/Todos.vue')
  },
  {
    path: '/about',
    name: 'About',
    component: () => import('../views/About.vue')
  },
  // 动态路由
  {
    path: '/user/:id',
    name: 'User',
    component: () => import('../views/User.vue'),
    props: true
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
```

#### 3.1.2 在应用中使用

```vue
<!-- src/App.vue -->
<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <nav>
    <RouterLink to="/">首页</RouterLink>
    <RouterLink to="/todos">待办</RouterLink>
    <RouterLink to="/about">关于</RouterLink>
  </nav>
  
  <RouterView />
</template>
```

#### 3.1.3 编程式导航

```vue
<script setup>
import { useRouter } from 'vue-router'

const router = useRouter()

// 导航到指定路径
const goToUser = (id) => {
  router.push(`/user/${id}`)
}

// 导航到命名的路由
const goToHome = () => {
  router.push({ name: 'Home' })
}

// 带查询参数
router.push({ path: '/search', query: { q: 'vue' } })

// 带状态
router.push({ path: '/home', state: { from: 'about' } })

// 返回上一页
router.back()

// 前进/后退
router.go(-1)  // 后退一页
router.go(1)   // 前进一页
</script>
```

### 3.2 学习 Pinia（状态管理）

#### 3.2.1 安装与配置

```bash
npm install pinia
```

```javascript
// src/main.js
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'

const app = createApp(App)
app.use(createPinia())
app.mount('#app')
```

#### 3.2.2 定义 Store

```javascript
// src/stores/todoStore.js
import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export const useTodoStore = defineStore('todo', () => {
  // State - 响应式状态
  const todos = ref([])
  const filter = ref('all')

  // Getters - 计算属性
  const filteredTodos = computed(() => {
    switch (filter.value) {
      case 'active':
        return todos.value.filter(t => !t.done)
      case 'completed':
        return todos.value.filter(t => t.done)
      default:
        return todos.value
    }
  })

  const activeCount = computed(() => 
    todos.value.filter(t => !t.done).length
  )

  // Actions - 方法
  const addTodo = (text) => {
    todos.value.push({
      id: Date.now(),
      text,
      done: false
    })
  }

  const toggleTodo = (id) => {
    const todo = todos.value.find(t => t.id === id)
    if (todo) todo.done = !todo.done
  }

  const removeTodo = (id) => {
    todos.value = todos.value.filter(t => t.id !== id)
  }

  const setFilter = (newFilter) => {
    filter.value = newFilter
  }

  return {
    todos,
    filter,
    filteredTodos,
    activeCount,
    addTodo,
    toggleTodo,
    removeTodo,
    setFilter
  }
})
```

#### 3.2.3 在组件中使用

```vue
<script setup>
import { useTodoStore } from '../stores/todoStore'

const todoStore = useTodoStore()

// 直接使用 store 的状态和方法
const newTodo = ref('')

const handleAdd = () => {
  if (newTodo.value.trim()) {
    todoStore.addTodo(newTodo.value)
    newTodo.value = ''
  }
}
</script>

<template>
  <div>
    <input v-model="newTodo" @keyup.enter="handleAdd" />
    <button @click="handleAdd">添加</button>
    
    <ul>
      <li 
        v-for="todo in todoStore.filteredTodos"
        :key="todo.id"
        @click="todoStore.toggleTodo(todo.id)"
      >
        {{ todo.text }} - {{ todo.done ? '完成' : '待办' }}
      </li>
    </ul>
    
    <p>待办数量: {{ todoStore.activeCount }}</p>
  </div>
</template>
```

### 3.3 学习组合式函数（Composables）

> 这是 Vue 3 中逻辑复用的方式，替代了 Vue 2 的 Mixins。

#### 3.3.1 useMouse（获取鼠标位置）

```javascript
// src/composables/useMouse.js
import { ref, onMounted, onUnmounted } from 'vue'

export function useMouse() {
  const x = ref(0)
  const y = ref(0)

  const updatePosition = (e) => {
    x.value = e.clientX
    y.value = e.clientY
  }

  onMounted(() => {
    window.addEventListener('mousemove', updatePosition)
  })

  onUnmounted(() => {
    window.removeEventListener('mousemove', updatePosition)
  })

  return { x, y }
}
```

#### 3.3.2 useLocalStorage（本地存储）

```javascript
// src/composables/useLocalStorage.js
import { ref, watch } from 'vue'

export function useLocalStorage(key, defaultValue) {
  const data = ref(defaultValue)
  
  // 初始化
  const saved = localStorage.getItem(key)
  if (saved) {
    try {
      data.value = JSON.parse(saved)
    } catch {
      data.value = saved
    }
  }

  // 监听变化
  watch(data, (newValue) => {
    const toSave = typeof newValue === 'object' 
      ? JSON.stringify(newValue) 
      : newValue
    localStorage.setItem(key, toSave)
  }, { deep: true })

  return data
}
```

#### 3.3.3 使用组合式函数

```vue
<script setup>
import { useMouse } from '../composables/useMouse'
import { useLocalStorage } from '../composables/useLocalStorage'

const { x, y } = useMouse()
const username = useLocalStorage('username', '')
const theme = useLocalStorage('theme', 'light')
</script>

<template>
  <p>鼠标位置: {{ x }}, {{ y }}</p>
  <input v-model="username" placeholder="用户名" />
  <p>当前主题: {{ theme }}</p>
</template>
```

### 3.4 综合性项目：个人博客前端

> 选择一个你感兴趣的主题，模仿一个现有网站的部分功能。

**项目功能规划：**

| 页面 | 功能 | 技术要点 |
|:---|:---|:---|
| 首页/文章列表 | 文章卡片展示、分页、搜索 | Vue Router、组件拆分 |
| 文章详情 | Markdown 渲染、评论、目录 | 路由传参、computed |
| 分类/标签页 | 按分类筛选文章 | 动态路由、计算属性 |
| 关于页面 | 个人介绍、联系方式 | 静态页面 |
| 用户中心 | 登录注册、收藏文章 | Pinia、组合式函数 |

**推荐 UI 库：**

| UI 库 | 特点 | 适用场景 |
|:---|:---|:---|
| **Element Plus** | 功能丰富、文档完善 | PC 端后台管理 |
| **Vant** | 轻量、移动端优先 | 移动端 H5 |
| **Naive UI** | TypeScript 支持好、设计现代 | 通用 |
| **Ant Design Vue** | 企业级组件丰富 | 中后台系统 |
| **UnoCSS + Headless UI** | 高度自定义 | 喜欢 Tailwind 的用户 |

**项目初始化：**

```bash
npm init vue@latest my-blog
# 选择: Vue Router ✓, Pinia ✓, TypeScript ✓

cd my-blog
npm install
npm install element-plus
npm install @element-plus/icons-vue
npm install vue-router
npm install pinia
npm install marked highlight.js
npm install axios
npm run dev
```

---

## 第四阶段：深化与工具链（第 22-28 天）

> **目标：** 提升工程化能力和代码质量。

### 4.1 TypeScript 与 Vue

#### 4.1.1 Props 类型定义

```vue
<script setup lang="ts">
interface Todo {
  id: number
  text: string
  done: boolean
}

interface Props {
  todo: Todo
  index: number
  readonly?: boolean
}

const props = defineProps<Props>()
const emit = defineEmits<{
  toggle: [id: number]
  remove: [id: number]
  edit: [id: number, text: string]
}>()
</script>
```

#### 4.1.2 响应式数据的类型

```vue
<script setup lang="ts">
import { ref, reactive, computed } from 'vue'

// 基本类型
const count = ref<number>(0)
const name = ref<string>('Vue')
const loading = ref<boolean>(false)

// 对象类型
const user = ref<User | null>(null)
const userInfo = reactive<User>({
  id: 1,
  name: '张三',
  email: 'zhang@example.com'
})

// 数组类型
const todos = ref<Todo[]>([])
const numbers = ref<number[]>([1, 2, 3])

// 计算属性
const doubleCount = computed<number>(() => count.value * 2)

// 泛型 ref
const dict = ref<Record<string, any>>({})
</script>
```

### 4.2 Vite 配置

```javascript
// vite.config.js
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'

export default defineConfig({
  plugins: [vue()],
  
  resolve: {
    alias: {
      '@': path.resolve(__dirname, 'src')
    }
  },
  
  server: {
    port: 3000,
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, '')
      }
    }
  },
  
  build: {
    outDir: 'dist',
    sourcemap: false
  }
})
```

### 4.3 环境变量

```
# .env - 所有环境
VITE_APP_TITLE=我的博客

# .env.development - 开发环境
VITE_API_BASE=http://localhost:8080/api

# .env.production - 生产环境
VITE_API_BASE=https://api.example.com

# .env.local - 本地覆盖（不提交）
VITE_DEBUG=true
```

```javascript
// 使用环境变量
console.log(import.meta.env.VITE_APP_TITLE)
console.log(import.meta.env.VITE_API_BASE)
```

### 4.4 测试与调试

#### 4.4.1 Vue Devtools

安装 Chrome 或 Firefox 扩展：**Vue Devtools**

功能：
- 查看组件树
- 检查组件状态
- 追踪 Vuex/Pinia 状态变化
- 性能分析

#### 4.4.2 Vitest 单元测试

```bash
npm install -D vitest @vue/test-utils jsdom
```

```javascript
// vite.config.js
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

export default defineConfig({
  plugins: [vue()],
  test: {
    environment: 'jsdom',
    globals: true
  }
})
```

```vue
<!-- src/components/__tests__/TodoItem.spec.js -->
import { describe, it, expect } from 'vitest'
import { mount } from '@vue/test-utils'
import TodoItem from '../TodoItem.vue'

describe('TodoItem', () => {
  it('renders todo text', () => {
    const todo = { id: 1, text: '测试任务', done: false }
    const wrapper = mount(TodoItem, {
      props: { todo, index: 0 }
    })
    
    expect(wrapper.text()).toContain('测试任务')
  })
  
  it('emits toggle event when clicked', async () => {
    const wrapper = mount(TodoItem, {
      props: { todo: { id: 1, text: '测试', done: false }, index: 0 }
    })
    
    await wrapper.find('input[type="checkbox"]').trigger('change')
    
    expect(wrapper.emitted()).toHaveProperty('toggle')
  })
})
```

### 4.5 进阶主题

#### 4.5.1 自定义指令

```javascript
// src/directives/focus.js
export const vFocus = {
  mounted: (el) => {
    el.focus()
  }
}

// src/directives/click-outside.js
export const vClickOutside = {
  mounted(el, binding) {
    el._clickOutside = (event) => {
      if (!el.contains(event.target)) {
        binding.value(event)
      }
    }
    document.addEventListener('click', el._clickOutside)
  },
  unmounted(el) {
    document.removeEventListener('click', el._clickOutside)
  }
}
```

```vue
<script setup>
const vFocus = { mounted: (el) => el.focus() }
const vClickOutside = {
  mounted(el, binding) {
    el._handler = (e) => !el.contains(e.target) && binding.value(e)
    document.addEventListener('click', el._handler)
  },
  unmounted(el) {
    document.removeEventListener('click', el._handler)
  }
}
</script>

<template>
  <input v-focus placeholder="自动聚焦" />
  <div v-click-outside="closeDropdown">
    下拉菜单内容
  </div>
</template>
```

#### 4.5.2 Teleport（传送门）

```vue
<template>
  <button @click="showModal = true">打开弹窗</button>
  
  <Teleport to="body">
    <div v-if="showModal" class="modal">
      <p>这是模态框内容</p>
      <button @click="showModal = false">关闭</button>
    </div>
  </Teleport>
</template>
```

#### 4.5.3 Suspense（异步组件）

```vue
<script setup>
import { defineAsyncComponent, onErrorCaptured } from 'vue'

const AsyncUserProfile = defineAsyncComponent(() =>
  import('./components/UserProfile.vue')
)

const error = ref(null)
onErrorCaptured((err) => {
  error.value = err
  return false
})
</script>

<template>
  <Suspense>
    <template #default>
      <AsyncUserProfile />
    </template>
    <template #fallback>
      <p>加载中...</p>
    </template>
  </Suspense>
  
  <div v-if="error">
    加载失败: {{ error.message }}
  </div>
</template>
```

#### 4.5.4 性能优化

```vue
<script setup>
// 组件懒加载
const HeavyComponent = defineAsyncComponent(() =>
  import('./HeavyComponent.vue')
)

// 路由懒加载
// 已在 router 中配置: component: () => import('./views/About.vue')
</script>

<template>
  <!-- v-once: 只渲染一次 -->
  <div v-once>{{ staticContent }}</div>
  
  <!-- v-memo: 条件性跳过更新 -->
  <div v-memo="[selectedId]">
    <ExpensiveTree :item="selectedItem" />
  </div>
  
  <!-- 异步组件 -->
  <Suspense>
    <template #default>
      <HeavyComponent />
    </template>
    <template #fallback>
      <LoadingSkeleton />
    </template>
  </Suspense>
</template>
```

---

## 高效学习的核心心法

### 1. 官方文档是第一资源

> Vue 的官方文档是公认的最佳入门教程，结构清晰，示例精准。遇到任何概念，**优先查阅官方文档**。

- [Vue 3 官方文档](https://cn.vuejs.org/)
- [Vue Router 文档](https://router.vuejs.org/zh/)
- [Pinia 文档](https://pinia.vuejs.org/zh/)

### 2. 80/20 法则

先掌握那 20% 最常用的核心语法（响应式、组件、指令、组合式 API），它们能解决 80% 的开发问题。

**必须掌握的：**
- ✅ `ref` 和 `reactive`
- ✅ `computed` 和 `watch`
- ✅ 组件 Props 和 Emit
- ✅ `v-if` 和 `v-for`
- ✅ `<script setup>` 语法
- ✅ Vue Router 基础
- ✅ Pinia Store 基础

**可以后续再学的：**
- 自定义指令
- Teleport / Suspense
- 高级 TypeScript 类型
- 单元测试

### 3. "抄"以致用

在 GitHub 或教程中找到优秀的开源项目，**本地运行起来，读懂每一行核心代码，然后尝试修改和添加功能**。

推荐学习项目：
- [Vitesse](https://github.com/antfu/vitesse) - 现代 Vue 3 项目模板
- [Vue3 官方示例](https://github.com/vuejs/core/tree/main/packages/vue/examples)
- [awesome-vue](https://github.com/vuejs/awesome-vue) - Vue 生态资源汇总

### 4. 教是最好的学

尝试把你学到的知识点，用自己的话写成笔记或博客：

- 📝 为每个核心概念写一篇笔记
- 📚 写一个项目总结
- 🗣️ 对着空气（或朋友）讲解某个概念
- 🎥 录制自己的项目演示视频

### 5. 拥抱错误

编程中遇到的每一个报错都是学习机会：

1. **仔细阅读错误信息** - 很多错误已经说明了解决方案
2. **使用搜索引擎** - 善用 Google、Stack Overflow
3. **查阅官方文档** - 很多问题文档中已有答案
4. **查看 GitHub Issues** - 可能找到相同问题的讨论
5. **提问时提供最小复现** - 帮助他人理解你的问题

---

## 推荐资源

### 文档

| 资源 | 链接 |
|:---|:---|
| Vue 3 官方文档 | https://cn.vuejs.org/ |
| Vue Router | https://router.vuejs.org/zh/ |
| Pinia | https://pinia.vuejs.org/zh/ |
| Vite | https://cn.vitejs.dev/ |
| TypeScript | https://www.typescriptlang.org/zh/ |

### 视频课程

在 B 站搜索以下关键词，选择播放量高、日期较新的课程：

- `Vue 3 组合式 API`
- `Vue 3 + Pinia + Router 项目实战`
- `Vue 3 + TypeScript`
- `Vite 从入门到精通`

### 社区

| 社区 | 链接 |
|:---|:---|
| Vue GitHub | https://github.com/vuejs/core |
| Vue Discord | https://discord.com/invite/HBherRA |
| SegmentFault | https://segmentfault.com/t/vue.js |
| 掘金 | https://juejin.cn/ |

---

## 四周学习总结

| 周次 | 完成标准 |
|:---:|:---|
| **第一周** | 搭建好环境，理解 Vue 核心概念（响应式、模板、组件），完成 Todo List 基础版 |
| **第二周** | 掌握计算属性、侦听器、生命周期，能独立拆分组件，完成 Todo List 完整版 |
| **第三周** | 熟练使用 Vue Router 和 Pinia，能进行多页面开发，完成博客项目核心功能 |
| **第四周** | 了解 TypeScript 基础用法，掌握 Vite 配置和调试技巧，尝试进阶主题 |

---

## 快速学习公式

```
快速学习 Vue = 官方文档 + 不间断地写代码 + 一个贯穿始终的实战项目
```

**四周后，你将具备：**
- ✅ 独立开发中等复杂度 Vue 应用的能力
- ✅ 熟练使用 Vue Router 管理页面
- ✅ 使用 Pinia 进行状态管理
- ✅ 编写可复用的组合式函数
- ✅ 基本的 TypeScript + Vue 开发能力

---

> 🚀 **祝你学习顺利！如有疑问，欢迎随时提问。**
