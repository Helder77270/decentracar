<template>
  <v-app>
    <v-main>
      <Navbar/>
      <v-col md="12">
        <v-row class="justify-center">
          <!-- <router-view/> -->
           <router-view/>
           <!-- <Dashboard/> -->
          <!-- <AdminActions/> -->
        </v-row>
      </v-col>
    </v-main>
  </v-app>
</template>

<script>
import Navbar from './components/Navbar';
// import Login from './components/Login';
// import Dashboard from './components/Dashboard';
// import AdminActions from './components/AdminActions';

export default {
  name: 'App',

  components: {
    Navbar,
    // AdminActions,
    // Login,
    // Dashboard
  },

  data: () => ({
    //
  }),
};
</script>
