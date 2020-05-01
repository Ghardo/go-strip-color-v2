import Vue from 'vue'
import VueRouter from 'vue-router'

import ViewDashboard from "@/components/view/dashboard.vue"
import ViewSettings from "@/components/view/settings.vue"

Vue.use(VueRouter)

const router = new VueRouter({
    routes: [
        {   
            path: '/', 
            name: 'dashboard',
            component: ViewDashboard 
        },
        {
            path: '/settings',
            name: 'settings',
            component: ViewSettings
        }
    ]
})

export default router