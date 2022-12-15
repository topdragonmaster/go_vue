import request from 'axios'

request.defaults.baseURL = 'http://localhost/8080'

export const login = ({ commit, state }) => {
  return request.post('login').then((response) => {
    if (response.statusText === 'OK') {
      commit('UPDATE_USERINFO', response.data)
    }
  }).catch((error) => {
    console.log(error)
  })
}