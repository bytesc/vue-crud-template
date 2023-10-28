<template>

  <!-- Form -->

  <el-dialog v-model="dialogFormVisible" title="声明" :show-close="true" :fullscreen="false">
    <span>This is a message</span>
    <template #footer>
      <span class="dialog-footer">
        <el-button type="primary" @click="dialogFormVisible = false">
          确认
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script lang="ts" setup>
import {onMounted, reactive, ref} from 'vue'

import {Plus} from "@element-plus/icons-vue";

import { defineProps } from 'vue';
const props = defineProps({
  dialogType:String,
})

const dialogFormVisible = ref(false)

dialogFormVisible.value=true


const formLabelWidth = '100px'

const form = reactive({
  name: '',
  pwd:"",
})
</script>
<style scoped>
.dialog-footer button:first-child {
  margin-right: 10px;
}
</style>
