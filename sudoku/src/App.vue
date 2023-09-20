<template>
  <div id="app" :style="backgroundColor">
    <audio loop autoplay ref="au">
      <source :src="currentSongUrl" type="audio/mp3">
    </audio>
    <router-view @changeBackgroundColor = updateBackgroundColor :bgc="backgroundColor" @music = toggleAudio />
  </div>
</template>

<script>
export default {
  data() {
    return {
      backgroundColor: '',
      playlist: [
        'soundbgm1.mp3',
        'soundbgm2.mp3',
        'soundbgm3.mp3',
        'soundbgm4.mp3',
      ]
    };
  },
  computed: {
    currentSongUrl() {
      const currentSongIndex = Math.floor(Math.random() * this.playlist.length);
      return require(`@/assets/sounds/${this.playlist[currentSongIndex]}`);
    },
  },
  methods: {
    updateBackgroundColor(selectedColor) {
      this.backgroundColor = selectedColor;
    },
    toggleAudio(val) {
      const au = this.$refs.au;
      if (val == true) {
        au.play();
        au.volume = 0.5;
      } else {
        au.pause()
      }
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
