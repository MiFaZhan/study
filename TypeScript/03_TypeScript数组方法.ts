// Array 数组方法在 TypeScript 中的使用

const numbers: number[] = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];

// 1. filter - 过滤
const evens: number[] = numbers.filter(n => n % 2 === 0);
console.log("filter偶数:", evens);

const words: string[] = ["apple", "banana", "cherry", "date"];
const longWords: string[] = words.filter(w => w.length > 5);
console.log("filter长单词:", longWords);

// 2. map - 映射/转换
const doubled: number[] = numbers.map(n => n * 2);
console.log("map翻倍:", doubled);

const strings: string[] = numbers.map(n => `数字: ${n}`);
console.log("map转字符串:", strings);

// 带类型的map
const users = [
    { name: "Alice", age: 25 },
    { name: "Bob", age: 30 }
];
const userNames: string[] = users.map(u => u.name);
console.log("map提取name:", userNames);

// 3. join - 合并为字符串
const joined: string = numbers.join("-");
console.log("join:", joined);

const wordsJoined: string = words.join(", ");
console.log("join单词:", wordsJoined);

// 4. 链式调用
const result: string = numbers
    .filter(n => n > 5)
    .map(n => n * 2)
    .join(", ");
console.log("链式调用:", result);

// 5. 综合示例
interface Product {
    name: string;
    price: number;
    category: string;
}

const products: Product[] = [
    { name: "笔记本", price: 5000, category: "电子产品" },
    { name: "手机", price: 3000, category: "电子产品" },
    { name: "T恤", price: 100, category: "服装" },
    { name: "裤子", price: 200, category: "服装" },
    { name: "耳机", price: 800, category: "电子产品" }
];

// 筛选电子产品并获取名称
const electronicsNames: string[] = products
    .filter(p => p.category === "电子产品")
    .map(p => p.name);
console.log("电子产品:", electronicsNames);

// 计算总价
const totalPrice: number = products.reduce((sum, p) => sum + p.price, 0);
console.log("总价:", totalPrice);

// 按分类分组
const grouped = products.reduce<Record<string, Product[]>>((acc, p) => {
    if (!acc[p.category]) acc[p.category] = [];
    acc[p.category].push(p);
    return acc;
}, {});
console.log("分组:", grouped);

// 查找最贵的产品
const expensiveProduct: Product | undefined = products
    .filter(p => p.category === "电子产品")
    .reduce((max, p) => p.price > max.price ? p : max);
console.log("最贵电子产品:", expensiveProduct);

console.log("Array方法TS示例完成");
