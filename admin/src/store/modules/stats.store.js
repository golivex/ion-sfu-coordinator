import Stats from '@/api/stats.api';

const state = {
    stats: {},
}


const getters = {
    getStats(state) {
        return state.stats;
    }
}


const actions = {
    async fetchStats({ commit }) {
        try {
            const res = await Stats.getStats();
            commit('SET_STATS', res.data);
        } catch (error) {
            console.log('%c --error', 'color: #ffa640', error);
        }
    }
}


const mutations = {
    SET_STATS(state, data) {
        state.stats = data;
    }
}
export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}