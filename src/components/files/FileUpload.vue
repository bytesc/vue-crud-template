<template>
  <div style="margin: 20px;">
    <el-upload
        class="upload-demo"
        drag
        :http-request="uploadFile"
        multiple
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
  </div>

</template>

<script setup>
import {request} from "../../utils/requests.js";
import {UploadFilled} from "@element-plus/icons-vue";
import {ElMessage} from "element-plus";

const uploadFile = async (file)=> {
  if(file.file.size > 1024 * 1024) { // 文件大小超过1mb
    ElMessage.error('文件大小不能超过1mb');
    return;
  }
  const formData = new FormData();
  formData.append('file', file.file);
  request.post("/files/upload",formData,{
    'Content-Type': 'multipart/form-data',
  })
}
</script>
