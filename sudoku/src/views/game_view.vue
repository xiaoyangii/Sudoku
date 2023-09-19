<template>
  <div class="game">
    <div class="game_header">
      <div class="game_header_return" @click="$router.go(-1)" title="返回！">
        <img src="../assets/images/return.png" alt="" title="返回！">
      </div>
      <div class="game_header_change" @click="changeColor()" title="更换主题！">
        <img src="../assets/images/change.png" alt="" title="更换主题！">
      </div>
    </div>
    <div class="game_body">
      <div class="maingrid">
        <div class="grid" :id="Sudoku.Sid" v-for="Sudoku in SUDOKU" :key="Sudoku.Sid">
          <div class="grid_block" :id="block.id" v-for="block in Sudoku.blocks" :key="block.id">
            <div class="grid_block_box" :id="box.sid" v-for="box in block.sudoku" :key="box.sid">{{ box.value }}</div>
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
export default {
  name: 'Game',
  components: {
    Footer
  },
  data () {
    return {
      SUDOKU: [
        {
          blocks: [
            { id: "one", sudoku: [ { value: "", sid: "one-1"}, { value: 9, sid: "one-2"}, { value: "", sid: "one-3"}, { value: 7, sid: "one-4"}, { value: 3, sid: "one-5"}, { value: "", sid: "one-6"}, { value: 8, sid: "one-7"}, { value: 6, sid: "one-8"}, { value: "", sid: "one-9"}] },
            { id: "two", sudoku: [ { value: "", sid: "two-1"}, { value: "3", sid: "two-2"}, { value: "1", sid: "two-3"}, { value: "", sid: "two-4"}, { value: "8", sid: "two-5"}, { value: "6", sid: "two-6"}, { value: "7", sid: "two-7"}, { value: "", sid: "two-8"}, { value: "2", sid: "two-9"}] },
            { id: "three", sudoku: [ { value: "", sid: "three-1"}, { value: "6", sid: "three-2"}, { value: "7", sid: "three-3"}, { value: "9", sid: "three-4"}, { value: "1", sid: "three-5"}, { value: "", sid: "three-6"}, { value: "3", sid: "three-7"}, { value: "4", sid: "three-8"}, { value: "", sid: "three-9"}] },
            { id: "four", sudoku: [ { value: "", sid: "four-1"}, { value: "7", sid: "four-2"}, { value: "9", sid: "four-3"}, { value: "5", sid: "four-4"}, { value: "", sid: "four-5"}, { value: "3", sid: "four-6"}, { value: "", sid: "four-7"}, { value: "", sid: "four-8"}, { value: "", sid: "four-9"}] },
            { id: "five", sudoku: [ { value: "", sid: "five-1"}, { value: "1", sid: "five-2"}, { value: "", sid: "five-3"}, { value: "", sid: "five-4"}, { value: "2", sid: "five-5"}, { value: "4", sid: "five-6"}, { value: "9", sid: "five-7"}, { value: "6", sid: "five-8"}, { value: "7", sid: "five-9"}] },
            { id: "six", sudoku: [ { value: "", sid: "six-1"}, { value: "2", sid: "six-2"}, { value: "8", sid: "six-3"}, { value: "", sid: "six-4"}, { value: "9", sid: "six-5"}, { value: "6", sid: "six-6"}, { value: "", sid: "six-7"}, { value: "5", sid: "six-8"}, { value: "3", sid: "six-9"}] },
            { id: "seven", sudoku: [ { value: "1", sid: "seven-1"}, { value: "", sid: "seven-2"}, { value: "7", sid: "seven-3"}, { value: "", sid: "seven-4"}, { value: "", sid: "seven-5"}, { value: "6", sid: "seven-6"}, { value: "9", sid: "seven-7"}, { value: "", sid: "seven-8"}, { value: "4", sid: "seven-9"}] },
            { id: "eight", sudoku: [ { value: "6", sid: "eight-1"}, { value: "5", sid: "eight-2"}, { value: "", sid: "eight-3"}, { value: "", sid: "eight-4"}, { value: "4", sid: "eight-5"}, { value: "", sid: "eight-6"}, { value: "2", sid: "eight-7"}, { value: 1, sid: "eight-8"}, { value: "", sid: "eight-9"}] },
            { id: "nine", sudoku: [ { value: "", sid: "nine-1"}, { value: "", sid: "nine-2"}, { value: "4", sid: "nine-3"}, { value: "5", sid: "nine-4"}, { value: "7", sid: "nine-5"}, { value: "9", sid: "nine-6"}, { value: "6", sid: "nine-7"}, { value: "", sid: "nine-8"}, { value: "", sid: "nine-9"}] },
          ],
          Sid: "first"
        },
        {
          blocks: [
            { id: "one", sudoku: [ { value: "", sid: "one-1"}, { value: 9, sid: "one-2"}, { value: "", sid: "one-3"}, { value: 7, sid: "one-4"}, { value: 3, sid: "one-5"}, { value: "", sid: "one-6"}, { value: 8, sid: "one-7"}, { value: 6, sid: "one-8"}, { value: "", sid: "one-9"}] },
            { id: "two", sudoku: [ { value: "", sid: "two-1"}, { value: "3", sid: "two-2"}, { value: "1", sid: "two-3"}, { value: "", sid: "two-4"}, { value: "8", sid: "two-5"}, { value: "6", sid: "two-6"}, { value: "7", sid: "two-7"}, { value: "", sid: "two-8"}, { value: "2", sid: "two-9"}] },
            { id: "three", sudoku: [ { value: "", sid: "three-1"}, { value: "6", sid: "three-2"}, { value: "7", sid: "three-3"}, { value: "9", sid: "three-4"}, { value: "1", sid: "three-5"}, { value: "", sid: "three-6"}, { value: "3", sid: "three-7"}, { value: "4", sid: "three-8"}, { value: "", sid: "three-9"}] },
            { id: "four", sudoku: [ { value: "", sid: "four-1"}, { value: "7", sid: "four-2"}, { value: "9", sid: "four-3"}, { value: "5", sid: "four-4"}, { value: "", sid: "four-5"}, { value: "3", sid: "four-6"}, { value: "", sid: "four-7"}, { value: "", sid: "four-8"}, { value: "", sid: "four-9"}] },
            { id: "five", sudoku: [ { value: "", sid: "five-1"}, { value: "1", sid: "five-2"}, { value: "", sid: "five-3"}, { value: "", sid: "five-4"}, { value: "2", sid: "five-5"}, { value: "4", sid: "five-6"}, { value: "9", sid: "five-7"}, { value: "6", sid: "five-8"}, { value: "7", sid: "five-9"}] },
            { id: "six", sudoku: [ { value: "", sid: "six-1"}, { value: "2", sid: "six-2"}, { value: "8", sid: "six-3"}, { value: "", sid: "six-4"}, { value: "9", sid: "six-5"}, { value: "6", sid: "six-6"}, { value: "", sid: "six-7"}, { value: "5", sid: "six-8"}, { value: "3", sid: "six-9"}] },
            { id: "seven", sudoku: [ { value: "1", sid: "seven-1"}, { value: "", sid: "seven-2"}, { value: "7", sid: "seven-3"}, { value: "", sid: "seven-4"}, { value: "", sid: "seven-5"}, { value: "6", sid: "seven-6"}, { value: "9", sid: "seven-7"}, { value: "", sid: "seven-8"}, { value: "4", sid: "seven-9"}] },
            { id: "eight", sudoku: [ { value: "6", sid: "eight-1"}, { value: "5", sid: "eight-2"}, { value: "", sid: "eight-3"}, { value: "", sid: "eight-4"}, { value: "4", sid: "eight-5"}, { value: "", sid: "eight-6"}, { value: "2", sid: "eight-7"}, { value: 1, sid: "eight-8"}, { value: "", sid: "eight-9"}] },
            { id: "nine", sudoku: [ { value: "", sid: "nine-1"}, { value: "", sid: "nine-2"}, { value: "4", sid: "nine-3"}, { value: "5", sid: "nine-4"}, { value: "7", sid: "nine-5"}, { value: "9", sid: "nine-6"}, { value: "6", sid: "nine-7"}, { value: "", sid: "nine-8"}, { value: "", sid: "nine-9"}] },
          ],
          Sid: "second"
        },
        {
          blocks: [
            { id: "one", sudoku: [ { value: "", sid: "one-1"}, { value: 9, sid: "one-2"}, { value: "", sid: "one-3"}, { value: 7, sid: "one-4"}, { value: 3, sid: "one-5"}, { value: "", sid: "one-6"}, { value: 8, sid: "one-7"}, { value: 6, sid: "one-8"}, { value: "", sid: "one-9"}] },
            { id: "two", sudoku: [ { value: "", sid: "two-1"}, { value: "3", sid: "two-2"}, { value: "1", sid: "two-3"}, { value: "", sid: "two-4"}, { value: "8", sid: "two-5"}, { value: "6", sid: "two-6"}, { value: "7", sid: "two-7"}, { value: "", sid: "two-8"}, { value: "2", sid: "two-9"}] },
            { id: "three", sudoku: [ { value: "", sid: "three-1"}, { value: "6", sid: "three-2"}, { value: "7", sid: "three-3"}, { value: "9", sid: "three-4"}, { value: "1", sid: "three-5"}, { value: "", sid: "three-6"}, { value: "3", sid: "three-7"}, { value: "4", sid: "three-8"}, { value: "", sid: "three-9"}] },
            { id: "four", sudoku: [ { value: "", sid: "four-1"}, { value: "7", sid: "four-2"}, { value: "9", sid: "four-3"}, { value: "5", sid: "four-4"}, { value: "", sid: "four-5"}, { value: "3", sid: "four-6"}, { value: "", sid: "four-7"}, { value: "", sid: "four-8"}, { value: "", sid: "four-9"}] },
            { id: "five", sudoku: [ { value: "", sid: "five-1"}, { value: "1", sid: "five-2"}, { value: "", sid: "five-3"}, { value: "", sid: "five-4"}, { value: "2", sid: "five-5"}, { value: "4", sid: "five-6"}, { value: "9", sid: "five-7"}, { value: "6", sid: "five-8"}, { value: "7", sid: "five-9"}] },
            { id: "six", sudoku: [ { value: "", sid: "six-1"}, { value: "2", sid: "six-2"}, { value: "8", sid: "six-3"}, { value: "", sid: "six-4"}, { value: "9", sid: "six-5"}, { value: "6", sid: "six-6"}, { value: "", sid: "six-7"}, { value: "5", sid: "six-8"}, { value: "3", sid: "six-9"}] },
            { id: "seven", sudoku: [ { value: "1", sid: "seven-1"}, { value: "", sid: "seven-2"}, { value: "7", sid: "seven-3"}, { value: "", sid: "seven-4"}, { value: "", sid: "seven-5"}, { value: "6", sid: "seven-6"}, { value: "9", sid: "seven-7"}, { value: "", sid: "seven-8"}, { value: "4", sid: "seven-9"}] },
            { id: "eight", sudoku: [ { value: "6", sid: "eight-1"}, { value: "5", sid: "eight-2"}, { value: "", sid: "eight-3"}, { value: "", sid: "eight-4"}, { value: "4", sid: "eight-5"}, { value: "", sid: "eight-6"}, { value: "2", sid: "eight-7"}, { value: 1, sid: "eight-8"}, { value: "", sid: "eight-9"}] },
            { id: "nine", sudoku: [ { value: "", sid: "nine-1"}, { value: "", sid: "nine-2"}, { value: "4", sid: "nine-3"}, { value: "5", sid: "nine-4"}, { value: "7", sid: "nine-5"}, { value: "9", sid: "nine-6"}, { value: "6", sid: "nine-7"}, { value: "", sid: "nine-8"}, { value: "", sid: "nine-9"}] },
          ],
          Sid: "third"
        },
        {
          blocks: [
            { id: "one", sudoku: [ { value: "", sid: "one-1"}, { value: 9, sid: "one-2"}, { value: "", sid: "one-3"}, { value: 7, sid: "one-4"}, { value: 3, sid: "one-5"}, { value: "", sid: "one-6"}, { value: 8, sid: "one-7"}, { value: 6, sid: "one-8"}, { value: "", sid: "one-9"}] },
            { id: "two", sudoku: [ { value: "", sid: "two-1"}, { value: "3", sid: "two-2"}, { value: "1", sid: "two-3"}, { value: "", sid: "two-4"}, { value: "8", sid: "two-5"}, { value: "6", sid: "two-6"}, { value: "7", sid: "two-7"}, { value: "", sid: "two-8"}, { value: "2", sid: "two-9"}] },
            { id: "three", sudoku: [ { value: "", sid: "three-1"}, { value: "6", sid: "three-2"}, { value: "7", sid: "three-3"}, { value: "9", sid: "three-4"}, { value: "1", sid: "three-5"}, { value: "", sid: "three-6"}, { value: "3", sid: "three-7"}, { value: "4", sid: "three-8"}, { value: "", sid: "three-9"}] },
            { id: "four", sudoku: [ { value: "", sid: "four-1"}, { value: "7", sid: "four-2"}, { value: "9", sid: "four-3"}, { value: "5", sid: "four-4"}, { value: "", sid: "four-5"}, { value: "3", sid: "four-6"}, { value: "", sid: "four-7"}, { value: "", sid: "four-8"}, { value: "", sid: "four-9"}] },
            { id: "five", sudoku: [ { value: "", sid: "five-1"}, { value: "1", sid: "five-2"}, { value: "", sid: "five-3"}, { value: "", sid: "five-4"}, { value: "2", sid: "five-5"}, { value: "4", sid: "five-6"}, { value: "9", sid: "five-7"}, { value: "6", sid: "five-8"}, { value: "7", sid: "five-9"}] },
            { id: "six", sudoku: [ { value: "", sid: "six-1"}, { value: "2", sid: "six-2"}, { value: "8", sid: "six-3"}, { value: "", sid: "six-4"}, { value: "9", sid: "six-5"}, { value: "6", sid: "six-6"}, { value: "", sid: "six-7"}, { value: "5", sid: "six-8"}, { value: "3", sid: "six-9"}] },
            { id: "seven", sudoku: [ { value: "1", sid: "seven-1"}, { value: "", sid: "seven-2"}, { value: "7", sid: "seven-3"}, { value: "", sid: "seven-4"}, { value: "", sid: "seven-5"}, { value: "6", sid: "seven-6"}, { value: "9", sid: "seven-7"}, { value: "", sid: "seven-8"}, { value: "4", sid: "seven-9"}] },
            { id: "eight", sudoku: [ { value: "6", sid: "eight-1"}, { value: "5", sid: "eight-2"}, { value: "", sid: "eight-3"}, { value: "", sid: "eight-4"}, { value: "4", sid: "eight-5"}, { value: "", sid: "eight-6"}, { value: "2", sid: "eight-7"}, { value: 1, sid: "eight-8"}, { value: "", sid: "eight-9"}] },
            { id: "nine", sudoku: [ { value: "", sid: "nine-1"}, { value: "", sid: "nine-2"}, { value: "4", sid: "nine-3"}, { value: "5", sid: "nine-4"}, { value: "7", sid: "nine-5"}, { value: "9", sid: "nine-6"}, { value: "6", sid: "nine-7"}, { value: "", sid: "nine-8"}, { value: "", sid: "nine-9"}] },
          ],
          Sid: "fourth"
        },
        {
          blocks: [
            { id: "one", sudoku: [ { value: "", sid: "one-1"}, { value: 9, sid: "one-2"}, { value: "", sid: "one-3"}, { value: 7, sid: "one-4"}, { value: 3, sid: "one-5"}, { value: "", sid: "one-6"}, { value: 8, sid: "one-7"}, { value: 6, sid: "one-8"}, { value: "", sid: "one-9"}] },
            { id: "two", sudoku: [ { value: "", sid: "two-1"}, { value: "3", sid: "two-2"}, { value: "1", sid: "two-3"}, { value: "", sid: "two-4"}, { value: "8", sid: "two-5"}, { value: "6", sid: "two-6"}, { value: "7", sid: "two-7"}, { value: "", sid: "two-8"}, { value: "2", sid: "two-9"}] },
            { id: "three", sudoku: [ { value: "", sid: "three-1"}, { value: "6", sid: "three-2"}, { value: "7", sid: "three-3"}, { value: "9", sid: "three-4"}, { value: "1", sid: "three-5"}, { value: "", sid: "three-6"}, { value: "3", sid: "three-7"}, { value: "4", sid: "three-8"}, { value: "", sid: "three-9"}] },
            { id: "four", sudoku: [ { value: "", sid: "four-1"}, { value: "7", sid: "four-2"}, { value: "9", sid: "four-3"}, { value: "5", sid: "four-4"}, { value: "", sid: "four-5"}, { value: "3", sid: "four-6"}, { value: "", sid: "four-7"}, { value: "", sid: "four-8"}, { value: "", sid: "four-9"}] },
            { id: "five", sudoku: [ { value: "", sid: "five-1"}, { value: "1", sid: "five-2"}, { value: "", sid: "five-3"}, { value: "", sid: "five-4"}, { value: "2", sid: "five-5"}, { value: "4", sid: "five-6"}, { value: "9", sid: "five-7"}, { value: "6", sid: "five-8"}, { value: "7", sid: "five-9"}] },
            { id: "six", sudoku: [ { value: "", sid: "six-1"}, { value: "2", sid: "six-2"}, { value: "8", sid: "six-3"}, { value: "", sid: "six-4"}, { value: "9", sid: "six-5"}, { value: "6", sid: "six-6"}, { value: "", sid: "six-7"}, { value: "5", sid: "six-8"}, { value: "3", sid: "six-9"}] },
            { id: "seven", sudoku: [ { value: "1", sid: "seven-1"}, { value: "", sid: "seven-2"}, { value: "7", sid: "seven-3"}, { value: "", sid: "seven-4"}, { value: "", sid: "seven-5"}, { value: "6", sid: "seven-6"}, { value: "9", sid: "seven-7"}, { value: "", sid: "seven-8"}, { value: "4", sid: "seven-9"}] },
            { id: "eight", sudoku: [ { value: "6", sid: "eight-1"}, { value: "5", sid: "eight-2"}, { value: "", sid: "eight-3"}, { value: "", sid: "eight-4"}, { value: "4", sid: "eight-5"}, { value: "", sid: "eight-6"}, { value: "2", sid: "eight-7"}, { value: 1, sid: "eight-8"}, { value: "", sid: "eight-9"}] },
            { id: "nine", sudoku: [ { value: "", sid: "nine-1"}, { value: "", sid: "nine-2"}, { value: "4", sid: "nine-3"}, { value: "5", sid: "nine-4"}, { value: "7", sid: "nine-5"}, { value: "9", sid: "nine-6"}, { value: "6", sid: "nine-7"}, { value: "", sid: "nine-8"}, { value: "", sid: "nine-9"}] },
          ],
          Sid: "fifth"
        },
        {
          blocks: [
            { id: "one", sudoku: [ { value: "", sid: "one-1"}, { value: 9, sid: "one-2"}, { value: "", sid: "one-3"}, { value: 7, sid: "one-4"}, { value: 3, sid: "one-5"}, { value: "", sid: "one-6"}, { value: 8, sid: "one-7"}, { value: 6, sid: "one-8"}, { value: "", sid: "one-9"}] },
            { id: "two", sudoku: [ { value: "", sid: "two-1"}, { value: "3", sid: "two-2"}, { value: "1", sid: "two-3"}, { value: "", sid: "two-4"}, { value: "8", sid: "two-5"}, { value: "6", sid: "two-6"}, { value: "7", sid: "two-7"}, { value: "", sid: "two-8"}, { value: "2", sid: "two-9"}] },
            { id: "three", sudoku: [ { value: "", sid: "three-1"}, { value: "6", sid: "three-2"}, { value: "7", sid: "three-3"}, { value: "9", sid: "three-4"}, { value: "1", sid: "three-5"}, { value: "", sid: "three-6"}, { value: "3", sid: "three-7"}, { value: "4", sid: "three-8"}, { value: "", sid: "three-9"}] },
            { id: "four", sudoku: [ { value: "", sid: "four-1"}, { value: "7", sid: "four-2"}, { value: "9", sid: "four-3"}, { value: "5", sid: "four-4"}, { value: "", sid: "four-5"}, { value: "3", sid: "four-6"}, { value: "", sid: "four-7"}, { value: "", sid: "four-8"}, { value: "", sid: "four-9"}] },
            { id: "five", sudoku: [ { value: "", sid: "five-1"}, { value: "1", sid: "five-2"}, { value: "", sid: "five-3"}, { value: "", sid: "five-4"}, { value: "2", sid: "five-5"}, { value: "4", sid: "five-6"}, { value: "9", sid: "five-7"}, { value: "6", sid: "five-8"}, { value: "7", sid: "five-9"}] },
            { id: "six", sudoku: [ { value: "", sid: "six-1"}, { value: "2", sid: "six-2"}, { value: "8", sid: "six-3"}, { value: "", sid: "six-4"}, { value: "9", sid: "six-5"}, { value: "6", sid: "six-6"}, { value: "", sid: "six-7"}, { value: "5", sid: "six-8"}, { value: "3", sid: "six-9"}] },
            { id: "seven", sudoku: [ { value: "1", sid: "seven-1"}, { value: "", sid: "seven-2"}, { value: "7", sid: "seven-3"}, { value: "", sid: "seven-4"}, { value: "", sid: "seven-5"}, { value: "6", sid: "seven-6"}, { value: "9", sid: "seven-7"}, { value: "", sid: "seven-8"}, { value: "4", sid: "seven-9"}] },
            { id: "eight", sudoku: [ { value: "6", sid: "eight-1"}, { value: "5", sid: "eight-2"}, { value: "", sid: "eight-3"}, { value: "", sid: "eight-4"}, { value: "4", sid: "eight-5"}, { value: "", sid: "eight-6"}, { value: "2", sid: "eight-7"}, { value: 1, sid: "eight-8"}, { value: "", sid: "eight-9"}] },
            { id: "nine", sudoku: [ { value: "", sid: "nine-1"}, { value: "", sid: "nine-2"}, { value: "4", sid: "nine-3"}, { value: "5", sid: "nine-4"}, { value: "7", sid: "nine-5"}, { value: "9", sid: "nine-6"}, { value: "6", sid: "nine-7"}, { value: "", sid: "nine-8"}, { value: "", sid: "nine-9"}] },
          ],
          Sid: "sixth"
        },
        {
          blocks: [
            { id: "one", sudoku: [ { value: "", sid: "one-1"}, { value: 9, sid: "one-2"}, { value: "", sid: "one-3"}, { value: 7, sid: "one-4"}, { value: 3, sid: "one-5"}, { value: "", sid: "one-6"}, { value: 8, sid: "one-7"}, { value: 6, sid: "one-8"}, { value: "", sid: "one-9"}] },
            { id: "two", sudoku: [ { value: "", sid: "two-1"}, { value: "3", sid: "two-2"}, { value: "1", sid: "two-3"}, { value: "", sid: "two-4"}, { value: "8", sid: "two-5"}, { value: "6", sid: "two-6"}, { value: "7", sid: "two-7"}, { value: "", sid: "two-8"}, { value: "2", sid: "two-9"}] },
            { id: "three", sudoku: [ { value: "", sid: "three-1"}, { value: "6", sid: "three-2"}, { value: "7", sid: "three-3"}, { value: "9", sid: "three-4"}, { value: "1", sid: "three-5"}, { value: "", sid: "three-6"}, { value: "3", sid: "three-7"}, { value: "4", sid: "three-8"}, { value: "", sid: "three-9"}] },
            { id: "four", sudoku: [ { value: "", sid: "four-1"}, { value: "7", sid: "four-2"}, { value: "9", sid: "four-3"}, { value: "5", sid: "four-4"}, { value: "", sid: "four-5"}, { value: "3", sid: "four-6"}, { value: "", sid: "four-7"}, { value: "", sid: "four-8"}, { value: "", sid: "four-9"}] },
            { id: "five", sudoku: [ { value: "", sid: "five-1"}, { value: "1", sid: "five-2"}, { value: "", sid: "five-3"}, { value: "", sid: "five-4"}, { value: "2", sid: "five-5"}, { value: "4", sid: "five-6"}, { value: "9", sid: "five-7"}, { value: "6", sid: "five-8"}, { value: "7", sid: "five-9"}] },
            { id: "six", sudoku: [ { value: "", sid: "six-1"}, { value: "2", sid: "six-2"}, { value: "8", sid: "six-3"}, { value: "", sid: "six-4"}, { value: "9", sid: "six-5"}, { value: "6", sid: "six-6"}, { value: "", sid: "six-7"}, { value: "5", sid: "six-8"}, { value: "3", sid: "six-9"}] },
            { id: "seven", sudoku: [ { value: "1", sid: "seven-1"}, { value: "", sid: "seven-2"}, { value: "7", sid: "seven-3"}, { value: "", sid: "seven-4"}, { value: "", sid: "seven-5"}, { value: "6", sid: "seven-6"}, { value: "9", sid: "seven-7"}, { value: "", sid: "seven-8"}, { value: "4", sid: "seven-9"}] },
            { id: "eight", sudoku: [ { value: "6", sid: "eight-1"}, { value: "5", sid: "eight-2"}, { value: "", sid: "eight-3"}, { value: "", sid: "eight-4"}, { value: "4", sid: "eight-5"}, { value: "", sid: "eight-6"}, { value: "2", sid: "eight-7"}, { value: 1, sid: "eight-8"}, { value: "", sid: "eight-9"}] },
            { id: "nine", sudoku: [ { value: "", sid: "nine-1"}, { value: "", sid: "nine-2"}, { value: "4", sid: "nine-3"}, { value: "5", sid: "nine-4"}, { value: "7", sid: "nine-5"}, { value: "9", sid: "nine-6"}, { value: "6", sid: "nine-7"}, { value: "", sid: "nine-8"}, { value: "", sid: "nine-9"}] },
          ],
          Sid: "seventh"
        },
        {
          blocks: [
            { id: "one", sudoku: [ { value: "", sid: "one-1"}, { value: 9, sid: "one-2"}, { value: "", sid: "one-3"}, { value: 7, sid: "one-4"}, { value: 3, sid: "one-5"}, { value: "", sid: "one-6"}, { value: 8, sid: "one-7"}, { value: 6, sid: "one-8"}, { value: "", sid: "one-9"}] },
            { id: "two", sudoku: [ { value: "", sid: "two-1"}, { value: "3", sid: "two-2"}, { value: "1", sid: "two-3"}, { value: "", sid: "two-4"}, { value: "8", sid: "two-5"}, { value: "6", sid: "two-6"}, { value: "7", sid: "two-7"}, { value: "", sid: "two-8"}, { value: "2", sid: "two-9"}] },
            { id: "three", sudoku: [ { value: "", sid: "three-1"}, { value: "6", sid: "three-2"}, { value: "7", sid: "three-3"}, { value: "9", sid: "three-4"}, { value: "1", sid: "three-5"}, { value: "", sid: "three-6"}, { value: "3", sid: "three-7"}, { value: "4", sid: "three-8"}, { value: "", sid: "three-9"}] },
            { id: "four", sudoku: [ { value: "", sid: "four-1"}, { value: "7", sid: "four-2"}, { value: "9", sid: "four-3"}, { value: "5", sid: "four-4"}, { value: "", sid: "four-5"}, { value: "3", sid: "four-6"}, { value: "", sid: "four-7"}, { value: "", sid: "four-8"}, { value: "", sid: "four-9"}] },
            { id: "five", sudoku: [ { value: "", sid: "five-1"}, { value: "1", sid: "five-2"}, { value: "", sid: "five-3"}, { value: "", sid: "five-4"}, { value: "2", sid: "five-5"}, { value: "4", sid: "five-6"}, { value: "9", sid: "five-7"}, { value: "6", sid: "five-8"}, { value: "7", sid: "five-9"}] },
            { id: "six", sudoku: [ { value: "", sid: "six-1"}, { value: "2", sid: "six-2"}, { value: "8", sid: "six-3"}, { value: "", sid: "six-4"}, { value: "9", sid: "six-5"}, { value: "6", sid: "six-6"}, { value: "", sid: "six-7"}, { value: "5", sid: "six-8"}, { value: "3", sid: "six-9"}] },
            { id: "seven", sudoku: [ { value: "1", sid: "seven-1"}, { value: "", sid: "seven-2"}, { value: "7", sid: "seven-3"}, { value: "", sid: "seven-4"}, { value: "", sid: "seven-5"}, { value: "6", sid: "seven-6"}, { value: "9", sid: "seven-7"}, { value: "", sid: "seven-8"}, { value: "4", sid: "seven-9"}] },
            { id: "eight", sudoku: [ { value: "6", sid: "eight-1"}, { value: "5", sid: "eight-2"}, { value: "", sid: "eight-3"}, { value: "", sid: "eight-4"}, { value: "4", sid: "eight-5"}, { value: "", sid: "eight-6"}, { value: "2", sid: "eight-7"}, { value: 1, sid: "eight-8"}, { value: "", sid: "eight-9"}] },
            { id: "nine", sudoku: [ { value: "", sid: "nine-1"}, { value: "", sid: "nine-2"}, { value: "4", sid: "nine-3"}, { value: "5", sid: "nine-4"}, { value: "7", sid: "nine-5"}, { value: "9", sid: "nine-6"}, { value: "6", sid: "nine-7"}, { value: "", sid: "nine-8"}, { value: "", sid: "nine-9"}] },
          ],
          Sid: "eighth"
        },
        {
          blocks: [
            { id: "one", sudoku: [ { value: "", sid: "one-1"}, { value: 9, sid: "one-2"}, { value: "", sid: "one-3"}, { value: 7, sid: "one-4"}, { value: 3, sid: "one-5"}, { value: "", sid: "one-6"}, { value: 8, sid: "one-7"}, { value: 6, sid: "one-8"}, { value: "", sid: "one-9"}] },
            { id: "two", sudoku: [ { value: "", sid: "two-1"}, { value: "3", sid: "two-2"}, { value: "1", sid: "two-3"}, { value: "", sid: "two-4"}, { value: "8", sid: "two-5"}, { value: "6", sid: "two-6"}, { value: "7", sid: "two-7"}, { value: "", sid: "two-8"}, { value: "2", sid: "two-9"}] },
            { id: "three", sudoku: [ { value: "", sid: "three-1"}, { value: "6", sid: "three-2"}, { value: "7", sid: "three-3"}, { value: "9", sid: "three-4"}, { value: "1", sid: "three-5"}, { value: "", sid: "three-6"}, { value: "3", sid: "three-7"}, { value: "4", sid: "three-8"}, { value: "", sid: "three-9"}] },
            { id: "four", sudoku: [ { value: "", sid: "four-1"}, { value: "7", sid: "four-2"}, { value: "9", sid: "four-3"}, { value: "5", sid: "four-4"}, { value: "", sid: "four-5"}, { value: "3", sid: "four-6"}, { value: "", sid: "four-7"}, { value: "", sid: "four-8"}, { value: "", sid: "four-9"}] },
            { id: "five", sudoku: [ { value: "", sid: "five-1"}, { value: "1", sid: "five-2"}, { value: "", sid: "five-3"}, { value: "", sid: "five-4"}, { value: "2", sid: "five-5"}, { value: "4", sid: "five-6"}, { value: "9", sid: "five-7"}, { value: "6", sid: "five-8"}, { value: "7", sid: "five-9"}] },
            { id: "six", sudoku: [ { value: "", sid: "six-1"}, { value: "2", sid: "six-2"}, { value: "8", sid: "six-3"}, { value: "", sid: "six-4"}, { value: "9", sid: "six-5"}, { value: "6", sid: "six-6"}, { value: "", sid: "six-7"}, { value: "5", sid: "six-8"}, { value: "3", sid: "six-9"}] },
            { id: "seven", sudoku: [ { value: "1", sid: "seven-1"}, { value: "", sid: "seven-2"}, { value: "7", sid: "seven-3"}, { value: "", sid: "seven-4"}, { value: "", sid: "seven-5"}, { value: "6", sid: "seven-6"}, { value: "9", sid: "seven-7"}, { value: "", sid: "seven-8"}, { value: "4", sid: "seven-9"}] },
            { id: "eight", sudoku: [ { value: "6", sid: "eight-1"}, { value: "5", sid: "eight-2"}, { value: "", sid: "eight-3"}, { value: "", sid: "eight-4"}, { value: "4", sid: "eight-5"}, { value: "", sid: "eight-6"}, { value: "2", sid: "eight-7"}, { value: 1, sid: "eight-8"}, { value: "", sid: "eight-9"}] },
            { id: "nine", sudoku: [ { value: "", sid: "nine-1"}, { value: "", sid: "nine-2"}, { value: "4", sid: "nine-3"}, { value: "5", sid: "nine-4"}, { value: "7", sid: "nine-5"}, { value: "9", sid: "nine-6"}, { value: "6", sid: "nine-7"}, { value: "", sid: "nine-8"}, { value: "", sid: "nine-9"}] },
          ],
          Sid: "nineth"
        }
      ]
    }
  },
  props: {
    bgc:{
      type: String,
      required: true,
      default:''     //默认值
    },
  },
  methods: {
    changeColor() {
      let currentIndex = backgroundStyles.indexOf(this.bgc);
      currentIndex += 1;
      if ( currentIndex >= backgroundStyles.length ) {
        currentIndex %= backgroundStyles.length;
      }
      const selectedColor = backgroundStyles[currentIndex];
      console.log(currentIndex);
      this.$emit('changeBackgroundColor', selectedColor);
    }
  },
  created () {},
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
    &_change, &_return {
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
#one-1 {
  border-top-left-radius: 1.7vmin;
}
#three-3 {
  border-top-right-radius: 1.7vmin;
}
#seven-7 {
  border-bottom-left-radius: 1.7vmin;
}
#nine-9 {
  border-bottom-right-radius: 1.7vmin;
}
</style>
