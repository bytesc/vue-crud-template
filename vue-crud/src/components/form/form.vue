

<template>
    <div class="table-box" style="margin: 5px auto;">
      <div class="title" style="text-align: center">
        <h2>crud demo</h2>
      </div>

    <div class="query-box" style="display: flex;justify-content: space-between;margin: 5px auto;">
      <el-input v-model="QueryInput" placeholder="Please input" />
      <Dialog/>
    </div>

    <el-table :data="TableData"
              ref="multipleTableRef"
              style="width: 100%"
              @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" />
      <el-table-column fixed prop="date" label="Date" width="150" />
      <el-table-column prop="name" label="Name" width="120" />
      <el-table-column prop="state" label="State" width="120" />
      <el-table-column prop="city" label="City" width="120" />
      <el-table-column prop="address" label="Address" width="600" />
      <el-table-column prop="zip" label="Zip" width="120" />
      <el-table-column fixed="right" label="Operations" width="120">
        <template #default>
          <el-button link type="primary" size="small" @click="handleRowClick"
          >Detail</el-button
          >
          <el-button link type="primary" size="small">Edit</el-button>
        </template>
      </el-table-column>
    </el-table>

    <div>

    </div>
  </div>
</template>

<script setup lang="ts">
import Dialog from "./dialog.vue"
import {ref} from "vue";
import {valueEquals} from "element-plus";
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
const TableData = [
  {
    date: '2016-05-03',
    name: 'Tom',
    state: 'California',
    city: 'Los Angeles',
    address: 'No. 189, Grove St, Los Angeles',
    zip: 'CA 90036',
    tag: 'Home',
  },
  {
    date: '2016-05-02',
    name: 'Tom',
    state: 'California',
    city: 'Los Angeles',
    address: 'No. 189, Grove St, Los Angeles',
    zip: 'CA 90036',
    tag: 'Office',
  },
  {
    date: '2016-05-04',
    name: 'Tom',
    state: 'California',
    city: 'Los Angeles',
    address: 'No. 189, Grove St, Los Angeles',
    zip: 'CA 90036',
    tag: 'Home',
  },
  {
    date: '2016-05-01',
    name: 'Tom',
    state: 'California',
    city: 'Los Angeles',
    address: 'No. 189, Grove St, Los Angeles',
    zip: 'CA 90036',
    tag: 'Office',
  },
]


// 方法

</script>

<style scoped>

</style>
