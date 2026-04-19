# CSS 渐变颜色文档

## 1. linear-gradient（线性渐变）

### 基本语法
```css
background: linear-gradient(方向/角度, 颜色1, 颜色2, ...);
```

### 角度说明
| 角度 | 效果 |
|------|------|
| `0deg` | 从下到上 |
| `90deg` | 从左到右 |
| `180deg` | 从上到下 |
| `135deg` | 左下到右上（对角线） |

### 示例
```css
/* 从上到下 */
background: linear-gradient(180deg, #667eea, #764ba2);

/* 从左到右 */
background: linear-gradient(90deg, red, blue);

/* 对角线 */
background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);

/* 多位置渐变 */
background: linear-gradient(90deg, red 0%, yellow 50%, blue 100%);
```

---

## 2. radial-gradient（径向渐变）

### 基本语法
```css
background: radial-gradient(形状 大小 at 位置, 颜色1, 颜色2);
```

### 形状
- `circle` - 圆形
- `ellipse` - 椭圆（默认）

### 位置
- `center`（默认）
- `top`, `bottom`, `left`, `right`
- 百分比或像素值

### 示例
```css
/* 圆形渐变 */
background: radial-gradient(circle, red, blue);

/* 椭圆渐变 */
background: radial-gradient(ellipse at center, red 0%, blue 100%);

/* 指定圆心位置 */
background: radial-gradient(circle at top left, red, blue);
```

---

## 3. conic-gradient（锥形渐变）

### 基本语法
```css
background: conic-gradient(从角度开始, 颜色1, 颜色2, ...);
```

### 示例
```css
/* 色相环 */
background: conic-gradient(red, yellow, green, blue, red);

/* 饼图效果 */
background: conic-gradient(#667eea 0% 25%, #764ba2 25% 50%, #f093fb 50% 75%, #f5576c 75% 100%);
```

---

## 4. repeating-linear-gradient（重复线性渐变）

### 示例
```css
/* 条纹图案 */
background: repeating-linear-gradient(
    45deg,
    #667eea,
    #667eea 10px,
    #764ba2 10px,
    #764ba2 20px
);
```

---

## 5. repeating-radial-gradient（重复径向渐变）

### 示例
```css
/* 圆环图案 */
background: repeating-radial-gradient(
    circle,
    #667eea,
    #667eea 10px,
    #764ba2 10px,
    #764ba2 20px
);
```

---

## 6. 常用工具

- [cssgradient.io](https://cssgradient.io) - 在线渐变生成器
- [uigradients.com](https://uigradients.com) - 预设渐变合集
- [hexcolors.co](https://hexcolors.co) - 渐变配色方案

---

## 7. 浏览器兼容前缀

```css
background: -webkit-linear-gradient(180deg, red, blue);  /* Safari */
background: -moz-linear-gradient(180deg, red, blue);     /* Firefox */
background: -o-linear-gradient(180deg, red, blue);       /* Opera */
background: linear-gradient(180deg, red, blue);           /* 标准语法 */
```

建议配合工具如 PostCSS 自动添加前缀。
