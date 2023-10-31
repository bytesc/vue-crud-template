<template>
  <div style="text-align: center" v-show="!hasChildRoute">
<!--    只在没有子路由的时候显示-->
    这是信息一  <a :href=url>详情</a>
  </div>
  <router-view></router-view>
</template>

<script setup>

import {useRoute} from "vue-router";
import {ref, watch} from "vue";

let url = window.location.href + "/detail1"
let route = useRoute()
let hasChildRoute = ref(false)

//通过路由数量判断是否有子路由
watch(() => route.matched, () => {
  hasChildRoute.value = route.matched.length > 2
  console.log(route.matched.length)
})

</script>

<style scoped>

</style>