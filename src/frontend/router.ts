import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";

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

// router.beforeEach((to, from, next) => {});
