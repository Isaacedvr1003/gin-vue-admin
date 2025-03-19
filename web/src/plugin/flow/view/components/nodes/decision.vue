<template>
  <div>
    <div class="box-border w-12 h-12 bg-orange-500 translate-y-2 translate-x-2 border-2 border-solid border-orange-600 transform rotate-45 flex items-center justify-center text-white">
      <span class="-rotate-45">决策</span>
    </div>
    <div class="w-16 text-sm text-center mt-4">
      {{ data.title }}
    </div>
   </div>
</template>

<script setup>
import {ref, inject, onMounted} from 'vue'

defineOptions({
  name: 'DecisionNode',
})

const data = ref({})
const getNode = inject('getNode')

onMounted(() => {
  const node = getNode()
  data.value = node.getData()
})

</script>
