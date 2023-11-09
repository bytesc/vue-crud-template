

<template>

    <div class="table-box" style="margin: 5px auto;">
      <div class="title" style="text-align: center">
        <h2>crud demo</h2>
      </div>

    <div class="query-box" style=" margin: 5px auto; ">
      <el-row :gutter="10">
        <el-col :xs="24" :sm="16" :md="16" :lg="20" :xl="20">
          <el-row justify="start">
            <el-input v-model="QueryInput"  style="margin: 5px; width: auto"
                      type="date" v-if="searchModeValue==='birthday'"/>
            <el-input v-model="QueryInput" placeholder="请输入搜索内容🔍"  style="margin: 5px; width: auto"
                      v-else/>
            <el-select v-model="searchModeValue" placeholder="选择搜索模式" style="margin: 5px; width: auto">
              <el-option
                  v-for="item in searchModeOptions"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
                  :disabled="item.disabled"
              />
            </el-select>
            <!--        <el-button type="primary" style="margin: 5px;" @click="HandleQuery"><el-icon><Search /></el-icon></el-button>-->
          </el-row>
        </el-col>
        <el-col :xs="24" :sm="8" :md="8" :lg="4" :xl="4">
          <el-row justify="end">
            <el-button type="danger" @click="handleListDelete" style="margin: 5px" v-if="multipleSelection.length>0">
              <el-icon><Delete /></el-icon> 删除
            </el-button>
            <div style="margin: 5px">
              <Dialog dialogType="add"  @MsgToFather="childEven"/>
            </div>
            <!--      子父组件传数据，父端-->
          </el-row>
       </el-col>
      </el-row>

    </div>

    <el-table :data="TableData"
              ref="multipleTableRef"
              style="width: 100%"
              @selection-change="handleSelectionChange"
              :row-class-name="tableRowClassName"
              :max-height="800">
      <el-table-column type="selection" width="55" />
<!--      <el-table-column fixed  prop="name" label="Name" width="120" />-->
      <el-table-column prop="name" label="Name" width="120"  sortable/>
      <el-table-column prop="ID" label="ID" width="120"  sortable/>
      <el-table-column prop="level" label="Level" width="120"  sortable/>
      <el-table-column prop="email" label="email" width="240"  sortable/>
      <el-table-column prop="phone" label="phone" width="160"  sortable/>
      <el-table-column prop="birthday" label="Birthday" width="120"  sortable/>
      <el-table-column fixed="right" label="Operations" width="120">
        <template #default="scope">
<!-- 这里必须是#default="scope"，表示在子组件el-table-column中的插槽slot-->
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
<!--      <el-pagination background layout="prev, pager, next"-->
<!--                     style="display: flex;justify-content: center; margin: 10px"-->
<!--                     :total="total"-->
<!--                     :page-size="pageSize"-->
<!--                     :current-page="curPage"-->
<!--                     @current-change="handleChangePage"-->
<!--      />-->
      <el-pagination
          background layout="total, sizes, prev, pager, next, jumper"
          style="display: flex;justify-content: center; margin: 10px"
          v-model:current-page="curPage"
          v-model:page-size="pageSize"
          :page-sizes="[100, 200, 300, 400,500, 10000]"
          :total="total"
          @size-change="handleSizeChange"
          @current-change="handleChangePage"
      />

  </div>
</template>

<script setup lang="ts">
import Dialog from "./dialog.vue" //对话框组件
import {computed, ref} from "vue";

import {Delete} from "@element-plus/icons-vue"; //引入icon

// import { useRoute } from 'vue-router'
// const route = useRoute()
// const routerNow = computed(()=>{  //计算属性
//   return route
// })
// // {{ routerNow }}
// import { onMounted } from 'vue'
// onMounted(() => {
//   console.log(`the component is now mounted.`)
//   console.log(route.path) //获取路由
// })

let TableData = ref([
  {
    ID:"",
    name: '',
    phone:"",
    email:"",
    birthday: '',
    level: '',
  }
])

// 方法

let total = ref(1)
let curPage = ref(1)
let pageSize= ref(20)

