// TypeScript 接口

// 1. 基本接口
interface Point {
    x: number;
    y: number;
}

const point: Point = { x: 10, y: 20 };
console.log("基本接口:", point);

// 2. 可选属性
interface User {
    name: string;
    age: number;
    email?: string;
    phone?: string;
}

const user1: User = { name: "Alice", age: 25 };
const user2: User = { name: "Bob", age: 30, email: "bob@example.com" };
console.log("可选属性:", user1, user2);

// 3. 只读属性
interface Config {
    readonly apiUrl: string;
    readonly timeout: number;
}

const config: Config = { apiUrl: "https://api.example.com", timeout: 5000 };
// config.apiUrl = "new"; // 错误：只读属性不能修改

// 4. 函数类型接口
interface SearchFunc {
    (source: string, subString: string): boolean;
}

const search: SearchFunc = (source, sub) => {
    return source.includes(sub);
};
console.log("函数类型接口:", search("hello world", "world"));

// 5. 索引签名
interface StringArray {
    [index: number]: string;
}

const arr: StringArray = ["a", "b", "c"];
console.log("索引签名:", arr[0]);

interface StringMap {
    [key: string]: number;
}

const scores: StringMap = { Alice: 90, Bob: 85 };
console.log("字符串索引:", scores);

// 6. 继承接口
interface Animal {
    name: string;
}

interface Dog extends Animal {
    breed: string;
}

const dog: Dog = { name: "旺财", breed: "金毛" };
console.log("继承接口:", dog);

// 7. 多重继承
interface A { a: string; }
interface B { b: number; }
interface C extends A, B { c: boolean; }

const c: C = { a: "test", b: 1, c: true };
console.log("多重继承:", c);

// 8. 接口合并（声明合并）
interface Window {
    ts: number;
}
// Window.ts 已添加

// 9. 接口描述类
interface Speakable {
    speak(): void;
}

class Person implements Speakable {
    speak(): void {
        console.log("Hello!");
    }
}
new Person().speak();

// 10. 接口描述函数
interface Calculate {
    (a: number, b: number): number;
}

const add: Calculate = (x, y) => x + y;
const multiply: Calculate = (x, y) => x * y;
console.log("接口函数:", add(2, 3), multiply(2, 3));

// 11. 接口描述构造函数
interface ClockConstructor {
    new(hour: number, minute: number): ClockInterface;
}

interface ClockInterface {
    tick(): void;
}

function createClock(ctor: ClockConstructor, hour: number, minute: number): ClockInterface {
    return new ctor(hour, minute);
}

class DigitalClock implements ClockInterface {
    constructor(h: number, m: number) {}
    tick(): void { console.log("beep beep"); }
}

const clock = createClock(DigitalClock, 10, 30);
clock.tick();

// 12. 接口作为类型
interface Response<T> {
    data: T;
    status: number;
    message: string;
}

const resp: Response<User> = {
    data: { name: "Alice", age: 25, id: 1 },
    status: 200,
    message: "success"
};
console.log("泛型接口:", resp);

console.log("接口示例完成");
