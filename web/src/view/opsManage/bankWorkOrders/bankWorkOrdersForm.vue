
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="银行名称:" prop="BankName">
           <el-select v-model="formData.BankName" placeholder="请选择银行名称" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in banklistsOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="工单编号:" prop="ID">
          <el-input v-model="formData.ID" :clearable="true"  placeholder="请输入工单编号" />
       </el-form-item>
        <el-form-item label="工单摘要:" prop="Summary">
          <el-input v-model="formData.Summary" :clearable="true"  placeholder="请输入工单摘要" />
       </el-form-item>
        <el-form-item label="业务联系人:" prop="BusinessContact">
          <el-input v-model="formData.BusinessContact" :clearable="true"  placeholder="请输入业务联系人" />
       </el-form-item>
        <el-form-item label="工单详情:" prop="Details">
          <RichEdit v-model="formData.Details"/>
       </el-form-item>
        <el-form-item label="工单进度:" prop="Status">
           <el-select v-model="formData.Status" placeholder="请选择工单进度" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in workOrderStatusOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="工单申请时间:" prop="SubmitTime">
          <el-date-picker v-model="formData.SubmitTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="工单负责人:" prop="SubmitPerson">
           <el-select v-model="formData.SubmitPerson" placeholder="请选择工单负责人" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in mattersAssignedOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="工单转交付时间:" prop="TransferTime">
          <el-date-picker v-model="formData.TransferTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="工单交付时间:" prop="DeliveryTime">
          <el-date-picker v-model="formData.DeliveryTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="交付企微审批编号:" prop="DeliveryID">
          <el-input v-model="formData.DeliveryID" :clearable="true"  placeholder="请输入交付企微审批编号" />
       </el-form-item>
        <el-form-item label="工单上线时间:" prop="CompleteTime">
          <el-date-picker v-model="formData.CompleteTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="上线堡垒机编号:" prop="CompleteID">
          <el-input v-model="formData.CompleteID" :clearable="true"  placeholder="请输入上线堡垒机编号" />
       </el-form-item>
        <el-form-item label="工单更新时间:" prop="UpdateTime">
          <el-date-picker v-model="formData.UpdateTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="工单备注:" prop="notes">
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
  createBankWorkOrders,
  updateBankWorkOrders,
  findBankWorkOrders
} from '@/api/opsManage/bankWorkOrders'

defineOptions({
    name: 'BankWorkOrdersForm'
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
const banklistsOptions = ref([])
const workOrderStatusOptions = ref([])
const mattersAssignedOptions = ref([])
const formData = ref({
            BankName: '',
            ID: '',
            Summary: '',
            BusinessContact: '',
            Details: '',
            Status: '',
            SubmitTime: new Date(),
            SubmitPerson: '',
            TransferTime: new Date(),
            DeliveryTime: new Date(),
            DeliveryID: '',
            CompleteTime: new Date(),
            CompleteID: '',
            UpdateTime: new Date(),
            notes: '',
        })
// 验证规则
const rule = reactive({
               BankName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               ID : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               Summary : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               BusinessContact : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               Details : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               Status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               SubmitTime : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               SubmitPerson : [{
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
      const res = await findBankWorkOrders({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    banklistsOptions.value = await getDictFunc('banklists')
    workOrderStatusOptions.value = await getDictFunc('workOrderStatus')
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
               res = await createBankWorkOrders(formData.value)
               break
             case 'update':
               res = await updateBankWorkOrders(formData.value)
               break
             default:
               res = await createBankWorkOrders(formData.value)
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
