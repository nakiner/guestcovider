import Vue from 'vue'
import Vuex from 'vuex'
import axiosInst from '../api';
import axios from 'axios';
import { SEARCH_USERS, UPDATE_USER } from '../api/routes';

const CancelToken = axios.CancelToken;
const source = CancelToken.source();

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    users: [],
    isSearching: false,
    isUpdating: false,
  },
  mutations: {
    setUsers(state, users) {
      state.users = users ? [...users] : [];
    },
    setIsSearching(state, payload) {
      state.isSearching = payload;
    },
    setIsUpdating(state, payload) {
      state.isUpdating = payload;
    },
    updateUser(state, user) {
      state.users = state.users.map((u) => {
        return u.id === user.id ? user : u;
      })
    }
  },
  actions: {
    async searchUsers({commit}, surname) {
      try {
        commit('setIsSearching', true);
        const res = await axiosInst.get(SEARCH_USERS(surname), {
          cancelToken: source.token
        });
        if (res) {
          commit('setUsers', res.data.data);
          commit('setIsSearching', false);
        }
      } catch (e) {
        if (axios.isCancel(e)) {
          console.log('Request canceled', e.message);
        } else {
          console.error(e);
        }
      }
    },
    async updateUserById({commit}, user) {
      try {
        commit('setIsUpdating', true);
        await axiosInst.put(UPDATE_USER(), {
          id: user.id,
          data: {
            covidPass: user.covidPass,
            checkin: user.checkin,
          }
        });
        commit('updateUser', user);
        return true;
      } catch(e) {
        console.error(e);
      } finally {
        commit('setIsUpdating', false);
      }
    },
  },
  modules: {
  }
})
