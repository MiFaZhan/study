// TypeScript 基础类型

// 1. 基本类型
let name: string = "Alice";
let age: number = 25;
let isActive: boolean = true;

// 2. 数组
let numbers: number[] = [1, 2, 3];
let names: Array<string> = ["a", "b", "c"];

// 3. 元组
let tuple: [string, number] = ["Alice", 25];
console.log("元组:", tuple[0], tuple[1]);

// 4. 枚举
enum Status {
    Pending = "pending",
    Active = "active",
    Completed = "completed"
}
let currentStatus: Status = Status.Active;
console.log("枚举值:", currentStatus);
console.log("枚举键:", Status.Active);

// 5. Any（尽量少用）
let anything: any = 4;
anything = "string";
anything = true;

// 6. Void（无返回值）
function logMessage(): void {
    console.log("无返回值函数");
}

// 7. Null 和 Undefined
let nullValue: null = null;
let undefinedValue: undefined = undefined;

// 8. Never（永不返回）
function throwError(): never {
    throw new Error("错误");
}
function infiniteLoop(): never {
    while (true) {}
}

// 9. Object
let person: object = { name: "Alice", age: 25 };
console.log("Object:", person);

// 10. 类型断言
let someValue: any = "hello";
let strLength: number = (someValue as string).length;
let strLength2: number = (<string>someValue).length;

// 11. 接口
interface User {
    name: string;
    age: number;
    email?: string; // 可选属性
    readonly id: number; // 只读属性
}

const user: User = {
    name: "Bob",
    age: 30,
    id: 1
};
console.log("接口:", user);

// 12. 类
class Animal {
    private name: string;
    protected age: number;
    public constructor(name: string, age: number) {
        this.name = name;
        this.age = age;
    }
    protected speak(): void {
        console.log(`${this.name}叫了一声`);
    }
}

class Dog extends Animal {
    constructor(name: string, age: number) {
        super(name, age);
    }
    public makeSound(): void {
        this.speak();
    }
}

const dog = new Dog("旺财", 3);
dog.makeSound();
console.log("类型基础示例完成");
