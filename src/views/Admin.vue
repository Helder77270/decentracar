<template>
  <v-app>
    <v-main>
      <Navbar />
      <v-row class="mb-5 ml-10">
          <router-link :to="{ name: 'Dashboard' }">
            <v-btn color="primary"> <v-icon>mdi-chevron-left</v-icon> Revenir au panneau principal</v-btn>
          </router-link>
      </v-row>
      
      <v-col md="12">
        <v-row class="justify-center">
          <AdminActions />
        </v-row>
      </v-col>
    </v-main>
  </v-app>
</template>

<script>
import Navbar from "@/components/Navbar";

import AdminActions from "@/components/AdminActions";

export default {
  name: "App",

  components: {
    Navbar,
    AdminActions,
  },

  data: () => ({
    //
  }),
};
</script>
