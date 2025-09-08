<script setup lang="ts">
  const hoge = () => {
    window.alert("Hoge!")
  }
</script>

<template>
  <el-button @click="hoge">Button!</el-button>
</template>

<style scoped></style>
