<template>
  <div id="app" :style="backgroundColor">
    <audio :src="currentSongUrl" loop autoplay ref="au"></audio>
    <router-view @changeBackgroundColor = updateBackgroundColor :bgc="backgroundColor"/>
  </div>
</template>

<script>
export default {
  data() {
    return {
      backgroundColor: '',
      playlist: [
        'sound_bgm_3.m4a',
        'sound_bgm_2.m4a',
        'sound_bgm_1.mp3',
      ],
      isPlaying: true
    };
  },
  computed: {
    currentSongUrl() {
      if (!this.isPlaying) {
        return null;
      }
      const currentSongIndex = Math.floor(Math.random() * this.playlist.length);
      return require(`@/assets/sounds/${this.playlist[currentSongIndex]}`);
    },
  },
  methods: {
    updateBackgroundColor(selectedColor) {
      this.backgroundColor = selectedColor;
    },
    toggleAudio() {
      this.isPlaying = !this.isPlaying;
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
