
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
           <el-form-item label="银行名称" prop="BankName">
            <el-select v-model="searchInfo.BankName" clearable placeholder="请选择" @clear="()=>{searchInfo.BankName=undefined}">
              <el-option v-for="(item,key) in banklistsOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
        <el-form-item label="工单编号" prop="ID">
         <el-input v-model="searchInfo.ID" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="工单摘要" prop="Summary">
         <el-input v-model="searchInfo.Summary" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="业务联系人" prop="BusinessContact">
         <el-input v-model="searchInfo.BusinessContact" placeholder="搜索条件" />
        </el-form-item>
           <el-form-item label="工单进度" prop="Status">
            <el-select v-model="searchInfo.Status" clearable placeholder="请选择" @clear="()=>{searchInfo.Status=undefined}">
              <el-option v-for="(item,key) in workOrderStatusOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
        <el-form-item label="工单申请时间" prop="SubmitTime">
            
            <template #label>
            <span>
              工单申请时间
              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </span>
          </template>
            <el-date-picker v-model="searchInfo.startSubmitTime" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endSubmitTime ? time.getTime() > searchInfo.endSubmitTime.getTime() : false"></el-date-picker>
            —
            <el-date-picker v-model="searchInfo.endSubmitTime" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startSubmitTime ? time.getTime() < searchInfo.startSubmitTime.getTime() : false"></el-date-picker>
        </el-form-item>
           <el-form-item label="工单负责人" prop="SubmitPerson">
            <el-select v-model="searchInfo.SubmitPerson" clearable placeholder="请选择" @clear="()=>{searchInfo.SubmitPerson=undefined}">
              <el-option v-for="(item,key) in mattersAssignedOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
                  <el-form-item label="工单详情" prop="Details">
                   <el-input v-model="searchInfo.Details" placeholder="搜索条件" />

                  </el-form-item>
          
                  <el-form-item label="工单转交付时间" prop="TransferTime">
                      
                      <template #label>
                      <span>
                        工单转交付时间
                        <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                          <el-icon><QuestionFilled /></el-icon>
                        </el-tooltip>
                      </span>
                    </template>
                      <el-date-picker v-model="searchInfo.startTransferTime" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endTransferTime ? time.getTime() > searchInfo.endTransferTime.getTime() : false"></el-date-picker>
                      —
                      <el-date-picker v-model="searchInfo.endTransferTime" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startTransferTime ? time.getTime() < searchInfo.startTransferTime.getTime() : false"></el-date-picker>

                  </el-form-item>
          
                  <el-form-item label="工单交付时间" prop="DeliveryTime">
                      
                      <template #label>
                      <span>
                        工单交付时间
                        <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                          <el-icon><QuestionFilled /></el-icon>
                        </el-tooltip>
                      </span>
                    </template>
                      <el-date-picker v-model="searchInfo.startDeliveryTime" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endDeliveryTime ? time.getTime() > searchInfo.endDeliveryTime.getTime() : false"></el-date-picker>
                      —
                      <el-date-picker v-model="searchInfo.endDeliveryTime" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startDeliveryTime ? time.getTime() < searchInfo.startDeliveryTime.getTime() : false"></el-date-picker>

                  </el-form-item>
          
                  <el-form-item label="交付企微审批编号" prop="DeliveryID">
                   <el-input v-model="searchInfo.DeliveryID" placeholder="搜索条件" />

                  </el-form-item>
          
                  <el-form-item label="工单上线时间" prop="CompleteTime">
                      
                      <template #label>
                      <span>
                        工单上线时间
                        <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                          <el-icon><QuestionFilled /></el-icon>
                        </el-tooltip>
                      </span>
                    </template>
                      <el-date-picker v-model="searchInfo.startCompleteTime" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCompleteTime ? time.getTime() > searchInfo.endCompleteTime.getTime() : false"></el-date-picker>
                      —
                      <el-date-picker v-model="searchInfo.endCompleteTime" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCompleteTime ? time.getTime() < searchInfo.startCompleteTime.getTime() : false"></el-date-picker>

                  </el-form-item>
          
                  <el-form-item label="上线堡垒机编号" prop="CompleteID">
                   <el-input v-model="searchInfo.CompleteID" placeholder="搜索条件" />

                  </el-form-item>
          
                  <el-form-item label="工单更新时间" prop="UpdateTime">
                      
                      <template #label>
                      <span>
                        工单更新时间
                        <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                          <el-icon><QuestionFilled /></el-icon>
                        </el-tooltip>
                      </span>
                    </template>
                      <el-date-picker v-model="searchInfo.startUpdateTime" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endUpdateTime ? time.getTime() > searchInfo.endUpdateTime.getTime() : false"></el-date-picker>
                      —
                      <el-date-picker v-model="searchInfo.endUpdateTime" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startUpdateTime ? time.getTime() < searchInfo.startUpdateTime.getTime() : false"></el-date-picker>

                  </el-form-item>
          
                  <el-form-item label="工单备注" prop="notes">
                   <el-input v-model="searchInfo.notes" placeholder="搜索条件" />

                  </el-form-item>
          
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
            <ExportTemplate  template-id="opsManage_BankWorkOrders" />
            <ExportExcel  template-id="opsManage_BankWorkOrders" />
            <ImportExcel  template-id="opsManage_BankWorkOrders" @on-success="getTableData" />
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        @sort-change="sortChange"
        >
        <el-table-column sortable align="left" label="银行名称" prop="BankName" width="160" show-overflow-tooltip>
            <template #default="scope">
                
                {{ filterDict(scope.row.BankName,banklistsOptions) }}
                
            </template>
        </el-table-column>
          <el-table-column sortable align="left" label="工单编号" prop="ID" width="160" />
          <el-table-column align="left" label="工单摘要" prop="Summary" width="420" show-overflow-tooltip/>
        <el-table-column sortable align="left" label="工单进度" prop="Status" width="110">
          <template #default="scope">
            <el-tag type="success" v-if="scope.row.Status === '01'">{{filterDict(scope.row.Status,workOrderStatusOptions)}}</el-tag>
            <el-tag type="primary" v-else-if="scope.row.Status === '02'">{{filterDict(scope.row.Status,workOrderStatusOptions)}}</el-tag>
            <el-tag type="danger" v-else-if="scope.row.Status === '03'">{{filterDict(scope.row.Status,workOrderStatusOptions)}}</el-tag>
            <el-tag type="primary" v-else>{{filterDict(scope.row.Status,workOrderStatusOptions)}}</el-tag>
          </template>
        </el-table-column>
         <el-table-column sortable align="left" label="工单申请时间" prop="SubmitTime" width="160">
            <template #default="scope">{{ formatDate(scope.row.SubmitTime) }}</template>
         </el-table-column>
        <el-table-column sortable align="left" label="工单负责人" prop="SubmitPerson" width="120">
            <template #default="scope">
                
                {{ filterDict(scope.row.SubmitPerson,mattersAssignedOptions) }}
                
            </template>
        </el-table-column>
          <el-table-column label="工单详情" prop="Details" width="120"  show-overflow-tooltip>
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
                  <RichView v-model="scope.row.Details" />
                </template>
              </el-popover>
            </template>
          </el-table-column>
          <el-table-column sortable align="left" label="业务联系人" prop="BusinessContact" width="120" />
          <el-table-column sortable align="left" label="工单转交付时间" prop="TransferTime" width="180">
            <template #default="scope">{{ formatDate(scope.row.TransferTime) }}</template>
         </el-table-column>
         <el-table-column sortable align="left" label="工单交付时间" prop="DeliveryTime" width="180">
            <template #default="scope">{{ formatDate(scope.row.DeliveryTime) }}</template>
         </el-table-column>
          <el-table-column sortable align="left" label="交付企微审批编号" prop="DeliveryID" width="180" />
         <el-table-column sortable align="left" label="工单上线时间" prop="CompleteTime" width="180">
            <template #default="scope">{{ formatDate(scope.row.CompleteTime) }}</template>
         </el-table-column>
          <el-table-column sortable align="left" label="上线堡垒机编号" prop="CompleteID" width="180" />
         <el-table-column sortable align="left" label="工单更新时间" prop="UpdateTime" width="180">
            <template #default="scope">{{ formatDate(scope.row.UpdateTime) }}</template>
         </el-table-column>
                      <el-table-column label="工单备注" prop="notes" width="200">
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
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateBankWorkOrdersFunc(scope.row)">编辑</el-button>
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
            <el-form-item label="银行名称:"  prop="BankName" >
              <el-select v-model="formData.BankName" placeholder="请选择银行名称" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in banklistsOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="工单编号:"  prop="ID" >
              <el-input v-model="formData.ID" :clearable="true"  placeholder="请输入工单编号" />
            </el-form-item>
            <el-form-item label="工单摘要:"  prop="Summary" >
              <el-input v-model="formData.Summary" :clearable="true"  placeholder="请输入工单摘要" />
            </el-form-item>
            <el-form-item label="业务联系人:"  prop="BusinessContact" >
              <el-input v-model="formData.BusinessContact" :clearable="true"  placeholder="请输入业务联系人" />
            </el-form-item>
            <el-form-item label="工单详情:"  prop="Details" >
              <RichEdit v-model="formData.Details"/>
            </el-form-item>
            <el-form-item label="工单进度:"  prop="Status" >
              <el-select v-model="formData.Status" placeholder="请选择工单进度" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in workOrderStatusOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="工单申请时间:"  prop="SubmitTime" >
              <el-date-picker v-model="formData.SubmitTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="工单负责人:"  prop="SubmitPerson" >
              <el-select v-model="formData.SubmitPerson" placeholder="请选择工单负责人" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in mattersAssignedOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="工单转交付时间:"  prop="TransferTime" >
              <el-date-picker v-model="formData.TransferTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="工单交付时间:"  prop="DeliveryTime" >
              <el-date-picker v-model="formData.DeliveryTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="交付企微审批编号:"  prop="DeliveryID" >
              <el-input v-model="formData.DeliveryID" :clearable="true"  placeholder="请输入交付企微审批编号" />
            </el-form-item>
            <el-form-item label="工单上线时间:"  prop="CompleteTime" >
              <el-date-picker v-model="formData.CompleteTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="上线堡垒机编号:"  prop="CompleteID" >
              <el-input v-model="formData.CompleteID" :clearable="true"  placeholder="请输入上线堡垒机编号" />
            </el-form-item>
            <el-form-item label="工单更新时间:"  prop="UpdateTime" >
              <el-date-picker v-model="formData.UpdateTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="工单备注:"  prop="notes" >
              <RichEdit v-model="formData.notes"/>
            </el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="银行名称">
                        
                        {{ filterDict(detailFrom.BankName,banklistsOptions) }}
                        
                    </el-descriptions-item>
                    <el-descriptions-item label="工单编号">
                        {{ detailFrom.ID }}
                    </el-descriptions-item>
                    <el-descriptions-item label="工单摘要">
                        {{ detailFrom.Summary }}
                    </el-descriptions-item>
                    <el-descriptions-item label="业务联系人">
                        {{ detailFrom.BusinessContact }}
                    </el-descriptions-item>
                    <el-descriptions-item label="工单详情">
                            <RichView v-model="detailFrom.Details" />
                    </el-descriptions-item>
                    <el-descriptions-item label="工单进度">
                        
                        {{ filterDict(detailFrom.Status,workOrderStatusOptions) }}
                        
                    </el-descriptions-item>
                    <el-descriptions-item label="工单申请时间">
                        {{ detailFrom.SubmitTime }}
                    </el-descriptions-item>
                    <el-descriptions-item label="工单负责人">
                        
                        {{ filterDict(detailFrom.SubmitPerson,mattersAssignedOptions) }}
                        
                    </el-descriptions-item>
                    <el-descriptions-item label="工单转交付时间">
                        {{ detailFrom.TransferTime }}
                    </el-descriptions-item>
                    <el-descriptions-item label="工单交付时间">
                        {{ detailFrom.DeliveryTime }}
                    </el-descriptions-item>
                    <el-descriptions-item label="交付企微审批编号">
                        {{ detailFrom.DeliveryID }}
                    </el-descriptions-item>
                    <el-descriptions-item label="工单上线时间">
                        {{ detailFrom.CompleteTime }}
                    </el-descriptions-item>
                    <el-descriptions-item label="上线堡垒机编号">
                        {{ detailFrom.CompleteID }}
                    </el-descriptions-item>
                    <el-descriptions-item label="工单更新时间">
                        {{ detailFrom.UpdateTime }}
                    </el-descriptions-item>
                    <el-descriptions-item label="工单备注">
                            <RichView v-model="detailFrom.notes" />
                    </el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createBankWorkOrders,
  deleteBankWorkOrders,
  deleteBankWorkOrdersByIds,
  updateBankWorkOrders,
  findBankWorkOrders,
  getBankWorkOrdersList
} from '@/api/opsManage/bankWorkOrders'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'
import RichView from '@/components/richtext/rich-view.vue'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useAppStore } from "@/pinia"

