<template>
  <div style="margin: 20px;">
<!--    <h3 style="text-align: center">多大文件上传</h3>-->
    <el-upload
        drag
        :http-request="uploadFile"
        multiple
        ref="uploadMul"
        :auto-upload="false"
        :on-success="UploadSuccess"
        :on-error="UploadError"
        @change="handleFileChange"
    >
      <el-icon class="el-icon--upload"><upload-filled /></el-icon>
      <div class="el-upload__text">
        文件拖拽到此处 <em>点击上传</em>
      </div>
      <template #tip>
        <div class="el-upload__tip">
          files with a size less than 500mb
        </div>
      </template>
    </el-upload>
    <el-progress v-if="uploadProgress!==0"
                 style="margin: 30px;"
                 :text-inside="false"
                 :stroke-width="22"
                 :percentage="uploadProgress"
    />
    <el-button class="ml-3" type="success" @click="submitMulUpload" v-if="hasFile"
    style="width: 100%">
      上传到服务器
    </el-button>
  </div>


<!--  <div style="margin: 20px;">-->
<!--    <h3 style="text-align: center">单文件上传</h3>-->
<!--    <el-upload-->
<!--        ref="upload"-->
<!--        :http-request="uploadFile"-->
<!--        :limit="1"-->
<!--        :on-exceed="handleExceed"-->
<!--        :auto-upload="false"-->
<!--        :on-success="UploadSuccess"-->
<!--        :on-error="UploadError"-->
<!--    >-->
<!--      <template #trigger>-->
<!--        <el-button type="primary">select file</el-button>-->
<!--      </template>-->
<!--      <template #tip>-->
<!--        <div class="el-upload__tip text-red">-->
<!--          limit 1 file, new file will cover the old file-->
<!--        </div>-->
<!--      </template>-->
<!--    </el-upload>-->
<!--    <el-button class="ml-3" type="success" @click="submitUpload">-->
<!--      upload to server-->
<!--    </el-button>-->
<!--  </div>-->
</template>

<script setup lang="ts">
import { ref } from 'vue'
import {ElMessage, genFileId} from 'element-plus'
import type { UploadInstance, UploadProps, UploadRawFile } from 'element-plus'

// const upload = ref<UploadInstance>()
const uploadMul = ref<UploadInstance>()

// const handleExceed = (files) => {
//   upload.value!.clearFiles()
//   const file = files[0] as UploadRawFile
//   file.uid = genFileId()
//   upload.value!.handleStart(file)
// }

const submitMulUpload = () => {
  uploadMul.value!.submit()
}

// const submitUpload = () => {
//   upload.value!.submit()
// }


import {requestPack} from "../../utils/requests.js";
import {UploadFilled} from "@element-plus/icons-vue";
requestPack.get("refresh")

let uploadProgress = ref(0)
const CHUNK_SIZE=1024*512;
// const uploadFile = async (file)=> {
//   const chunkList = []; // 文件分片列表
//   let start = 0; // 分片开始索引
//   while (start < file.file.size) {
//     chunkList.push({ index: start, file: file.file.slice(start, start + CHUNK_SIZE) });
//     start += CHUNK_SIZE;
//   }
//   const timestamp = Date.now();
//   // 上传每个分片
//   for (let chunk of chunkList) {
//     uploadProgress.value=parseInt((chunk.index/file.file.size*100).toString())
//     const formData = new FormData();
//     formData.append('file', chunk.file);
//     formData.append('filename', `${timestamp}_${file.file.name}`);
//     formData.append('index', chunk.index.toString());
//     let res = await requestPack.post("/files/big_upload",formData,{
//       'Content-Type': 'multipart/form-data',
//     })
//   }
//   uploadProgress.value=0
// }
import axios from "axios";
const uploadFile = async (file)=> {
  await requestPack.get("refresh");
  let token = localStorage.getItem('token');
  const chunkList = []; // 文件分片列表
  let start = 0; // 分片开始索引
  while (start < file.file.size) {
    chunkList.push({ index: start, file: file.file.slice(start, start + CHUNK_SIZE) });
    start += CHUNK_SIZE;
  }
  const timestamp = Date.now();
  // 上传每个分片
  for (let chunk of chunkList) {
    uploadProgress.value=Number((chunk.index / file.file.size * 100).toFixed(2))
    const filename = `${timestamp}_${file.file.name}`;
    let filenameEncoded = encodeURIComponent(filename); //非ascii需要编码
    const index = chunk.index.toString();
    let res = await axios.post("api/files/big_upload", chunk.file, {
      headers: {
        'Content-Type': 'application/octet-stream',
        'filename': filenameEncoded,
        'index': index,
        "token":token
      },
      timeout: 120000
    });
  }
  uploadProgress.value=0
}


const UploadSuccess = (res, file, fileList)=>{
  ElMessage.success("上传成功")
}

const UploadError= (res, file, fileList)=>{
  ElMessage.error("上传失败")
}

const hasFile = ref(false)
const handleFileChange = (file, fileList) => {
  const notUploaded = fileList.filter(file => file.status !== 'success').length
  hasFile.value = notUploaded > 0
}

</script>
