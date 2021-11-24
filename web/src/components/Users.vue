<template>
  <v-col cols="12">
    <v-simple-table v-if="!isGettingUsers" style="width: 100%">
      <template v-slot:default>
        <thead>
        <tr>
          <th class="text-left">Фамилия</th>
          <th class="text-left">Имя</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="u in users" :key="u.id" @click="editUser(u)">
          <td>{{ u.surname }}</td>
          <td>{{ u.name }}</td>
        </tr>
        </tbody>
      </template>
    </v-simple-table>

    <v-layout v-if="isGettingUsers" justify-center>
      <v-progress-circular
        indeterminate
        :size="50"
        color="primary"
      ></v-progress-circular>
    </v-layout>

    <v-dialog
      v-model="dialog"
      persistent
      max-width="600px"
    >
      <v-card>
        <v-card-text>
          <v-container>
            <v-row>
              <v-col cols="12">
                <v-text-field
                  label="Фамилия"
                  :value="dUser.surname"
                  readonly
                />
                <v-text-field
                  label="Имя"
                  :value="dUser.name"
                  readonly
                />
                <v-checkbox
                  label="Пришел на мероприятие"
                  v-model="dUser.checkin"
                />
                <v-radio-group
                  v-model="dUser.covidPass"
                >
                  <v-radio
                    label="ПЦР"
                    value="ПЦР"
                  />
                  <v-radio
                    label="QR"
                    value="QR"
                  />
                  <v-radio
                    label="Экспресс"
                    value="Экспресс"
                  />
                  <v-radio
                    label="Антитела"
                    value="Антитела"
                  />
                </v-radio-group>
              </v-col>
            </v-row>
          </v-container>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
            color="blue darken-1"
            text
            :loading="isUpdating"
            @click="cancelEdit"
          >
            Отменить
          </v-btn>
          <v-btn
            color="blue darken-1"
            text
            :loading="isUpdating"
            @click="editUserRequest"
          >
            Сохранить
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-col>
</template>

<script>
  export default {
    name: 'Users',

    data: () => ({
      dialog: false,
      dUser: {},
    }),

    computed: {
      users() {
        return this.$store.state.users;
      },
      isGettingUsers() {
        return this.$store.state.isSearching;
      },
      isUpdating() {
        return this.$store.state.isUpdating;
      },
    },
    methods: {
      editUser(user) {
        this.dUser = _.cloneDeep(user);
        this.dialog = true;
      },
      cancelEdit() {
        this.dUser = {};
        this.dialog = false;
      },
      async editUserRequest() {
        const result = await this.$store.dispatch('updateUserById', this.dUser);
        if (result) {
          this.dialog = false;
          this.$emit('clear-and-focus-search')
        }
      }
    }
  }
</script>
