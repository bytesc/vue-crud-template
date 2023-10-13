

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
      <div style="margin: 5px"><Dialog dialogType="add"  @MsgToFather="childEven"/></div>
<!--      子父组件传数据，父端-->
    </div>

    <el-table :data="TableData"
              ref="multipleTableRef"
              style="width: 100%"
              @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" />
<!--      <el-table-column fixed  prop="name" label="Name" width="120" />-->
      <el-table-column prop="name" label="Name" width="120" />
      <el-table-column prop="id" label="id" width="120" />
      <el-table-column prop="level" label="Level" width="150" />
      <el-table-column prop="email" label="email" width="160" />
      <el-table-column prop="phone" label="phone" width="160" />
      <el-table-column prop="birthday" label="Birthday" width="150" />
      <el-table-column fixed="right" label="Operations" width="120">
        <template #default>
          <el-button link type="primary" size="small" @click="handleRowClick"
          >Detail</el-button
          >
          <Dialog dialogType="edit"></Dialog>
        </template>
      </el-table-column>
    </el-table>

    <div>

    </div>
  </div>
</template>

<script setup lang="ts">
import Dialog from "./dialog.vue"
import {ref, reactive} from "vue";
import {valueEquals} from "element-plus";
import {Search} from "@element-plus/icons-vue";
// 数据
let QueryInput = ref("")
// let TableData= ref("")
const multipleSelection = ref([])

const handleSelectionChange = (val) => {
  multipleSelection.value = val
  console.log(val)
}

let handleRowClick = () => {
  console.log('click')
}

// 父传子，父端
const childEven=(val)=>{
  console.log(val);
  TableData.push({
    id: (TableData.length+1).toString(),
    ...val
  })
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
