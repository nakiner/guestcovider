<template>
  <v-container fluid style="height: 100vh;" >
    <v-layout justify-center>
      <v-col cols="10">
        <v-text-field
          ref="search"
          v-model="surname"
          label="Фамилия"
          autofocus
          @input="searchUsersRequest"
        />
      </v-col>
    </v-layout>
    <v-layout fill-height>
      <users @clear-and-focus-search="clearAndFocusSearch"></users>
    </v-layout>
  </v-container>
</template>

<script>

import Users from '../components/Users';

export default {
  name: 'Main',

  data: () => ({
    zIndex: 0,
    overlay: false,
    dark: false,
    currentItemPost: null,
    surname: '',
  }),

  components: {
    'users': Users,
  },

  methods: {
    searchUsersRequest: _.debounce(function () {
      this.$store.dispatch('searchUsers', this.surname);
    }, 200),
    clearAndFocusSearch() {
      this.surname = '';
      this.$nextTick(() => this.$refs.search.focus());
    }
  }
}
</script>
