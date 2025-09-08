<script setup lang="ts">
  import type { ReviewCustomizeOption } from '@/pages/AppRoot/types';
  import { ref } from 'vue';

  const { options, value } = defineProps<{
    options: ReviewCustomizeOption,
    value: string,
  }>()
  const valueRef = ref(value);

</script>

<template>
  <el-select v-model="valueRef">
    <el-option
      v-for="[key, {value, label}] of Object.entries(options)"
      :key
      :value
      :label
    />
  </el-select>
</template>

<style scoped></style>
