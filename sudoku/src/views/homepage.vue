<template>
  <div class="home">
    <div class="home_header">
      <div class="home_header_change">
        <div class="home_header_change_bg"></div>
      </div>
    </div>
    <div class="home_body">
      <div class="home_body_title">
        <h1>数独</h1>
        <h2>SUDOKU</h2>
      </div>
      <div class="home_body_btn">
        START
      </div>
    </div>
    <div class="home_footer">
      Design by Xiaoyangii
      <a href="https://github.com/xiaoyangii/Sudoku">Github</a>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Homepage',
  components: {},
  data () {
    return {}
  },
  computed: {},
  watch: {},
  methods: {},
  created () {},
}
</script>
<style scoped lang='less'>
.home {
  width: 100%;
  height: 1080px;
  background: linear-gradient(180deg, #CA6ECE 0%, #232361 100%);
  &_header {
    display: flex;
    justify-content: right;
    width: 100%;
    height: 130px;
    &_change {
      display: flex;
      margin-top: 32px;
      margin-right: 32px;
      width: 66px;
      height: 66px;
      line-height: 66px;
      border-radius: 10px;
      border-bottom: 5px solid rgba(0, 0, 0, 0.20);
      background-color: rgba(255, 255, 255, 0.60);
      &_bg {
        margin: 11px auto;
        width: 40px;
        height: 40px;
        background: url('@/assets/images/change.jpg') no-repeat center center;
      }
      
    }
  }
  &_body {
    margin: 0 auto;
    width: 100%;
    height: calc(100% - 240px);
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
        top: -220px;
        left: 0;
        color: rgba(255, 255, 255, 0.80);
        text-align: center;
        font-family: Microsoft YaHei;
        font-size: 83.4px;
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
      }
    }
  }
  &_footer {
    display: flex;
    flex-direction: column;
    width: 100%;
    height: 110px;
    justify-content: center;
    color: #6A5E5E;
    text-align: center;
    font-family: Microsoft YaHei UI;
    font-size: 24px;
    font-style: normal;
    font-weight: 400;
    line-height: normal;
  }
}

</style>
