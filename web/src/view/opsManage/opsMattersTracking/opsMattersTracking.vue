
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
        <el-form-item label="事项ID" prop="id">
            
             <el-input v-model.number="searchInfo.id" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="事项简述" prop="description">
         <el-input v-model="searchInfo.description" placeholder="搜索条件" />
        </el-form-item>
           <el-form-item label="事项级别" prop="priority">
            <el-select v-model="searchInfo.priority" clearable placeholder="请选择" @clear="()=>{searchInfo.priority=undefined}">
              <el-option v-for="(item,key) in mattersPriorityOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
           <el-form-item label="事项状态" prop="status">
            <el-select v-model="searchInfo.status" clearable placeholder="请选择" @clear="()=>{searchInfo.status=undefined}">
              <el-option v-for="(item,key) in mattersStatusOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
           <el-form-item label="跟踪人员" prop="assignedTo">
            <el-select v-model="searchInfo.assignedTo" clearable placeholder="请选择" @clear="()=>{searchInfo.assignedTo=undefined}">
              <el-option v-for="(item,key) in mattersAssignedOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button  type="primary" icon="plus" @click="openDialog()">新增</el-button>
            <ExportTemplate  template-id="opsManage_OpsMattersTracking" />
            <ExportExcel  template-id="opsManage_OpsMattersTracking" />
            <ImportExcel  template-id="opsManage_OpsMattersTracking" @on-success="getTableData" />
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="id"
        @selection-change="handleSelectionChange"
        @sort-change="sortChange"
        >
          <el-table-column sortable align="left" label="创建时间" prop="createdAt" width="180">
            <template #default="scope">{{ formatDate(scope.row.createdAt) }}</template>
          </el-table-column>
          <el-table-column sortable align="left" label="事项简述" prop="description" width="500" show-overflow-tooltip/>
        <el-table-column sortable align="left" label="事项级别" prop="priority" width="120">
            <template #default="scope">
              <el-rate v-model="scope.row.priority" :max="3" size="large" disabled/>
            </template>
        </el-table-column>
        <el-table-column sortable align="left" label="事项状态" prop="status" width="120">
            <template #default="scope">
              <el-tag type="success" v-if="scope.row.status === '01'">{{filterDict(scope.row.status,mattersStatusOptions)}}</el-tag>
              <el-tag type="primary" v-else-if="scope.row.status === '02'">{{filterDict(scope.row.status,mattersStatusOptions)}}</el-tag>
              <el-tag type="danger" v-else-if="scope.row.status === '03'">{{filterDict(scope.row.status,mattersStatusOptions)}}</el-tag>
              <el-tag type="primary" v-else>{{filterDict(scope.row.status,mattersStatusOptions)}}</el-tag>
            </template>
        </el-table-column>
                      <el-table-column label="解决过程" prop="resolutionSteps" width="140">
                         <template #default="scope">
                           <el-popover
                             placement="top"
                             :width="400"
                             trigger="hover"
                           >
                             <template #reference>
                               <el-text>[富文本内容]</el-text>
                             </template>
                             <template #default>
                               <RichView v-model="scope.row.resolutionSteps" />
                             </template>
                           </el-popover>
                         </template>
                      </el-table-column>
        <el-table-column sortable align="left" label="跟踪人员" prop="assignedTo" width="120">
            <template #default="scope">
                
                {{ filterDict(scope.row.assignedTo,mattersAssignedOptions) }}
                
            </template>
        </el-table-column>
         <el-table-column sortable align="left" label="更新时间" prop="updatedAt" width="180">
            <template #default="scope">{{ formatDate(scope.row.updatedAt) }}</template>
         </el-table-column>
                      <el-table-column label="事项备注" prop="notes" width="200">
                        <template #default="scope">
                          <el-popover
                            placement="top"
                            :width="400"
                            trigger="hover"
                          >
                            <template #reference>
                              <el-text>[富文本内容]</el-text>
                            </template>
                            <template #default>
                              <RichView v-model="scope.row.notes" />
                            </template>
                          </el-popover>
                        </template>
                      </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" :min-width="150">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateOpsMattersTrackingFunc(scope.row)">编辑</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'新增':'编辑'}}</span>
                <div>
                  <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="事项简述:"  prop="description" >
              <el-input v-model="formData.description" :clearable="true" :disabled="type!=='create'"  placeholder="请输入事项简述" />
            </el-form-item>
            <el-form-item label="事项级别:"  prop="priority" >
              <el-select v-model="formData.priority" placeholder="请选择事项级别" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in mattersPriorityOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="事项状态:"  prop="status" >
              <el-select v-model="formData.status" placeholder="请选择事项状态" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in mattersStatusOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="解决过程:"  prop="resolutionSteps" >
              <RichEdit v-model="formData.resolutionSteps"/>
            </el-form-item>
            <el-form-item label="跟踪人员:"  prop="assignedTo" >
              <el-select v-model="formData.assignedTo" placeholder="请选择跟踪人员" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in mattersAssignedOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="创建时间:"  prop="createdAt" >
              <el-date-picker v-model="formData.createdAt" :disabled="type!=='create'" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="更新时间:"  prop="updatedAt" >
              <el-date-picker v-model="formData.updatedAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="事项备注:"  prop="notes" >
              <RichEdit v-model="formData.notes"/>
            </el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border label-width="90">
                    <el-descriptions-item label="事项ID">
                        {{ detailFrom.id }}
                    </el-descriptions-item>
                    <el-descriptions-item label="事项简述">
                        {{ detailFrom.description }}
                    </el-descriptions-item>
                    <el-descriptions-item label="事项级别">
                        
                        {{ filterDict(detailFrom.priority,mattersPriorityOptions) }}
                        
                    </el-descriptions-item>
                    <el-descriptions-item label="事项状态">
                        
                        {{ filterDict(detailFrom.status,mattersStatusOptions) }}
                        
                    </el-descriptions-item>
                    <el-descriptions-item label="解决过程">
                            <RichView v-model="detailFrom.resolutionSteps" />
                    </el-descriptions-item>
                    <el-descriptions-item label="跟踪人员">
                        
                        {{ filterDict(detailFrom.assignedTo,mattersAssignedOptions) }}
                        
                    </el-descriptions-item>
                    <el-descriptions-item label="创建时间">
                        {{ detailFrom.createdAt }}
                    </el-descriptions-item>
                    <el-descriptions-item label="更新时间">
                        {{ detailFrom.updatedAt }}
                    </el-descriptions-item>
                    <el-descriptions-item label="事项备注">
                            <RichView v-model="detailFrom.notes" />
                    </el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createOpsMattersTracking,
  deleteOpsMattersTracking,
  deleteOpsMattersTrackingByIds,
  updateOpsMattersTracking,
  findOpsMattersTracking,
  getOpsMattersTrackingList
} from '@/api/opsManage/opsMattersTracking'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'
import RichView from '@/components/richtext/rich-view.vue'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive, computed } from 'vue'
import { useAppStore } from "@/pinia"

