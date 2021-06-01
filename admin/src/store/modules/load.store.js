import Load from '@/api/load.api';
import Vue from 'vue';

const state = {
    loadTestingStatus: localStorage.getItem("loadTest") || false,
    loadStatus: []
}


const getters = {
    getLoadTestingStatus(state) {
        return state.loadTestingStatus;
    },

    getLoadStatus(state) {
        return state.loadStatus;
    }
}


const actions = {
    async startLoadTest({ commit }, payload) {
        const { sessionName, file, clients, role, cycle, rooms } = payload

        try {
            const res = await Load.startLoadTest(file, clients, role, cycle, rooms);
            console.log('%c --res[startLoadTest]', 'color: #bfffc8', res);
            const { status, data } = res
            if (status >= 200 && status < 400) {
                localStorage.setItem('loadTest', data)
                commit('SET_LOAD_TESTING_STATUS', true)
            }
            return res
        } catch (error) {
            console.log('%c --error[startLoadTest]', 'color: #619c1d', error);
        }
    },


    async stopLoadTest({ commit }) {
        try {
            const res = await Load.stopLoadTest();
            console.log('%c --res[stopLoadTest]', 'color: #1d3f73', res);
            const { status } = res
            if (status >= 200 && status < 400) {
                localStorage.removeItem('loadTest')
                commit('SET_LOAD_TESTING_STATUS', false)
                commit('SET_LOAD_STATUS', [])
            }
            return res
        } catch (error) {
            console.log('%c --error[stopLoadTest]', 'color: #cc0088', error);
        }
    },


    async fetchLoadStats({ commit }) {
        try {
            const res = await Load.loadStats();
            console.log('%c --res[fetchLoadStats]', 'color: #735656', res);
            const { Response } = res.data || []
            commit('SET_LOAD_STATUS', Response)
        } catch (error) {
            console.log('%c --error[fetchLoadStats]', 'color: #ffcc00', error);
        }
    }
}


const mutations = {
    SET_LOAD_TESTING_STATUS(state, data) {
        state.loadTestingStatus = data;
    },

    SET_LOAD_STATUS(state, data) {
        Vue.set(state, 'loadStatus', data)
    }
}
export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}