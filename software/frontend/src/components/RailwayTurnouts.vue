<template>
  <v-expansion-panel>
    <template v-slot:title>
      <v-icon class="mr-2" v-if="!isCustomActiveRoute">mdi-lock</v-icon>
      <v-icon class="mr-2" v-if="isCustomActiveRoute">mdi-lock-open-variant-outline</v-icon>
      {{title}}
    </template>
    <v-expansion-panel-text>
      <v-row>
        <v-col cols="4" v-for="i in indexes" :key="i">
          <v-tooltip
              v-if="!isCustomActiveRoute"
              activator="parent"
              location="top"
          >Use `Custom` to adjust the railway turnout
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
        <v-col v-if="isCustomActiveRoute" cols="12" class="text-right">
          <v-btn
              class="mx-2"
              color="amber-lighten-1"
              @click="(e) => {reset()}"
          >
            <v-icon>mdi-autorenew</v-icon>
            Reset
          </v-btn>

          <v-btn
              class=""
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

  const store = useControllerStore()
  const rules = {
    required: value => !!value || 'Field is required.',
    max: value => value.length <= 50 || 'Max 50 characters.',
    min: value => value.length > 2 || 'Min 3 characters.',
    exists: value => !railwayModule.value.routes[value] || 'Please choose a different name.'
  };

  const saveAsForm = ref(false)
  const saveAsDialog = ref(false)
  const newRoute = ref({})
  const props = defineProps(['name', 'title', 'indexes'])
  const railwayModule = computed(() => store.railwayModules[props.name])
  const isCustomActiveRoute = computed(() => store.railwayModules[props.name].activeRoute === 'Custom')

  function saveAs() {
    railwayModule.value.routes[newRoute.value.name] = {
      name: newRoute.value.name
    }
    for (let i = 0; i <= 32; i++) {
      railwayModule.value.routes[newRoute.value.name][`f${i}`] = railwayModule.value.routes["Custom"][`f${i}`]
    }
    update()
  }

  function reset() {
    for (let i = 0; i <= 32; i++) {
      railwayModule.value.routes["Custom"][`f${i}`] = false
    }
    update()
  }

  function update() {
    store.sendDataToServer()
  }
</script>

<style scoped>

</style>