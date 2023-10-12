<template>
  <el-row :gutter="20">
    <el-col :span="16" class="foot-item"><div class="grid-content ep-bg-purple" ></div></el-col>
    <el-col :span="8" class="foot-item"><div class="grid-content ep-bg-purple"></div></el-col>
  </el-row>
  <el-row :gutter="20">
    <el-col :span="8" class="foot-item"><div class="grid-content ep-bg-purple" ></div></el-col>
    <el-col :span="8" class="foot-item"><div class="grid-content ep-bg-purple" ></div></el-col>
    <el-col :span="4" class="foot-item"><div class="grid-content ep-bg-purple" ></div></el-col>
    <el-col :span="4" class="foot-item"><div class="grid-content ep-bg-purple" ></div></el-col>
  </el-row>
  <el-row :gutter="20" >
      <el-col :span="4" class="foot-bottom"><div class="grid-content ep-bg-purple" ></div></el-col>
      <el-col :span="16" class="foot-bottom"><div class="grid-content ep-bg-purple" >
        <p style="text-align: center; color: #888888"><strong>© 2023 Copyright: bytesc</strong></p>
      </div></el-col>
      <el-col :span="4" class="foot-bottom"><div class="grid-content ep-bg-purple" ></div></el-col>
  </el-row>
</template>

<style>
.foot-item{
  padding: 0 !important;
  background: #dadada;
}
.foot-bottom{
  padding: 0 !important;
  background: #ebebeb;
}
.grid-content {
  min-height: 36px;
}

</style>


<script>
export default {
  name: "Footer"
}
</script>

