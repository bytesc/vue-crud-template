<template>
  <el-row style="justify-content: flex-end;">
    <a href="#/files/big_upload">
      <el-button type="primary" style="margin-right: 20px"><el-icon><Upload /></el-icon>上传</el-button>
    </a>
  </el-row>
  <router-view></router-view>
  <div style="margin: 30px">
    <el-row :gutter="12">
      <el-col :xs="6" :sm="6" :md="4" :lg="4" :xl="4"
          v-for="(file, index) in files"
          :key="index"
      >
        <el-card shadow="hover" style="margin-bottom: 10px">
          <div>
            <div style="height: 100px; overflow: hidden; text-overflow: ellipsis; word-wrap: break-word">
              {{ file.Name }}
            </div>
            <time class="time">{{ file.Time }}</time>
            <el-row>
              <el-button text class="button" style="color: #409EFF;"><el-icon><Download /></el-icon></el-button>
              <el-button text class="button" style="color: #F56C6C;"><el-icon><DeleteFilled /></el-icon></el-button>
            </el-row>
          </div>
        </el-card>
      </el-col>
    </el-row>

  </div>
</template>

<script lang="ts" setup>
import { ref ,onMounted} from 'vue'

const files = ref([])

import {request} from "../../utils/requests.js";
import {DeleteFilled, Download, Upload} from "@element-plus/icons-vue";

onMounted(async () => {
  files.value = await request.get("/files/list")
  console.log(files.value)
})

</script>

<style>
.time {
  font-size: 12px;
  color: #999;
}

.bottom {
  margin-top: 13px;
  line-height: 12px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.button {
  padding: 0;
  min-height: auto;
}

.image {
  width: 100%;
  display: block;
}
</style>
