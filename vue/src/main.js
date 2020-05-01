import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify';
import astor from './plugins/astor';
import store from './store.js'
import router from './router.js'

Vue.config.productionTip = false

Vue.use(
  astor,  {
    debug: false
  }
)

document.addEventListener('astilectron-ready', function() {
  new Vue({
    router,
    vuetify,
    store,
    render: h => h(App)
  }).$mount('#app')
})

