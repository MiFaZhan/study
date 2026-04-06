<script setup>
import{ref,reactive} from 'vue'

// ref 用于基本类型
const message = ref('hello vue3')
const count = ref(0)

// reactive 用于对象类型
const user = reactive({
    name:'Vue',
    age:18
})

// 修改值
const updateMessage = () => {
    message.value = 'hello vue3'
    count.value++
}
</script>

<template>
    <div>
        <h1>{{message}}</h1>
        <h2>{{user.name}}</h2>
        <p>当前年龄: {{user.age}}</p>
        <p>当前计数: {{count}}</p>
        <button @click="updateMessage">更新</button>
    </div>
</template>
