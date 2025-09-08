<script setup lang="ts">
  import { ref } from 'vue';

  const { value } = defineProps<{
    value: string,
  }>()
  const valueRef = ref(value);

</script>

<template>
  <el-input
    v-model="valueRef"
    :rows="2"
    type="textarea"
  />
</template>

<style scoped></style>
