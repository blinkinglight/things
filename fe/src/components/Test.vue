<template>
  <v-col no-gutters :align-self="left">
    <v-sheet>
      <h2>color: <input type="text" v-model="store" /></h2>
      <p>size: {{ size }}</p>
      <p>{{ store }}</p>
    </v-sheet>
  </v-col>
</template>
  
<script>
export default {
  name: "Test",
  props: {
    color: String,
    size: String,
    left: String,
  },
  data: (props) => ({
    alignments: [
      'start',
      'center',
      'end'
    ],
    store: props.color,
  }),
  computed: {
    copy(data) {
      return data.store;
    }
  },
  watch: {
    store: function (val) {
      this.copys = val;
    },
  },

};
</script>