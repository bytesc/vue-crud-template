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
          <el-form-item label="密码" :label-width="formLabelWidth">
            <el-input v-model="form.password" autocomplete="off"
                      type="password" :show-password="true"
                      placeholder="输入密码"/>
          </el-form-item>
        </el-form>
      </el-col>
    </el-row>

    <template #footer>
      <span class="dialog-footer">
        <a href="#/user/signup">
          <el-button>注册</el-button>
        </a>
        <el-button type="primary" @click="HandleLogin">
          登录
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

const form = ref({
  username: '',
  password:""
})

import {rsaEncrypt} from "../../utils/rsa.js";
const HandleLogin = async ()=>{
  let res = request.post("/user/login",{
    "name":form.value.username,
    "password":await rsaEncrypt(form.value.password)
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