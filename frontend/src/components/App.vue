<template>
  <v-app>
    <v-container fluid class="px-0 py-0">
    <v-row :justify="'center'">
      <v-col cols="12" sm="12" md="6" lg="4">
        <v-card>
          <v-toolbar
            color="black"
            dark
            flat
          >
            <v-app-bar-nav-icon></v-app-bar-nav-icon>

            <v-toolbar-title>Command Module</v-toolbar-title>

            <v-spacer></v-spacer>

            <v-btn icon>
              <v-icon>mdi-dots-vertical</v-icon>
            </v-btn>

            <template v-slot:extension>
              <v-tabs
                v-model="tab"
                align-with-title
              >
                <v-tabs-slider color="yellow"></v-tabs-slider>

                <v-tab
                  v-for="item in items"
                  :key="item"
                >
                  {{ item }}
                </v-tab>
              </v-tabs>
            </template>
          </v-toolbar>

          <v-tabs-items v-model="tab">
            <v-tab-item
              v-for="item in items"
              :key="item"
            >
              <v-card flat>
                <v-card-text v-text="text"></v-card-text>
              </v-card>
              <v-row :justify="'center'">
                <v-col>
                  <train></train>
                </v-col>
              </v-row>
            </v-tab-item>
          </v-tabs-items>
        </v-card>
      </v-col>  
    </v-row>  
  </v-container>
  </v-app>
</template>

<script>
import Train from './Train.vue';
export default {
  components: { Train },
  data() {
    return {
      message: 'Hello World! This is command module',
      tab: null,
      items: [
        'Train 1', 'Train 2', 'Train 3', 'Train 4', 'Train 5',
      ],
      text: '',
    };
  },
};
</script>

<style>
p {
  font-size: 18px;
  font-family: 'Roboto', sans-serif;
  color: blue;
}
</style>