<template>
  <v-app>
    <v-overlay
        location-strategy="connected"
        scroll-strategy="block"
        close-delay="1000"
        class="align-center justify-center"
        v-model="isDisconnected"
        o
    >
      <v-snackbar
          v-model="isDisconnected"
          multi-line
          variant="elevated"
          location="center"
          color="error"
      >
        <v-icon>mdi-wifi-alert</v-icon>
        Connection lost.
        <br>
        <br>
        Please check your Wi-Fi connection and try again.
        <br>
        <br>
        Alternatively, you can reload the page to reconnect.
        <template v-slot:actions>
          <v-btn
              color="green"
              variant="elevated"
              @click="() => (reloadPage())"
          >
            Reload
            <template v-slot:prepend>
              <v-icon>mdi-reload</v-icon>
            </template>
          </v-btn>
        </template>
      </v-snackbar>
    </v-overlay>
    <v-card class="overflow-hidden">
      <v-app-bar
          color="deep-purple-darken-1"
          density="compact"
      >
        <v-app-bar-nav-icon @click="drawer = true"></v-app-bar-nav-icon>

        <v-toolbar-title>Train Commander</v-toolbar-title>

        <v-btn
            :variant="started ? 'flat': 'outlined'"
            :color="started ? 'red-accent-2': ''"
            :disabled="!started"
            @click="
            () => {
              stopAll();
            }
          "
        >
          STOP ALL
        </v-btn>

        <template v-slot:extension>
          <v-tabs v-model="tab" align-with-title show-arrows>
            <v-tab v-for="item in ordered(getEnabledLocomotives)" :key="item.address" :value="item.name"
                   prepend-icon="mdi-tram"
                   :append-icon="item.speed > 0 ? (item.direction === 1 ? 'mdi-arrow-right-bold' : 'mdi-arrow-left-bold' ) : 'mdi-alert-octagon'">
              {{ item.name }}
            </v-tab>
            <v-tab v-for="item in ordered(getEnabledRailwayModules)"
                   :key="item.address"
                   :value="item.name"
                   prepend-icon="mdi-routes"
            >
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
          <v-window v-model="tab" class="px-0 py-15">
            <v-window-item
                v-for="item in ordered(getEnabledLocomotives)"
                :key="item.address"
                :value="item.name"
            >
              <v-row :justify="'center'">
                <v-col md="4" sm="12" xs="12">
                  <train :name="item.name"></train>
                </v-col>
              </v-row>
            </v-window-item>
            <v-window-item
                v-for="item in ordered(getEnabledRailwayModules)"
                :key="item.address"
                :value="item.name"
            >
              <v-row :justify="'center'">
                <v-col md="4" sm="12" xs="12">
                  <railway-module :name="item.name"></railway-module>
                </v-col>
              </v-row>
            </v-window-item>
          </v-window>
          <v-navigation-drawer v-model="drawer" absolute temporary>
            <v-list density="compact">
              <v-list-subheader>TRAINS</v-list-subheader>
              <v-list-item density="compact" v-if="locomotives && Object.keys(locomotives).length > 0"
                           v-for="item in ordered(locomotives)" :key="item.address">
                <template v-slot:title>
                  <v-row>
                    <v-col cols="2" class="my-2">
                      <v-icon icon="mdi-tram"></v-icon>
                    </v-col>
                    <v-col cols="6" class="my-2">
                      {{ item.name }}
                    </v-col>
                    <v-col cols="4" class="mr-0">
                      <v-switch
                          density="compact"
                          :color="'green'"
                          :true-value="true"
                          :false-value="false"
                          v-model="item.enabled"
                          @change="(e) => {update(); item.enabled ? tab=item.name : (Object.keys(getEnabledLocomotives).length > 0 ? tab=Object.values(getEnabledLocomotives)[0].name : tab=null);}"
                      ></v-switch>
                    </v-col>
                  </v-row>
                </template>
              </v-list-item>
              <v-list-subheader>RAILWAY MODULES</v-list-subheader>
              <v-list-item density="compact" v-if="railwayModules && Object.keys(railwayModules).length > 0"
                           v-for="item in ordered(railwayModules)" :key="item.address">
                <template v-slot:title>
                  <v-row>
                    <v-col cols="2" class="my-2">
                      <v-icon icon="mdi-routes"></v-icon>
                    </v-col>
                    <v-col cols="6" class="my-2">
                      {{ item.name }}
                    </v-col>
                    <v-col cols="4" class="mr-0">
                      <v-switch
                          density="compact"
                          :color="'green'"
                          :true-value="true"
                          :false-value="false"
                          v-model="item.enabled"
                          @change="(e) => {update()}"
                      ></v-switch>
                    </v-col>
                  </v-row>
                </template>
              </v-list-item>
              <v-list-subheader>THEME</v-list-subheader>
              <v-list-item density="compact" key="theme">
                <template v-slot:title>
                  <v-row>
                    <v-col cols="2" class="my-2">
                      <v-icon v-if="theme.global.current.value.dark" icon="mdi-weather-night"></v-icon>
                      <v-icon v-if="!theme.global.current.value.dark" icon="mdi-weather-sunny"></v-icon>
                    </v-col>
                    <v-col cols="6" class="my-2">
                      {{theme.global.current.value.dark ? 'Dark':'Light'}} theme
                    </v-col>
                    <v-col cols="4" class="mr-0">
                      <v-switch
                          density="compact"
                          color="yellow-lighten-1"
                          v-model="lightThemeOn"
                          @change="(e) => {toggleTheme()}"
                      ></v-switch>
                    </v-col>
                  </v-row>
                </template>
              </v-list-item>
              <v-list-subheader>SYSTEM MANAGER</v-list-subheader>
              <v-list-item density="compact" @click="reboot" append-icon="mdi-restart">
                <v-list-item-title>Reboot</v-list-item-title>
              </v-list-item>
              <v-list-item density="compact" @click="poweroff" append-icon="mdi-power">
                <v-list-item-title>Shutdown</v-list-item-title>
              </v-list-item>
              <v-list-subheader>INFO</v-list-subheader>
              <v-list-item density="compact" @click="showInfo" append-icon="mdi-information-outline">
                <v-list-item-title>Show info</v-list-item-title>
              </v-list-item>
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
    <v-dialog
        v-model="infoDialog"
        persistent
        width="auto"
    >
      <v-card>
        <v-card-title class="text-h5">
          <v-icon>mdi-information-outline</v-icon>
          Info
        </v-card-title>
        <v-card-text>
          <b>Version: </b> {{ version }}
          <br>
          <br>
          <b>Built at: </b> {{ builtAt }}
          <br>
          <br>
          <b>Author: </b> {{ author }}
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
              color="grey-darken-1"
              variant="outlined"
              @click="infoDialog = false"
          >
            OK
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <v-footer bottom padless>
      <v-container fluid class="px-0 py-10">
        <v-row>
          <v-col>
            <span class="creds">Created and Designed by:</span>
            <a href="https://begoonlab.tech" target="_blank">
              <v-img v-if="lightThemeOn" class="mx-auto" max-width="300" width="60%" contain src="begoon-lab-logo.png"></v-img>
              <v-img v-if="!lightThemeOn" class="mx-auto" max-width="300" width="60%" contain src="begoon-lab-logo-dark-theme.png"></v-img>
            </a>
          </v-col>
        </v-row>
      </v-container>
    </v-footer>
  </v-app>
