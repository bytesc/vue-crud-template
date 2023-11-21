<template>
  <el-row style="justify-content: flex-end;">
    <el-button type="primary" style="margin-right: 20px" @click="dialogUpload=true"
    ><el-icon><Upload /></el-icon>上传</el-button>
  </el-row>

  <router-view></router-view>

  <el-dialog v-model="dialogUpload" title="Tips" width="70%" draggable
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

  <el-dialog v-model="dialogDownload" title="Tips" width="70%" draggable
             :before-close="handleCloseUploadDialog">
    <template #header="{ close, titleId, titleClass }">
      下载中
    </template>
    <el-progress
                 style="margin: 30px;"
                 :text-inside="true"
                 :stroke-width="22"
                 :percentage="percentDownloadCompleted"
    />
  </el-dialog>

  <el-empty v-if="files === null || files === ''" description="没有文件"></el-empty>

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
            <br/>
            <span class="time">{{(file.size/1024).toFixed(2)}}KB</span>
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

const dialogUpload = ref(false)
const dialogDownload = ref(false)

import {requestPack} from "../../utils/requests.js";
import {CircleCloseFilled, DeleteFilled, Download, Upload} from "@element-plus/icons-vue";
import BigFileUpload from "./BigFileUpload.vue";
import {ElMessage} from "element-plus";

const handleCloseUploadDialog = async ()=>{
  dialogUpload.value=false
  dialogDownload.value=false
  files.value = await requestPack.get("/files/list")
}

const handleDel = async (filename)=>{
  let res = await requestPack.post(`/files/delete/${filename}`)
  files.value = await requestPack.get("/files/list")
}

import axios from "axios";
let percentDownloadCompleted = ref(0)
const handleDownload = async (filename) => {
  await requestPack.get("refresh");
  dialogDownload.value=true
  let token = localStorage.getItem('token');
  let headers = { 'token': token };
  axios({
    url: 'api/files/download/' + filename,
    method: 'GET',
    responseType: 'blob', // important
    headers: headers,
    timeout: 9999999999,
    onDownloadProgress: function(progressEvent) {
      percentDownloadCompleted.value = Math.round((progressEvent.loaded * 100) / progressEvent.total);
      console.log(dialogDownload.value)
    }
  }).then((response) => {
    const url = URL.createObjectURL(new Blob([response.data]));
    const link = document.createElement('a');
    link.href = url;
    link.setAttribute('download', filename);
    document.body.appendChild(link);
    link.click();
    link.remove();
    // dialogDownload.value=false
  }).catch(error => {
    console.error('Error:', error);
    ElMessage.error("下载失败");
  });
}

onMounted(async () => {
  files.value = await requestPack.get("/files/list")
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
