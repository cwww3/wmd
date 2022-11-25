import { createWebHistory, createRouter } from "vue-router";
import Edit from "../views/Edit.vue";
import Config from "../views/Config.vue";

const routes = [
  {
    path: "/",
    name: "Edit",
    component: Edit,
  },
  {
    path: "/config",
    name: "Config",
    component: Config,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
