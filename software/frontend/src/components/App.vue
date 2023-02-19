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
      <v-app-bar app dense fixed dark flat>
        <v-app-bar-nav-icon @click="drawer = true"></v-app-bar-nav-icon>

        <v-toolbar-title>Train Commander</v-toolbar-title>

        <v-spacer></v-spacer>

        <v-btn
            depressed
            color="error"
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
          <v-tabs v-model="tab" align-with-title>
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
          <v-window v-model="tab" class="px-0 py-15">
            <v-window-item
                v-for="item in getEnabledLocomotives"
                :key="item.address"
            >
              <v-row :justify="'center'">
                <v-col md="4" sm="12" xs="12">
                  <train :name="item.name"></train>
                </v-col>
              </v-row>
            </v-window-item>
          </v-window>
          <v-navigation-drawer v-model="drawer" absolute temporary>
            <v-list nav dense>
                <v-list-item v-if="locomotives && Object.keys(locomotives).length > 0" v-for="item in locomotives" :key="item.address" append-icon="mdi-tram">
                  <v-list-item-action>
                    <v-switch
                        :color="'green'"
                        :true-value="true"
                        :false-value="false"
                        v-model="item.enabled"
                        @change="(e) => update()"
                    ></v-switch>
                  </v-list-item-action>
                  <v-list-item-title>{{ item.name }}</v-list-item-title>
                </v-list-item>
                <v-list-item @click="reboot" append-icon="mdi-restart">
                  <v-list-item-title>Reboot</v-list-item-title>
                </v-list-item>
                <v-list-item @click="poweroff" append-icon="mdi-power">
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
</script>