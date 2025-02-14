<template>
  <div class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-8 py-2 gap-4 md:gap-2 gva-container2">
    <gva-card custom-class="col-span-1 lg:col-span-2 h-32">
      <gva-chart :type="1" title="访问人数" />
    </gva-card>
    <gva-card custom-class="col-span-1 lg:col-span-2 h-32 ">
      <gva-chart :type="2" title="新增客户" />
    </gva-card>
    <gva-card custom-class="col-span-1 lg:col-span-2 h-32">
      <gva-chart :type="3" title="解决数量" />
    </gva-card>
    <gva-card custom-class="col-span-1 lg:col-span-2 h-32">
      <gva-chart :type="2" title="好评数量" />
    </gva-card>
    <gva-card title="内容数据" custom-class="col-span-1 md:col-span-2 md:row-start-2 lg:col-span-8 col-start-1 row-span-2">
      <gva-chart :type="4" />
    </gva-card>
  </div>
</template>

<script setup>
import { GvaChart, GvaCard  } from "./components"
defineOptions({
  name: 'DashBoard'
})
</script>

<style lang="scss" scoped>
</style>
