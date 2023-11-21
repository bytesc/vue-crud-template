<template>
  <el-menu
      :default-active="activeIndex"
      class="el-menu-demo"
      mode="horizontal"
      :ellipsis="false"
      @select="handleSelect"
      background-color="#545c64"
      text-color="#fff"
      active-text-color="#ffd04b"
  >
    <el-menu-item index="0">LOGO</el-menu-item>
    <div class="flex-grow" />
    <el-menu-item index="1">Processing Center</el-menu-item>
    <el-sub-menu index="2">
      <template #title><el-icon><Tools /></el-icon>{{ username }}</template>
      <a  href="/#/user/login" v-if="username===''">
        <el-menu-item index="2-1"><el-icon><UserFilled /></el-icon>登录</el-menu-item>
      </a>
      <div v-else>
        <a @click="HandleLogout" >
          <el-menu-item index="2-1"><el-icon><UserFilled /></el-icon>注销</el-menu-item>
        </a>
        <a href="/#/user/change_pwd" >
          <el-menu-item index="2-2"><el-icon><UserFilled /></el-icon>修改密码</el-menu-item>
        </a>
      </div>
      <el-menu-item index="2-3">item three</el-menu-item>
      <el-sub-menu index="2-4">
        <template #title>item four</template>
        <el-menu-item index="2-4-1">item one</el-menu-item>
        <el-menu-item index="2-4-2">item two</el-menu-item>
        <el-menu-item index="2-4-3">item three</el-menu-item>
      </el-sub-menu>
    </el-sub-menu>
  </el-menu>
</template>

<script lang="ts" setup>
import {computed, ref} from 'vue'
import {Tools, UserFilled} from "@element-plus/icons-vue";

const activeIndex = ref('1')
const handleSelect = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}

import { useStore } from 'vuex'
const store = useStore()
let username = computed(() => store.getters.getName || '')

import {requestPack} from "../../utils/requests.js";
const HandleLogout = ()=>{
  let res = requestPack.get(`/user/logout/`,{
  })
}

</script>

<style>
.flex-grow {
  flex-grow: 1;
}
</style>