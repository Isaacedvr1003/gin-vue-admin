<template>
  <div>
    <div class="box-border w-12 h-12 bg-amber-500 border-2 border-solid border-amber-600 transform rotate-45 translate-y-2 translate-x-2 flex items-center justify-center text-white">
      <span class="-rotate-45">合并</span>
    </div>
    <div class="w-12 text-sm text-center mt-2">
      {{ data.title }}
    </div>
  </div>
</template>

<script setup>
import {ref, inject, onMounted} from 'vue'

defineOptions({
  name: 'MergeNode',
})

const data = ref({})
const getNode = inject('getNode')

onMounted(() => {
  const node = getNode()
  data.value = node.getData()
})

</script>
