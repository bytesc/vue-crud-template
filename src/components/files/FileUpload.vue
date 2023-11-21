<template>
  <div style="margin: 20px;">
    <h3 style="text-align: center">多文件上传</h3>
    <el-upload
        drag
        :http-request="uploadFile"
        multiple
        ref="uploadMul"
        :on-exceed="handleExceed"
        :auto-upload="false"
        :before-upload="checkBeforeUpload"
        :on-success="UploadSuccess"
        :on-error="UploadError"
    >
      <el-icon class="el-icon--upload"><upload-filled /></el-icon>
      <div class="el-upload__text">
        文件拖拽到此处 <em>点击上传</em>
      </div>
      <template #tip>
        <div class="el-upload__tip">
          files with a size less than 500kb
        </div>
      </template>
    </el-upload>
    <el-button class="ml-3" type="success" @click="submitMulUpload">
      upload to server
    </el-button>
  </div>

  <div style="margin: 20px;">
    <h3 style="text-align: center">单文件上传</h3>
    <el-upload
        ref="uploadSingle"
        :http-request="uploadFile"
        :limit="1"
        :on-exceed="handleExceed"
        :before-upload="checkBeforeUpload"
        :auto-upload="false"
    >
      <template #trigger>
        <el-button type="primary">select file</el-button>
      </template>
      <template #tip>
        <div class="el-upload__tip text-red">
          limit 1 file, new file will cover the old file
        </div>
      </template>
    </el-upload>
    <el-button class="ml-3" type="success" @click="submitSingleUpload">
      upload to server
    </el-button>
  </div>

</template>

<script lang="ts" setup>

import {UploadFilled} from "@element-plus/icons-vue";
import {ElMessage} from "element-plus";
import {ref} from "vue";

import { genFileId } from 'element-plus'
import type { UploadInstance, UploadProps, UploadRawFile } from 'element-plus'

const uploadMul = ref<UploadInstance>()
const uploadSingle = ref<UploadInstance>()

const handleExceed = (files) => {
  uploadSingle.value!.clearFiles()
  const file = files[0] as UploadRawFile
  file.uid = genFileId()
  uploadSingle.value!.handleStart(file)
}

const submitSingleUpload = () => {
  uploadSingle.value!.submit()
}

const submitMulUpload = () => {
  uploadMul.value!.submit()
}

const checkBeforeUpload =  (rawFile) => {
  if(rawFile.size > 500 * 1024 ) {
    ElMessage.error('文件大小不能超过500KB');
    return false
  }
  return true
}

import {requestPack} from "../../utils/requests.js";
requestPack.get("refresh")
const uploadFile = async (file)=> {
  const formData = new FormData();
  formData.append('file', file.file);
  let res = await requestPack.post("/files/upload",formData,{
    'Content-Type': 'multipart/form-data',
  })
}

const UploadSuccess = (res, file, fileList)=>{
  ElMessage.success("上传成功")
}

const UploadError= (res, file, fileList)=>{
  ElMessage.error("上传失败")
}
</script>
