// Object 对象方法

const obj = { name: "Alice", age: 25, city: "Beijing" };

// 判断属性
console.log("判断属性:");
console.log("in:", "name" in obj);                   // true
console.log("hasOwnProperty:", obj.hasOwnProperty("age")); // true
console.log("hasOwnProperty (不存在的属性):", obj.hasOwnProperty("test")); // false

// 获取属性描述符
console.log("\n属性描述符:");
const descriptor = Object.getOwnPropertyDescriptor(obj, "name");
console.log("name属性描述符:", descriptor);
// { value: 'Alice', writable: true, enumerable: true, configurable: true }

// 定义属性
console.log("\n定义属性:");
Object.defineProperty(obj, "gender", {
    value: "female",
    writable: true,
    enumerable: true,
    configurable: true
});
console.log("添加gender后:", obj);                    // { name: 'Alice', age: 25, city: 'Beijing', gender: 'female' }

// 定义多个属性
Object.defineProperties(obj, {
    phone: { value: "123456", writable: true },
    email: { value: "alice@example.com", writable: true }
});
console.log("defineProperties:", obj);

// 添加修改属性
console.log("\n添加修改属性:");
obj.address = "朝阳区";                              // 添加新属性
obj.age = 26;                                         // 修改属性
console.log("添加修改后:", obj);

// 获取属性
console.log("\n获取属性:");
console.log("keys:", Object.keys(obj));               // 所有可枚举属性名
console.log("values:", Object.values(obj));           // 所有属性值
console.log("entries:", Object.entries(obj));         // [[key, value], ...]

// 冻结与密封
console.log("\n冻结与密封:");
const frozen = Object.freeze({ x: 1 });
// frozen.x = 2; // 无效（严格模式下会报错）
console.log("freeze后修改x:", frozen.x);             // 1
console.log("isFrozen:", Object.isFrozen(frozen));    // true

const sealed = Object.seal({ y: 2 });
// sealed.y = 3; // 可以修改
// sealed.z = 4; // 无效
console.log("seal后修改y:", sealed.y);                // 2
console.log("isSealed:", Object.isSealed(sealed));   // true

// 遍历
console.log("\n遍历:");
for (let key in obj) {
    if (obj.hasOwnProperty(key)) {
        console.log(`${key}: ${obj[key]}`);
    }
}

// 拷贝
console.log("\n拷贝:");
const shallowCopy = Object.assign({}, obj);
console.log("浅拷贝:", shallowCopy);
const deepClone = JSON.parse(JSON.stringify(obj));
console.log("深拷贝:", deepClone);

// 创建对象
console.log("\n创建对象:");
const newObj = Object.create(obj);
console.log("create (原型为obj):", newObj);
console.log("create原型:", Object.getPrototypeOf(newObj));

// 比较
console.log("\n比较:");
console.log("is (===行为):", Object.is(1, 1));        // true
console.log("is (NaN):", Object.is(NaN, NaN));        // true
console.log("is (0, -0):", Object.is(0, -0));         // false

// 其他方法
console.log("\n其他:");
console.log("getOwnPropertyNames:", Object.getOwnPropertyNames(obj)); // 包括不可枚举
console.log("getPrototypeOf:", Object.getPrototypeOf(obj));
console.log("setPrototypeOf:", Object.setPrototypeOf(obj, null));

// 解构赋值
const { name, age, city = "默认城市" } = obj;
console.log("\n解构赋值:", name, age, city);

// 展开运算符
const cloned = { ...obj };
console.log("展开运算符拷贝:", cloned);

const merged = { ...obj, country: "China" };
console.log("展开运算符合并:", merged);

console.log("Object示例完成");
