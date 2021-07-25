import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import router from './router';
import Vuex from 'vuex';
import store from '@/store';

Vue.config.productionTip = false
Vue.use(Vuex);

const vuex = new Vuex.Store({...store});

new Vue({
  vuetify,
  render: h => h(App),
  router,
  store: vuex
}).$mount('#app')
