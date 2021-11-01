export const state = () => ({
    roomId: '',
    userId: '',
    ws: {}
})

export const mutations = {
    setConfig(state, config) {
        state.userId = config.userId
        state.roomId = config.roomId
        state.ws = new WebSocket(`ws://${window.location.host}/backend/api/chat/${state.roomId}`)
    },
}

export const actions = {
    setConfig({ commit }, config) {
        commit('setConfig', config)
    }
}

export const getters = {
    getConfig(state) {
        return state
    },
    getRoomId(state) {
        return state.roomId
    },
    getUserId(state) {
        return state.userId
    },
    getWebSocket(state) {
        return state.ws
    }
}