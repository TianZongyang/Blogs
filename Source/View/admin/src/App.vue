<template>
  <div id="app">
    <el-container>
      <el-header><Header></Header></el-header>
      <el-container>
        <el-aside width="200px"><Menu></Menu></el-aside>
        <el-main>Main</el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script>
import Menu from '@/components/Menu.vue'
import Header from "@/components/Header";
export default {
  name: 'app',
  components:{
    Menu,
    Header
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
body{
  padding: 0;
  margin: 0;
}
</style>
