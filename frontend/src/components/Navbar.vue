<template>
  <div id="navbar">
    <Tree :fileTree="contentStore.fileTree"></Tree>
  </div>
</template>

<script setup>
import { ref, onBeforeUnmount } from "vue";
import Tree from "./Tree.vue";
import { sha256 } from "js-sha256";
import bus from "../eventbus/index";
import { SaveFile, SaveFiles } from "../../wailsjs/go/main/App";
import { useContentStore } from "../store/content";
const contentStore = useContentStore();

bus.on("onChange", onChange);
bus.on("onSave", onSave);

function onChange(data) {
  contentStore.updateOnlyFileContent(data.totalPath, data.content);
}

function onSave(data) {
  SaveFile(data.totalPath, data.content).then(() => {
    // 保存成功 更新hash
    contentStore.updateOnlyFileHash(data.totalPath, data.content);
  });
}

// 在组件卸载之前移除监听
onBeforeUnmount(() => {
  bus.off("onSave", onSave);
  bus.off("onChange", onChange);
});
</script>

<style scoped>
#navbar {
  width: 200px;
  flex-shrink: 0;
  background-color: rgba(66, 66, 66, 0.1);
  height: calc(100vh - 50px);
  overflow: auto;
  flex-wrap: nowrap;
}
</style>
