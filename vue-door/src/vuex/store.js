import { createStore } from 'vuex'
import * as actions from './actions'
import * as getters from './getters'
const defaultState = {
  userInfo: [],
}

const inBrowser = typeof window !== 'undefined'

// if in browser, use pre-fetched state injected by SSR
const state = (inBrowser && window.__INITIAL_STATE__) || defaultState

const mutations = {
  UPDATE_USERINFO: (state, userInfo) => {
    state.userInfo = userInfo
  },
}

export default createStore({
  state,
  actions,
  mutations,
  getters
})
