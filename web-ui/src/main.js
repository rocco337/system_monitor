import Vue from 'vue'
import App from './App.vue'
import './quasar'
import { Dark } from 'quasar'

// set status
Dark.set(true) // or false or "auto"

Vue.config.productionTip = false

new Vue({
  render: h => h(App),
}).$mount('#app')
