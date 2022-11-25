<template>
  <div id="content">
    <md-editor
      v-if="contentStore.totalPath.length > 0"
      v-model="contentStore.content"
      @onSave="onSave"
      @onChange="onChange"
      @onUploadImg="onUploadImg"
    />
    <div v-else id="empty"></div>
  </div>
</template>

<script setup>
import MdEditor from "md-editor-v3";
import "md-editor-v3/lib/style.css";
import { useContentStore } from "../store/content";
import { useConfigStore } from "../store/config";
import bus from "../eventbus/index";
import { UploadFiles } from "../../wailsjs/go/main/App";
import log from "../log/log";

const contentStore = useContentStore();
const configStore = useConfigStore();
// 保存
const onSave = function (content) {
  bus.emit("onSave", { totalPath: contentStore.totalPath, content: content });
};

// 更新
const onChange = function (content) {
  bus.emit("onChange", { totalPath: contentStore.totalPath, content: content });
};

const onUploadImg = async (files, callback) => {
  try {
    var reader = new FileReader();
    var fileByteArray = [];
    reader.readAsArrayBuffer(files[0]);
    reader.onloadend = function (evt) {
      if (evt?.target?.readyState == FileReader.DONE) {
        const arrayBuffer = evt.target.result;
        const array = new Uint8Array(arrayBuffer);
        for (var i = 0; i < array.length; i++) {
          fileByteArray.push(array[i]);
        }
        UploadFiles({ content: fileByteArray }).then((url) => {
          callback([url]);
        });
      }
    };
  } catch (err) {
    log("picgo upload failed", err);
  }
};
</script>

<style scoped>
#content {
  flex: 1;
}
.md {
  height: calc(100vh - 50px);
}
#empty {
  height: calc(100vh - 50px);
  width: 100%;
  background: url("../assets/images/logo-universal.png") no-repeat;
  background-size: 100% 100%;
}
</style>
