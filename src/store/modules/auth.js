const state = {
    isConnected: false,
}

const getters = {
    isConnected: state => {
        return state.isConnected
    },
    
}

const mutations = {
    SET_CONNECTED (state, isConnected) {
        state.isConnected = isConnected
    },
    
}

const actions = {
    setConnected: ({commit}, isConnected) => {
        commit('SET_CONNECTED', isConnected);
    },
    
}

export default {
    state,
    getters,
    mutations,
    actions
}