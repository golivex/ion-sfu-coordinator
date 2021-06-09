import Rtmp from '@/api/rtmp.api';

const state = {
    liveRtmpStatus: localStorage.getItem("liveRtmp") || false
}


const getters = {
    getliveRtmpStatus(state) {
        return state.liveRtmpStatus;
    }
}


const actions = {
    async startLiveRtmp({ commit }, payload) {
        const session = payload
        try {
            const res = await Rtmp.startLiveRtmp(session);
            const { status } = res
            if (status >= 200 && status < 400) {
                localStorage.setItem('liveRtmp', session)
                commit('SET_LIVE_RTMP_STATUS', true)
            }
            return res
        } catch (error) {
            console.log('%c --error[startliveRtmp]', 'color: #619c1d', error);
        }
    },


    async stopLiveRtmp({ commit }) {
        const session = localStorage.getItem("liveRtmp")
        try {
            const res = await Rtmp.stopLiveRtmp(session);
            const { status } = res
            if (status >= 200 && status < 400) {
                localStorage.removeItem('liveRtmp')
                commit('SET_LIVE_RTMP_STATUS', false)
            }
            return res
        } catch (error) {
            console.log('%c --error[stopliveRtmp]', 'color: #cc0088', error);
        }
    },
    async demoLiveRtmp({ commit }, payload) {
        const session = payload 
        try {
            const res = await Rtmp.demoLiveRtmp(session);
            console.log(res,"demoLiveRtmp")
            return res
        } catch (error) {
            console.log('%c --error[demoLiveRtmp]', 'color: #cc0088', error);
        }
    }
}


const mutations = {
    SET_LIVE_RTMP_STATUS(state, data) {
        state.LiveRtmpStatus = data;
    }
}
export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}