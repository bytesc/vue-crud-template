<template>
  <div style="text-align: center">
    <a :href=urlNow>返回</a>这是详情一
  </div>
</template>

<script setup>
//接收上一级url来实现返回功能
const props = defineProps({
  urlNow:String,
})

</script>

<style scoped>

</style>