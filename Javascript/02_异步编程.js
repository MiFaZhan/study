// 异步编程：回调函数、Promise、async/await

// 1. 回调函数
function fetchData(callback) {
    setTimeout(() => {
        const data = { id: 1, name: "Alice" };
        callback(null, data);
    }, 1000);
}

fetchData((err, data) => {
    if (err) {
        console.log("错误:", err);
    } else {
        console.log("回调获取数据:", data);
    }
});

// 2. Promise
function promiseFetchData() {
    return new Promise((resolve, reject) => {
        setTimeout(() => {
            const success = true;
            if (success) {
                resolve({ id: 2, name: "Bob" });
            } else {
                reject(new Error("获取数据失败"));
            }
        }, 1000);
    });
}

promiseFetchData()
    .then(data => console.log("Promise成功:", data))
    .catch(err => console.log("Promise错误:", err.message));

// 3. Promise链式调用
function delay(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

delay(500)
    .then(() => { console.log("步骤1完成"); return delay(500); })
    .then(() => { console.log("步骤2完成"); return delay(500); })
    .then(() => console.log("步骤3完成"));

// 4. async/await
async function asyncFetchData() {
    try {
        const data = await promiseFetchData();
        console.log("async/await获取数据:", data);
        return data;
    } catch (err) {
        console.log("async/await错误:", err.message);
    }
}

asyncFetchData();

// 5. 并行执行 Promise.all
async function parallelFetch() {
    const promises = [
        Promise.resolve(1),
        Promise.resolve(2),
        Promise.resolve(3)
    ];
    
    const results = await Promise.all(promises);
    console.log("Promise.all结果:", results);
}

parallelFetch();

// 6. Promise.race（竞态）
Promise.race([
    new Promise(resolve => setTimeout(() => resolve("快速"), 100)),
    new Promise(resolve => setTimeout(() => resolve("慢速"), 500))
]).then(result => console.log("Promise.race获胜:", result));

console.log("异步代码开始执行...");
