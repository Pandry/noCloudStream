import { createStore } from 'vuex'

export const store = createStore({
    state () {
      return {
        user: {
            name: "",
            loggedIn: false,
            token: {
                key: "",
                expiry: ""
            }
        }
      }
    },
    mutations: { //mutate the data - sync
      login (st,v) {
        //state.user = s.user
        st.user.loggedIn = v
      }
    },
    actions:{}, // dispatch async actions to commit mutations
    modules:{}, // separate modules for sub-stores
    getters: {
        userIsLoggedIn(state){
            return state.user.loggedIn
        }
    }
  })