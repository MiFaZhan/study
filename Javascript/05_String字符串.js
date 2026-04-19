// String 字符串方法

const str = "Hello, World!";

// 查找
console.log("查找:");
console.log("indexOf:", str.indexOf("World"));        // 7
console.log("lastIndexOf:", str.lastIndexOf("o"));     // 8
console.log("includes:", str.includes("World"));      // true
console.log("startsWith:", str.startsWith("Hello"));  // true
console.log("endsWith:", str.endsWith("!"));          // true
console.log("search:", str.search(/World/));          // 7

// 替换
console.log("\n替换:");
console.log("replace:", str.replace("World", "JS"));  // Hello, JS!
console.log("replaceAll:", str.replaceAll("o", "O")); // Hello, WOrld!
console.log("replace正则:", str.replace(/o/g, "O"));   // Hello, WOrld!

// 分割
console.log("\n分割:");
const email = "user@example.com";
console.log("split:", email.split("@"));              // ["user", "example.com"]
const path = "a/b/c/d";
console.log("split限制:", path.split("/", 2));         // ["a", "b"]

// 子字符串
console.log("\n子字符串:");
console.log("substring:", str.substring(0, 5));       // Hello
console.log("slice:", str.slice(0, 5));                // Hello
console.log("slice负数:", str.slice(-6));              // World!
console.log("substr已废弃:", str.substr(0, 5));        // Hello

// 大小写
console.log("\n大小写:");
console.log("toUpperCase:", str.toUpperCase());       // HELLO, WORLD!
console.log("toLowerCase:", str.toLowerCase());       // hello, world!

// 去除空白
console.log("\n去除空白:");
const padded = "  Hello  ";
console.log("trim:", padded.trim());                  // Hello
console.log("trimStart:", padded.trimStart());         // "Hello  "
console.log("trimEnd:", padded.trimEnd());             // "  Hello"

// 连接与填充
console.log("\n连接与填充:");
console.log("concat:", str.concat(" ", "Good!"));      // Hello, World! Good!
console.log("padStart:", "5".padStart(4, "0"));        // 0005
console.log("padEnd:", "5".padEnd(4, "0"));            // 5000

// 字符访问
console.log("\n字符访问:");
console.log("charAt:", str.charAt(1));                 // e
console.log("charCodeAt:", str.charCodeAt(1));         // 101
console.log("at:", str.at(-1));                        // !

// 其他方法
console.log("\n其他:");
console.log("length:", str.length);                   // 13
console.log("match:", str.match(/[A-Z]/g));            // ["H", "W"]
console.log("matchAll:", [...str.matchAll(/[a-z]/g)]); // 遍历所有匹配
console.log("localeCompare:", "apple".localeCompare("banana")); // -1
console.log("repeat:", "ab".repeat(3));                // ababab

// 模板字符串（反引号）
const name = "Alice";
const age = 25;
const template = `Name: ${name}, Age: ${age}`;
console.log("模板字符串:", template);

console.log("String示例完成");
