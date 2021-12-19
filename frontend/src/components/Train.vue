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
          :input-value="locomotive.direction"
          @change="(v) => update(v, 'direction')"
        ></v-switch>
        <v-switch
          :label="`Light: ${locomotive.fl === true ? 'On' : 'Off'}`"
          :true-value="true"
          :false-value="false"
          :input-value="locomotive.fl"
          @change="(v) => update(v, 'fl')"
        ></v-switch>
      </v-col>
      <v-col cols="4" class="text-center">
        <v-slider
          justify-center
          vertical
          hint="Im a hint"
          height="270"
          max="31"
          min="0"
          :thumb-color="`rgb(${targetSpeed * 8} 100 50)`"
          :thumb-size="48"
          :thumb-label="true"
          :value="locomotive.speed"
          @input="(v) => targetSpeed = v"
          @change="(v) => update(v, 'speed')"
        ></v-slider>
        <v-btn
          :disabled="locomotive.speed == 0"
          depressed
          color="error"
          @click="
            () => {
              update(0, 'speed');
            }
          "
        >
          STOP
        </v-btn>
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="4" v-for="func in funcNumber" :key="func">
        <div class="f-checkbox">
          <v-checkbox 
            :label="`F${func}`"
            :true-value="true"
            :false-value="false"
            :input-value="locomotive['f'+func]"
            @change="(v) => update(v, 'f'+func)"
          ></v-checkbox>
        </div>
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
      rules: [
        (v) => v <= 31 || "Max speed is 31",
        (v) => v >= 0 || "Min speed is 0",
      ],
      funcNumber: 21,
      targetSpeed: 0,
    };
  },
  methods: {
    update(value, where) {
      this.$store.dispatch("controller/setLocomotiveState", {
        name: this.name,
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
.f-checkbox {
  height: 20px;
  display: flex;
  justify-content: center;
}
.v-slider__thumb-label {
  transform: translateY(20%) translateY(15px) translateX(-200%) rotate(-45deg) !important;
}
.v-slider__thumb-label div {
  transform: rotate(45deg) !important;
}
.v-slider__thumb-label span {
  font-size: 2em;
}
</style>