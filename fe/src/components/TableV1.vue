<template>
  <v-col no-gutters :align-self="left">
    <v-sheet>
      <v-table>
        <tr  v-for="row in tdata"> 
          <component v-for="field in row" :is="field.type" v-bind="field.props"></component>
        </tr>
      </v-table>
    </v-sheet>
  </v-col>
</template>
  
<script>
import eventbus from "@/eventbus";
import TableRowV1 from "./TableRowV1.vue"
export default {
  name: "InputV1",
  components: {
    TableRowV1
  },
  props: {
    command: String,
    function: String,
    left: String,
  },
  created() {
    eventbus.$on("fe.rerender-table", this.updateValue);

    // this.$socketClient.onMessage = msg => {
    //   console.log("jee", JSON.parse(msg.data))
    //   this.data = JSON.parse(msg.data);
    //   if(this.data.metadata?.command === "svc.rerender-table") {
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