</template>

<script setup>
import Train from "./Train.vue";
import {onMounted, ref} from "vue";
import {storeToRefs} from "pinia";
import {useControllerStore} from "../store/modules/controller";
import {useWsStore} from "../store/modules/ws";
import { useTheme } from 'vuetify';

const theme = useTheme()
const controller = useControllerStore()
const ws = useWsStore()
const {
  railwayModules,
  locomotives,
  started,
  isDisconnected,
  getEnabledLocomotives,
  getEnabledRailwayModules,
} = storeToRefs(controller)

onMounted(() => {
  ws.connectToWebsocket();
  lightThemeOn.value = localStorage.getItem("train_commander_theme") !== 'dark';
  theme.global.name.value = localStorage.getItem("train_commander_theme") === 'dark' ? 'dark' : 'light'
})

const version = __VERSION__
const builtAt = __BUILT__
const author = __AUTHOR__
const tab = ref(null)
const drawer = ref(false)
const infoDialog = ref(false)
const lightThemeOn = ref(false)

function toggleTheme() {
  theme.global.name.value = theme.global.current.value.dark ? 'light' : 'dark'
  lightThemeOn.value = theme.global.name.value === 'light'
  localStorage.setItem("train_commander_theme", theme.global.name.value)
}

function reloadPage() {
  window.location.reload()
}

function stopAll() {
  controller.stopAll()
}

function reboot() {
  controller.setReboot()
}

function poweroff() {
  controller.setPowerOff()
}

function update() {
  controller.sendDataToServer()
}

function showInfo() {
  infoDialog.value = true;
}

function ordered(unordered) {
  return Object.keys(unordered).sort().reduce(
      (obj, key) => {
        obj[key] = unordered[key];
        return obj;
      },
      {}
  );
}
</script>