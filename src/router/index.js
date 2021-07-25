import VueRouter from 'vue-router';
import Vue from 'vue';

Vue.use(VueRouter);
const routes = [
     {
         path: "/",
         name: "Login",
         component: () => import('@/views/Login'),
         props: true
     },
     {
        path: "/dashboard",
        name: "Dashboard",
        component: () => import('@/views/Dashboard'),
        props: true
     },
     {
        path: "/admin_actions",
        name: "Admin",
        component: () => import('@/views/Admin'),
        props: true
     },
];

const router = new VueRouter({
    mode: "history",
    base: "/Decentracar",
    routes
});

export default router;


