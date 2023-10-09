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
      su_copy1: [],
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
    },
    bgcolor() {
      let idx = this.bgc.indexOf('#')
      return this.bgc.slice(idx, idx + 7)
    },
  },
  methods: {
    inputText(e, id1, id2, id3) {
      this.SUDOKU[id1][id2][id3] = e.target.innerText
      let children = e.target.parentNode.children
      if(this.isSudokuFull(children)) {
        let newArr = []
        for (let i = 0; i < 9; i++) {
          newArr.push(parseInt(children[i].innerText))
        }
        this.isSudokuOk(newArr, e.target.style)
      } else {
        let index = this.isRepeatNumber(e, id1, id2, id3)
        if(index != -1) {
          e.target.style.backgroundColor = "rgba(255, 0, 0, 0.7)"
          console.log("sudoku_repeat")
        } else {
          e.target.style.backgroundColor = "rgba(255, 255, 255, 0.9)"
        }
      }
      let ind = this.isRowRepeat(id1, id2, id3)
      var ind2 = 0
      var ind3 = 0
      if(ind != -1) {
        if(id2 >= 0 && id2 <= 2) {
          ind2 = parseInt(ind/3)
        } else if(id2 >= 3 && id2 <= 5) {
          ind2 = parseInt(ind/3) + 3
        } else {
          ind2 = parseInt(ind/3) + 6
        }
        if(id3 >= 0 && id3 <= 2) {
          ind3 = ind%3
        } else if(id3 >= 3 && id3 <= 5) {
          ind3 = ind%3 + 3
        } else {
          ind3 = ind%3 + 6
        }
        console.log(ind, "row_repeat");
        this.$refs[this.idx(id1, ind2, ind3)][0].style.backgroundColor = "rgba(255, 0, 0, 0.7)"
      } else {
        this.$refs[this.idx(id1, ind2, ind3)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
      }

      let inde = this.isColRepeat(id1, id2, id3)
      var ind4 = 0
      var ind5 = 0
      let a = [0, 3, 6]
      let b = [1, 4, 7]
      let c = [2, 5, 8]
      if(inde != -1) {
        if(id2 == 0 || id2 == 3 || id2 == 6) {
          ind4 = a[parseInt(inde/3)]
        } else if(id2 == 1 || id2 == 4 || id2 == 7) {
          ind4 = b[parseInt(inde/3)]
        } else {
          ind4 = c[parseInt(inde/3)]
        }
        if(id3 == 0 || id3 == 3 || id3 == 6) {
          ind5 = a[inde%3]
        } else if(id3 == 1 || id3 == 4 || id3 == 7) {
          ind5 = b[inde%3]
        } else {
          ind5 = c[inde%3]
        }
        console.log(inde, "col_repeat");
        this.$refs[this.idx(id1, ind4, ind5)][0].style.backgroundColor = "rgba(255, 0, 0, 0.7)"
      } else {
        this.$refs[this.idx(id1, ind4, ind5)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
      }
      
    },
    async solve() {
      let res = await solveSudoku()
      let data = JSON.parse(res.data)
      console.log(data)
      let arr = []
      let arr1 = []
      let arr2 = []
      for (let i = 0; i < 9; i++) { // 处理九个数独
        for(let j = 0; j < 3; j++) { // 处理1个数独的1个九宫格
          for(let k = 0; k < 3; k++) {
            arr2.push(data[i][j][k])
          }
        }
        arr1.push(arr2)
        arr2 = []
        for(let j = 0; j < 3; j++) { // 处理1个数独的1个九宫格
          for(let k = 3; k < 6; k++) {
            arr2.push(data[i][j][k])
          }
        }
        arr1.push(arr2)
        arr2 = []
        for(let j = 0; j < 3; j++) { // 处理1个数独的1个九宫格
          for(let k = 6; k < 9; k++) {
            arr2.push(data[i][j][k])
          }
        }
        arr1.push(arr2)
        arr2 = []
        for(let j = 3; j < 6; j++) { // 处理1个数独的1个九宫格
          for(let k = 0; k < 3; k++) {
            arr2.push(data[i][j][k])
          }
        }
        arr1.push(arr2)
        arr2 = []
        for(let j = 3; j < 6; j++) { // 处理1个数独的1个九宫格
          for(let k = 3; k < 6; k++) {
            arr2.push(data[i][j][k])
          }
        }
        arr1.push(arr2)
        arr2 = []
        for(let j = 3; j < 6; j++) { // 处理1个数独的1个九宫格
          for(let k = 6; k < 9; k++) {
            arr2.push(data[i][j][k])
          }
        }
        arr1.push(arr2)
        arr2 = []
        for(let j = 6; j < 9; j++) { // 处理1个数独的1个九宫格
          for(let k = 0; k < 3; k++) {
            arr2.push(data[i][j][k])
          }
        }
        arr1.push(arr2)
        arr2 = []
        for(let j = 6; j < 9; j++) { // 处理1个数独的1个九宫格
          for(let k = 3; k < 6; k++) {
            arr2.push(data[i][j][k])
          }
        }
        arr1.push(arr2)
        arr2 = []
        for(let j = 6; j < 9; j++) { // 处理1个数独的1个九宫格
          for(let k = 6; k < 9; k++) {
            arr2.push(data[i][j][k])
          }
        }
        arr1.push(arr2)
        arr2 = []
        arr.push(arr1)
        arr1 = []
      }
      this.again()
      this.SUDOKU = arr
    },
    isInput(e, id1, id2, id3) {
      if(e.target.innerText != "") {
        return
      } else {
        e.target.contentEditable = true
        e.target.style.color = this.bgcolor
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
    isRepeatNumber(e, id1, id2, id3) {
      let number = parseInt(e.target.innerText)
      let arr = this.SudokutoArray(id1, id2)
      let index = arr.indexOf(number)
      if(index == id3) {
        arr[index] = 0
        index = arr.indexOf(number)
      }
      if(index == -1){
        return -1
      } else {
        return index
      }
    },
    SudokutoArray(id1, id2) {
      let arr = []
      for (let i = 0; i < 9; i++) {
        if (this.$refs[this.idx(id1, id2, i)][0].innerText == "") {
          arr.push(0)
        } else {
          arr.push(parseInt(this.$refs[this.idx(id1, id2, i)][0].innerText))
        }
      }
      return arr
    },
    isRowRepeat(id1, id2, id3) {
      let arr = this.RowtoArray(id1, id2, id3)
      let number = parseInt(this.$refs[this.idx(id1, id2, id3)][0].innerText)
      let index = arr.indexOf(number)
      let cmp = 0
      if(id3 >= 0 && id3 <= 2) {
        cmp = index%3
      } else if(id3 >= 3 && id3 <= 5) {
        cmp = index%3 + 3
      } else {
        cmp = index%3 + 6
      }
      console.log(cmp, "rowcmp")
      if(id3 == cmp) {
        arr[index] = 0
        index = arr.indexOf(number)
        console.log(index, "rowindex");
      }
      if(index == -1){
        return -1
      } else {
        return index
      }
    },
    RowtoArray(id1, id2, id3) {
      let arr = []
      let a = 6
      let b = 6
      if(id2 >= 0 && id2 <=2) {
        a = 0
      } else if(id2 >= 3 && id2 <= 5) {
        a = 3
      }
      if(id3 >= 0 && id3 <=2) {
        b = 0
      } else if(id3 >= 3 && id3 <= 5) {
        b = 3
      }
      for (let i = a; i < a + 3; i++) {
        for (let j = b; j < b+3; j++) {
          if (this.$refs[this.idx(id1, i, j)][0].innerText == "") {
            arr.push(0)
          } else {
            arr.push(parseInt(this.$refs[this.idx(id1, i, j)][0].innerText))
          }
        }
      }
      return arr
    },
    isColRepeat(id1, id2, id3) {
      let arr = this.ColtoArray(id1, id2, id3)
      let number = parseInt(this.$refs[this.idx(id1, id2, id3)][0].innerText)
      let index = arr.indexOf(number)
      let cmp = 0
      if(id2 == 0 || id2 == 1 || id2 == 2) {
        cmp = id3/3
      } else if(id2 == 3 || id2 == 4 || id2 == 5) {
        cmp = id3/3 + 3
      } else {
        cmp = id3/3 + 6
      }
      console.log(cmp, "colcmp")
      if(index == cmp) {
        arr[index] = 0
        index = arr.indexOf(number)
        console.log(arr)
        console.log(index, "colindex")
      }
      if(index == -1){
        return -1
      } else {
        return index
      }
    },
    ColtoArray(id1, id2, id3) {
      let arr = []
      let a = 0
      let b = 0
      if(id2 == 0 || id2 == 3 || id2 == 6) {
        a = 0
      } else if(id2 == 1 || id2 == 4 || id2 == 7) {
        a = 1
      } else {
        a = 2
      }
      if(id3 == 0 || id3 == 3 || id3 == 6) {
        b = 0
      } else if(id3 == 1 || id3 == 4 || id3 == 7) {
        b = 1
      } else {
        b = 2
      }
      for (let i = a; i < 9; i += 3) {
        for (let j = b; j < 9; j += 3) {
          if (this.$refs[this.idx(id1, i, j)][0].innerText == "") {
            arr.push(0)
          } else {
            arr.push(parseInt(this.$refs[this.idx(id1, i, j)][0].innerText))
          }
        }
      }
      return arr
    },
    isError(id1, id2, id3) {
      return this.$refs[this.idx(id1, id2, id3)][0].style.backgroundColor === "rgba(255, 0, 0, 0.7)"
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
        if (id3 >= 0 && id3 <= 2) {
          for (let i = 0; i <= 2; i++) {
            for (let j = 0; j <= 2; j++) {
              if(this.isError(id1, i, j)) {
                continue
              }
              this.$refs[this.idx(id1, i, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
            }
          }
        } else if (id3 >= 3 && id3 <= 5) {
          for (let i = 0; i <= 2; i++) {
            for (let j = 3; j <= 5; j++) {
              if(this.isError(id1, i, j)) {
                continue
              }
              this.$refs[this.idx(id1, i, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
            }
          }
        } else if (id3 >= 6 && id3 <=8) {
          for (let i = 0; i <= 2; i++) {
            for (let j = 6; j <= 8; j++) {
              if(this.isError(id1, i, j)) {
                continue
              }
              this.$refs[this.idx(id1, i, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
            }
          }
        }
      } else if (id2 >= 3 && id2 <= 5) {
        if (id3 >= 0 && id3 <=2) {
          for (let i = 3; i <= 5; i++) {
            for (let j = 0; j <= 2; j++) {
              if(this.isError(id1, i, j)) {
                continue
              }
              this.$refs[this.idx(id1, i, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
            }
          }
        } else if (id3 >= 3 && id3 <= 5) {
          for (let i = 3; i <= 5; i++) {
            for (let j = 3; j <= 5; j++) {
              if(this.isError(id1, i, j)) {
                continue
              }
              this.$refs[this.idx(id1, i, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
            }
          }
        } else if (id3 >= 6 && id3 <= 8) {
          for (let i = 3; i <= 5; i++) {
            for (let j = 6; j <= 8; j++) {
              if(this.isError(id1, i, j)) {
                continue
              }
              this.$refs[this.idx(id1, i, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
            }
          }
        }
      } else if (id2 >= 6 && id2 <= 8) {
        if (id3 >= 0 && id3 <=2) {
          for (let i = 6; i <= 8; i++) {
            for (let j = 0; j <= 2; j++) {
              if(this.isError(id1, i, j)) {
                continue
              }
              this.$refs[this.idx(id1, i, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
            }
          }
        } else if (id3 >= 3 && id3 <= 5) {
          for (let i = 6; i <= 8; i++) {
            for (let j = 3; j <= 5; j++) {
              if(this.isError(id1, i, j)) {
                continue
              }
              this.$refs[this.idx(id1, i, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
            }
          }
        } else if (id3 >= 6 && id3 <= 8) {
          for (let i = 6; i <= 8; i++) {
            for (let j = 6; j <= 8; j++) {
              if(this.isError(id1, i, j)) {
                continue
              }
              this.$refs[this.idx(id1, i, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
            }
          }
        }
      }

      if (id2 == 0 || id2 == 3 || id2 == 6) {
          if (id3 == 0 || id3 == 3 || id3 == 6) {
            for (let j = 0; j <= 8; j += 3) {
              if(this.isError(id1, 0, j)) {
                this.$refs[this.idx(id1, 3, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                this.$refs[this.idx(id1, 6, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                continue
              } else if(this.isError(id1, 3, j)) {
                this.$refs[this.idx(id1, 0, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                this.$refs[this.idx(id1, 6, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                continue
              } else if(this.isError(id1, 6, j)) {
                this.$refs[this.idx(id1, 3, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                this.$refs[this.idx(id1, 0, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                continue
              }
              this.$refs[this.idx(id1, 0, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
              this.$refs[this.idx(id1, 3, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
              this.$refs[this.idx(id1, 6, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
            }
          } else if (id3 == 1 || id3 == 4 || id3 == 7) {
            for (let j = 1; j <= 8; j += 3) {
              if(this.isError(id1, 0, j)) {
                this.$refs[this.idx(id1, 3, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                this.$refs[this.idx(id1, 6, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                continue
              } else if(this.isError(id1, 3, j)) {
                this.$refs[this.idx(id1, 0, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                this.$refs[this.idx(id1, 6, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                continue
              } else if(this.isError(id1, 6, j)) {
                this.$refs[this.idx(id1, 3, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                this.$refs[this.idx(id1, 0, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                continue
              }
              this.$refs[this.idx(id1, 0, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
              this.$refs[this.idx(id1, 3, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
              this.$refs[this.idx(id1, 6, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
            }
          } else {
            for (let j = 2; j <= 8; j += 3) {
              if(this.isError(id1, 0, j)) {
                this.$refs[this.idx(id1, 3, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                this.$refs[this.idx(id1, 6, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                continue
              } else if(this.isError(id1, 3, j)) {
                this.$refs[this.idx(id1, 0, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                this.$refs[this.idx(id1, 6, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                continue
              } else if(this.isError(id1, 6, j)) {
                this.$refs[this.idx(id1, 3, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                this.$refs[this.idx(id1, 0, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                continue
              }
              this.$refs[this.idx(id1, 0, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
              this.$refs[this.idx(id1, 3, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
              this.$refs[this.idx(id1, 6, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
            }
          }
      } else if (id2 == 1 || id2 == 4 || id2 == 7) {
          if (id3 == 0 || id3 == 3 || id3 == 6) {
            for (let j = 0; j <= 8; j += 3) {
              if(this.isError(id1, 1, j)) {
                this.$refs[this.idx(id1, 4, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                this.$refs[this.idx(id1, 7, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                continue
              } else if(this.isError(id1, 4, j)) {
                this.$refs[this.idx(id1, 1, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                this.$refs[this.idx(id1, 7, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                continue
              } else if(this.isError(id1, 7, j)) {
                this.$refs[this.idx(id1, 1, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                this.$refs[this.idx(id1, 4, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                continue
              }
              this.$refs[this.idx(id1, 1, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
              this.$refs[this.idx(id1, 4, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
              this.$refs[this.idx(id1, 7, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
            }
          } else if (id3 == 1 || id3 == 4 || id3 == 7) {
            for (let j = 1; j <= 8; j += 3) {
              if(this.isError(id1, 1, j)) {
                this.$refs[this.idx(id1, 4, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                this.$refs[this.idx(id1, 7, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                continue
              } else if(this.isError(id1, 4, j)) {
                this.$refs[this.idx(id1, 1, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                this.$refs[this.idx(id1, 7, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                continue
              } else if(this.isError(id1, 7, j)) {
                this.$refs[this.idx(id1, 1, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                this.$refs[this.idx(id1, 4, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                continue
              }
              this.$refs[this.idx(id1, 1, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
              this.$refs[this.idx(id1, 4, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
              this.$refs[this.idx(id1, 7, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
            }
          } else if (id3 == 2 || id3 == 5 || id3 == 8) {
            for (let j = 2; j <= 8; j += 3) {
              if(this.isError(id1, 1, j)) {
                this.$refs[this.idx(id1, 4, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                this.$refs[this.idx(id1, 7, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                continue
              } else if(this.isError(id1, 4, j)) {
                this.$refs[this.idx(id1, 1, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                this.$refs[this.idx(id1, 7, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                continue
              } else if(this.isError(id1, 7, j)) {
                this.$refs[this.idx(id1, 1, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                this.$refs[this.idx(id1, 4, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                continue
              }
              this.$refs[this.idx(id1, 1, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
              this.$refs[this.idx(id1, 4, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
              this.$refs[this.idx(id1, 7, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
            }
          }
      } else {
          if (id3 == 0 || id3 == 3 || id3 == 6) {
            for (let j = 0; j <= 8; j += 3) {
              if(this.isError(id1, 2, j)) {
                this.$refs[this.idx(id1, 5, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                this.$refs[this.idx(id1, 8, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                continue
              } else if(this.isError(id1, 5, j)) {
                this.$refs[this.idx(id1, 2, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                this.$refs[this.idx(id1, 8, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                continue
              } else if(this.isError(id1, 8, j)) {
                this.$refs[this.idx(id1, 2, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                this.$refs[this.idx(id1, 5, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                continue
              }
              this.$refs[this.idx(id1, 2, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
              this.$refs[this.idx(id1, 5, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
              this.$refs[this.idx(id1, 8, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
            }
          } else if (id3 == 1 || id3 == 4 || id3 == 7) {
            for (let j = 1; j <= 8; j += 3) {
              if(this.isError(id1, 2, j)) {
                this.$refs[this.idx(id1, 5, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                this.$refs[this.idx(id1, 8, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                continue
              } else if(this.isError(id1, 5, j)) {
                this.$refs[this.idx(id1, 2, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                this.$refs[this.idx(id1, 8, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                continue
              } else if(this.isError(id1, 8, j)) {
                this.$refs[this.idx(id1, 2, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                this.$refs[this.idx(id1, 5, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                continue
              }
              this.$refs[this.idx(id1, 2, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
              this.$refs[this.idx(id1, 5, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
              this.$refs[this.idx(id1, 8, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
            }
          } else {
            for (let j = 2; j <= 8; j += 3) {
              if(this.isError(id1, 2, j)) {
                this.$refs[this.idx(id1, 5, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                this.$refs[this.idx(id1, 8, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                continue
              } else if(this.isError(id1, 5, j)) {
                this.$refs[this.idx(id1, 2, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                this.$refs[this.idx(id1, 8, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                continue
              } else if(this.isError(id1, 8, j)) {
                this.$refs[this.idx(id1, 2, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                this.$refs[this.idx(id1, 5, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.7)"
                continue
              }
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
                if(this.isError(id1, i, j)) {
                  continue
                }
                this.$refs[this.idx(id1, i, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              }
            }
          } else if (id3 >= 3 && id3 <=5) {
            for (let i = 0; i <= 2; i++) {
              for (let j = 3; j <= 5; j++) {
                if(this.isError(id1, i, j)) {
                  continue
                }
                this.$refs[this.idx(id1, i, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              }
            }
          } else if (id3 >= 6 && id3 <=8) {
            for (let i = 0; i <= 2; i++) {
              for (let j = 6; j <= 8; j++) {
                if(this.isError(id1, i, j)) {
                  continue
                }
                this.$refs[this.idx(id1, i, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              }
            }
          }
      } else if (id2 >= 3 && id2 <=5) {
        if (id3 >= 0 && id3 <=2) {
            for (let i = 3; i <= 5; i++) {
              for (let j = 0; j <= 2; j++) {
                if(this.isError(id1, i, j)) {
                  continue
                }
                this.$refs[this.idx(id1, i, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              }
            }
          } else if (id3 >= 3 && id3 <=5) {
            for (let i = 3; i <= 5; i++) {
              for (let j = 3; j <= 5; j++) {
                if(this.isError(id1, i, j)) {
                  continue
                }
                this.$refs[this.idx(id1, i, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              }
            }
          } else if (id3 >= 6 && id3 <=8) {
            for (let i = 3; i <= 5; i++) {
              for (let j = 6; j <= 8; j++) {
                if(this.isError(id1, i, j)) {
                  continue
                }
                this.$refs[this.idx(id1, i, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              }
            }
          }
      } else if (id2 >= 6 && id2 <=8) {
        if (id3 >= 0 && id3 <=2) {
            for (let i = 6; i <= 8; i++) {
              for (let j = 0; j <= 2; j++) {
                if(this.isError(id1, i, j)) {
                  continue
                }
                this.$refs[this.idx(id1, i, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              }
            }
          } else if (id3 >= 3 && id3 <=5) {
            for (let i = 6; i <= 8; i++) {
              for (let j = 3; j <= 5; j++) {
                if(this.isError(id1, i, j)) {
                  continue
                }
                this.$refs[this.idx(id1, i, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              }
            }
          } else if (id3 >= 6 && id3 <=8) {
            for (let i = 6; i <= 8; i++) {
              for (let j = 6; j <= 8; j++) {
                if(this.isError(id1, i, j)) {
                  continue
                }
                this.$refs[this.idx(id1, i, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              }
            }
          }
      }

      if (id2 == 0 || id2 == 3 || id2 == 6) {
        if (id3 == 0 || id3 == 3 || id3 == 6) {
          for (let j = 0; j <= 8; j += 3) {
            if(this.isError(id1, 0, j)) {
              this.$refs[this.idx(id1, 3, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              this.$refs[this.idx(id1, 6, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              continue
            } else if(this.isError(id1, 3, j)) {
              this.$refs[this.idx(id1, 0, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              this.$refs[this.idx(id1, 6, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              continue
            } else if(this.isError(id1, 6, j)) {
              this.$refs[this.idx(id1, 3, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              this.$refs[this.idx(id1, 0, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              continue
            }
            this.$refs[this.idx(id1, 0, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
            this.$refs[this.idx(id1, 3, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
            this.$refs[this.idx(id1, 6, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
          }
        } else if (id3 == 1 || id3 == 4 || id3 == 7) {
          for (let j = 1; j <= 8; j += 3) {
            if(this.isError(id1, 0, j)) {
              this.$refs[this.idx(id1, 3, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              this.$refs[this.idx(id1, 6, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              continue
            } else if(this.isError(id1, 3, j)) {
              this.$refs[this.idx(id1, 0, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              this.$refs[this.idx(id1, 6, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              continue
            } else if(this.isError(id1, 6, j)) {
              this.$refs[this.idx(id1, 3, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              this.$refs[this.idx(id1, 0, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              continue
            }
            this.$refs[this.idx(id1, 0, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
            this.$refs[this.idx(id1, 3, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
            this.$refs[this.idx(id1, 6, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
          }
        } else {
          for (let j = 2; j <= 8; j += 3) {
            if(this.isError(id1, 0, j)) {
              this.$refs[this.idx(id1, 3, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              this.$refs[this.idx(id1, 6, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              continue
            } else if(this.isError(id1, 3, j)) {
              this.$refs[this.idx(id1, 0, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              this.$refs[this.idx(id1, 6, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              continue
            } else if(this.isError(id1, 6, j)) {
              this.$refs[this.idx(id1, 3, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              this.$refs[this.idx(id1, 0, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              continue
            }
            this.$refs[this.idx(id1, 0, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
            this.$refs[this.idx(id1, 3, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
            this.$refs[this.idx(id1, 6, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
          }
        }
      } else if (id2 == 1 || id2 == 4 || id2 == 7) {
        if (id3 == 0 || id3 == 3 || id3 == 6) {
          for (let j = 0; j <= 8; j += 3) {
            if(this.isError(id1, 1, j)) {
              this.$refs[this.idx(id1, 4, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              this.$refs[this.idx(id1, 7, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              continue
            } else if(this.isError(id1, 4, j)) {
              this.$refs[this.idx(id1, 1, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              this.$refs[this.idx(id1, 7, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              continue
            } else if(this.isError(id1, 7, j)) {
              this.$refs[this.idx(id1, 1, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              this.$refs[this.idx(id1, 4, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              continue
            }
            this.$refs[this.idx(id1, 1, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
            this.$refs[this.idx(id1, 4, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
            this.$refs[this.idx(id1, 7, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
          }
        } else if (id3 == 1 || id3 == 4 || id3 == 7) {
          for (let j = 1; j <= 8; j += 3) {
            if(this.isError(id1, 1, j)) {
              this.$refs[this.idx(id1, 4, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              this.$refs[this.idx(id1, 7, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              continue
            } else if(this.isError(id1, 4, j)) {
              this.$refs[this.idx(id1, 1, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              this.$refs[this.idx(id1, 7, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              continue
            } else if(this.isError(id1, 7, j)) {
              this.$refs[this.idx(id1, 1, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              this.$refs[this.idx(id1, 4, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              continue
            }
            this.$refs[this.idx(id1, 1, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
            this.$refs[this.idx(id1, 4, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
            this.$refs[this.idx(id1, 7, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
          }
        } else if (id3 == 2 || id3 == 5 || id3 == 8) {
          for (let j = 2; j <= 8; j += 3) {
            if(this.isError(id1, 1, j)) {
              this.$refs[this.idx(id1, 4, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              this.$refs[this.idx(id1, 7, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              continue
            } else if(this.isError(id1, 4, j)) {
              this.$refs[this.idx(id1, 1, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              this.$refs[this.idx(id1, 7, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              continue
            } else if(this.isError(id1, 7, j)) {
              this.$refs[this.idx(id1, 1, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              this.$refs[this.idx(id1, 4, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              continue
            }
            this.$refs[this.idx(id1, 1, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
            this.$refs[this.idx(id1, 4, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
            this.$refs[this.idx(id1, 7, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
          }
        }
      } else {
        if (id3 == 0 || id3 == 3 || id3 == 6) {
          for (let j = 0; j <= 8; j += 3) {
            if(this.isError(id1, 2, j)) {
              this.$refs[this.idx(id1, 5, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              this.$refs[this.idx(id1, 8, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              continue
            } else if(this.isError(id1, 5, j)) {
              this.$refs[this.idx(id1, 2, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              this.$refs[this.idx(id1, 8, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              continue
            } else if(this.isError(id1, 8, j)) {
              this.$refs[this.idx(id1, 2, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              this.$refs[this.idx(id1, 5, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              continue
            }
            this.$refs[this.idx(id1, 2, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
            this.$refs[this.idx(id1, 5, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
            this.$refs[this.idx(id1, 8, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
          }
        } else if (id3 == 1 || id3 == 4 || id3 == 7) {
          for (let j = 1; j <= 8; j += 3) {
            if(this.isError(id1, 2, j)) {
              this.$refs[this.idx(id1, 5, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              this.$refs[this.idx(id1, 8, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              continue
            } else if(this.isError(id1, 5, j)) {
              this.$refs[this.idx(id1, 2, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              this.$refs[this.idx(id1, 8, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              continue
            } else if(this.isError(id1, 8, j)) {
              this.$refs[this.idx(id1, 2, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              this.$refs[this.idx(id1, 5, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              continue
            }
            this.$refs[this.idx(id1, 2, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
            this.$refs[this.idx(id1, 5, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
            this.$refs[this.idx(id1, 8, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
          }
        } else {
          for (let j = 2; j <= 8; j += 3) {
            if(this.isError(id1, 2, j)) {
              this.$refs[this.idx(id1, 5, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              this.$refs[this.idx(id1, 8, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              continue
            } else if(this.isError(id1, 5, j)) {
              this.$refs[this.idx(id1, 2, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              this.$refs[this.idx(id1, 8, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              continue
            } else if(this.isError(id1, 8, j)) {
              this.$refs[this.idx(id1, 2, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              this.$refs[this.idx(id1, 5, j)][0].style.backgroundColor = "rgba(255, 255, 255, 0.9)"
              continue
            }
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
      let data = res.data
      let arr = []
      let arr1 = []
      let arr2 = []
      for (let i = 0; i < 9; i++) { // 处理九个数独
        for(let j = 0; j < 3; j++) { // 处理1个数独的1个九宫格
          for(let k = 0; k < 3; k++) {
            arr2.push(data[i][j][k])
          }
        }
        arr1.push(arr2)
        arr2 = []
        for(let j = 0; j < 3; j++) { // 处理1个数独的1个九宫格
          for(let k = 3; k < 6; k++) {
            arr2.push(data[i][j][k])
          }
        }
        arr1.push(arr2)
        arr2 = []
        for(let j = 0; j < 3; j++) { // 处理1个数独的1个九宫格
          for(let k = 6; k < 9; k++) {
            arr2.push(data[i][j][k])
          }
        }
        arr1.push(arr2)
        arr2 = []
        for(let j = 3; j < 6; j++) { // 处理1个数独的1个九宫格
          for(let k = 0; k < 3; k++) {
            arr2.push(data[i][j][k])
          }
        }
        arr1.push(arr2)
        arr2 = []
        for(let j = 3; j < 6; j++) { // 处理1个数独的1个九宫格
          for(let k = 3; k < 6; k++) {
            arr2.push(data[i][j][k])
          }
        }
        arr1.push(arr2)
        arr2 = []
        for(let j = 3; j < 6; j++) { // 处理1个数独的1个九宫格
          for(let k = 6; k < 9; k++) {
            arr2.push(data[i][j][k])
          }
        }
        arr1.push(arr2)
        arr2 = []
        for(let j = 6; j < 9; j++) { // 处理1个数独的1个九宫格
          for(let k = 0; k < 3; k++) {
            arr2.push(data[i][j][k])
          }
        }
        arr1.push(arr2)
        arr2 = []
        for(let j = 6; j < 9; j++) { // 处理1个数独的1个九宫格
          for(let k = 3; k < 6; k++) {
            arr2.push(data[i][j][k])
          }
        }
        arr1.push(arr2)
        arr2 = []
        for(let j = 6; j < 9; j++) { // 处理1个数独的1个九宫格
          for(let k = 6; k < 9; k++) {
            arr2.push(data[i][j][k])
          }
        }
        arr1.push(arr2)
        arr2 = []
        arr.push(arr1)
        arr1 = []
      }
      this.SUDOKU = arr
      this.su_copy = JSON.parse(JSON.stringify(arr))
      this.su_copy1 = JSON.parse(JSON.stringify(this.su_copy))
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
