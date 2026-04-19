# Vue 最佳实践

---

## 1. 组件命名

```javascript
// ✅ PascalCase 命名
// MyComponent.vue
// UserCard.vue

// ✅ 目录结构
components/
├── AppHeader.vue
├── user/
│   ├── UserCard.vue
│   └── UserList.vue
```

---

## 2. 组件结构顺序

```vue
<template>
    <div class="user-card">
        <h2>{{ userName }}</h2>
    </div>
</template>

<script>
import { defineComponent } from 'vue';
import { mapState } from 'vuex';
import UserAvatar from './UserAvatar.vue';

export default defineComponent({
    name: 'UserCard',
    
    components: { UserAvatar },
    
    props: {
        userName: {
            type: String,
            required: true
        }
    },
    
    emits: ['update'],
    
    setup() {
        // setup()
    },
    
    data() {
        return { localCount: 0 };
    },
    
    computed: {
        ...mapState(['user'])
    },
    
    watch: {
        userName(newVal) {
            this.localCount = newVal.length;
        }
    },
    
    created() { },
    mounted() { },
    
    methods: {
        handleClick() { }
    }
});
</script>

<style scoped>
.user-card {
    padding: 1rem;
}
</style>
```

---

## 3. Props 定义

```typescript
// ✅ 用对象形式定义，带类型校验
props: {
    title: {
        type: String,
        required: true
    },
    count: {
        type: Number,
        default: 0
    },
    items: {
        type: Array,
        default: () => []  // 数组默认值必须用函数
    },
    config: {
        type: Object,
        default: () => ({})  // 对象默认值必须用函数
    }
}

// ✅ TypeScript 写法
interface Props {
    title: string;
    count?: number;
    items?: string[];
}

const props = withDefaults(defineProps<Props>(), {
    count: 0,
    items: () => []
});
```

---

## 4. Emit 事件

```typescript
// ✅ 对象形式定义（带验证）
emits: {
    'update:value': (val: string) => typeof val === 'string',
    'click': null
}

// ✅ 组合式 API
const emit = defineEmits(['update:value', 'click']);
emit('update:value', newValue);
```

---

## 5. 响应式数据

```typescript
// ✅ ref 用于基本类型
const count = ref(0);
const name = ref<string>('');

// ✅ reactive 用于对象
const state = reactive({
    user: null,
    loading: false,
    error: null
});

// ✅ computed
const doubleCount = computed(() => count.value * 2);

// ✅ watch
watch(count, (newVal, oldVal) => {
    console.log(`count 从 ${oldVal} 变为 ${newVal}`);
});

// ✅ 避免响应式开销
// ❌ 整个大对象响应式
const state = reactive(bigObject);

// ✅ 只响应需要的字段
const userName = computed(() => bigObject.user.name);
```

---

## 6. 生命周期顺序

```
beforeCreate → created → beforeMount → mounted
→ beforeUpdate → updated → beforeUnmount → unmounted
```

```typescript
// ✅ 典型用法
onMounted(() => {
    // DOM 已挂载，开始网络请求
    fetchData();
});

onUnmounted(() => {
    // 清理定时器、事件监听
    clearInterval(timer);
    window.removeEventListener('resize', handleResize);
});
```

---

## 7. v-if vs v-show

```html
<!-- v-if 真删除（条件很少改变时） -->
<v-if v-if="isShow">内容</v-if>

<!-- v-show 假隐藏（频繁切换时） -->
<v-show v-show="isShow">内容</v-show>
```

---

## 8. v-for 规范

```html
<!-- ✅ 必须加 key，用稳定唯一值 -->
<li v-for="user in users" :key="user.id">

<!-- ❌ 避免用 index 作为 key -->
<li v-for="(user, index) in users" :key="index">

<!-- ✅ 配合 template 遍历多元素 -->
<template v-for="user in users" :key="user.id">
    <li>{{ user.name }}</li>
    <li>{{ user.email }}</li>
</template>
```

---

## 9. 组件通信

