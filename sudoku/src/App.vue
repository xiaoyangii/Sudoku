<template>
  <div id="app" :style="backgroundColor">
    <router-view @changeBackgroundColor = updateBackgroundColor />
  </div>
</template>

<script>
export default {
  data() {
    return {
      backgroundColor: ''
    };
  },
  methods: {
    updateBackgroundColor(selectedColor) {
      this.backgroundColor = selectedColor;
    }
  }
}
</script>
<style lang="less">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  background: linear-gradient(180deg, #CA6ECE 0%, #232361 100%);
}
::-webkit-scrollbar {
  width: 0 !important;
}
::-webkit-scrollbar {
  width: 0 !important;height: 0;
}
</style>
