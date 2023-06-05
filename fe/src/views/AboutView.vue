<template>
  <div class="about">
    <h1>This is an about page</h1>
  </div>
  <v-container>
    <v-row v-for="row in rows">
      <component v-for="item in row.components" :is="item.name" v-bind="item.props"></component>
    </v-row>
  </v-container>
</template>

<script>

import Test from '../components/Test.vue';
import InputV1 from '../components/InputV1.vue';
import TableV1 from '@/components/TableV1.vue';
import eventbus from "@/eventbus";
import axios from 'axios';

export default {
  created() {
    // eventbus.$on("svc.rerender-table", () => {
    //   alert("A");
  },
  components: {
    Test, InputV1, TableV1
  },
  data() {
    return {
      data: {},
      rows: []
    }
  },
  async mounted() {
    const response = await axios.get('http://localhost:3000/api');
    this.data = response.data;
    this.rows = response.data.rows;
    console.log(this.rows);
    if (this.data.onload) {
      for (let onload in this.data.onload) {
        onload = this.data.onload[onload]
        console.log("onload", onload);
        fetch(`http://localhost:3000/pipe?subject=${onload.subject}&type=command`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ data: onload.payload }),
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
    }
  },
  methods: {
    async updateValue(event) {
      console.log("event", event)
    }
  }
}
</script>