// Array 数组方法

const arr = [1, 2, 3, 4, 5];
const fruits = ["apple", "banana", "orange"];

// 添加和删除（会修改原数组）
console.log("添加和删除:");
console.log("push:", arr.push(6));         // 6, arr变成[1,2,3,4,5,6]
console.log("pop:", arr.pop());           // 6, arr变回[1,2,3,4,5]
console.log("unshift:", arr.unshift(0));   // 6, arr变成[0,1,2,3,4,5]
console.log("shift:", arr.shift());       // 0, arr变回[1,2,3,4,5]

// 合并与切片
console.log("\n合并与切片:");
const merged = fruits.concat(["grape", "melon"]);
console.log("concat:", merged);           // ["apple","banana","orange","grape","melon"]
console.log("slice:", arr.slice(1, 3));    // [2, 3]
console.log("splice:", arr.splice(1, 2, "a", "b")); // 删除[2,3], arr变成[1,"a","b",4,5]

// 查找
console.log("\n查找:");
console.log("indexOf:", fruits.indexOf("banana"));     // 1
console.log("lastIndexOf:", [1,2,1,2].lastIndexOf(1)); // 2
console.log("includes:", fruits.includes("apple"));   // true
console.log("find:", arr.find(x => x > 3));           // 4
console.log("findIndex:", arr.findIndex(x => x > 3)); // 3

// 遍历
console.log("\n遍历:");
fruits.forEach((item, index) => console.log(`${index}: ${item}`));

console.log("map:", arr.map(x => x * 2));              // [2, 4, 6, 8, 10]
console.log("filter:", arr.filter(x => x > 2));       // [3, 4, 5]
console.log("some:", arr.some(x => x > 4));           // true
console.log("every:", arr.every(x => x > 0));         // true

// 排序与反转
console.log("\n排序与反转:");
const nums = [3, 1, 4, 1, 5];
console.log("sort:", nums.sort((a, b) => a - b));     // [1,1,3,4,5]
console.log("reverse:", nums.reverse());             // [5,4,3,1,1]

// 转换
console.log("\n转换:");
console.log("join:", fruits.join(" - "));             // apple - banana - orange
console.log("toString:", fruits.toString());          // apple,banana,orange

// 归并
console.log("\n归并:");
console.log("reduce:", arr.reduce((sum, x) => sum + x, 0));  // 15
console.log("reduceRight:", arr.reduceRight((sum, x) => sum + x, 0)); // 15

// 检查
console.log("\n检查:");
console.log("Array.isArray:", Array.isArray(arr));     // true

// 迭代器方法
console.log("\n迭代器:");
const iterator = arr.keys();
console.log("keys:", [...iterator]);                  // [0,1,2,3,4]
const values = arr.values();
console.log("values:", [...values]);                  // [1,2,3,4,5]
const entries = arr.entries();
console.log("entries:", [...entries]);                 // [[0,1],[1,2],[2,3],[3,4],[4,5]]

// 填充与复制
console.log("\n填充与复制:");
console.log("fill:", new Array(3).fill(0));           // [0, 0, 0]
console.log("copyWithin:", [1,2,3,4,5].copyWithin(0, 3)); // [4,5,3,4,5]

// flat 与 flatMap
console.log("\n扁平化:");
console.log("flat:", [1,[2,[3]]].flat(2));            // [1, 2, 3]
console.log("flatMap:", [1,2,3].flatMap(x => [x, x * 2])); // [1,2,2,4,3,6]

console.log("Array示例完成");
