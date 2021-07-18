<template>
  <v-app>
    <v-main>
      <Navbar/>
      <v-col md="12">
        <v-row class="justify-center">
           <Dashboard/>
         
        </v-row>
      </v-col>
    </v-main>
  </v-app>
</template>

<script>
import Navbar from '../components/Navbar';
// import Login from './components/Login';
import Dashboard from '../components/Dashboard';

export default {
  name: 'App',

  components: {
    Navbar,
    // Login
    Dashboard
  },

  data: () => ({
    //
  }),
};
</script>
