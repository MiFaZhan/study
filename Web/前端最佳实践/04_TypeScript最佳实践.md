# TypeScript 最佳实践

---

## 1. 类型基础

```typescript
// ✅ 基础类型
let name: string = '张三';
let age: number = 25;
let isActive: boolean = true;

// ✅ 数组
let nums: number[] = [1, 2, 3];
let names: Array<string> = ['a', 'b'];

// ✅ 元组
let tuple: [string, number] = ['name', 25];

// ✅ 联合类型
let status: string | number;
status = 'pending';
status = 0;

// ✅ 字面量类型
let direction: 'left' | 'right' | 'up' | 'down';

// ✅ void 用于无返回值函数
function log(message: string): void {
    console.log(message);
}

// ✅ never 用于永不返回的函数
function throwError(msg: string): never {
    throw new Error(msg);
}
```

---

## 2. 接口 vs 类型别名

```typescript
// ✅ 接口可扩展，推荐用于数据模型
interface User {
    id: number;
    name: string;
    email?: string;  // 可选
    readonly createdAt: Date;  // 只读
}

interface Admin extends User {
    permissions: string[];
}

// ✅ 类型别名用于联合、交叉等复杂类型
type Status = 'pending' | 'active' | 'deleted';
type Point = { x: number } & { y: number };
```

---

## 3. 泛型

```typescript
// ✅ 泛型函数
function identity<T>(arg: T): T {
    return arg;
}
identity<string>('hello');
identity(123);  // 类型推断

// ✅ 泛型接口
interface ApiResponse<T> {
    code: number;
    data: T;
    message: string;
}

// ✅ 泛型约束
function getProperty<T, K extends keyof T>(obj: T, key: K): T[K] {
    return obj[key];
}
getProperty({ name: '张三', age: 25 }, 'name');  // '张三'

// ✅ 多泛型
function pair<K, V>(k: K, v: V): [K, V] {
    return [k, v];
}
```

---

## 4. 数组方法类型

```typescript
// ✅ filter - 返回数组
const users: User[] = [{ name: '张三', age: 25 }, { name: '李四', age: 30 }];
const adults = users.filter((u: User): boolean => u.age >= 18);

// ✅ map - 转换数组
const names = users.map((u: User): string => u.name);

// ✅ reduce - 汇总
const totalAge = users.reduce((sum: number, u: User): number => sum + u.age, 0);

// ✅ find - 查找单个
const user = users.find((u: User): boolean => u.name === '张三');
```

---

## 5. 枚举

```typescript
// ✅ 数字枚举
enum Status {
    Pending,   // 0
    Active,    // 1
    Deleted   // 2
}

// ✅ 字符串枚举
enum UserRole {
    Admin = 'ADMIN',
    User = 'USER',
    Guest = 'GUEST'
}

// ✅ const 枚举（性能更好）
const enum Direction {
    Up = 'UP',
    Down = 'DOWN'
}
```

---

## 6. 解构与类型

```typescript
// ✅ 对象解构
interface Config { url: string; timeout: number; }
function request({ url, timeout }: Config) { }

// ✅ 带默认值的解构
function connect({ host = 'localhost', port = 8080 }: { host?: string; port?: number }) { }

// ✅ 数组解构
function parseFirstTwo([a, b]: number[]): [number, number] {
    return [a, b];
}
```

---

## 7. 常用工具类型

```typescript
// ✅ Partial - 全部属性可选
type PartialUser = Partial<User>;

// ✅ Required - 全部属性必填
type RequiredUser = Required<User>;

// ✅ Pick - 选取部分属性
type UserPreview = Pick<User, 'id' | 'name'>;

// ✅ Omit - 排除部分属性
type UserWithoutEmail = Omit<User, 'email'>;

// ✅ Record - 键值对
type UserMap = Record<string, User>;

// ✅ Exclude / Extract
type Status = 'pending' | 'active' | 'deleted';
type NonDeletedStatus = Exclude<Status, 'deleted'>;
type ActiveStatus = Extract<Status, 'active'>;

// ✅ ReturnType - 获取函数返回值
function createUser() { return { id: 1, name: '张三' }; }
type User = ReturnType<typeof createUser>;
```

---

## 8. 类与类型

```typescript
// ✅ 类实现接口
interface Printable {
    print(): void;
}

class Report implements Printable {
    print(): void {
        console.log('打印报表');
    }
}

// ✅ 抽象类
abstract class Shape {
    abstract area(): number;
    
    describe(): void {
        console.log('这是一个形状');
    }
}

class Circle extends Shape {
    constructor(public radius: number) {
        super();
    }
    
    area(): number {
        return Math.PI * this.radius ** 2;
    }
}
```

---

## 9. 严格类型

```typescript
// ✅ 避免 any，用 unknown 代替
function parseJSON(input: string): unknown {
    return JSON.parse(input);
}

// 使用时需要类型检查
const data = parseJSON(str);
if (typeof data === 'string') {
    console.log(data.toUpperCase());
}

// ✅ 类型守卫
function isUser(obj: unknown): obj is User {
    return typeof obj === 'object' && obj !== null && 'name' in obj;
}

// ✅ 非空断言（谨慎使用）
const el = document.querySelector('#app')!;
```

---

## 10. tsconfig 关键配置

```json
{
    "compilerOptions": {
        "strict": true,
        "noImplicitAny": true,
        "strictNullChecks": true,
        "strictFunctionTypes": true,
        "noUnusedLocals": true,
        "noUnusedParameters": true,
        "noImplicitReturns": true,
        "esModuleInterop": true,
        "skipLibCheck": true,
        "forceConsistentCasingInFileNames": true
    }
}
```

---

## 11. 常见类型错误处理

```typescript
// ✅ 可选链
const city = user?.address?.city;

// ✅ 空值合并
const name = userName ?? '匿名';

// ✅ 类型断言（谨慎使用）
const el = document.getElementById('app') as HTMLElement;
const input = event.target as HTMLInputElement;

// ✅ 双类型断言（any 过渡）
const str: string = data as any as string;
```

---

## 速查表

| 规范 | 示例 |
|------|------|
| 接口用于数据模型 | `interface User { }` |
| 类型别名用于复杂类型 | `type Status = 'a' \| 'b'` |
| 泛型约束 | `K extends keyof T` |
| 避免 any | 用 `unknown` |
| 工具类型 | `Partial`, `Pick`, `Omit` |
| 严格模式 | `strict: true` |
| 可选链 | `user?.address?.city` |
