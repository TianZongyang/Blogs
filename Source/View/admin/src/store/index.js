import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    isMenuCollapse:true
  },
  mutations: {
    changeMenuCollapseStatus(state){
      state.isMenuCollapse = !state.isMenuCollapse;
    }
  },
  actions: {
  },
  modules: {
  },
  getters:{
  }
})
