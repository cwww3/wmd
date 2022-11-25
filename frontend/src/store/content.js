import { defineStore } from "pinia";
import { sha256 } from "js-sha256";
import log from "../log/log";

export const useContentStore = defineStore("content", {
  state: () => {
    return {
      fileTree: {},
      totalPath: "",
      content: "",
    };
  },
  getters: {},
  actions: {
    collapseDir(path) {
      const fileTree = findFileTree(this.fileTree, path);
      if (fileTree) {
        if (fileTree.isShow) {
          const f = function (fileTree) {
            if (fileTree.type === "dir") {
              fileTree.isShow = false;
            }
            fileTree.children?.forEach((child) => {
              f(child, path);
            });
          };
          f(fileTree);
        } else {
          fileTree.isShow = !fileTree.isShow;
        }
      } else {
        log("not found", this.fileTree, path);
      }
    },
    updateFileContent(path, content) {
      const fileTree = findFileTree(this.fileTree, path);
      if (fileTree) {
        fileTree.content = content;
        fileTree.contentHash = sha256(content);
      } else {
        log("not found", this.fileTree, path);
      }
    },
    updateOnlyFileContent(path, content) {
      const fileTree = findFileTree(this.fileTree, path);
      if (fileTree) {
        fileTree.content = content;
      } else {
        log("not found", this.fileTree, path);
      }
    },
    updateOnlyFileHash(path, content) {
      const fileTree = findFileTree(this.fileTree, path);
      if (fileTree) {
        fileTree.contentHash = sha256(content);
      } else {
        log("not found", this.fileTree, path);
      }
    },
    swithContent(path, content) {
      this.totalPath = path;
      this.content = content;
    },
  },
});

function findFileTree(fileTree, path) {
  if (fileTree.totalPath === path) {
    return fileTree;
  }
  if (fileTree.type === "dir") {
    let g;
    for (let index = 0; index < fileTree.children?.length; index++) {
      const child = fileTree.children[index];
      const f = findFileTree(child, path);
      if (f) {
        g = f;
      }
    }
    return g;
  }
  return undefined;
}
