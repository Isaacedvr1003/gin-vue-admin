<template>
  <div>
    <div class="box-border w-14 h-14 rounded-full border-2 border-solid border-green-600 bg-green-500 flex items-center justify-center text-white">
      开始
    </div>
    <div class="w-14 text-sm text-center mt-2">{{data.title}}</div>
  </div>
</template>

<script setup>
import { ref,inject,onMounted } from 'vue'

defineOptions({
  name: 'StartNode',
})

const getNode = inject('getNode')

const data = ref({})

onMounted(() => {
  const node = getNode()
  data.value = node.getData()
})

</script>
