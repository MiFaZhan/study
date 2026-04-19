# JavaScript 最佳实践

---

## 1. 变量命名

```javascript
// ✅ 常量全大写下划线分隔
const MAX_RETRY_COUNT = 3;
const API_BASE_URL = 'https://api.example.com';

// ✅ 变量/函数 camelCase
let userName = '张三';
function getUserInfo() { }

// ✅ 布尔值用 is/has/can 前缀
let isActive = true;
let hasPermission = false;
let canEdit = true;

// ✅ 数组用复数名词或 _list 后缀
let users = [];
let userList = [];

// ✅ 函数名用动宾短语
function getUserById(id) { }
function validateEmail(email) { }
function handleClick() { }
```

---

## 2. 函数设计

```javascript
// ✅ 单一职责
function getUser(id) { }
function sendEmail(to, content) { }

// ✅ 默认参数
function greet(name = '游客') {
    return `你好, ${name}`;
}

// ✅ 箭头函数简写
const add = (a, b) => a + b;
const getName = user => user.name;

// ✅ 避免回调地狱，用 async/await
// ❌
api.getUser(id, (err, user) => {
    api.getOrders(user.id, (err, orders) => {
        // 地狱
    });
});

// ✅
const user = await api.getUser(id);
const orders = await api.getOrders(user.id);
```

---

## 3. 字符串处理

```javascript
// ✅ 模板字符串
const message = `用户 ${userName} 已登录`;

// ✅ 字符串方法
const email = 'TEST@EXAMPLE.COM';
email.toLowerCase();                    // 'test@example.com'
email.includes('@');                     // true
email.startsWith('test');               // true
'Hello World'.replace('World', 'JS');  // 'Hello JS'
'hello,world'.split(',');              // ['hello', 'world']
'Hello'.slice(0, 2);                   // 'He'

// ✅ 判断空字符串
if (str.length === 0) { }
if (!str.trim()) { }  // 包含空白的空字符串
```

---

## 4. 数组操作

```javascript
// ✅ 常用方法
const nums = [1, 2, 3, 4, 5];

nums.map(n => n * 2);          // [2, 4, 6, 8, 10]
nums.filter(n => n > 2);      // [3, 4, 5]
nums.reduce((sum, n) => sum + n, 0);  // 15
nums.find(n => n > 2);         // 3
nums.findIndex(n => n > 2);   // 2
nums.some(n => n > 4);        // true
nums.every(n => n > 0);       // true

// ✅ 展开运算符合并
const arr1 = [1, 2];
const arr2 = [3, 4];
const merged = [...arr1, ...arr2];  // [1, 2, 3, 4]

// ✅ 判断数组
Array.isArray([]);  // true
```

---

## 5. 对象操作

```javascript
// ✅ 解构赋值
const { name, age } = user;
const { data: result } = response;

// ✅ 展开运算符合并
const defaults = { theme: 'light', lang: 'zh' };
const config = { ...defaults, theme: 'dark' };  // { theme: 'dark', lang: 'zh' }

// ✅ 动态属性
const key = 'email';
const obj = { [key]: 'test@example.com' };

// ✅ 判断属性存在
obj.hasOwnProperty('name');
'name' in obj;
obj.name !== undefined;
```

---

## 6. 条件判断

```javascript
// ✅ 短路运算
const name = userInput || '默认名称';
const isValid = value !== null && value !== undefined;

// ✅ 三元运算简单情况
const label = isActive ? '启用' : '禁用';

// ✅ 多条件用数组 includes
if (status === 'pending' || status === 'processing') { }
// 改为
if (['pending', 'processing'].includes(status)) { }

// ✅ 早期返回减少嵌套
function processUser(user) {
    if (!user) return;
    if (!user.isActive) return;
    // 正式处理
}
```

---

## 7. 错误处理

```javascript
// ✅ try-catch
try {
    const data = JSON.parse(str);
} catch (err) {
    console.error('解析失败:', err.message);
}

// ✅ 自定义错误类
class ValidationError extends Error {
    constructor(message, field) {
        super(message);
        this.name = 'ValidationError';
        this.field = field;
    }
}

// ✅ async await 错误处理
async function fetchData() {
    try {
        const res = await fetch(url);
        return await res.json();
    } catch (err) {
        console.error('请求失败:', err);
        throw err;
    }
}
```

---

## 8. DOM 操作

```javascript
// ✅ 缓存 DOM 查询
const container = document.querySelector('.container');
container.addEventListener('click', handleClick);

// ✅ 事件委托（减少监听器）
document.querySelector('.list').addEventListener('click', (e) => {
    if (e.target.matches('.item')) {
        handleItemClick(e.target);
    }
});

// ✅ 避免频繁操作 DOM
// ❌
for (let i = 0; i < 100; i++) {
    el.innerHTML += '<span>item</span>';
}

// ✅ 拼接后一次性写入
const fragment = document.createDocumentFragment();
for (let i = 0; i < 100; i++) {
    fragment.appendChild(createSpan());
}
el.appendChild(fragment);
```

---

## 9. 类型转换

```javascript
// ✅ 转数字
Number('123');        // 123
parseInt('123', 10);  // 123
parseFloat('1.5');    // 1.5
+'123';               // 123 (一元加号)

// ✅ 转字符串
String(123);          // '123'
(123).toString();     // '123'
'' + 123;             // '123' (不推荐)

// ✅ 转布尔
Boolean(x);
!!x;
!x;  // 取反后再取反

// ✅ 判断
typeof 123;                    // 'number'
Array.isArray([]);            // true
obj instanceof Array;         // true
```

---

## 10. 模块化

```javascript
// ✅ 命名导出
export const PI = 3.14;
export function add(a, b) { return a + b; }

// ✅ 默认导出
export default class UserService { }

// ✅ 导入时解构
import { add, PI } from './math';
import UserService from './UserService';
```

---

## 11. 安全注意

```javascript
// ❌ 避免 eval
eval(userInput);  // 安全风险

// ✅ DOM XSS 防范
element.textContent = userInput;  // 而非 innerHTML
```

---

## 12. 性能注意

```javascript
// ✅ 节流/防抖
function throttle(fn, delay) {
    let last = 0;
    return (...args) => {
        const now = Date.now();
        if (now - last >= delay) {
            last = now;
            fn(...args);
        }
    };
}

// ✅ 懒加载
const loadImage = (src) => {
    const img = new Image();
    img.src = src;
};

// ✅ 避免全局查找
// ❌
function update() {
    document.querySelector('#id').textContent = value;
}

// ✅ 缓存
const el = document.querySelector('#id');
function update() {
    el.textContent = value;
}
```

---

## 速查表

| 规范 | 示例 |
|------|------|
| 常量命名 | `MAX_COUNT` |
| 变量命名 | `userName` |
| 布尔命名 | `isActive`, `hasError` |
| 函数命名 | `getUserById()` |
| 条件判断 | `[a, b].includes(x)` |
| 异步处理 | async/await |
| DOM 缓存 | `const el = ...` |
| 错误处理 | try-catch + 自定义错误 |
