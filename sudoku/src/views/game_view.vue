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
      <div class="maingrid" :key="key">
        <div class="grid" v-for="(Sudoku, index1) in SUDOKU" :key="index1">
          <div class="grid_block" v-for="(block, index2) in Sudoku" :key="index2">
            <div class="grid_block_box" 
              v-for="(box, index3) in block" :key="index3"
              :ref="idx(index1, index2, index3)"
              @click="isInput($event, index1, index2, index3)"
              @input="inputText($event, index1, index2, index3)"
              @blur="inputBlur"
              @focus="inputFocus"
              @mouseover="mouseover($event, index1, index2, index3)"
              @mouseout="mouseout($event, index1, index2, index3)"
              :style="toChangeStyle(index1, index2, index3, bgc)"
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
      su_copy: [],
      isBlur: true, // 解决赋值时光标自动定位到起始位置
      key: "",
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
      this.SUDOKU[id1][id2][id3] = e.target.innerText
      // console.log(parseInt(e.target.innerText))
      // console.log(e)
      // console.log(id1, id2, id3)
      // console.log(e.target.parentNode.children)
      let children = e.target.parentNode.children
      if(this.isSudokuFull(children)) {
        let newArr = []
        for (let i = 0; i < 9; i++) {
          newArr.push(parseInt(children[i].innerText))
        }
        this.isSudokuOk(newArr, e.target.style)
      }
    },
    async solve() {
      let res = await solveSudoku()
      this.SUDOKU = JSON.parse(res.data)
    },
    isInput(e, id1, id2, id3) {
      if(e.target.innerText != "") {
        return
      } else {
        e.target.contentEditable = true
      }
    },
    toChangeStyle(id1, id2, id3, bgc) {
      return {
        backgroundColor: bgc,
      }
    },
    isSudokuFull(children) {
      // 判断九宫格是否填满
      for (let i = 0; i < 9; i++) {
        if(! (parseInt(children[i].innerText) > 0 && parseInt(children[i].innerText) < 10)) {
          return false
        }
      }
      return true
    },
    isSudokuOk(arr, style) {
      // 判断该9宫格是否正确
      var sum = arr.reduce((pre, cur) => {
        return pre + cur;
      })
      if (sum != 45) {
        style.backgroundColor = "rgba(255, 0, 0, 0.7)"
        return
      } else {
        style.backgroundColor = "rgba(255, 255, 255, 0.9)"
        return
      }
    },
    mouseover(e, id1, id2, id3) {
      // console.log(e, id1, id2, id3)
      // console.log(e.target.parentNode.children[0].style);
      let flag = 0
      let ind = 0
      for (let i = 0; i < 9; i++) {
        if (e.target.parentNode.children[i].style.backgroundColor === "rgba(255, 0, 0, 0.7)") {
          flag = 1
          ind = i
        }
      }
      for (let i = 0; i < 9; i++) {
        if(flag === 1) {
          e.target.parentNode.children[i].style.backgroundColor =  "rgba(255, 255, 255, 0.7)"
          e.target.parentNode.children[ind].style.backgroundColor = "rgba(255, 0, 0, 0.7)"
          continue
        }
        e.target.parentNode.children[i].style.backgroundColor =  "rgba(255, 255, 255, 0.7)"
      }
      if (id2 >= 0 && id2 <= 2) {
        if (id3 >= 0 && id3 <=2) {
          for (let i = 0; i <= 2; i++) {
            for (let j = 0; j <= 2; j++) {
              this.$refs[this.idx(id1, i, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
            }
          }
        } else if (id3 >= 3 && id3 <=5) {
          for (let i = 0; i <= 2; i++) {
            for (let j = 3; j <= 5; j++) {
              this.$refs[this.idx(id1, i, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
            }
          }
        } else if (id3 >= 6 && id3 <=8) {
          for (let i = 0; i <= 2; i++) {
            for (let j = 6; j <= 8; j++) {
              this.$refs[this.idx(id1, i, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
            }
          }
        }
      } else if (id2 >= 3 && id2 <=5) {
        if (id3 >= 0 && id3 <=2) {
          for (let i = 3; i <= 5; i++) {
            for (let j = 0; j <= 2; j++) {
              this.$refs[this.idx(id1, i, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
            }
          }
        } else if (id3 >= 3 && id3 <=5) {
          for (let i = 3; i <= 5; i++) {
            for (let j = 3; j <= 5; j++) {
              this.$refs[this.idx(id1, i, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
            }
          }
        } else if (id3 >= 6 && id3 <=8) {
          for (let i = 3; i <= 5; i++) {
            for (let j = 6; j <= 8; j++) {
              this.$refs[this.idx(id1, i, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
            }
          }
        }
      } else if (id2 >= 6 && id2 <=8) {
        if (id3 >= 0 && id3 <=2) {
          for (let i = 6; i <= 8; i++) {
            for (let j = 0; j <= 2; j++) {
              this.$refs[this.idx(id1, i, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
            }
          }
        } else if (id3 >= 3 && id3 <=5) {
          for (let i = 6; i <= 8; i++) {
            for (let j = 3; j <= 5; j++) {
              this.$refs[this.idx(id1, i, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
            }
          }
        } else if (id3 >= 6 && id3 <=8) {
          for (let i = 6; i <= 8; i++) {
            for (let j = 6; j <= 8; j++) {
              this.$refs[this.idx(id1, i, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
            }
          }
        }
      }

      if (id2 == 0 || id2 == 3 || id2 == 6) {
          if (id3 == 0 || id3 == 3 || id3 == 6) {
            for (let j = 0; j <= 8; j += 3) {
              this.$refs[this.idx(id1, 0, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
              this.$refs[this.idx(id1, 3, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
              this.$refs[this.idx(id1, 6, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
            }
          } else if (id3 == 1 || id3 == 4 || id3 == 7) {
            for (let j = 1; j <= 8; j += 3) {
              this.$refs[this.idx(id1, 0, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
              this.$refs[this.idx(id1, 3, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
              this.$refs[this.idx(id1, 6, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
            }
          } else {
            for (let j = 2; j <= 8; j += 3) {
              this.$refs[this.idx(id1, 0, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
              this.$refs[this.idx(id1, 3, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
              this.$refs[this.idx(id1, 6, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
            }
          }
      } else if (id2 == 1 || id2 == 4 || id2 == 7) {
          if (id3 == 0 || id3 == 3 || id3 == 6) {
            for (let j = 0; j <= 8; j += 3) {
              this.$refs[this.idx(id1, 1, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
              this.$refs[this.idx(id1, 4, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
              this.$refs[this.idx(id1, 7, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
            }
          } else if (id3 == 1 || id3 == 4 || id3 == 7) {
            for (let j = 0; j <= 8; j += 3) {
              this.$refs[this.idx(id1, 1, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
              this.$refs[this.idx(id1, 4, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
              this.$refs[this.idx(id1, 7, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
            }
          } else if (id3 == 2 || id3 == 5 || id3 == 8) {
            for (let j = 0; j <= 8; j += 3) {
              this.$refs[this.idx(id1, 1, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
              this.$refs[this.idx(id1, 4, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
              this.$refs[this.idx(id1, 7, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
            }
          }
      } else {
          if (id3 == 0 || id3 == 3 || id3 == 6) {
            for (let j = 0; j <= 8; j += 3) {
              this.$refs[this.idx(id1, 2, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
              this.$refs[this.idx(id1, 5, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
              this.$refs[this.idx(id1, 8, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
            }
          } else if (id3 == 1 || id3 == 4 || id3 == 7) {
            for (let j = 0; j <= 8; j += 3) {
              this.$refs[this.idx(id1, 2, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
              this.$refs[this.idx(id1, 5, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
              this.$refs[this.idx(id1, 8, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
            }
          } else {
            for (let j = 0; j <= 8; j += 3) {
              this.$refs[this.idx(id1, 2, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
              this.$refs[this.idx(id1, 5, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
              this.$refs[this.idx(id1, 8, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
            }
          }
      }
      // this.SUDOKU[id1][id2].forEach(function (item, index, array) {
      //   // item数组中的当前项, index当前项的索引, array原始数组；
      // })
    },
    mouseout(e, id1, id2, id3) {
      let flag = 0
      let ind = 0
      for (let i = 0; i < 9; i++) {
        if (e.target.parentNode.children[i].style.backgroundColor === "rgba(255, 0, 0, 0.7)") {
          flag = 1
          ind = i
        }
      }
      for (let i = 0; i < 9; i++) {
        if(flag === 1) {
          e.target.parentNode.children[i].style.backgroundColor =  "rgba(255, 255, 255, 0.9)"
          e.target.parentNode.children[ind].style.backgroundColor = "rgba(255, 0, 0, 0.7)"
          continue
        }
        e.target.parentNode.children[i].style.backgroundColor =  "rgba(255, 255, 255, 0.9)"
      }
      if (id2 >= 0 && id2 <=2) {
          if (id3 >= 0 && id3 <=2) {
            for (let i = 0; i <= 2; i++) {
              for (let j = 0; j <= 2; j++) {
                this.$refs[this.idx(id1, i, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              }
            }
          } else if (id3 >= 3 && id3 <=5) {
            for (let i = 0; i <= 2; i++) {
              for (let j = 3; j <= 5; j++) {
                this.$refs[this.idx(id1, i, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              }
            }
          } else if (id3 >= 6 && id3 <=8) {
            for (let i = 0; i <= 2; i++) {
              for (let j = 6; j <= 8; j++) {
                this.$refs[this.idx(id1, i, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              }
            }
          }
      } else if (id2 >= 3 && id2 <=5) {
        if (id3 >= 0 && id3 <=2) {
            for (let i = 3; i <= 5; i++) {
              for (let j = 0; j <= 2; j++) {
                this.$refs[this.idx(id1, i, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              }
            }
          } else if (id3 >= 3 && id3 <=5) {
            for (let i = 3; i <= 5; i++) {
              for (let j = 3; j <= 5; j++) {
                this.$refs[this.idx(id1, i, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              }
            }
          } else if (id3 >= 6 && id3 <=8) {
            for (let i = 3; i <= 5; i++) {
              for (let j = 6; j <= 8; j++) {
                this.$refs[this.idx(id1, i, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              }
            }
          }
      } else if (id2 >= 6 && id2 <=8) {
        if (id3 >= 0 && id3 <=2) {
            for (let i = 6; i <= 8; i++) {
              for (let j = 0; j <= 2; j++) {
                this.$refs[this.idx(id1, i, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              }
            }
          } else if (id3 >= 3 && id3 <=5) {
            for (let i = 6; i <= 8; i++) {
              for (let j = 3; j <= 5; j++) {
                this.$refs[this.idx(id1, i, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              }
            }
          } else if (id3 >= 6 && id3 <=8) {
            for (let i = 6; i <= 8; i++) {
              for (let j = 6; j <= 8; j++) {
                this.$refs[this.idx(id1, i, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              }
            }
          }
      }

      if (id2 == 0 || id2 == 3 || id2 == 6) {
        if (id3 == 0 || id3 == 3 || id3 == 6) {
          for (let j = 0; j <= 8; j += 3) {
            this.$refs[this.idx(id1, 0, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
            this.$refs[this.idx(id1, 3, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
            this.$refs[this.idx(id1, 6, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
          }
        } else if (id3 == 1 || id3 == 4 || id3 == 7) {
          for (let j = 1; j <= 8; j += 3) {
            this.$refs[this.idx(id1, 0, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
            this.$refs[this.idx(id1, 3, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
            this.$refs[this.idx(id1, 6, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
          }
        } else {
          for (let j = 2; j <= 8; j += 3) {
            this.$refs[this.idx(id1, 0, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
            this.$refs[this.idx(id1, 3, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
            this.$refs[this.idx(id1, 6, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
          }
        }
      } else if (id2 == 1 || id2 == 4 || id2 == 7) {
        if (id3 == 0 || id3 == 3 || id3 == 6) {
          for (let j = 0; j <= 8; j += 3) {
            this.$refs[this.idx(id1, 1, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
            this.$refs[this.idx(id1, 4, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
            this.$refs[this.idx(id1, 7, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
          }
        } else if (id3 == 1 || id3 == 4 || id3 == 7) {
          for (let j = 0; j <= 8; j += 3) {
            this.$refs[this.idx(id1, 1, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
            this.$refs[this.idx(id1, 4, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
            this.$refs[this.idx(id1, 7, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
          }
        } else if (id3 == 2 || id3 == 5 || id3 == 8) {
          for (let j = 0; j <= 8; j += 3) {
            this.$refs[this.idx(id1, 1, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
            this.$refs[this.idx(id1, 4, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
            this.$refs[this.idx(id1, 7, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
          }
        }
      } else {
        if (id3 == 0 || id3 == 3 || id3 == 6) {
          for (let j = 0; j <= 8; j += 3) {
            this.$refs[this.idx(id1, 2, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
            this.$refs[this.idx(id1, 5, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
            this.$refs[this.idx(id1, 8, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
          }
        } else if (id3 == 1 || id3 == 4 || id3 == 7) {
          for (let j = 0; j <= 8; j += 3) {
            this.$refs[this.idx(id1, 2, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
            this.$refs[this.idx(id1, 5, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
            this.$refs[this.idx(id1, 8, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
          }
        } else {
          for (let j = 0; j <= 8; j += 3) {
            this.$refs[this.idx(id1, 2, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
            this.$refs[this.idx(id1, 5, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
            this.$refs[this.idx(id1, 8, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
          }
        }
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
    idx(id1, id2, id3) {
      return id1+ "-" + id2 + "-" + id3
    },
    async getsudoku () {
      let res = await getSudoku(this.lever)
      this.SUDOKU = res.data
      this.su_copy = JSON.parse(JSON.stringify(res.data))
    },
    again() {
      this.SUDOKU = []
      this.SUDOKU = JSON.parse(JSON.stringify(this.su_copy))
      this.key = new Date().getTime()
    },
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
