<template>
  <v-col no-gutters :align-self="left">
    <v-sheet>
      <v-text-field v-model="store" @change="updateValue" />
    </v-sheet>
  </v-col>
</template>
  
<script>
export default {
  name: "InputV1",
  props: {
    command: String,
    left: String,
  },
  data: (props) => ({
    alignments: [
      'start',
      'center',
      'end'
    ],
    store: "",
    cmd: props.command
  }),
  computed: {
    copy(data) {
      return data.store;
    }
  },
  methods: {
    updateValue: function (event) {
      console.log(event, this.cmd);
      fetch(`http://localhost:3000/pipe?subject=${this.cmd}&type=command`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ data: {todo: event.target.value }}),
      })
        .then(response => response.json())
        .then(data => {
          console.log('Success:', data);
          this.store = data.color;
        })
        .catch((error) => {
          console.error('Error:', error);
        });
    }
  },
  watch: {
    store: function (val) {
      this.copys = val;
    },
  },

};
</script>