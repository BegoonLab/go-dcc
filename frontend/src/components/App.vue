<template>
  <v-app>
    <v-card class="overflow-hidden">
      <v-app-bar app dense fixed dark flat>
        <v-app-bar-nav-icon @click="drawer = true"></v-app-bar-nav-icon>

        <v-toolbar-title>Train Commander</v-toolbar-title>

        <v-spacer></v-spacer>

        <v-btn
          depressed
          color="error"
          @click="
            () => {
              stopAll();
            }
          "
        >
          STOP ALL
        </v-btn>

        <template v-slot:extension>
          <v-tabs v-model="tab" align-with-title>
            <v-tabs-slider color="yellow"></v-tabs-slider>

            <v-tab v-for="item in locomotives" :key="item.address">
              {{ item.name }}
            </v-tab>
          </v-tabs>
        </template>
      </v-app-bar>
      <v-sheet
        id="scrolling-techniques-3"
        class="overflow-y-auto"
        height="1000px"
      >
        <v-container fluid class="px-0 py-10">
          <v-tabs-items v-model="tab" class="px-0 py-15">
            <v-tab-item v-for="item in locomotives" :key="item.address">
              <v-row :justify="'center'">
                <v-col md="4" sm="12" xs="12">
                  <train :name="item.name"></train>
                </v-col>
              </v-row>
            </v-tab-item>
          </v-tabs-items>
          <v-navigation-drawer v-model="drawer" absolute temporary>
            <v-list nav dense>
              <v-list-item-group
                v-model="group"
                active-class="deep-purple--text text--accent-4"
              >
                <v-list-item>
                  <v-list-item-icon>
                    <v-icon>mdi-home</v-icon>
                  </v-list-item-icon>
                  <v-list-item-title>Home</v-list-item-title>
                </v-list-item>

                <v-list-item>
                  <v-list-item-icon>
                    <v-icon>mdi-account</v-icon>
                  </v-list-item-icon>
                  <v-list-item-title>Account</v-list-item-title>
                </v-list-item>
              </v-list-item-group>
            </v-list>
          </v-navigation-drawer>
        </v-container>
      </v-sheet>
    </v-card>
  </v-app>
</template>

<script>
import Train from "./Train.vue";
import { mapGetters, mapState } from "vuex";
export default {
  components: { Train },

  computed: {
    ...mapState({
      locomotives: (state) => state.controller.locomotives,
    }),
  },

  data() {
    return {
      tab: null,
      drawer: false,
      group: null,
    };
  },

  mounted: function () {},
  methods: {
    stopAll() {
        this.$store.dispatch("controller/stopAll");
    }
  },
};
</script>

<style>
p {
  font-size: 18px;
  font-family: "Roboto", sans-serif;
  color: blue;
}
</style>