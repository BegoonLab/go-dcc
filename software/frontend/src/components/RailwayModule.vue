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
            <v-icon>mdi-arrow-decision-outline</v-icon>
            &nbsp;
            Routes of {{railwayModule.name}}
          </template>
          <v-expansion-panel-text>
            <v-radio-group v-model="railwayModule.activeRoute" @change="(e) => update()">
              <v-radio value="Custom" class="my-2">
                <template v-slot:label>
                  <v-row>
                    <v-col cols="8" class="my-2"><b>Custom</b></v-col>
                    <v-col cols="4">
                      <v-btn color="amber-lighten-1">
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
        <v-expansion-panel>
          <template v-slot:title>
            <v-icon v-if="!isCustomActiveRoute">mdi-lock</v-icon>
            <v-icon v-if="isCustomActiveRoute">mdi-lock-open-variant-outline</v-icon>
            &nbsp;
            Left railroad switches
          </template>
          <v-expansion-panel-text>
            <v-row>
              <v-col cols="4" v-for="i in left" :key="i">
                <v-tooltip
                    v-if="!isCustomActiveRoute"
                    activator="parent"
                    location="top"
                >Use `Custom` to adjust the railway switch
                </v-tooltip>
                <v-switch
                    v-if="isCustomActiveRoute"
                    :label="`# ${i}`"
                    :true-value="true"
                    v-model="railwayModule.routes['Custom'][`f${i}`]"
                    color="success"
                    density="comfortable"
                    @change="(e) => {update()}"
                ></v-switch>
                <v-switch
                    v-if="!isCustomActiveRoute"
                    readonly
                    :label="`# ${i}`"
                    :true-value="true"
                    v-model="railwayModule.routes[railwayModule.activeRoute][`f${i}`]"
                    color="green-lighten-3"
                    density="comfortable"
                ></v-switch>
              </v-col>
              <v-col v-if="isCustomActiveRoute" cols="3" class="text-right">
                <v-btn
                    class="my-1"
                    color="amber-lighten-1"
                    @click="(e) => {reset()}"
                >
                  <v-icon>mdi-autorenew</v-icon>
                </v-btn>
              </v-col>
              <v-col v-if="isCustomActiveRoute" cols="5">
                <v-btn
                    class="my-1"
                    color="success"
                    @click="(e) => {saveAsDialog = true; newRoute.name = null}"
                >
                  <v-icon>mdi-check-circle</v-icon>
                  Save As...
                </v-btn>
              </v-col>
            </v-row>
          </v-expansion-panel-text>
        </v-expansion-panel>
        <v-expansion-panel>
          <template v-slot:title>
            <v-icon v-if="!isCustomActiveRoute">mdi-lock</v-icon>
            <v-icon v-if="isCustomActiveRoute">mdi-lock-open-variant-outline</v-icon>
            &nbsp;
            Right railroad switches
          </template>
          <v-expansion-panel-text>
            <v-row>
              <v-col cols="4" v-for="i in right" :key="i">
                <v-tooltip
                    v-if="!isCustomActiveRoute"
                    activator="parent"
                    location="top"
                >Use `Custom` to adjust the railway switch
                </v-tooltip>
                <v-switch
                    v-if="isCustomActiveRoute"
                    :label="`# ${i}`"
                    :true-value="true"
                    v-model="railwayModule.routes['Custom'][`f${i}`]"
                    color="success"
                    density="comfortable"
                    @change="(e) => {update()}"
                ></v-switch>
                <v-switch
                    v-if="!isCustomActiveRoute"
                    readonly
                    :label="`# ${i}`"
                    :true-value="true"
                    v-model="railwayModule.routes[railwayModule.activeRoute][`f${i}`]"
                    color="green-lighten-3"
                    density="comfortable"
                ></v-switch>
              </v-col>
              <v-col v-if="isCustomActiveRoute" cols="3" class="text-right">
                <v-btn
                    class="my-1"
                    color="amber-lighten-1"
                    @click="(e) => {reset()}"
                >
                  <v-icon>mdi-autorenew</v-icon>
                </v-btn>
              </v-col>
              <v-col v-if="isCustomActiveRoute" cols="5">
                <v-btn
                    class="my-1"
                    color="success"
                    @click="(e) => {saveAsDialog = true; newRoute.name = null}"
                >
                  <v-icon>mdi-check-circle</v-icon>
                  Save As...
                </v-btn>
              </v-col>
            </v-row>
          </v-expansion-panel-text>
        </v-expansion-panel>
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
  <v-dialog
      v-model="saveAsDialog"
      persistent
      width="1024"
  >
    <v-form
        v-model="saveAsForm"
    >
      <v-card>
        <v-card-title class="text-h5">
          Save new route
        </v-card-title>
        <v-card-text>
          <v-text-field
              v-model="newRoute.name"
              :rules="[rules.required, rules.min, rules.max, rules.exists]"
              label="Route name*"
              required
          ></v-text-field>
          <small>*indicates required field</small>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
              color="grey-darken-1"
              variant="outlined"
              @click="saveAsDialog = false"
          >
            Close
          </v-btn>
          <v-btn
              :disabled="!saveAsForm"
              color="success"
              variant="outlined"
              @click="() => {saveAs(); saveAsDialog = false}"
          >
            <v-icon>mdi-check-circle</v-icon>
            Save
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-form>
  </v-dialog>
</template>

<script setup>
import {computed, defineProps, ref} from "vue";
import {useControllerStore} from "../store/modules/controller";

const props = defineProps(['name'])
const panel = ref([0])
const confirmDeletionDialog = ref(false)
const saveAsForm = ref(false)
const saveAsDialog = ref(false)
const newRoute = ref({})
const routeToDelete = ref(null)
const left = ref([1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31])
const right = ref([0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30])
const rules = {
  required: value => !!value || 'Field is required.',
  max: value => value.length <= 50 || 'Max 50 characters.',
  min: value => value.length > 2 || 'Min 3 characters.',
  exists: value => !railwayModule.value.routes[value] || 'Please choose a different name.'
};

const store = useControllerStore()
const railwayModule = computed(() => store.railwayModules[props.name])
const isCustomActiveRoute = computed(() => store.railwayModules[props.name].activeRoute === 'Custom')

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
}

function saveAs() {
  railwayModule.value.routes[newRoute.value.name] = {
    name: newRoute.value.name
  }
  for (let i = 0; i <= 32; i++) {
    railwayModule.value.routes[newRoute.value.name][`f${i}`] = railwayModule.value.routes["Custom"][`f${i}`]
  }
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