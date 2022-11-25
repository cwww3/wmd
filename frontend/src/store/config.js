import { defineStore } from "pinia";

export const useConfigStore = defineStore("config", {
  state: () => {
    return {
      config: {},
    };
  },
  getters: {},
  actions: {},
});
