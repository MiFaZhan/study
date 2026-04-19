// 事件处理：事件冒泡、捕获、添加和移除事件监听器

// 1. 基本事件监听
const button = document.getElementById("myButton");
button.addEventListener("click", function(event) {
    console.log("按钮被点击", event.type);
});

// 2. 事件对象常用属性和方法
button.addEventListener("click", function(event) {
    console.log("target:", event.target);           // 触发事件的目标元素
    console.log("currentTarget:", event.currentTarget); // 绑定事件的元素
    console.log("type:", event.type);               // 事件类型
    console.log("bubbles:", event.bubbles);         // 是否冒泡
});

// 3. 事件冒泡
// 内层元素先触发，然后向外传播
document.getElementById("inner").addEventListener("click", () => {
    console.log("inner click 冒泡阶段");
});
document.getElementById("outer").addEventListener("click", () => {
    console.log("outer click 冒泡阶段");
});

// 4. 事件捕获
// 外层元素先触发，然后向内传播
document.getElementById("inner").addEventListener("click", () => {
    console.log("inner click 捕获阶段");
}, true); // true 表示捕获阶段
document.getElementById("outer").addEventListener("click", () => {
    console.log("outer click 捕获阶段");
}, true);

// 5. 阻止冒泡和捕获
document.getElementById("stopBtn").addEventListener("click", function(event) {
    event.stopPropagation(); // 阻止事件继续传播
    console.log("事件已停止传播");
});

// 6. 阻止默认行为
document.getElementById("link").addEventListener("click", function(event) {
    event.preventDefault(); // 阻止默认行为（如跳转链接）
    console.log("默认行为被阻止");
});

// 7. 移除事件监听器
function handleClick() {
    console.log("被移除的监听器");
}
button.addEventListener("click", handleClick);
button.removeEventListener("click", handleClick);

// 8. 事件委托（利用冒泡）
document.getElementById("list").addEventListener("click", function(event) {
    if (event.target.tagName === "LI") {
        console.log("点击了:", event.target.textContent);
    }
});

// 9. 一次性事件
button.addEventListener("click", function handler() {
    console.log("只执行一次");
    button.removeEventListener("click", handler);
}, { once: true });

// 10. 常见事件类型
button.addEventListener("mouseenter", () => {});  // 鼠标进入
button.addEventListener("mouseleave", () => {}); // 鼠标离开
button.addEventListener("mousemove", () => {});  // 鼠标移动
button.addEventListener("mousedown", () => {});  // 鼠标按下
button.addEventListener("mouseup", () => {});    // 鼠标释放
button.addEventListener("dblclick", () => {});   // 双击

// 键盘事件
document.addEventListener("keydown", (e) => {
    console.log("按键:", e.key, "代码:", e.code);
});
document.addEventListener("keyup", () => {});

// 表单事件
const form = document.getElementById("myForm");
form.addEventListener("submit", (e) => {
    e.preventDefault();
    console.log("表单提交");
});
form.addEventListener("reset", () => {});

// 输入事件
const input = document.getElementById("myInput");
input.addEventListener("focus", () => console.log("获得焦点"));
input.addEventListener("blur", () => console.log("失去焦点"));
input.addEventListener("input", (e) => console.log("输入:", e.target.value));
input.addEventListener("change", () => console.log("值改变"));

// 窗口事件
window.addEventListener("resize", () => console.log("窗口调整大小"));
window.addEventListener("scroll", () => console.log("滚动"));
window.addEventListener("load", () => console.log("页面加载完成"));
document.addEventListener("DOMContentLoaded", () => console.log("DOM加载完成"));

console.log("事件处理示例完成");
