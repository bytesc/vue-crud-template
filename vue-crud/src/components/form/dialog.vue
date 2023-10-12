<template>

  <!-- Form -->
  <el-button type="primary" @click="handleAdd">
    <el-icon><Plus /></el-icon> 添加
  </el-button>

  <el-dialog v-model="dialogFormVisible" :title="props.dialogType === 'add' ? '新增':'编辑'">
    <el-form :model="form">
      <el-form-item label="name" :label-width="formLabelWidth">
        <el-input v-model="form.name" autocomplete="off" />
      </el-form-item>
      <el-form-item label="phone" :label-width="formLabelWidth">
        <el-input v-model="form.phone" autocomplete="off" />
      </el-form-item>
      <el-form-item label="email" :label-width="formLabelWidth">
        <el-input v-model="form.email" autocomplete="off" />
      </el-form-item>
      <el-form-item label="level" :label-width="formLabelWidth">
        <el-input v-model="form.level" autocomplete="off" />
      </el-form-item>
      <el-form-item label="birthdate" :label-width="formLabelWidth">
        <el-input v-model="form.birthday" autocomplete="off" type="date"/>
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取消</el-button>
        <el-button type="primary" @click="dialogFormVisible = false">
          确认
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script lang="ts" setup>
import { reactive, ref } from 'vue'

import {Plus} from "@element-plus/icons-vue";

import { defineProps } from 'vue';
const props = defineProps({
  dialogType:String,
})

const dialogFormVisible = ref(false)
const handleAdd = ()=>{
  dialogFormVisible.value=true //注意，这是要.value
  console.log(dialogFormVisible.value)
}

const formLabelWidth = '140px'

const form = reactive({
  name: '',
  phone:"",
  email:"",
  birthday: '',
  level: '',
})
</script>
<style scoped>
.dialog-footer button:first-child {
  margin-right: 10px;
}
</style>
