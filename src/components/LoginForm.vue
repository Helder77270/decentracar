<template>
  <v-card elevation="2">
    <h1>Log into Decentracar</h1>
    <v-col md="12">
      <v-row>
        <v-text-field
          label="Login"
          placeholder="Your login"
          :rules="[rules.required]"
        >
        </v-text-field>
      </v-row>

      <v-row>
        <v-text-field
          v-model="password"
          :append-icon="show1 ? 'mdi-eye' : 'mdi-eye-off'"
          :rules="[rules.required, rules.min]"
          :type="show1 ? 'text' : 'password'"
          name="input-10-1"
          label="Password"
          hint="At least 8 characters"
          counter
          @click:append="show1 = !show1"
        ></v-text-field>
      </v-row>

      <router-link :to="{ name: 'Dashboard' }" class="login_button">
        <v-btn color="primary" > Submit</v-btn>
      </router-link>
    </v-col>
  </v-card>
</template>

      
<script>
import { mapActions } from 'vuex';

export default {
  data() {
    return {
      show1: false,
      password: "Password",
      rules: {
        required: (value) => !!value || "Required.",
        min: (v) => v.length >= 8 || "Min 8 characters",
      },
    };
  },
   methods:{
    ...mapActions("setConnected")
  },
};
</script>

<style scoped>
.v-card {
  padding: 20px;
  margin-top: 20px;
}

.v-btn {
  border: solid 1px black;
}
</style>

