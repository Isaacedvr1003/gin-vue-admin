<template>
  <div class="p-2">
    <div class="bg-white rounded p-2 flex flex-col gap-2">
      <div>
        名称： {{ formData.name }}({{ formData.uuid }})
      </div>
      <div>
        介绍： {{ formData.desc }}
      </div>
    </div>
    <div class="h-[calc(100vh-240px)] p-2 bg-white mt-2 rounded">
      <FlowGraph :flowUUID="route.params.uuid"/>
    </div>
  </div>
</template>

<script setup>
import {
  findFlow
} from '@/plugin/flow/api/flow'


defineOptions({
    name: 'FlowForm'
})

import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
import FlowGraph from "@/plugin/flow/view/components/graph.vue";

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            name: '',
            uuid: '',
            desc: '',
        })


// 初始化方法
const init = async () => {
  console.log(route.params.uuid)
    if (route.params.uuid) {
      const res = await findFlow({ uuid: route.params.uuid })
      if (res.code === 0) {
        formData.value = res.data.reflow
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
  router.go(-1)
}

</script>

