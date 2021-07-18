<template>
  <v-col md="12">
    <v-row>
      <v-col md="4">
        <v-card elevation="2">
          <v-card-text>
            <h1>Chiffres de ventes</h1>

            <div class="text-primary">Total : 21,8 millions</div>
          </v-card-text>
        </v-card>
      </v-col>

      <v-col md="4">
        <v-card elevation="2">
          <v-card-text>
            <h1>Chiffre d'affaire</h1>
            <p>Total : 18,9 millions</p>
          </v-card-text>
        </v-card>
      </v-col>

      <v-col md="4">
        <v-card elevation="2">
          <v-card-text>
            <h1>Cout de réparation</h1>
            <div class="text-primary">Total : 2,9 millions</div>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>

    <v-row>
      <v-col md="12">
        <v-list subheader>
          <v-subheader>Listes des voitures</v-subheader>

          <v-list-item>
            <v-list-item-avatar>
              <v-img :src="assets.image_path" alt="voiture"></v-img>
            </v-list-item-avatar>

            <v-list-item-content>
              <v-card elevation="2">
                  <v-list-item-title> Je suis un titre</v-list-item-title>
              </v-card>
            </v-list-item-content>

          </v-list-item>
        </v-list>
      </v-col>
    </v-row>
  </v-col>
</template>


<script>
export default {
  name: 'Dashboard',

  components: {

  },

  data: () => ({
    assets: [
        {
            image_path: "@/assets/voiture.svg"
        }
    ]
  }),
};
</script>

<style scoped>
.v-subheader {
  display: flex;
  flex-direction: column;
  align-content: center;
  align-items: center;
  justify-content: center;
}
</style>