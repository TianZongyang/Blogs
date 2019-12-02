<template>
    <div id="app">
        <el-container>
            <el-aside style="width: auto">
                <Menu></Menu>
            </el-aside>
            <el-container>
                <el-header>
                    <Header></Header>
                </el-header>
                <el-main style="background-color: #f8f8f8">
                    <router-view></router-view>
                </el-main>
            </el-container>
        </el-container>
    </div>
</template>

<script>
    import Menu from '@/components/Menu.vue'
    import Header from "@/components/Header";

    export default {
        name: 'app',
        components: {
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

    body {
        padding: 0;
        margin: 0;
    }

    .el-header {
        padding: 0 !important;
    }
</style>
