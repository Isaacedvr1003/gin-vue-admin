<template>
  <div>
    <div class="box-border w-14 h-14 rounded-full border-2 border-solid border-red-500 bg-red-400 flex items-center justify-center text-white">
      结束
    </div>
    <div class="w-14 text-sm text-center mt-2">
      {{ data.title }}
    </div>
  </div>

</template>

<script setup>
import { ref,inject,onMounted } from 'vue'

defineOptions({
  name: 'ProgressNode',
})

const data = ref({})
const getNode = inject('getNode')

onMounted(() => {
  const node = getNode()
  data.value = node.getData()
})

</script>
