<template>
    <div class="header-container">
        <i :class="[isCollapse?'el-icon-upload2':'el-icon-download', 'menu-show'].join(' ')" v-on:click="changeMenuCollapseStatus"></i>
        <div>
            <i class="el-icon-full-screen"></i>
            <el-dropdown class="dropdown-box">
              <span class="el-dropdown-link">
                <i class="el-icon-user"></i>
              </span>
                <el-dropdown-menu class="dropdown" :split-button="true" :show-timeout="10" slot="dropdown">
                    <el-dropdown-item>黄金糕</el-dropdown-item>
                    <el-dropdown-item>狮子头</el-dropdown-item>
                    <el-dropdown-item>螺蛳粉</el-dropdown-item>
                </el-dropdown-menu>
            </el-dropdown>
        </div>
    </div>
</template>

<script>
    export default {
        name: "Header",
        methods: {
            changeMenuCollapseStatus() {
                this.$store.commit("changeMenuCollapseStatus");
            }
        },
        computed:{
            isCollapse(){
                return this.$store.state.isMenuCollapse;
            }
        }

    }
</script>

<style scoped>
    .header-container {
        background-color: #1e88e5;
        width: 100%;
        height: 100%;
        color: white;
        display: flex;
        align-items: center;
        justify-content: space-between;
    }

    .header-container .menu-show {
        margin-left: 18px;
        transform: rotate(90deg);
    }

    .header-container div {
        margin-right: 16px;
    }

    i {
        font-size: 26px;
        cursor: pointer;
        color: white;
    }

    .dropdown-box {
        margin-left: 10px;
    }

    .dropdown {
        width: 200px;
    }
</style>
