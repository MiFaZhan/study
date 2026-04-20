// DOM 操作：原生API进行文档对象模型的操作

// 获取元素
const element = document.getElementById("app");
const elements = document.getElementsByClassName("container");
const firstElement = document.querySelector(".container");
const allElements = document.querySelectorAll("div");

// 创建元素
const newDiv = document.createElement("div");
newDiv.textContent = "新创建的div";
newDiv.className = "new-div";

// 添加到DOM
document.body.appendChild(newDiv);

// 插入到指定位置
const parent = document.getElementById("parent");
parent.insertBefore(newDiv, parent.firstChild);

// 移除元素
const toRemove = document.getElementById("to-remove");
if (toRemove && toRemove.parentNode) {
    toRemove.parentNode.removeChild(toRemove);
}
// ES6+
toRemove.remove();  

// 修改元素内容
element.textContent = "新文本内容";
element.innerHTML = "<span>HTML内容</span>";

// 修改元素属性
element.setAttribute("data-id", "123");
element.getAttribute("data-id");
element.removeAttribute("data-id");

// 修改元素样式
element.style.color = "red";
element.style.fontSize = "16px";

// 修改元素类名
element.classList.add("active");
element.classList.remove("hidden");
element.classList.toggle("expanded");
element.classList.contains("active");

// 获取/设置数据属性
element.dataset.userId = "456";
console.log(element.dataset.userId);

// 遍历元素
const childNodes = parent.childNodes; // 包含文本节点
const children = parent.children; // 仅元素节点
const firstChild = parent.firstChild;
const lastChild = parent.lastChild;
const nextSibling = element.nextSibling;
const previousSibling = element.previousSibling;
const parentNode = element.parentNode;

// 创建文本节点
const textNode = document.createTextNode("文本内容");
parent.appendChild(textNode);

// 克隆元素
const clone = element.cloneNode(true); // 深克隆
const shallowClone = element.cloneNode(false); // 浅克隆

// 文档片段（性能优化）
const fragment = document.createDocumentFragment();
for (let i = 0; i < 5; i++) {
    const item = document.createElement("div");
    item.textContent = `项 ${i}`;
    fragment.appendChild(item);
}
parent.appendChild(fragment);

console.log("DOM操作示例完成");
