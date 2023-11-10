<template>
  <div style="text-align: center">
    这是文件一 <a href="#/file1/info1">信息</a>
  </div>
  <router-view></router-view>
</template>

<script lang="ts" setup>

import {onMounted, ref, watch, computed} from 'vue';
import { useRoute } from 'vue-router';
const currentActive = ref()
const route = useRoute()

watch(
    () => route.matched,
    (newVal) => {
      currentActive.value = newVal
    }
)

</script>

<style scoped>

</style>


