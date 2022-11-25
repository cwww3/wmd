<template>
  <el-form :model="form" label-width="120px">
    <el-form-item label="开启图传">
      <el-switch v-model="form.enable" />
    </el-form-item>

    <div v-show="form.enable">
      <el-form-item label="图床类型">
        <el-select
          v-model="picgo.picBed.uploader"
          placeholder="please select your type"
        >
          <el-option label="无" value="" />
          <el-option label="github" value="github" />
        </el-select>
      </el-form-item>
      <div v-show="picgo.picBed.uploader === 'github'">
        <el-form-item label="token">
          <el-input v-model="picgo.picBed.github.token" show-password />
        </el-form-item>
        <el-form-item label="路径">
          <el-input v-model="picgo.picBed.github.path" />
        </el-form-item>
        <el-form-item label="仓库">
          <el-input v-model="picgo.picBed.github.repo" />
        </el-form-item>
        <el-form-item label="分支">
          <el-input v-model="picgo.picBed.github.branch" />
        </el-form-item>
      </div>
    </div>

    <el-form-item label="图片存储路径">
      <el-input v-model="form.localImagePath" :disabled="true" />
      <el-button type="primary" @click="onSelect">选择目录</el-button>
    </el-form-item>

    <el-form-item>
      <el-button type="primary" @click="onSubmit">保存</el-button>
    </el-form-item>
  </el-form>
</template>

<script setup>
import { SaveConfig, GetDir } from "../../wailsjs/go/main/App";
import { useConfigStore } from "../store/config";
import log from "../log/log";

const configStore = useConfigStore();

const form = configStore.config;
const picgo = form.picgo;

const onSelect = async () => {
  try {
    const dir = await GetDir();
    configStore.$patch((state) => {
      state.config.localImagePath = dir;
    });
  } catch (err) {
    log("select dir failed", err);
  }
};

const onSubmit = async () => {
  await SaveConfig(form);
};
</script>
