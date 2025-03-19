<template>
  <div>
    <div class="box-border w-20 h-14 rounded-2xl border-2 border-solid border-blue-600 bg-blue-500 flex items-center justify-center text-white">
      事件
    </div>
    <div class="w-20 text-sm text-center mt-2">
      {{ data.title }}
    </div>
  </div>
</template>

<script setup>
import { ref,inject,onMounted } from 'vue'

defineOptions({
  name: 'EventNode',
})

const getNode = inject('getNode')
const data = ref('')

onMounted(() => {
  const node = getNode()
  data.value = node.getData()
})

</script>
