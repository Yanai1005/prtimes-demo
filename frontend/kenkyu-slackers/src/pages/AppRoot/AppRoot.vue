<script setup lang="ts">
  import type { ReviewCustomizeOption } from "@/pages/AppRoot/types";
  import OptionSelector from "./OptionSelector.vue"
  import { onMounted, ref } from "vue";

  const getOptions = (optionType: string): ReviewCustomizeOption => {
    // TODO: 本番実装置き換え
    return {
      1: {
        value: `${optionType}-value1`,
        label: `${optionType}-label1`,
      },
      2: {
        value: `${optionType}-value2`,
        label: `${optionType}-label2`,
      },
      3: {
        value: `${optionType}-value3`,
        label: `${optionType}-label3`,
      },
    }
  }

  const hogeOptions = ref<ReviewCustomizeOption>({});
  const hogeValue = ref("");
  const fugaOptions = ref<ReviewCustomizeOption>({});
  const fugaValue = ref("");
  onMounted(() => {
    hogeOptions.value = getOptions("hoge");
    fugaOptions.value = getOptions("fuga");
  })
</script>

<template>
  <option-selector :value="hogeValue" :options="hogeOptions" />
  <option-selector :value="fugaValue" :options="fugaOptions" />
</template>

<style scoped></style>
