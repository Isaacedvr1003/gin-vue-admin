<!--
    @auther: bypanghu<bypanghu@163.com>
    @date: 2024/5/7
!-->

<template>
  <div class="flex items-center mx-4 gap-4">
     <el-tooltip class="" effect="dark" content="搜索" placement="bottom">
      <el-icon
        @click="handleCommand"
        class="w-8 h-8 shadow rounded-full border border-gray-200 dark:border-gray-600 cursor-pointer border-solid"
      >
        <Search />
      </el-icon>
    </el-tooltip>

    <el-tooltip class="" effect="dark" content="系统设置" placement="bottom">
      <el-icon
        class="w-8 h-8 shadow rounded-full border border-gray-200 dark:border-gray-600 cursor-pointer border-solid"
        @click="toggleSetting"
      >
        <Setting />
      </el-icon>
    </el-tooltip>
    <gva-setting v-model:drawer="showSettingDrawer"></gva-setting>
    <command-menu ref="command" />
  </div>
</template>

<script setup>
  import GvaSetting from '@/view/layout/setting/index.vue'
  import { ref } from 'vue'
  import CommandMenu from '@/components/commandMenu/index.vue'

  const showSettingDrawer = ref(false)

  const toggleSetting = () => {
    showSettingDrawer.value = true
  }

  const first = ref('')
  const command = ref()

  const handleCommand = () => {
    command.value.open()
  }
  const initPage = () => {
    // 判断当前用户的操作系统
    if (window.localStorage.getItem('osType') === 'WIN') {
      first.value = 'Ctrl'
    } else {
      first.value = '⌘'
    }
    // 当用户同时按下ctrl和k键的时候
    const handleKeyDown = (e) => {
      if (e.ctrlKey && e.key === 'k') {
        // 阻止浏览器默认事件
        e.preventDefault()
        handleCommand()
      }
    }
    window.addEventListener('keydown', handleKeyDown)
  }
  initPage()
</script>

<style scoped lang="scss"></style>
