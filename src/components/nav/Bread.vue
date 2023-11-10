<template>
  <el-breadcrumb separator="/">
    <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
    <el-breadcrumb-item v-for="item in breadcrumbList">
      <a :href="'#' + item.path" >{{ item.meta.title }}</a>
    </el-breadcrumb-item>
  </el-breadcrumb>
</template>


<!--<script lang="ts">-->

<!--import {onMounted, ref, watch} from 'vue';-->
<!--import { useRoute } from 'vue-router';-->
<!--const currentActive = ref()-->

<!--export default {-->

<!--  setup() {-->
<!--    const route = useRoute()-->

<!--    watch(-->
<!--        () => route.matched,-->
<!--        (newVal) => {-->
<!--          currentActive.value = newVal-->
<!--          console.log(newVal)-->
<!--        }-->
<!--    )-->
<!--  },-->

<!--  data(){-->
<!--    return{-->
<!--      breadcrumbList:currentActive-->
<!--    }-->
<!--  },-->

<!--};-->

<!--</script>-->

<script lang="ts" setup>

import {onMounted, ref, watch, computed} from 'vue';
import { useRoute } from 'vue-router';
const currentActive = ref()
const route = useRoute()

watch(
        () => route.matched,
        (newVal) => {
          currentActive.value = newVal
          // console.log(newVal)
        }
    )

let breadcrumbList = computed(() => currentActive.value)

</script>

<style scoped>

</style>