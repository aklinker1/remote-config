import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import { getAuthToken } from "./state/auth-token";

const routes: RouteRecordRaw[] = [
  {
    path: "/",
    component: () => import("./components/Dashboard.vue"),
    children: [
      {
        path: "",
        component: () => import("./components/AppList.vue"),
      },
      {
        path: ":app",
        component: () => import("./components/EditApp.vue"),
        meta: {
          requiresAuth: true,
        },
      },
    ],
  },
  {
    path: "/login",
    component: () => import("./components/Login.vue"),
  },
];

export const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach(async (to, from, next) => {
  if (!getAuthToken() && to.meta.requiresAuth) {
    router.replace(`/login?redirect=${to.fullPath}`);
    return;
  }
  next();
});
