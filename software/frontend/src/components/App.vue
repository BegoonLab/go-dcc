<template>
  <v-app>
    <v-overlay
        :absolute="false"
        :value="!isConnected"
    >
      <v-row>
        <v-col class="text-center">
          Connection lost&nbsp;&nbsp;
          <v-icon color="red">mdi-connection</v-icon>
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <v-btn
              color="success"
              @click="() => (reloadPage())"
          >
            <v-icon>mdi-reload</v-icon>&nbsp;&nbsp;Reload page
          </v-btn>
        </v-col>
      </v-row>
    </v-overlay>
    <v-card class="overflow-hidden">
      <v-app-bar app dense fixed dark flat>
        <v-app-bar-nav-icon @click="drawer = true"></v-app-bar-nav-icon>

        <v-toolbar-title>Train Commander</v-toolbar-title>

        <v-spacer></v-spacer>

        <v-btn
            depressed
            color="error"
            :disabled="!isStarted"
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

            <v-tab v-for="item in getEnabledLocomotives" :key="item.address">
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
        <v-container fluid class="px-0 py-10 overflow-x-hidden">
          <v-tabs-items v-model="tab" class="px-0 py-15">
            <v-tab-item
                v-for="item in getEnabledLocomotives"
                :key="item.address"
            >
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

                <v-list-item v-for="item in locomotives" :key="item.address">
                  <v-list-item-action>
                    <v-switch
                        :color="'green'"
                        :true-value="true"
                        :false-value="false"
                        :input-value="item.enabled"
                        @change="(v) => update(item.name, v, 'enabled')"
                    ></v-switch>
                  </v-list-item-action>
                  <v-list-item-icon>
                    <v-icon>mdi-tram</v-icon>
                  </v-list-item-icon>
                  <v-list-item-title>{{ item.name }}</v-list-item-title>
                </v-list-item>
                <v-list-item @click="reboot">
                  <v-list-item-icon>
                    <v-icon>mdi-restart</v-icon>
                  </v-list-item-icon>
                  <v-list-item-title>Reboot</v-list-item-title>
                </v-list-item>
                <v-list-item @click="poweroff">
                  <v-list-item-icon>
                    <v-icon>mdi-power</v-icon>
                  </v-list-item-icon>
                  <v-list-item-title>Shutdown</v-list-item-title>
                </v-list-item>
              </v-list-item-group>
            </v-list>
          </v-navigation-drawer>
          <v-row
              v-if="getEnabledLocomotives && Object.keys(getEnabledLocomotives).length === 0 && Object.getPrototypeOf(getEnabledLocomotives) === Object.prototype">
            <v-col>
              <v-row :justify="'center'">
                <v-col md="4" sm="12" xs="12" class="text-center">
                  No active locomotives found
                </v-col>
              </v-row>
              <v-row :justify="'center'">
                <v-col md="4" sm="12" xs="12" class="text-center">
                  <v-btn
                      color="success"
                      @click="() => (drawer = true)"
                  >
                    <v-icon>mdi-tram</v-icon>&nbsp;&nbsp;Activate locomotives
                  </v-btn>
                </v-col>
              </v-row>
            </v-col>
          </v-row>

        </v-container>
      </v-sheet>
    </v-card>
    <v-footer bottom padless>
      <v-container fluid class="px-0 py-10">
        <v-row>
          <v-col>
            <span class="creds">Created and Designed by:</span>
            <v-img class="mx-auto" max-width="300" width="60%" contain src="begoon-lab-logo.png"></v-img>
          </v-col>
        </v-row>
      </v-container>
    </v-footer>
  </v-app>
</template>

<script>
import Train from "./Train.vue";
import {mapGetters, mapState} from "vuex";

export default {
  components: {Train},

  computed: {
    ...mapGetters({getEnabledLocomotives: "controller/getEnabledLocomotives"}),
    ...mapState({
      locomotives: (state) => (state.controller.locomotives),
      isStarted: (state) => (state.controller.started),
      isConnected: (state) => (state.controller.connected),
    })
  },

  data() {
    return {
      tab: null,
      drawer: false,
      group: null,
      overlay: true,
    };
  },

  mounted: function () {
  },
  methods: {
    reloadPage() {
      window.location.reload()
    },
    stopAll() {
      this.$store.dispatch("controller/stopAll");
    },
    reboot() {
      this.$store.dispatch("controller/reboot");
    },
    poweroff() {
      this.$store.dispatch("controller/poweroff");
    },
    update(name, value, where) {
      this.$store.dispatch("controller/setLocomotiveState", {
        name: name,
        value: value,
        where: where,
      });
    },
  },
};
</script>

<style>
p {
  font-size: 18px;
  font-family: "Roboto", sans-serif;
  color: blue;
}
.container {
  background: #FFF;
}
footer {
  font-size: 8pt;
}
.creds {
  padding-left: 29%;
}
</style>