import {request} from "../../utils/requests.js";
const getTableData = async (cur = 1)=>{
  // let res= request.get("/list",{
  //   pageSize:10,
  //   pageNum:cur
  // })
  let res= await request.get(`user/list/?pageSize=${pageSize.value}&pageNum=${cur}`)

  console.log(res)
  total.value = res.total

  TableData.value = res.list
  console.log(TableData.value)
}
getTableData()

const handleChangePage = (val)=>{
  getTableData(val)
}
const handleSizeChange = (val)=>{
  getTableData(curPage.value)
  pageSize.value = val
}


// 单行删除
const handleRowDelete = (id) =>{
  // console.log(row.id)
  let index = TableData.value.findIndex((item) => item.ID === id)
  // console.log(index)
  TableData.value.splice(index,1)
}

//多行选择
const multipleSelection = ref([])
const handleSelectionChange = (val) => {
  // console.log(val)
  // multipleSelection.value = val
  multipleSelection.value=[]
  val.forEach(item=>{
    multipleSelection.value.push(item.ID)
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
const childEven=async (val, dialogType) => {
  // console.log(val);
  if (dialogType === "add") {
    /* //这是前端操作
    let newRow = {...val.value}
    let newId = parseInt(
        TableData.value.reduce((max, current) => (
        parseInt(current.id) > parseInt(max.id) ? current : max)
        ).id)+1  // 找出现有最大id+1
    newRow.id = newId.toString()
    TableData.value.push(newRow)
    */
    console.log(val.value.name)
    let res = await request.post("/user/add/", {
      "name":val.value.name,
      "phone":val.value.phone,
      "email":val.value.email,
      "birthday":val.value.birthday,
      "level":val.value.level
    })
    await getTableData(curPage.value)
  }

  if (dialogType === "edit") {
    // TableData.forEach((item) => {
    //   if (item.id === val.id) {
    //     item.name = val.name;
    //     item.level = val.level;
    //     item.email = val.email;
    //     item.phone = val.phone;
    //     item.birthday = val.birthday;
    //   }
    // });
    let index = TableData.value.findIndex((item) => item.ID === val.value.ID)
    TableData.value[index] = val.value
  }
}

const searchModeValue = ref('id')
const searchModeOptions = [
  {
    value: 'id',
    label: '按 id',
  },
  {
    value: 'name',
    label: '按名字',
  },
  {
    value: 'phone_email',
    label: '按手机号邮箱号',
  },
  {
    value: 'level',
    label: '按用户等级',
    disabled: false
  },
  {
    value: 'birthday',
    label: '按生日',
    disabled: false
  }
]

let QueryInput = ref("")

/*控制行可见*/
const tableRowClassName = ({row, index})=>{
    if (searchModeValue.value==="phone_email"){
      if (row.phone.toLowerCase().match(QueryInput.value.toLowerCase()) || row.email.toLowerCase().match(QueryInput.value.toLowerCase())) {
        return '';
      }
      else{return 'hidden-row'}
    }
    if (searchModeValue.value==="id"){
      console.log(row)
      if (row.ID.toString().toLowerCase().match(QueryInput.value.toLowerCase())) {
        return '';
      }
      else{return 'hidden-row'}
    }
    if (searchModeValue.value==="name"){
      if (row.name.toLowerCase().match(QueryInput.value.toLowerCase())) {
        return '';
      }
      else{return 'hidden-row'}
    }
    if (searchModeValue.value==="level"){
      if (row.level.toLowerCase().match(QueryInput.value.toLowerCase())) {
        return '';
      }
      else{return 'hidden-row'}
    }
  if (searchModeValue.value==="birthday"){
    if (row.birthday.toLowerCase().match(QueryInput.value.toLowerCase())) {
      return '';
    }
    else{return 'hidden-row'}
  }
    return ''
}

    // reactive刷新方法
    // TableData.value.splice(0,TableData.value.length)
    // TableData.value.unshift(...newTableData)




// 深拷贝
// let copiedData = JSON.parse(JSON.stringify(TableData))
// let TableDataCopy = ref(copiedData)
//浅拷贝
// let TableDataCopy = Object.assign(TableData)



</script>

<!--<style scoped>-->
<!--不能scoped，否则tableRowClassName出错-->
<style>

/*控制行可见*/
.el-table .hidden-row {
  display: none;
}

</style>
