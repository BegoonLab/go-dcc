<template>
  <v-row class="px-3 py-2">
    <v-col class="v-col-4">
      <v-text-field
          label="Address"
          v-model="railwayModule.address"
          variant="outlined"
          readonly
          density="compact"
      ></v-text-field>
    </v-col>
    <v-col class="v-col-4">
      <v-text-field
          label="Name"
          v-model="railwayModule.name"
          variant="outlined"
          readonly
          density="compact"
      ></v-text-field>
    </v-col>
    <v-col class="v-col-4">
      <v-select
          label="Route"
          :items="['Custom']"
          v-model="railwayModule.route"
          variant="outlined"
          density="compact"
          @change="(e) => update()"
      ></v-select>
    </v-col>
  </v-row>
  <v-row class="px-3">
    <v-col cols="12">
      <v-expansion-panels variant="accordion" color="deep-purple-lighten-2" v-model="panel">
        <v-expansion-panel title="Routes">
          <v-expansion-panel-text>
            <v-radio-group v-model="railwayModule.route">
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
              <v-radio :value="`Custom${i}`" v-for="i in 10" :key="i" class="my-2">
                <template v-slot:label>
                  <v-row>
                    <v-col cols="8" class="my-2">Custom {{ i }}</v-col>
                    <v-col cols="4">
                      <v-btn color="red-darken-2"
                             @click="confirmDeletionDialog = true"
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
        <v-expansion-panel title="Left">
          <v-expansion-panel-text>
            <v-row>
              <v-col cols="4" v-for="i in left" :key="i">
                <v-switch
                    :label="`#${i}`"
                    :true-value="true"
                    v-model="railwayModule[`f${i}`]"
                    color="success"
                    density="comfortable"
                    @change="(e) => update()"
                ></v-switch>
              </v-col>
            </v-row>
          </v-expansion-panel-text>
        </v-expansion-panel>
        <v-expansion-panel title="Right">
          <v-expansion-panel-text>
            <v-row>
              <v-col cols="4" v-for="i in right" :key="i">
                <v-switch
                    :label="`#${i}`"
                    :true-value="true"
                    v-model="railwayModule[`f${i}`]"
                    color="success"
                    density="comfortable"
                    @change="(e) => update()"
                ></v-switch>
              </v-col>
              <v-col cols="3" class="text-right">
                <v-btn
                    class="my-1"
                    color="amber-lighten-1"
                    @click="(e) => {reset()}"
                >
                  <v-icon>mdi-autorenew</v-icon>
                </v-btn>
              </v-col>
              <v-col cols="5">
                <v-btn
                    class="my-1"
                    color="success"
                    @click="(e) => {saveAsDialog = true}"
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
        <b>Warning:</b> You are about to remove the route 'Main loop'.
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
            @click="confirmDeletionDialog = false"
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
    <v-card>
      <v-card-title class="text-h5">
        Save new route
      </v-card-title>
      <v-card-text>
          <v-text-field
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
            color="success"
            variant="outlined"
            @click="saveAsDialog = false"
        >
          <v-icon>mdi-check-circle</v-icon>
          Save
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script setup>
import {computed, defineProps, ref} from "vue";
import {useControllerStore} from "../store/modules/controller";

const props = defineProps(['name'])
const panel = ref([0])
const confirmDeletionDialog = ref(false)
const saveAsDialog = ref(false)
const left = ref([1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31])
const right = ref([2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32])

const store = useControllerStore()
const railwayModule = computed(() => store.railwayModules[props.name])

function reset() {

}

function saveAs() {

}

function update() {
  store.sendDataToServer()
}
</script>