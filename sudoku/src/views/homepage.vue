<template>
  <div class="home">
    <div class="home_header">
      <div class="home_header_music" @click="music()" title="播放音乐！" ref="music_button">
        <div class="home_header_music_before" id="au_before" v-show="isShow"></div>
        <img src="../assets/images/music.png" alt="" title="播放音乐！">
      </div>
      <div class="home_header_change" @click="changeColor()" title="更换主题！">
        <img src="../assets/images/change.png" alt="" title="更换主题！">
      </div>
    </div>
    <div class="home_body">
      <div class="home_body_title">
        <h1>数独</h1>
        <h2>SUDOKU</h2>
      </div>
      <div class="home_body_btn" @click="$router.push('/select')">
        START
      </div>
    </div>
    <Footer></Footer>
  </div>
</template>

<script>
import Footer from '@/views/footer.vue'
import { backgroundStyles } from '@/assets/js/style.js'
export default {
  name: 'Homepage',
  components: {
    Footer
  },
  data () {
    return {}
  },
  props: {
    bgc:{
      type: String,
      required: true,
      default:'linear-gradient(180deg, #CA6ECE 0%, #232361 100%);'     //默认值
    },
  },
  computed: {
    isShow() {
      console.log(this.$parent.$refs.au.paused);
      if(this.$parent.$refs.au.paused) {
        return true;
      } else {
        return false;
      }
    },
  },
  methods: {
    changeColor () {
      let currentIndex = backgroundStyles.indexOf(this.bgc);
      currentIndex += 1;
      if ( currentIndex >= backgroundStyles.length ) {
        currentIndex %= backgroundStyles.length;
      }
      const selectedColor = backgroundStyles[currentIndex];
      this.$emit('changeBackgroundColor', selectedColor);
    },
    music() {
      const au = this.$parent.$refs.au
      var audio = document.getElementById("au_before");
      if(au.paused) {
        this.$emit('music', true);
        audio.style.display = "none";
      } else {
        this.$emit('music', false);
        audio.style.display = "block";
      }
    }
  }
}
</script>
<style scoped lang='less'>
.home {
  width: 100vw;
  height: 100vh;
  // background: linear-gradient(180deg, #CA6ECE 0%, #232361 100%);
  &_header {
    display: flex;
    justify-content: right;
    width: 100%;
    height: 130px;
    &_change , &_music {
      display: flex;
      margin-top: 32px;
      margin-right: 32px;
      width: 66px;
      height: 66px;
      line-height: 66px;
      border-radius: 10px;
      border-bottom: 5px solid rgba(0, 0, 0, 0.20);
      background-color: rgba(255, 255, 255, 0.60);
      transition: all 0.3s ease 0s;
      &:hover {
        cursor: pointer;
        transform: scale(1.05);
        transition: all 0.3s linear;
      }
      img {
        margin: 11px auto;
        width: 40px;
        height: 40px;
      }
    }
    &_music {
      position: relative;
      &_before {
        display: block;
        content: "";
        position: absolute;
        left: 18px;
        top: 0px;
        width: 50%;
        height: 33px;
        box-sizing: border-box;
        border-bottom: 4px solid rgba(255, 255, 255, 0.80);
        transform-origin: bottom center;
        transform: rotateZ(45deg) scale(1.414);
      }
    }
  }
  &_body {
    margin: 0 auto;
    width: 100%;
    height: calc(100% - 220px);
    &_title {
      display: flex;
      flex-direction: column;
      h1 {
        margin-top: 150px;
        color: rgba(0, 0, 0, 0.20);
        text-align: center;
        // font-family: Microsoft YaHei;
        font-size: 270px;
        font-weight: 700;
        letter-spacing: 2px;
      }
      h2 {
        position: relative;
        height: 0;
        top: -235px;
        left: 0;
        color: rgba(255, 255, 255, 0.80);
        text-align: center;
        font-family: Microsoft YaHei;
        font-size: 85px;
        font-weight: 700;
        letter-spacing: 2px;
      }
    }
    &_btn {
      display: flex;
      flex-direction: column;
      margin: 0 auto;
      margin-top: 42px;
      width: 542px;
      height: 110px;
      border-radius: 12px;
      border: 1px solid rgba(255, 255, 255, 0.20);
      color: rgba(255, 255, 255, 0.60);
      text-align: center;
      font-family: Microsoft YaHei;
      font-size: 37.8px;
      font-style: normal;
      font-weight: 700;
      line-height: 110px; /* 285.714% */
      letter-spacing: 10.8px;
      transition: all 0.4s ease 0s;
      &:hover {
        cursor: pointer;
        border-color: rgba(255, 255, 255, 0.4);
        background-color: rgba(255, 255, 255, 0.1);
        color: rgba(255, 255, 255, 0.8);
        transition: all 0.4s linear;
        font-size: 42px;
        transform: scale(1.05);
      }
    }
  }
}

</style>