// 导出组件
import ExportExcel from '@/components/exportExcel/exportExcel.vue'
// 导入组件
import ImportExcel from '@/components/exportExcel/importExcel.vue'
// 导出模板组件
import ExportTemplate from '@/components/exportExcel/exportTemplate.vue'


defineOptions({
    name: 'BankWorkOrders'
})

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
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
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               ID : [{
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
               Summary : [{
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
               BusinessContact : [{
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
               Details : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               Status : [{
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
               SubmitTime : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               SubmitPerson : [{
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
        SubmitTime : [{ validator: (rule, value, callback) => {
        if (searchInfo.value.startSubmitTime && !searchInfo.value.endSubmitTime) {
          callback(new Error('请填写结束日期'))
        } else if (!searchInfo.value.startSubmitTime && searchInfo.value.endSubmitTime) {
          callback(new Error('请填写开始日期'))
        } else if (searchInfo.value.startSubmitTime && searchInfo.value.endSubmitTime && (searchInfo.value.startSubmitTime.getTime() === searchInfo.value.endSubmitTime.getTime() || searchInfo.value.startSubmitTime.getTime() > searchInfo.value.endSubmitTime.getTime())) {
          callback(new Error('开始日期应当早于结束日期'))
        } else {
          callback()
        }
      }, trigger: 'change' }],
        TransferTime : [{ validator: (rule, value, callback) => {
        if (searchInfo.value.startTransferTime && !searchInfo.value.endTransferTime) {
          callback(new Error('请填写结束日期'))
        } else if (!searchInfo.value.startTransferTime && searchInfo.value.endTransferTime) {
          callback(new Error('请填写开始日期'))
        } else if (searchInfo.value.startTransferTime && searchInfo.value.endTransferTime && (searchInfo.value.startTransferTime.getTime() === searchInfo.value.endTransferTime.getTime() || searchInfo.value.startTransferTime.getTime() > searchInfo.value.endTransferTime.getTime())) {
          callback(new Error('开始日期应当早于结束日期'))
        } else {
          callback()
        }
      }, trigger: 'change' }],
        DeliveryTime : [{ validator: (rule, value, callback) => {
        if (searchInfo.value.startDeliveryTime && !searchInfo.value.endDeliveryTime) {
          callback(new Error('请填写结束日期'))
        } else if (!searchInfo.value.startDeliveryTime && searchInfo.value.endDeliveryTime) {
          callback(new Error('请填写开始日期'))
        } else if (searchInfo.value.startDeliveryTime && searchInfo.value.endDeliveryTime && (searchInfo.value.startDeliveryTime.getTime() === searchInfo.value.endDeliveryTime.getTime() || searchInfo.value.startDeliveryTime.getTime() > searchInfo.value.endDeliveryTime.getTime())) {
          callback(new Error('开始日期应当早于结束日期'))
        } else {
          callback()
        }
      }, trigger: 'change' }],
        CompleteTime : [{ validator: (rule, value, callback) => {
        if (searchInfo.value.startCompleteTime && !searchInfo.value.endCompleteTime) {
          callback(new Error('请填写结束日期'))
        } else if (!searchInfo.value.startCompleteTime && searchInfo.value.endCompleteTime) {
          callback(new Error('请填写开始日期'))
        } else if (searchInfo.value.startCompleteTime && searchInfo.value.endCompleteTime && (searchInfo.value.startCompleteTime.getTime() === searchInfo.value.endCompleteTime.getTime() || searchInfo.value.startCompleteTime.getTime() > searchInfo.value.endCompleteTime.getTime())) {
          callback(new Error('开始日期应当早于结束日期'))
        } else {
          callback()
        }
      }, trigger: 'change' }],
        UpdateTime : [{ validator: (rule, value, callback) => {
        if (searchInfo.value.startUpdateTime && !searchInfo.value.endUpdateTime) {
          callback(new Error('请填写结束日期'))
        } else if (!searchInfo.value.startUpdateTime && searchInfo.value.endUpdateTime) {
          callback(new Error('请填写开始日期'))
        } else if (searchInfo.value.startUpdateTime && searchInfo.value.endUpdateTime && (searchInfo.value.startUpdateTime.getTime() === searchInfo.value.endUpdateTime.getTime() || searchInfo.value.startUpdateTime.getTime() > searchInfo.value.endUpdateTime.getTime())) {
          callback(new Error('开始日期应当早于结束日期'))
        } else {
          callback()
        }
      }, trigger: 'change' }],
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
            BankName: 'BankName',
            ID: 'ID',
            BusinessContact: 'BusinessContact',
            Status: 'Status',
            SubmitTime: 'SubmitTime',
            SubmitPerson: 'SubmitPerson',
            TransferTime: 'TransferTime',
            DeliveryTime: 'DeliveryTime',
            DeliveryID: 'DeliveryID',
            CompleteTime: 'CompleteTime',
            CompleteID: 'CompleteID',
            UpdateTime: 'UpdateTime',
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
  const table = await getBankWorkOrdersList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    banklistsOptions.value = await getDictFunc('banklists')
    workOrderStatusOptions.value = await getDictFunc('workOrderStatus')
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
            deleteBankWorkOrdersFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteBankWorkOrdersByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateBankWorkOrdersFunc = async(row) => {
    const res = await findBankWorkOrders({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteBankWorkOrdersFunc = async (row) => {
    const res = await deleteBankWorkOrders({ ID: row.ID })
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
  const res = await findBankWorkOrders({ ID: row.ID })
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


</script>

<style lang="scss" scoped>
  :deep(.el-popper) {
    @apply max-w-lg bg-white text-gray-600 text-base;
  }
</style>