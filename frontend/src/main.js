import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import axios from 'axios';

Vue.config.productionTip = false;
Vue.prototype.$http = axios
Vue.prototype.$api = 'http://localhost:7777/v1/'

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount("#app");
