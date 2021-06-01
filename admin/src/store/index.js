import Vue from 'vue'
import Vuex from 'vuex'
import Stats from './modules/stats.store';
import Load from './modules/load.store';

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
  },
  mutations: {
  },
  actions: {
  },
  modules: {
    Stats,
    Load
  }
})
