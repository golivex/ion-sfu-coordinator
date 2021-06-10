import Stream from '@/api/stream.api';

const state = {
    liveStreamStatus: localStorage.getItem("liveStream") || false
}


const getters = {
    getliveStreamingStatus(state) {
        return state.liveStreamStatus;
    }
}

const actions = {
    async startLiveStream({ commit }, payload) {
        const session = payload
        try {
            const res = await Stream.startLiveStream(session);
            const { status } = res
            if (status >= 200 && status < 400) {
                localStorage.setItem('liveStream', session)
                commit('SET_LIVE_STREAM_STATUS', true)
            }
            return res
        } catch (error) {
            console.log('%c --error[startliveStream]', 'color: #619c1d', error);
        }
    },


    async stopLiveStream({ commit }) {
        const session = localStorage.getItem("liveStream")
        try {
            const res = await Stream.stopLiveStream(session);
            const { status } = res
            if (status >= 200 && status < 400) {
                localStorage.removeItem('liveStream')
                commit('SET_LIVE_STREAM_STATUS', false)
            }
            return res
        } catch (error) {
            console.log('%c --error[stopLiveStream]', 'color: #cc0088', error);
        }
    },
    async demoLiveStream({ commit }, payload) {
        const session = payload
        try {
            const res = await Stream.demoLiveStream(session);
            console.log(res,"demoLiveStream")
            return res
        } catch (error) {
            console.log('%c --error[demoLiveStream]', 'color: #cc0088', error);
        }
    }
}


const mutations = {
    SET_LIVE_STREAM_STATUS(state, data) {
        state.liveStreamStatus = data;
    }
}
export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}