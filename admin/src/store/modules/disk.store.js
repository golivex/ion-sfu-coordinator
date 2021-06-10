import Disk from '@/api/disk.api';

const state = {
    diskSavingStatus: localStorage.getItem("diskSave") || false
}


const getters = {
    getdiskSavingStatus(state) {
        return state.diskSavingStatus;
    }
}


const actions = {
    async startDiskSave({ commit }, payload) {
        const { session, filename } = payload
        try {
            const res = await Disk.startDiskSave(session, filename);
            const { status } = res
            if (status >= 200 && status < 400) {
                localStorage.setItem('diskSave', session)
                commit('SET_DISK_SAVING_STATUS', true)
            }
            return res
        } catch (error) {
            console.log('%c --error[startDiskSave]', 'color: #619c1d', error);
        }
    },


    async stopDiskSave({ commit }) {
        const session = localStorage.getItem("diskSave")
        try {
            const res = await Disk.stopDiskSave(session);
            const { status } = res
            if (status >= 200 && status < 400) {
                localStorage.removeItem('diskSave')
                commit('SET_DISK_SAVING_STATUS', false)
            }
            return res
        } catch (error) {
            console.log('%c --error[stopDiskSave]', 'color: #cc0088', error);
        }
    }
}


const mutations = {
    SET_DISK_SAVING_STATUS(state, data) {
        state.diskSavingStatus = data;
    }
}
export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}