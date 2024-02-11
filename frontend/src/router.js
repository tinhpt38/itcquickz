import { createRouter, createWebHashHistory } from "vue-router"; import Login from "./view/Login.vue"; import Landing
from "./view/Landing.vue"; import Quiz from "./view/Quiz.vue"; const router = createRouter({ history:
createWebHashHistory(), routes: [ { path: "/", name: "landing", component: Landing }, { path: "/login/:user", name:
"login", component: Login }, { path: "/quiz/:data", name: "quiz", component: Quiz }, ], }); export default router;
