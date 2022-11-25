<template>
  <div id="main">
    <div id="header">
      <router-link to="/"> 首页 </router-link>
      <router-link to="/config"> 配置 </router-link>
    </div>
    <div id="body">
      <router-view />
    </div>
  </div>
</template>

<script setup>
import { useContentStore } from "../src/store/content";
import { useConfigStore } from "../src/store/config";
import { SaveFiles, ReadFile, GetConfig } from "../wailsjs/go/main/App";
import log from "../src/log/log";
import { sha256 } from "js-sha256";
import { onBeforeUnmount } from "vue";

const contentStore = useContentStore();
const configStore = useConfigStore();

(async () => {
  try {
    const config = await GetConfig();
    log("get config", config);
    configStore.$patch((state) => {
      state.config = config;
    });
  } catch (err) {
    log("get config failed", err);
  }
})();

// 监听打开文件操作;
window.runtime.EventsOn("opendir", async (fileTree) => {
  let set = false;

  // // 判断是否是第一次打开
  if (contentStore.fileTree.name) {
    // 判断是否有未保存的
    const arr = checkUnSave(contentStore.fileTree);
    if (arr.length > 0) {
      // 提示是否保存
      log("保存", arr);
      SaveFiles(arr).then(() => {
        log("保存成功");
      });
    }

    // 如果打开的是一个文件 保存的文件中包含该打开的文件 那么把对应的内容赋值给新打开的文件
    if (fileTree.type === "file") {
      const file = arr.find((item) => {
        if (item.totalPath === fileTree.totalPath) {
          return true;
        }
      });
      if (file) {
        fileTree.content = file.content;
        set = true;
      }
    }

    //  重置contentStore
    contentStore.$patch((state) => {
      state.content = "";
      state.totalPath = "";
      state.fileTree = {};
    });
  }
  // 如果打开的是文件且没有拿到文件内容
  if (fileTree.type === "file" && !set) {
    try {
      const result = await ReadFile(fileTree.totalPath);
      fileTree.content = result;
    } catch (err) {
      log("readfile err", err);
    }
  }

  // 存储fileTree
  contentStore.$patch((state) => {
    state.fileTree = fileTree;
  });
  // 如果是文件直接展示内容
  if (fileTree.type === "file") {
    contentStore.updateOnlyFileHash(fileTree.totalPath, fileTree.content);
    contentStore.swithContent(fileTree.totalPath, fileTree.content);
  }
});

onBeforeUnmount(() => {
  window.runtime.EventsOff("opendir");
  window.runtime.EventsOff("log");
});

function checkUnSave(ft) {
  let arr = [];
  const getUnSaveFiles = (ft) => {
    if (ft.type === "file") {
      if (ft.contentHash && ft.contentHash !== sha256(ft.content)) {
        arr.push({ totalPath: ft.totalPath, content: ft.content });
      }
    }
    for (var i = 0; i < ft.children?.length; i++) {
      getUnSaveFiles(ft.children[i]);
    }
  };
  getUnSaveFiles(ft);
  return arr;
}
</script>

<style>
#main {
  display: flex;
  flex-direction: column;
}
#header {
  height: 50px;
  background-color: white;
}
</style>
