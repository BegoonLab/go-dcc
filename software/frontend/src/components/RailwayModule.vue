<template>
  <v-row class="px-3 py-2">
    <v-col cols="3">
      <v-text-field
          label="Address"
          v-model="railwayModule.address"
          :hint="railwayModule.name"
          variant="outlined"
          readonly
          density="compact"
      ></v-text-field>
    </v-col>
    <v-col cols="9">
      <v-select
          label="Route"
          :items="Object.keys(getOrderedRoutes(railwayModule.routes))"
          v-model="railwayModule.activeRoute"
          variant="outlined"
          density="compact"
          @update:modelValue="() => {update()}"
      >
        <template v-slot:prepend-item>
          <v-list-item
              @click="() => {railwayModule.activeRoute = 'Custom'; update()}"
          >
            <template v-slot:title><b>Custom</b></template>
          </v-list-item>
          <v-divider class="mt-2"></v-divider>
        </template>
      </v-select>
    </v-col>
  </v-row>
  <v-row class="px-3">
    <v-col cols="12">
      <v-expansion-panels variant="accordion" color="deep-purple-lighten-2" v-model="panel">
        <v-expansion-panel>
          <template v-slot:title>
            <v-icon class="mr-2">mdi-source-branch</v-icon>
            Routes of {{ railwayModule.name }}
          </template>
          <v-expansion-panel-text>
            <v-radio-group v-model="railwayModule.activeRoute" @change="(e) => update()">
              <v-radio value="Custom" class="my-2">
                <template v-slot:label>
                  <v-row>
                    <v-col cols="8" class="my-2"><b>Custom</b></v-col>
                    <v-col cols="4">
                      <v-btn color="amber-lighten-1" @click="reset()">
                        <v-icon>mdi-autorenew</v-icon>
                      </v-btn>
                    </v-col>
                  </v-row>
                </template>
              </v-radio>
              <v-radio
                  :value="item.name"
                  v-for="item in getOrderedRoutes(railwayModule.routes)"
                  :key="item.address"
                  class="my-2"
              >
                <template v-slot:label>
                  <v-row>
                    <v-col cols="8" class="my-2">{{ item.name }}</v-col>
                    <v-col cols="4">
                      <v-btn color="red-darken-2"
                             @click="() => {confirmDeletionDialog = true; routeToDelete = item.name}"
                      >
                        <v-icon>mdi-trash-can-outline</v-icon>
                      </v-btn>
                    </v-col>
                  </v-row>
                </template>
              </v-radio>
            </v-radio-group>
          </v-expansion-panel-text>
        </v-expansion-panel>
        <railway-turnouts
            :name="name"
            :title="'Left railway turnouts'"
            :indexes="left"
        ></railway-turnouts>
        <railway-turnouts
          :name="name"
          :title="'Right railway turnouts'"
          :indexes="right"
        ></railway-turnouts>
      </v-expansion-panels>
    </v-col>
  </v-row>
  <v-dialog
      v-model="confirmDeletionDialog"
      persistent
      width="auto"
  >
    <v-card>
      <v-card-title class="text-h5">
        Are you sure?
      </v-card-title>
      <v-card-text>
        <b>Warning:</b> You are about to remove the route '{{ routeToDelete }}'.
        This action cannot be undone. Are you sure you want to proceed?
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn
            color="grey-darken-1"
            variant="outlined"
            @click="confirmDeletionDialog = false"
        >
          No, Cancel
        </v-btn>
        <v-btn
            color="red-darken-4"
            variant="outlined"
            @click="deleteRoute()"
        >
          Yes, delete
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>

</template>

<script setup>
import {computed, defineProps, ref} from "vue";
import {useControllerStore} from "../store/modules/controller";

const props = defineProps(['name'])
const routeToDelete = ref(null)
const confirmDeletionDialog = ref(false)
const panel = ref([0])
const left = ref([1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31])
const right = ref([0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30])
const store = useControllerStore()
const railwayModule = computed(() => store.railwayModules[props.name])

function deleteRoute() {
  railwayModule.value.activeRoute = 'Custom'
  delete railwayModule.value.routes[routeToDelete.value]
  update()
  confirmDeletionDialog.value = false
}

function reset() {
  for (let i = 0; i <= 32; i++) {
    railwayModule.value.routes["Custom"][`f${i}`] = false
  }
  update()
}

function getOrderedRoutes(unordered) {
  return Object.keys(unordered).sort().reduce(
      (obj, key) => {
        if (key === "Custom") {
          return obj;
        }
        obj[key] = unordered[key];
        return obj;
      },
      {}
  );
}

function update() {
  store.sendDataToServer()
}
</script>