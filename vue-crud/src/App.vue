<script setup lang="ts">

//app.vue是整个项目的基底，这个文件里主要负责总布局

//三个边栏组件
import SideNav from "./components/nav/SideNav.vue"
import TopNav from "./components/nav/TopNav.vue";
import Footer from "./components/nav/Footer.vue";
import Bread from "./components/nav/Bread.vue";

//声明弹窗
import Statement from "./components/form/Statement.vue";
</script>

<template>
  <Statement/>
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

        <el-container style="min-height: 450px">
          <el-main>
            <Bread></Bread>
            <router-view></router-view>  <!--这里是路由区域，点击连接会在这里显示子页面-->
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
