<template>
  <div class="gva-table-box">
    <el-timeline style="max-width: 600px">
      <el-timeline-item v-for="step in steps" :timestamp="step.CreatedAt" placement="top">
        <el-card>
          <h4>节点状态：{{step.status}}</h4>
          <h4>{{ step.attr || "当前节点无需数据" }}</h4>
        </el-card>
      </el-timeline-item>
    </el-timeline>
    <div>
      <el-form v-model="formData" label-position="top">
        <el-form-item label="姓名">
          <el-input v-model="formData.name" placeholder="请输入姓名" />
        </el-form-item>
        <el-form-item label="请假原因">
          <el-input v-model="formData.info" type="textarea" placeholder="请输入请假原因" />
        </el-form-item>
      </el-form>
    </div>
      <el-button @click="next()">
        下一步
      </el-button>
      <el-button @click="next('1')">
        下一步（同意）
      </el-button>
      <el-button @click="next('')">
        下一步（拒绝）
      </el-button>
  </div>
</template>

<script setup>

import {getTransitionExampleNext, getTransitionExampleStep} from '@/plugin/flow/api/transition'
import {useRoute} from "vue-router";

const route = useRoute()

import { ref } from 'vue'
import {ElMessage} from "element-plus";

defineOptions({
  name: 'FlowTest'
})

const steps = ref([])

const getStep = async () =>{
 const res = await getTransitionExampleStep({
    workUUID:route.params.workUUID
  })
  if(res.code === 0){
    steps.value = res.data.step
  }
}

getStep()

const formData = ref({
  name: '',
  info: '',
})

const next = async (logic) =>{

  const step = steps.value.find(item => item.status === 'active')
  const res = await getTransitionExampleNext({
    stepID: step.ID,
    workUUID:route.params.workUUID,
    attr: formData.value,
    nodeID: step.nodeID,
    logic
  })
  if(res.code === 0){
    ElMessage({
      type: 'success',
      message: '操作成功'
    })
    getStep()
  }
}


</script>

<style>

</style>

