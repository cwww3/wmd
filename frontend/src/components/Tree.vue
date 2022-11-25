<template>
  <div class="main">
    <div v-if="fileTree.name" class="fd" @click.stop="entry">
      <i :class="fileTreeStyle"></i>
      <span id="filename">{{ fileTree.name }}</span>
      <i v-if="fileTree.type === 'file'" :class="saveStyle"></i>
    </div>
    <Tree
      v-show="fileTree.isShow"
      v-for="tree in fileTree.children"
      :fileTree="tree"
    />
  </div>
</template>

<script setup>
import { defineProps, toRefs } from "vue";
import { ReadFile } from "../../wailsjs/go/main/App";

import { sha256 } from "js-sha256";
import { useContentStore } from "../store/content";
import { computed } from "@vue/reactivity";
const props = defineProps({
  fileTree: Object,
});
const { fileTree } = toRefs(props);

const contentStore = useContentStore();

const contentHash = computed({
  get() {
    return sha256(fileTree.value.content || "");
  },
});

const fileTreeStyle = computed(() => {
  return {
    "icon-opendir": !!(fileTree.value.type === "dir" && fileTree.value.isShow),
    "icon-closedir": fileTree.value.type === "dir" && !fileTree.value.isShow,
    "icon-file": fileTree.value.type === "file",
  };
});

const saveStyle = computed(() => {
  return {
    "icon-save": fileTree.value.contentHash
      ? contentHash.value !== fileTree.value.contentHash
      : false,
  };
});

function entry() {
  const isDir = fileTree.value?.type === "dir";
  if (isDir) {
    contentStore.collapseDir(fileTree.value.totalPath);
  } else {
    readFile();
  }
}

function readFile() {
  //判断是否是第一次读文件内容
  if (fileTree.value.contentHash) {
    contentStore.swithContent(fileTree.value.totalPath, fileTree.value.content);
    return;
  }

  ReadFile(fileTree.value.totalPath).then((result) => {
    contentStore.updateFileContent(fileTree.value.totalPath, result);
    contentStore.swithContent(fileTree.value.totalPath, result);
  });
}
</script>

<style scoped>
.fd {
  display: flex;
  align-items: center;
}
/* .icon {
  background: url("../assets/images/file.svg") no-repeat center;
  width: 20px;
  height: 20px;
  background-size: cover;
} */

.icon-opendir {
  background: url("../assets/images/open_dir.svg") no-repeat center;
  width: 20px;
  height: 20px;
  background-size: cover;
}
.icon-closedir {
  background: url("../assets/images/dir.svg") no-repeat center;
  width: 20px;
  height: 20px;
  background-size: cover;
}
.icon-file {
  background: url("../assets/images/file.svg") no-repeat center;
  width: 20px;
  height: 20px;
  background-size: cover;
}
.icon-save {
  background: url("../assets/images/dot.svg") no-repeat center;
  width: 20px;
  height: 20px;
  background-size: cover;
}

.main:nth-child(n + 2) {
  padding-left: 20px;
}

#filename {
  display: block;
}
</style>
