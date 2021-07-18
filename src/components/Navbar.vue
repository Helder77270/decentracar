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
      <span>
        <router-link :to="{ name: 'Admin' }">
          <v-btn text > Admin actions </v-btn>
        </router-link>
      </span>
    </div>

    <v-spacer></v-spacer>

    <router-link :to="{ name: 'Login' }">
      <v-btn text> Logout </v-btn>
    </router-link>

    <v-btn href="https://github.com/HelderSalvador" target="_blank" text>
      <span class="mr-2">Our Github</span>
      <v-icon>mdi-open-in-new</v-icon>
    </v-btn>
  </v-app-bar>
</template>

<script>
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
};
</script>