// 导出组件
import ExportExcel from '@/components/exportExcel/exportExcel.vue'
// 导入组件
import ImportExcel from '@/components/exportExcel/importExcel.vue'
// 导出模板组件
import ExportTemplate from '@/components/exportExcel/exportTemplate.vue'


defineOptions({
    name: 'OpsMattersTracking'
})
const ratePriority = ref(3)
// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
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
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               priority : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               resolutionSteps : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               assignedTo : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               createdAt : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               updatedAt : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               deletedAt : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 排序
const sortChange = ({ prop, order }) => {
  const sortMap = {
            id: 'id',
            description: 'description',
            priority: 'priority',
            status: 'status',
            assignedTo: 'assigned_to',
            createdAt: 'created_at',
            updatedAt: 'updated_at',
            deletedAt: 'deleted_at',
  }

  let sort = sortMap[prop]
  if(!sort){
   sort = prop.replace(/[A-Z]/g, match => `_${match.toLowerCase()}`)
  }

  searchInfo.value.sort = sort
  searchInfo.value.order = order
  getTableData()
}
// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getOpsMattersTrackingList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    mattersPriorityOptions.value = await getDictFunc('mattersPriority')
    mattersStatusOptions.value = await getDictFunc('mattersStatus')
    mattersAssignedOptions.value = await getDictFunc('mattersAssigned')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteOpsMattersTrackingFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.id)
        })
      const res = await deleteOpsMattersTrackingByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateOpsMattersTrackingFunc = async(row) => {
    const res = await findOpsMattersTracking({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteOpsMattersTrackingFunc = async (row) => {
    const res = await deleteOpsMattersTracking({ id: row.id })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
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
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}


const detailFrom = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findOpsMattersTracking({ id: row.id })
  if (res.code === 0) {
    detailFrom.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailFrom.value = {}
}
const getRatePriority = (row) => {
  return computed(() => row.priority - 96);
};

</script>

<style lang="scss" scoped>
:deep(.el-popper) {
  @apply max-w-lg bg-white text-gray-600 text-base;
}
</style>
