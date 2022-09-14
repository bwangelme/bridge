export default {
  template: `<template>
    <el-table :data="tableData">
       <el-table-colume prop="id" label="ID" width="180"></el-table-colume>     
    </el-table>
  </template>`,
  data: {
    "tableData": [
      {"id": 1,},
      {"id": 2,},
      {"id": 3,},
      {"id": 4,},
    ]
  }
}
