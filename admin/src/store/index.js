import Vue from 'vue'
import Vuex from 'vuex'
import Stats from './modules/stats.store';
import Load from './modules/load.store';
import Disk from './modules/disk.store';
import Stream from './modules/stream.store';
import Rtmp from './modules/rtmp.store';

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
    Load,
    Disk,
    Stream,
    Rtmp
  }
})
