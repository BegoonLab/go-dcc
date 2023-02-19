<template>
  <div class="px-2">
    <v-row>
      <v-col cols="4">
          <v-text-field
              label="Address"
              :value="locomotive.address"
              readonly
          ></v-text-field>
      </v-col>
      <v-col cols="4">
        <v-text-field
            label="Name"
            :value="locomotive.name"
            readonly
        ></v-text-field>
      </v-col>
      <v-col cols="4">
        <v-text-field
            label="Speed"
            :value="locomotive.speed"
            readonly
            type="number"
        ></v-text-field>
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="8">
        <v-switch
          :label="`Direction: ${
            locomotive.direction === 1 ? 'Forward >>' : '<< Backward'
          }`"
          :true-value="1"
          :false-value="0"
          v-model="locomotive.direction"
          @change="(e) => update()"
        ></v-switch>
        <v-switch
          :label="`Light: ${locomotive.fl === true ? 'On' : 'Off'}`"
          :true-value="true"
          :false-value="false"
          v-model="locomotive.fl"
          @change="(e) => update()"
        ></v-switch>
      </v-col>
      <v-col cols="4" class="text-center">
          <v-slider
          justify-center
          direction="vertical"
          height="270"
          max="31"
          min="0"
          step="1"
          :focused="false"
          :track-color="`rgb(${locomotive.speed * 10} 100 50)`"
          :thumb-color="`rgb(${locomotive.speed * 10} 100 50)`"
          :thumb-size="48"
          :thumb-label="true"
          v-model="locomotive.speed"
          @update:model-value="(e) => {processSlider()}"
        ></v-slider>
        <v-btn
          :disabled="locomotive.speed == 0"
          depressed
          color="error"
          @click="
            () => {
              locomotive.speed = 0;
              update();
            }
          "
        >
          STOP
        </v-btn>
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="4" v-for="(func, index) in funcNumber" :key="func">
        <div class="f-checkbox">
          <v-checkbox 
            :label="`F${func}`"
            :true-value="true"
            :false-value="false"
            v-model="locomotive['f'+func]"
            @change="(e) => update()"
          ></v-checkbox>
        </div>
        <br v-if="(index + 2) % 9 === 0" />
      </v-col>
    </v-row>
  </div>
</template>

<script setup>
import {computed, defineProps, ref} from "vue";
import {useControllerStore} from "../store/modules/controller";

const props = defineProps(['name'])

const store = useControllerStore()
const locomotive = computed(() => store.locomotives[props.name])

const rules = [
  (v) => v <= 31 || "Max speed is 31",
  (v) => v >= 0 || "Min speed is 0",
]

const funcNumber = 28
const targetSpeed = ref(0)

function update() {
  store.sendDataToServer()
}

function debounce(func, timeout = 500){
  let timer;
  return (...args) => {
    clearTimeout(timer);
    timer = setTimeout(() => { func.apply(this, args); }, timeout);
  };
}
const processSlider = debounce(() => update());
</script>