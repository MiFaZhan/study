const proto = {
  baseMethod() { return "from proto"; }
};

const key = "dynamic";
let counter = 0;

const handler = function() {
  return "I'm a handler";
};

const obj = {
  __proto__: proto,
  handler,                          // 简写 等同于 handler: handler
  counter,                          // 简写

  greet(name) {                     // 方法简写 等同于 greet: function(name) { return ... }
    return `Hello, ${name}!`;
  },

  // super 调用
  baseMethod() {
    return "override + " + super.baseMethod();
  },

  // 计算属性名
  [`${key}_prop`]: 42,
  [`id_${++counter}`]: "auto-1",
};

console.log(obj.greet("MiMo"));       // "Hello, MiMo!"
console.log(obj.baseMethod());        // "override + from proto"
console.log(obj.dynamic_prop);        // 42
console.log(obj.id_1);                // "auto-1"
