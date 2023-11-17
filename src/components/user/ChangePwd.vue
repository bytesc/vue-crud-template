<template>
  <!--  <el-button text @click="dialogFormVisible = true">-->
  <!--    Click to open Dialog-->
  <!--  </el-button>-->
  <el-dialog v-model="dialogFormVisible" title="用户登录"
             :fullscreen="true" :center="true" :show-close="false" :align-center="true"
  >
    <el-row :justify="'center'">
      <el-col :span="12">
        <el-form :model="form" >
          <el-form-item label="用户名" :label-width="formLabelWidth">
            <el-input v-model="form.username" autocomplete="off"
                      placeholder="输入用户名"/>
          </el-form-item>
          <el-form-item label="旧密码" :label-width="formLabelWidth">
            <el-input v-model="form.old_password" autocomplete="off"
                      type="password" :show-password="true"
                      placeholder="旧密码"/>
          </el-form-item>
          <el-form-item label="新密码" :label-width="formLabelWidth">
            <el-input v-model="form.new_password" autocomplete="off"
                      type="password" :show-password="true"
                      placeholder="新密码"/>
          </el-form-item>
          <el-form-item label="重复新密码" :label-width="formLabelWidth">
            <el-input v-model="form.again_password" autocomplete="off"
                      type="password" :show-password="true"
                      placeholder="重复新密码"/>
          </el-form-item>
        </el-form>
      </el-col>
    </el-row>

    <template #footer>
      <span class="dialog-footer">
        <a href="#/">
          <el-button>取消</el-button>
        </a>
        <el-button type="primary" @click="HandleChangePwd">
          修改
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script lang="ts" setup>
import {ref} from 'vue'
import {request} from "../../utils/requests.js";

const dialogFormVisible = ref(true)
const formLabelWidth = '140px'

import { useStore } from 'vuex'
const store = useStore()
const form = ref({
  username: store.getters.getName || '',
  old_password:"",
  again_password:"",
  new_password:""
})

import {rsaEncrypt} from "../../utils/rsa.js";
import {ElMessage} from "element-plus";
const HandleChangePwd = async ()=>{
  if(form.value.new_password.length<8 || form.value.new_password.length>50){
    ElMessage.error("密码长度必须在 8 到 50 之间")
    return
  }
  if(form.value.new_password!=form.value.again_password){
    ElMessage.error("密码不一致")
    return
  }
  let res = await request.post("/user/change_pwd",{
    "name":form.value.username,
    "old_password":await rsaEncrypt(form.value.old_password),
    "new_password":await rsaEncrypt(form.value.new_password)
  })
  request.get(`/user/logout/`,{
  })
}

</script>
<style scoped>
.el-button--text {
  margin-right: 15px;
}
.el-select {
  width: 300px;
}
.el-input {
  width: 300px;
}
.dialog-footer button:first-child {
  margin-right: 10px;
}
</style>