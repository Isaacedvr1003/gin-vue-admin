<template>
  <div>
    <div class="w-18 h-10 bg-blue-500 border-2 border-solid border-blue-600 transform skew-x-12 flex items-center justify-center text-white">
      <span class="-skew-x-12">自动</span>
    </div>
    <div class="w-18 text-sm text-center mt-2">
      {{ data.title }}
    </div>
  </div>
</template>

<script setup>
import { ref,inject,onMounted } from 'vue'

defineOptions({
  name: 'AutoNode',
})

const data = ref({})
const getNode = inject('getNode')

onMounted(() => {
  const node = getNode()
  data.value = node.getData()
})

</script>
