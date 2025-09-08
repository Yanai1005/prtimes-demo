<script setup lang="ts">
  import type { ReviewCustomizeOption } from "@/pages/AppRoot/types";
  import OptionSelector from "./OptionSelector.vue"
  import { onMounted, ref } from "vue";
import MarkdownEditor from "@/pages/AppRoot/MarkdownEditor.vue";

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
  const fugaOptions = ref<ReviewCustomizeOption>({});
  const hogeValue = ref("");
  const fugaValue = ref("");
  const markdownContent = ref("");
  onMounted(() => {
    hogeOptions.value = getOptions("hoge");
    fugaOptions.value = getOptions("fuga");
  })
</script>

<template>
  <div class="container">
    <el-form>
      <el-form-item label="選択肢1">
        <option-selector :value="hogeValue" :options="hogeOptions" />
      </el-form-item>
      <el-form-item label="選択肢2">
        <option-selector :value="fugaValue" :options="fugaOptions" />
      </el-form-item>
      <el-form-item label="本文（マークダウン）">
        <markdown-editor :value="markdownContent" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary">
          送信
        </el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<style scoped>
  .container {
    padding: 0 20vw;
  }
</style>
