<template>
  <div class="select">
    <div class="select_header">
      <div class="select_header_return" @click="$router.go(-1)" title="返回！">
        <img src="../assets/images/return.png" alt="" title="返回！">
      </div>
      <div class="select_header_change" @click="changeColor()" title="更换主题！">
        <img src="../assets/images/change.png" alt="" title="更换主题！">
      </div>
    </div>
    <div class="select_wrap">
      <ul class="select_wrap_difficulty">
        <li @click="game(1)">Easy</li>
        <li @click="game(2)">Normal</li>
        <li @click="game(3)">Hard</li>
        <li @click="game(4)">Hell</li>
      </ul>
    </div>
    <Footer></Footer>
  </div>
</template>

<script>
import Footer from '@/views/footer.vue'
export default {
  name: 'Select',
  data () {
    return {}
  },
  components: {
    Footer
  },
  methods: {
    changeColor() {

    },
    game() {
      this.$router.push('/game')
    }
  },
}
</script>
<style scoped lang='less'>
.select {
  width: 100vw;
  height: 100vh;
  
  &_header {
    display: flex;
    justify-content: right;
    width: 100vw;
    height: 130px;
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
  &_wrap {
    display: flex;
    justify-content: center;
    align-items: center;
    margin: 0 auto;
    width: 50vh;
    height: calc(100vh - 240px);
    &_difficulty {
      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: center;
      margin: 0 auto;
      width: 427px;
      height: 650px;
      gap: 40px;
      li {
        display: flex;
        width: 430px;
        height: 100px;
        padding: 20.61px 0px 21.39px 0px;
        justify-content: center;
        align-items: center;
        border-radius: 48px;
        border-bottom: 9.594px solid rgba(0, 0, 0, 0.20);
        background: #FFF;
        color: rgba(0, 0, 0, 0.70);
        text-align: center;
        font-family: Microsoft YaHei;
        font-size: 32px;
        font-weight: 700;
        transition: all 0.3s ease 0s;
        &:hover {
          cursor: pointer;
          transform: scale(1.05);
          font-size: 38px;
          transition: all 0.3s linear;
        }
      }
    }
  }
}
</style>
