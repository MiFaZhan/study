// TypeScript 泛型

// 1. 基本泛型函数
function identity<T>(arg: T): T {
    return arg;
}

console.log("identity<string>:", identity<string>("hello"));
console.log("identity<number>:", identity<number>(42));

// 2. 泛型数组
function logArray<T>(arr: T[]): void {
    arr.forEach(item => console.log(item));
}

logArray<number>([1, 2, 3]);
logArray<string>(["a", "b", "c"]);

// 3. 泛型接口
interface Container<T> {
    value: T;
    getValue(): T;
}

const numContainer: Container<number> = {
    value: 100,
    getValue() { return this.value; }
};
console.log("泛型接口:", numContainer.getValue());

// 4. 泛型类
class Stack<T> {
    private items: T[] = [];
    
    push(item: T): void {
        this.items.push(item);
    }
    
    pop(): T | undefined {
        return this.items.pop();
    }
    
    peek(): T | undefined {
        return this.items[this.items.length - 1];
    }
    
    isEmpty(): boolean {
        return this.items.length === 0;
    }
}

const stack = new Stack<number>();
stack.push(1);
stack.push(2);
stack.push(3);
console.log("Stack pop:", stack.pop());
console.log("Stack peek:", stack.peek());

// 5. 泛型约束
interface Lengthwise {
    length: number;
}

function logLength<T extends Lengthwise>(arg: T): number {
    return arg.length;
}

console.log("泛型约束:", logLength("hello"));
console.log("泛型约束数组:", logLength([1, 2, 3]));
console.log("泛型约束对象:", logLength({ length: 10 }));

// 6. 多个类型参数
function pair<K, V>(key: K, value: V): [K, V] {
    return [key, value];
}

const p = pair<string, number>("age", 25);
console.log("多个类型参数:", p);

// 7. 泛型约束另一个参数
function merge<T extends object, U extends object>(obj1: T, obj2: U): T & U {
    return { ...obj1, ...obj2 };
}

console.log("merge:", merge({ name: "Alice" }, { age: 25 }));

// 8. 使用 keyof 约束
function getProperty<T, K extends keyof T>(obj: T, key: K): T[K] {
    return obj[key];
}

const user = { name: "Bob", age: 30 };
console.log("getProperty:", getProperty(user, "name"));

// 9. 泛型默认值
interface Container2<T = string> {
    value: T;
}

const c1: Container2 = { value: "default" };
const c2: Container2<number> = { value: 123 };
console.log("默认值:", c1, c2);

// 10. 条件类型与泛型
type NonNullable<T> = T extends null | undefined ? never : T;
type Res1 = NonNullable<string | null | undefined>; // string

// 11. 泛型工具类型（实现）
type MyPartial<T> = { [P in keyof T]?: T[P] };
type MyRequired<T> = { [P in keyof T]-?: T[P] };
type MyPick<T, K extends keyof T> = { [P in K]: T[P] };
type MyRecord<K extends keyof any, T> = { [P in K]: T };

interface User2 { name: string; age: number; }
type PartialUser = MyPartial<User2>;
type UserNameOnly = MyPick<User2, "name">;

// 12. 多重泛型
function curry<T, U, V>(fn: (a: T, b: U) => V): (a: T) => (b: U) => V {
    return (a: T) => (b: U) => fn(a, b);
}

const add = (a: number, b: number): number => a + b;
const curriedAdd = curry(add);
console.log("curry:", curriedAdd(1)(2));

// 13. 泛型枚举
enum ResponseType {
    Success = "success",
    Error = "error"
}

interface ApiResponse<T> {
    type: ResponseType;
    data: T;
    timestamp: number;
}

const successResp: ApiResponse<User[]> = {
    type: ResponseType.Success,
    data: [{ name: "Alice", age: 25, id: 1 }],
    timestamp: Date.now()
};
console.log("泛型枚举:", successResp);

console.log("泛型示例完成");
