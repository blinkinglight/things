<template>
  <v-col no-gutters :align-self="left">
    <v-sheet>
      <!-- <h2>color: <input type="text" v-model="store" /></h2> -->
      <p>check: {{ value }}</p>
    </v-sheet>
  </v-col>
</template>
  
<script>
export default {
  name: "InputV1",
  props: {
    command: String,
    function: String,
    value: String,
    left: String,
  },
  created() {
    // this.$socketClient.onMessage = msg => {
    //   console.log("jee1", JSON.parse(msg.data))
    //   this.data = JSON.parse(msg.data);
    //   if(this.data.metadata?.command === "svc.rerender-table") {
    //     console.log("rerender")
    //     this.updateValue(this.data.data)
    //   }
    // }
  }, 
  data: (props) => ({
    alignments: [
      'start',
      'center',
      'end'
    ],
    store: "",
    cmd: props.command,
    tdata: [],
  }),
  computed: {
    copy(data) {
      return data.store;
    }
  },
  methods: {
    updateValue: function (event) {
      console.log("mh", event.data, this.cmd);
      this.tdata = event.data?.rows;
    }
  },
  watch: {
    store: function (val) {
      this.copys = val;
    },
  },

};
</script>