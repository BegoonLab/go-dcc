<template>
  <div class="px-2">
    <v-row>
      <v-col cols="8">
        <h3>Address: {{ locomotive.address }}</h3>
        <v-switch
          v-model="dir"
          :label="`Direction: ${dir}`"
          :true-value="'Forward'"
          :false-value="'Backward'"
        ></v-switch>

        <v-switch
          v-model="light"
          :label="`Light: ${light}`"
          :true-value="'On'"
          :false-value="'Off'"
        ></v-switch>
      </v-col>
      <v-col cols="4" class="text-center">
        <h3>Speed: {{ locomotive.speed }}</h3>
        <v-slider
          justify-center
          vertical
          hint="Im a hint"
          height="270"
          max="255"
          min="0"
          :value="locomotive.speed"
          @end="v => update(v, 'speed')"
        ></v-slider>
        <v-btn
          :disabled="locomotive.speed == 0"
          depressed
          color="error"
          @click="() => {update(0, 'speed')}"
        >
          STOP
        </v-btn>
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <v-checkbox v-model="f1" :label="`F1`" class="align-center justify-center"></v-checkbox>
        <v-checkbox v-model="f2" :label="`F2`"></v-checkbox>
        <v-checkbox v-model="f3" :label="`F3`"></v-checkbox>
        <v-checkbox v-model="f4" :label="`F4`"></v-checkbox>
        <v-checkbox v-model="f5" :label="`F5`"></v-checkbox>
        <v-checkbox v-model="f6" :label="`F6`"></v-checkbox>
        <v-checkbox v-model="f7" :label="`F7`"></v-checkbox>
      </v-col>
      <v-col>
        <v-checkbox v-model="f8" :label="`F8`"></v-checkbox>
        <v-checkbox v-model="f9" :label="`F9`"></v-checkbox>
        <v-checkbox v-model="f10" :label="`F10`"></v-checkbox>
        <v-checkbox v-model="f11" :label="`F11`"></v-checkbox>
        <v-checkbox v-model="f12" :label="`F12`"></v-checkbox>
        <v-checkbox v-model="f13" :label="`F13`"></v-checkbox>
        <v-checkbox v-model="f14" :label="`F14`"></v-checkbox>
      </v-col>
      <v-col>
        <v-checkbox v-model="f15" :label="`F15`"></v-checkbox>
        <v-checkbox v-model="f16" :label="`F16`"></v-checkbox>
        <v-checkbox v-model="f17" :label="`F17`"></v-checkbox>
        <v-checkbox v-model="f18" :label="`F18`"></v-checkbox>
        <v-checkbox v-model="f19" :label="`F19`"></v-checkbox>
        <v-checkbox v-model="f20" :label="`F20`"></v-checkbox>
        <v-checkbox v-model="f21" :label="`F21`"></v-checkbox>
      </v-col>
    </v-row>
  </div>
</template>

<script>
import { mapGetters, mapState } from "vuex";
export default {
  props: ["name"],
  computed: {
    ...mapState({
      locomotive: function (state) {
        return state.controller.locomotives[this.name];
      },
    }),
  },
  data() {
    return {
      dir: "Forward",
      light: "Off",
      f1: false,
      f2: false,
      f3: false,
      f4: false,
      f5: false,
      f6: false,
      f7: false,
      f8: false,
      f9: false,
      f10: false,
      f11: false,
      f12: false,
      f13: false,
      f14: false,
      f15: false,
      f16: false,
      f17: false,
      f18: false,
      f19: false,
      f20: false,
      f21: false,
    };
  },
  methods: {
    update(value, where) {
      
      this.$store.dispatch("controller/setLocomotiveState", {name: this.name, value: value, where: where});
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
</style>