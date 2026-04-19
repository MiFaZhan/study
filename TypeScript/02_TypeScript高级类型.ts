// TypeScript 高级类型

// 1. 交叉类型（&）
interface Person {
    name: string;
    age: number;
}
interface Employee {
    company: string;
    salary: number;
}
type Worker = Person & Employee;
const worker: Worker = {
    name: "Alice",
    age: 25,
    company: "TechCorp",
    salary: 10000
};
console.log("交叉类型:", worker);

// 2. 联合类型（|）
type StringOrNumber = string | number;
function printId(id: StringOrNumber): void {
    console.log("ID:", id);
}
printId("123");
printId(456);

// 3. 类型保护
interface Cat {
    meow(): void;
    run(): void;
}
interface Dog {
    bark(): void;
    run(): void;
}

function isCat(animal: Cat | Dog): animal is Cat {
    return (animal as Cat).meow !== undefined;
}

function speak(animal: Cat | Dog): void {
    if (isCat(animal)) {
        animal.meow();
    } else {
        animal.bark();
    }
}

// 4. 类型别名
type Point = { x: number; y: number };
type Callback = (data: string) => void;

// 5. 映射类型
type Readonly<T> = { readonly [P in keyof T]: T[P] };
type Partial<T> = { [P in keyof T]?: T[P] };
type OptionalUser = Partial<User>;
type FrozenUser = Readonly<User>;

// 6. 条件类型
type NonNullable<T> = T extends null | undefined ? never : T;
type Result = NonNullable<string | null | undefined>; // string

// 7. infer（推断）
type ReturnType<T> = T extends (...args: any[]) => infer R ? R : never;
function getString(): string { return "hello"; }
type MyReturnType = ReturnType<typeof getString>; // string

// 8. 索引访问类型
interface Users {
    [key: string]: number;
}
type UserAge = Users[string]; // number

// 9. 模板字面量类型
type EventName = `on${string}`;
type CSSUnit = `${number}px` | `${number}rem` | `${number}%`;

// 10. 分布式条件类型
type ToArray<T> = T extends any ? T[] : never;
type StrArr = ToArray<string | number>; // string[] | number[]

// 11. keyof 与 typeof
const userObj = { name: "Alice", age: 25 };
type UserKeys = keyof typeof userObj; // "name" | "age"

// 12. Exclude 与 Extract
type T1 = Exclude<"a" | "b" | "c", "a" | "b">; // "c"
type T2 = Extract<"a" | "b" | "c", "a" | "b">; // "a" | "b"

// 13. Required 与 Pick
type Required<T> = { [P in keyof T]-?: T[P] };
type Pick<T, K extends keyof T> = { [P in K]: T[P] };

// 14. Record
type UserRecord = Record<string, User>;
const users: UserRecord = {
    "user1": { name: "Alice", age: 25, id: 1 },
    "user2": { name: "Bob", age: 30, id: 2 }
};
console.log("Record:", users);

console.log("高级类型示例完成");
