
<template>
  <div style="width: 100%">

  </div>

  <div>
    <el-menu
        default-active="2"
        class="el-menu-vertical-demo"
        :collapse="isCollapse"
        @open="handleOpen"
        @close="handleClose"
    >


      <el-sub-menu index="1">
        <template #title>
          <el-icon><icon-menu /></el-icon>
          <span>Menu</span>
        </template>
        <el-menu-item-group>
          <template #title><span>Group One</span></template>
          <a href="#/page1"><el-menu-item index="1-1">页面1</el-menu-item></a>
<!-- 面链接需要加#，子页面会显示在app.vue的<router-view></router-view>区域-->
          <a href="#/page2"><el-menu-item index="1-2">页面2</el-menu-item></a>
        </el-menu-item-group>
        <el-menu-item-group title="Group Two">
          <a href="#/files/"><el-menu-item index="2-1">文件列表</el-menu-item></a>
          <el-menu-item index="2-2">item three</el-menu-item>
        </el-menu-item-group>
        <el-sub-menu index="2-3">
          <template #title><span>item four</span></template>
          <el-menu-item index="2-3-1">item one</el-menu-item>
          <el-menu-item index="2-3-2">item one</el-menu-item>
        </el-sub-menu>
      </el-sub-menu>

      <a href="#/form">
        <el-menu-item index="2">
          <el-icon><document /></el-icon>
          <template #title>表单</template>
        </el-menu-item>
      </a>

      <el-menu-item index="3" disabled>
        <el-icon><document /></el-icon>
        <template #title>Navigator Three</template>
      </el-menu-item>

      <el-menu-item index="4">
        <el-icon><location /></el-icon>
        <template #title>Navigator Four</template>
      </el-menu-item>

      <el-menu-item index="5">
        <el-icon><setting /></el-icon>
        <template #title>Navigator Four</template>
      </el-menu-item>
    </el-menu>
  </div>

</template>


<script lang="ts" setup>
import { ref } from 'vue'
import {
  Document,
  Menu as IconMenu,
  Location,
  Setting,
} from '@element-plus/icons-vue'

const isCollapse = ref(true)
const handleOpen = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}
const handleClose = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}

const activeIndex = ref('1')
const activeIndex2 = ref('1')
const handleSelect = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}
</script>

<style scoped>
.el-menu-vertical-demo:not(.el-menu--collapse) {
  width: 200px;
  min-height: 400px;
}
</style>