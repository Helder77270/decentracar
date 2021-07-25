<template>
  <v-app-bar app color="primary" dark>
    <div class="d-flex align-center">
      <v-img
        alt="Vuetify Logo"
        class="shrink mr-2"
        contain
        src="https://cdn.vuetifyjs.com/images/logos/vuetify-logo-dark.png"
        transition="scale-transition"
        width="40"
      />

      <span> Hyperledger project </span>
      <span v-if="isConnected">
        <router-link :to="{ name: 'Admin' }">
          <v-btn text > Admin actions </v-btn>
        </router-link>
      </span>
      <span>{{isConnected}}</span>
    </div>

    <v-spacer></v-spacer>

<div v-if="isConnected">
  <router-link :to="{ name: 'Login' }">
      <v-btn text> Logout </v-btn>
    </router-link>
</div>
    

    <v-btn href="https://github.com/HelderSalvador" target="_blank" text>
      <span class="mr-2">Our Github</span>
      <v-icon>mdi-open-in-new</v-icon>
    </v-btn>
  </v-app-bar>
</template>

<script>
import { mapGetters, mapActions } from "vuex";

export default {
  data() {
    return {
      items: [
        { title: "Dashboard", icon: "mdi-view-dashboard" },
        { title: "Account", icon: "mdi-account-box" },
        { title: "Admin", icon: "mdi-gavel" },
      ],
    };
  },
  methods:{
    ...mapActions("setConnected")
  },
  computed:{
    ...mapGetters("isConnected")
  }
};
</script>