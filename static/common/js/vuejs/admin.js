import Vue from 'vue/dist/vue'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-default/index.css'
import login from './admin/login/login.vue'
Vue.use(ElementUI)
new Vue({
	el:'#app',
	components:{
		'xh-login':login,
	}
})
