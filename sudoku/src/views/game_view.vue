<template>
  <div class="game">
    <div class="game_header">
      <div class="game_header_again" @click="solve()" title="一键求解！">
        <div>一键求解</div>
      </div>
      <div class="game_header_again" @click="again()" title="重新游玩！">
        <div>重新游玩</div>
      </div>
      <div class="game_header_return" @click="$router.go(-1)" title="返回！">
        <img src="../assets/images/return.png" alt="" title="返回！">
      </div>
      <div class="game_header_music" @click="music()" title="播放音乐！">
        <div class="game_header_music_before" id="au_before" v-show="isShow"></div>
        <img src="../assets/images/music.png" alt="" title="播放音乐！">
      </div>
      <div class="game_header_change" @click="changeColor()" title="更换主题！">
        <img src="../assets/images/change.png" alt="" title="更换主题！">
      </div>
    </div>
    <div class="game_body">
      <div class="maingrid">
        <div class="grid" v-for="(Sudoku, index1) in SUDOKU" :key="index1">
          <div class="grid_block" v-for="(block, index2) in Sudoku" :key="index2">
            <div class="grid_block_box" 
              v-for="(box, index3) in block" :key="index3"
              ref="editor"
              @click="isInput($event, index1, index2, index3)"
              @input="inputText($event, index1, index2, index3)"
              @blur="inputBlur"
              @focus="inputFocus" 
              contenteditable="false"
            >
              {{ box===0?"":box }}
            </div>
          </div>
        </div>
      </div>
    </div>
    <Footer></Footer>
  </div>
</template>

<script>
import Footer from '@/views/footer.vue';
import { backgroundStyles } from '@/assets/js/style.js'
import { getSudoku, solveSudoku } from '@/api/getSudoku.js'
export default {
  name: 'Game',
  components: {
    Footer
  },
  data () {
    return {
      SUDOKU: [],
      SUDOKU_copy: [],
      isBlur: true, // 解决赋值时光标自动定位到起始位置
    }
  },
  props: {
    bgc:{
      type: String,
      required: true,
      default:'linear-gradient(180deg, #CA6ECE 0%, #232361 100%);'//默认值
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
    lever () {
      return this.$route.params.lever
    }
  },
  methods: {
    inputText(e, id1, id2, id3) {
      this.SUDOKU[id1][id2][id3] = parseInt(e.target.innerText)
      // console.log(parseInt(e.target.innerText))
      console.log(e);
    },
    async solve() {
      let res = await solveSudoku(this.lever)
    },
    isInput(e, id1, id2, id3) {
      console.log(e, id1, id2, id3)
      if(e.target.innerText != "") {
        return
      } else {
        e.target.contentEditable = true
      }
    },
    inputFocus() {
      this.isBlur = false;
    },
    inputBlur() {
      this.isBlur = true;
    },
    changeColor() {
      let currentIndex = backgroundStyles.indexOf(this.bgc);
      currentIndex += 1;
      if ( currentIndex >= backgroundStyles.length ) {
        currentIndex %= backgroundStyles.length;
      }
      const selectedColor = backgroundStyles[currentIndex];
      console.log(currentIndex);
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
    },
    async getsudoku () {
      let res = await getSudoku(this.lever)
      this.SUDOKU = res.data
      this.SUDOKU_copy = this.SUDOKU
    },
    again() {
      this.SUDOKU = this.SUDOKU_copy
    },
    forbidenter(e) {
      // 禁止用户输入回车
      console.log(e);
      _code = e.keyCode;
      // if (_code == 13) {
      //   e.returnValue = false;
      // }
      if((_code < 48 || _code > 57) && _code != 190 && _code != 8) {
        e.preventDefault();
      }
    }
  },
  created () {
    this.getsudoku()
  },
}
</script>
<style scoped lang='less'>
.game {
  width: 100vw;
  height: 220vh;
  // background: linear-gradient(180deg, #CA6ECE 0%, #232361 100%);
  &_header {
    display: flex;
    justify-content: right;
    width: 100vw;
    height: 100px;
    &_change, &_return, &_music, &_again {
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
    &_again {
      width: 150px;
      color: #fff;
      div {
        display: flex;
        justify-content: center;
        margin: 0 auto;
        font-weight: 700;
        font-size: 25px;
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
    &_return {
      img {
        margin: 15px auto;
        width: 32px;
        height: 32px;
      }
    }
  }
  &_body {
    display: flex;
    justify-content: center;
    align-items: center;
    margin: 0 auto;
    width: 100%;
    height: calc(100% - 190px);
  }
}

.maingrid {
  width: 197.5vmin;
  height: 200vmin;
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;
  align-content: space-between;
}
.grid {
  width: 64vmin;
  height: 64vmin;
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;
  align-content: space-between;
  &_block {
    width: 21vmin;
    height: 21vmin;
    display: flex;
    flex-wrap: wrap;
    justify-content: space-between;
    align-content: space-between;
    &_box {
      display: flex;
      justify-content: center;
      align-items: center;
      width: 7vmin;
      height: 7vmin;
      border: 0.2vmin solid transparent;
      background-color: rgba(255, 255, 255, 0.9);
      background-clip: content-box;
      font-size: 3.5vmin;
      color: rgba(0, 0, 0, 0.8);
      font-weight: bold;
      // position: relative;
    }
  }
}
// #one-1 {
//   border-top-left-radius: 1.7vmin;
// }
// #three-3 {
//   border-top-right-radius: 1.7vmin;
// }
// #seven-7 {
//   border-bottom-left-radius: 1.7vmin;
// }
// #nine-9 {
//   border-bottom-right-radius: 1.7vmin;
// }

.grid_block {
  &:first-child {
    .grid_block_box {
      &:first-child {
        border-top-left-radius: 1.7vmin;
      }
    }
  }
  &:nth-child(3) {
    .grid_block_box {
      &:nth-child(3) {
        border-top-right-radius: 1.7vmin;
      }
    }
  }
  &:nth-child(7) {
    .grid_block_box {
      &:nth-child(7) {
        border-bottom-left-radius: 1.7vmin;
      }
    }
  }
  &:nth-child(9) {
    .grid_block_box {
      &:nth-child(9) {
        border-bottom-right-radius: 1.7vmin;
      }
    }
  }
}
</style>
