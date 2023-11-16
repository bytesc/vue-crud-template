<template>
  <!--  <el-button text @click="dialogFormVisible = true">-->
  <!--    Click to open Dialog-->
  <!--  </el-button>-->
  <el-dialog v-model="dialogFormVisible" title="用户注册"
             :fullscreen="true" :center="true" :show-close="false" :align-center="true"
  >
    <el-row :justify="'center'">
      <el-col :span="12">
        <el-form :model="form" >
          <el-form-item label="用户名" :label-width="formLabelWidth">
            <el-input v-model="form.username"/>
          </el-form-item>
          <el-form-item label="邮箱" :label-width="formLabelWidth ">
            <el-input v-model="form.email"/>
          </el-form-item>
          <el-form-item label="密码" :label-width="formLabelWidth" >
            <el-input v-model="form.password"
                      type="password" :show-password="true"
                      placeholder="Please input password"/>
          </el-form-item>
        </el-form>
      </el-col>
    </el-row>

    <template #footer>
      <span class="dialog-footer">
        <a href="#/user/login">
          <el-button>取消</el-button>
        </a>
        <el-button type="primary" @click="HandleSignup">
          注册
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script lang="ts" setup>
import {computed, onMounted, ref} from 'vue'

const dialogFormVisible = ref(true)
const formLabelWidth = '140px'

const form = ref({
  username: '',
  email:"",
  password:""
})

import {request} from "../../utils/requests.js";
const HandleSignup = async ()=>{
  let res = request.post("/user/signup",{
    "name":form.value.username,
    "password":form.value.password,
    "email":form.value.email
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