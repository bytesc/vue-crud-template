<template>
  <el-row style="justify-content: flex-end;">
    <el-button type="primary" style="margin-right: 20px" @click="dialogVisible=true"
    ><el-icon><Upload /></el-icon>上传</el-button>
  </el-row>

  <router-view></router-view>

  <el-dialog v-model="dialogVisible" title="Tips" width="70%" draggable
             :before-close="handleCloseUploadDialog">
    <template #header="{ close, titleId, titleClass }">
      <div class="my-header">
        <el-button type="danger" @click="close">
          <el-icon class="el-icon--left"><CircleCloseFilled /></el-icon>
          关闭
        </el-button>
      </div>
    </template>
    <BigFileUpload></BigFileUpload>
  </el-dialog>

  <el-empty v-if="files === null"
            description="没有文件"></el-empty>
  <div style="margin: 30px">
    <el-row :gutter="12">
      <el-col :xs="6" :sm="6" :md="4" :lg="4" :xl="4"
          v-for="(file, index) in files"
          :key="index"
      >
        <el-card shadow="hover" style="margin-bottom: 10px">
          <div>
            <div style="height: 100px; overflow: hidden; text-overflow: ellipsis; word-wrap: break-word">
              {{ file.name }}
            </div>
            <time class="time">{{ file.time }}</time>
            <el-row>
              <el-button text class="button" style="color: #409EFF;"
                @click="handleDownload(file.raw_name)"><el-icon><Download /></el-icon></el-button>
              <el-button text class="button" style="color: #F56C6C;"
                         @click="handleDel(file.raw_name)"><el-icon><DeleteFilled /></el-icon></el-button>
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

const dialogVisible = ref(false)

import {request} from "../../utils/requests.js";
import {CircleCloseFilled, DeleteFilled, Download, Upload} from "@element-plus/icons-vue";
import BigFileUpload from "./BigFileUpload.vue";

const handleCloseUploadDialog = async ()=>{
  dialogVisible.value=false
  files.value = await request.get("/files/list")
}

const handleDel = async (filename)=>{
  let res = await request.post(`/files/delete/${filename}`)
  files.value = await request.get("/files/list")
}

const handleDownload  = async (filename)=> {
  // 从localStorage或其他地方获取token
  let token = localStorage.getItem('token');
  // 构造请求头
  let headers = new Headers();
  headers.append('token', token);
  // 发起fetch请求
  await fetch('api/files/download/' + filename, { headers: headers })
      .then(response => response.blob())
      .then(blob => {
        // 创建隐藏的可下载链接
        let url = URL.createObjectURL(blob);
        let link = document.createElement('a');
        link.href = url;
        link.setAttribute('download', filename);
        // 自动在浏览器中点击链接
        document.body.appendChild(link);
        link.click();
        link.remove();
      })
      .catch(error => console.error('Error:', error));
}

onMounted(async () => {
  files.value = await request.get("/files/list")
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