```typescript
// ✅ 父子通信：props 向下，emit 向上
// Parent
<ChildComponent 
    :title="pageTitle" 
    @update="handleUpdate" 
/>

// Child
const props = defineProps<{ title: string }>();
const emit = defineEmits(['update']);
emit('update', newValue);

// ✅ provide/inject（跨级通信）
// 父级
provide('theme', 'dark');

// 子孙级
const theme = inject<string>('theme');

// ✅ 事件总线（简单场景）
// 初始化
const bus = mitt();
// 使用
bus.on('event', handler);
bus.emit('event', payload);
// 清理
bus.off('event', handler);
```

---

## 10. 路由规范

```typescript
// ✅ 路由按模块组织
const routes = [
    {
        path: '/',
        name: 'Home',
        component: () => import('@/views/Home.vue')
    },
    {
        path: '/user',
        name: 'UserLayout',
        component: () => import('@/layouts/UserLayout.vue'),
        children: [
            { path: 'list', name: 'UserList', component: () => import('@/views/user/List.vue') },
            { path: ':id', name: 'UserDetail', component: () => import('@/views/user/Detail.vue') }
        ]
    }
];

// ✅ 路由守卫顺序
router.beforeEach((to, from, next) => {
    // 1. 登录验证
    if (to.meta.requiresAuth && !isLoggedIn) {
        return next({ name: 'Login', query: { redirect: to.fullPath } });
    }
    // 2. 权限验证
    if (to.meta.permissions && !hasPermission(to.meta.permissions)) {
        return next({ name: 'Forbidden' });
    }
    // 3. 放行
    next();
});
```

---

## 11. Vuex/Pinia 状态管理

```typescript
// ✅ Pinia 示例（推荐 Vue3）
// stores/user.ts
export const useUserStore = defineStore('user', {
    state: () => ({
        name: '',
        token: ''
    }),
    
    getters: {
        isLoggedIn: (state) => !!state.token
    },
    
    actions: {
        async login(credentials) {
            const user = await api.login(credentials);
            this.name = user.name;
            this.token = user.token;
        },
        
        logout() {
            this.name = '';
            this.token = '';
        }
    }
});

// 组件中使用
const userStore = useUserStore();
if (userStore.isLoggedIn) {
    // ...
}
```

---

## 12. 性能优化

```typescript
// ✅ 路由懒加载
const UserDetail = () => import('@/views/user/Detail.vue');

// ✅ 组件懒加载
const HeavyChart = defineAsyncComponent(() => import('./HeavyChart.vue'));

// ✅ 防抖
import { debounce } from 'lodash-es';
const debouncedSearch = debounce(search, 300);

// ✅ v-memo（长列表优化）
<li v-for="item in list" v-memo="[item.id, item.name]">

// ✅ 保持静态内容不变
// 用 shallowRef 而不是 ref
const list = shallowRef([...veryLargeArray]);
```

---

## 13. 样式规范

```vue
<!-- ✅ scoped 隔离样式 -->
<style scoped>
.card { ... }
</style>

<!-- ✅ 深度选择器 -->
<style scoped>
:deep(.el-button) { ... }  /* Vue3 */
::v-deep .el-button { ... }  /* Vue2 */
</style>

<!-- ✅ CSS 变量主题 -->
<style scoped>
:root {
    --primary-color: v-bind(primaryColor);
}
</style>

<!-- ✅ 避免太多全局样式，优先用 CSS Modules -->
```

---

## 14. 常用指令缩写

```html
<!-- ✅ v-bind: → : -->
<input :value="name" :class="{ active: isActive }">

<!-- ✅ v-on: → @ -->
<button @click="handleClick" @mouseenter="onHover">

<!-- ✅ v-slot: → # -->
<template #header>
    <h1>标题</h1>
</template>
```

---

## 速查表

| 规范 | 推荐 |
|------|------|
| 组件名 | PascalCase |
| Props | 对象形式+类型 |
| v-for | 必须加 key |
| 响应式 | ref/reactive |
| 生命周期 | onMounted/onUnmounted |
| 路由 | 懒加载 |
| 状态 | Pinia/Vuex |
| 样式 | scoped |
| API | async/await |
