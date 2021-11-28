<template>
  <div class="px-2">
    <v-row>
      <v-col cols="8">
        <v-text-field
          label="Address"
          :value="locomotive.address"
          readonly
        ></v-text-field>
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
        <v-text-field
          label="Speed"
          :value="locomotive.speed"
          readonly
          type="number"
        ></v-text-field>
        <v-slider
          justify-center
          vertical
          hint="Im a hint"
          height="270"
          max="255"
          min="0"
          :value="locomotive.speed"
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
        (v) => v <= 255 || "Max speed is 255",
        (v) => v >= 0 || "Min speed is 0",
      ],
      funcNumber: 21,
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
</style>