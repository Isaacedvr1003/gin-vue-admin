
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="事项ID:" prop="id">
          <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="事项简述:" prop="description">
          <el-input v-model="formData.description" :clearable="true"  placeholder="请输入事项简述" />
       </el-form-item>
        <el-form-item label="事项级别:" prop="priority">
           <el-select v-model="formData.priority" placeholder="请选择事项级别" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in mattersPriorityOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="事项状态:" prop="status">
           <el-select v-model="formData.status" placeholder="请选择事项状态" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in mattersStatusOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="解决过程:" prop="resolutionSteps">
          <RichEdit v-model="formData.resolutionSteps"/>
       </el-form-item>
        <el-form-item label="跟踪人员:" prop="assignedTo">
           <el-select v-model="formData.assignedTo" placeholder="请选择跟踪人员" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in mattersAssignedOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="创建时间:" prop="createdAt">
          <el-date-picker v-model="formData.createdAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="更新时间:" prop="updatedAt">
          <el-date-picker v-model="formData.updatedAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="删除时间:" prop="deletedAt">
          <el-date-picker v-model="formData.deletedAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="事项备注:" prop="notes">
          <RichEdit v-model="formData.notes"/>
       </el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createOpsMattersTracking,
  updateOpsMattersTracking,
  findOpsMattersTracking
} from '@/api/opsManage/opsMattersTracking'

defineOptions({
    name: 'OpsMattersTrackingForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const mattersPriorityOptions = ref([])
const mattersStatusOptions = ref([])
const mattersAssignedOptions = ref([])
const formData = ref({
            id: undefined,
            description: '',
            priority: '',
            status: '',
            resolutionSteps: '',
            assignedTo: '',
            createdAt: new Date(),
            updatedAt: new Date(),
            deletedAt: new Date(),
            notes: '',
        })
// 验证规则
const rule = reactive({
               description : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               priority : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               resolutionSteps : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               assignedTo : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               createdAt : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               updatedAt : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               deletedAt : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findOpsMattersTracking({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    mattersPriorityOptions.value = await getDictFunc('mattersPriority')
    mattersStatusOptions.value = await getDictFunc('mattersStatus')
    mattersAssignedOptions.value = await getDictFunc('mattersAssigned')
}

init()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createOpsMattersTracking(formData.value)
               break
             case 'update':
               res = await updateOpsMattersTracking(formData.value)
               break
             default:
               res = await createOpsMattersTracking(formData.value)
               break
           }
           btnLoading.value = false
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
