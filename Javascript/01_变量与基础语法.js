// 变量声明
let userName = "Alice";

const age = 25;
let isLogin = true;;
let user = null;

let data;
console.log("data:", data);

let uniqueId = Symbol();
console.log("uniqueId:", uniqueId);

let bigNum = 9007199254740991n;
console.log("bigNum:", bigNum);


// 条件
if (age >= 18) {
    console.log("已成年");
} else {
    console.log("未成年");
}

let key = "value";
switch (key) {
    case "value":
        console.log("switch匹配到value");
        break;
    default:
        console.log("switch未匹配");
}

// 循环
console.log("for循环:");
for (let i = 0; i < 5; i++) {
    console.log(i);
}

let condition = true;
let count = 0;
console.log("while循环:");
while (condition) {
    console.log(count);
    count++;
    if (count >= 3) condition = false;
}

let array = [10, 20, 30];
console.log("for...of遍历数组:");
// for...of — 遍历值
for (let item of array) {
    console.log(item);
}

let object = { name: "Alice", age: 25 };
console.log("for...in遍历对象:");
// for...in — 遍历索引
for (let k in object) {
    console.log(k, ":", object[k]);
}

// 函数 (三种定义方式)
function add(a, b) { return a + b; }
console.log("add(2, 3) =", add(2, 3));

const subtract = (a, b) => a - b;
console.log("subtract(5, 2) =", subtract(5, 2));

// 没有提升：必须在定义后调用
const multiply = function(a, b) { return a * b; };
console.log("multiply(3, 4) =", multiply(3, 4));

// 类 (ES6+)
class Person {
    constructor(name) { this.name = name; }
    greet() { console.log(`Hello, I'm ${this.name}`); }
}
const p = new Person("Bob");
p.greet();

console.log("程序执行完毕!");
