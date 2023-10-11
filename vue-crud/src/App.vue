<script setup lang="ts">

import SideNav from "./components/nav/SideNav.vue"
import TopNav from "./components/nav/TopNav.vue";
import Footer from "./components/nav/Footer.vue";
</script>

<template>
  <div class="common-layout">
    <el-container>
      <el-header style="padding: 0">
        <TopNav></TopNav>
      </el-header>
      <el-container>
        <div style="width: 60px">
          <el-aside>
            <SideNav></SideNav>
          </el-aside>
        </div>

        <el-container>
          <el-main>
            <router-view></router-view>
          </el-main>

        </el-container>

      </el-container>
      <el-footer style="padding: 0">
        <Footer></Footer>
      </el-footer>
    </el-container>
  </div>
</template>

<style scoped>

</style>
