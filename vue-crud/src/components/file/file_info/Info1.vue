<template>
  <div style="text-align: center" v-show="!hasChildRoute">
<!--    只在没有子路由的时候显示-->
    这是信息一  <a :href=url_detail1>详情</a>
  </div>
<!--  向下一级传递本级url来实现返回功能-->
  <router-view :urlNow="urlNow"></router-view>
</template>

<script setup>

import {useRoute} from "vue-router";
import {ref, watch} from "vue";

let url_detail1 = window.location.href + "/detail1"
let urlNow = window.location.href
let route = useRoute()
let hasChildRoute = ref(false)

//通过路由数量route.matched.length判断是否有子路由
watch(() => route.matched, () => {
  hasChildRoute.value = route.matched.length > 2
  // console.log(route.path)
  // console.log(route.matched.length)
})

</script>

<style scoped>

</style>