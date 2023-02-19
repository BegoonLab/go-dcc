<template>
  <v-app>
    <v-overlay
        :absolute="false"
        :value="!connected"
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
      <v-app-bar
           color="deep-purple-darken-1"
           density="compact"
      >
        <v-app-bar-nav-icon @click="drawer = true"></v-app-bar-nav-icon>

        <v-toolbar-title>Train Commander</v-toolbar-title>

        <v-btn
            variant="flat"
            color="red-accent-2"
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
            <v-tab v-for="item in ordered(getEnabledLocomotives)" :key="item.address" :value="item.name" :append-icon="item.speed > 0 ? (item.direction === 1 ? 'mdi-arrow-right-bold' : 'mdi-arrow-left-bold' ) : 'mdi-alert-octagon'">
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
          </v-window>
          <v-navigation-drawer v-model="drawer" absolute temporary>
            <v-list density="compact">
              <v-list-subheader>TRAINS</v-list-subheader>
              <v-list-item density="compact" v-if="locomotives && Object.keys(locomotives).length > 0" v-for="item in ordered(locomotives)" :key="item.address">
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
              <v-list-subheader>SYSTEM MANAGER</v-list-subheader>
              <v-list-item density="compact" @click="reboot" append-icon="mdi-restart">
                <v-list-item-title>Reboot</v-list-item-title>
              </v-list-item>
              <v-list-item density="compact" @click="poweroff" append-icon="mdi-power">
                <v-list-item-title>Shutdown</v-list-item-title>
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

<script setup>
import Train from "./Train.vue";
import {computed, reactive, onMounted, ref, watch} from "vue";
import {storeToRefs} from "pinia";
import {useControllerStore} from "../store/modules/controller";
import {useWsStore} from "../store/modules/ws";

const controller = useControllerStore()
const ws = useWsStore()
// const getEnabledLocomotives = controller.getEnabledLocomotives
const { locomotives, started, connected, getEnabledLocomotives } = storeToRefs(controller)

onMounted(() => {
  ws.connectToWebsocket();
})

const tab = ref(null)
const drawer = ref(false)
const group = ref(null)
const overlay = ref(true)

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

function ordered (unordered) {
  return Object.keys(unordered).sort().reduce(
      (obj, key) => {
        obj[key] = unordered[key];
        return obj;
      },
      {}
  );
}
</script>