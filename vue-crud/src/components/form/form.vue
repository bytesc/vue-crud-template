

<template>

    <div class="table-box" style="margin: 5px auto;">
      <div class="title" style="text-align: center">
        <h2>crud demo</h2>
      </div>

    <div class="query-box" style="display: flex;justify-content: space-between; margin: 5px auto;">

      <div style=" display: flex; float: left">
        <el-input v-model="QueryInput" placeholder="Please input" style="margin: 5px"/>
        <el-button type="primary" style="margin: 5px;"><el-icon><Search /></el-icon></el-button>
      </div>
      <div style=" display: flex; float: right">
        <el-button type="danger" @click="handleListDelete" style="margin: 5px" v-if="multipleSelection.length>0">
          <el-icon><Delete /></el-icon> 删除
        </el-button>
        <div style="margin: 5px">
          <Dialog dialogType="add"  @MsgToFather="childEven"/>
        </div>
        <!--      子父组件传数据，父端-->
      </div>
    </div>

    <el-table :data="TableData"
              ref="multipleTableRef"
              style="width: 100%"
              @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" />
<!--      <el-table-column fixed  prop="name" label="Name" width="120" />-->
      <el-table-column prop="name" label="Name" width="120" />
      <el-table-column prop="id" label="id" width="120" />
      <el-table-column prop="level" label="Level" width="80" />
      <el-table-column prop="email" label="email" width="240" />
      <el-table-column prop="phone" label="phone" width="160" />
      <el-table-column prop="birthday" label="Birthday" width="150" />
      <el-table-column fixed="right" label="Operations" width="120">
        <template #default="scope">
<!-- 这里必须是#default="scope"，表示在子组件el-table中的插槽slot-->
          <el-button link type="danger" size="small" @click="handleRowDelete(scope.row.id)">
            删除
          </el-button>
          <!--      子父组件传数据，父端-->
          <Dialog dialogType="edit"  :TableRow=scope.row @MsgToFather="childEven"></Dialog>
        </template>
      </el-table-column>
    </el-table>

    <div>

    </div>
  </div>
</template>

<script setup lang="ts">
import Dialog from "./dialog.vue" //对话框组件

import {ref, reactive} from "vue";

import {Delete, Search} from "@element-plus/icons-vue"; //引入icon

// 数据
let QueryInput = ref("")
// let TableData= ref("")

// 单行删除
const handleRowDelete = (id) =>{
  // console.log(row.id)
  let index = TableData.findIndex((item) => item.id === id)
  // console.log(index)
  TableData.splice(index,1)
}

//多行选择
const multipleSelection = ref([])
const handleSelectionChange = (val) => {
  // console.log(val)
  // multipleSelection.value = val
  multipleSelection.value=[]
  val.forEach(item=>{
    multipleSelection.value.push(item.id)
  })
  // console.log(multipleSelection.value)
}

//多行删除
const handleListDelete=()=>{
  multipleSelection.value.forEach((id)=>{
    // console.log(id)
    handleRowDelete(id)
  })
  multipleSelection.value=[]
}

// 父传子，父端
// 对话框子组件事件，编辑或删除
const childEven=(val, dialogType)=>{
  // console.log(val);
  if (dialogType==="add"){
    let newRow = {...val}
    val.id = (TableData.length+1).toString()
    TableData.push(newRow)
  }

  if (dialogType==="edit"){
    // TableData.forEach((item) => {
    //   if (item.id === val.id) {
    //     item.name = val.name;
    //     item.level = val.level;
    //     item.email = val.email;
    //     item.phone = val.phone;
    //     item.birthday = val.birthday;
    //   }
    // });
    let index = TableData.findIndex((item)=>item.id === val.id)
    TableData[index] = val
  }
}

// reactive才可刷新
let TableData = reactive([
  {
    id:"1",
    name: 'Tom',
    level:"1",
    email: '123456@xx.com',
    phone:"18812341234",
    birthday: "2023-10-12"
  },
  {
    id:"2",
    name: 'Tom',
    level:"1",
    email: '123456@xx.com',
    phone:"18812341234",
    birthday: "2023-10-12"
  },
  {
    id:"3",
    name: 'Tom',
    level:"1",
    email: '123456@xx.com',
    phone:"18812341234",
    birthday: "2023-10-12"
  },
  {
    id:"4",
    name: 'Tom',
    level:"1",
    email: '123456@xx.com',
    phone:"18812341234",
    birthday: "2023-10-12"
  },
])


// 方法

</script>

<style scoped>

</style>
