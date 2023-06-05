<template>
  <nav>
    <router-link to="/">Home</router-link> |
    <router-link to="/about">About</router-link>
  </nav>
  <router-view/>
</template>


<script>
import eventbus from './eventbus';

export default {
  name: "App",
  mounted() {
    this.$socketClient.onMessage = msg => {
      console.log("je111e", JSON.parse(msg.data))
      this.data = JSON.parse(msg.data);
      if(this.data.metadata?.command) {
        eventbus.$emit(this.data.metadata.command, this.data.data)
      }
    }
  },
};
</script>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

nav {
  padding: 30px;

  a {
    font-weight: bold;
    color: #2c3e50;

    &.router-link-exact-active {
      color: #42b983;
    }
  }
}
</style>
