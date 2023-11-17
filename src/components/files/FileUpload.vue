<template>
  <el-upload
      class="upload-demo"
      drag
      :http-request="uploadFile"
      multiple
  >
    <el-icon class="el-icon--upload"><upload-filled /></el-icon>
    <div class="el-upload__text">
      Drop file here or <em>click to upload</em>
    </div>
    <template #tip>
      <div class="el-upload__tip">
         files with a size less than 500kb
      </div>
    </template>
  </el-upload>
</template>

<script setup>
import {request} from "../../utils/requests.js";
import {UploadFilled} from "@element-plus/icons-vue";

const uploadFile = async (file)=> {
  const formData = new FormData();
  formData.append('file', file.file);
  request.post("/files/upload",formData,{
    'Content-Type': 'multipart/form-data',
  })
}
</script>
