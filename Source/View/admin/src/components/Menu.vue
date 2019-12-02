<template>
    <div class="menu-container">
        <el-menu :default-active="menuList[0].path" class="menu-vertical" :collapse="this.$store.state.isMenuCollapse"
                 :unique-opened="true" :collapse-transition="false" :router="true" active-text-color="#1e88e5">
            <template v-for="(item1,index1) in menuList">
                <el-submenu :index="item1.path" v-bind:key="index1" v-if="item1.child !== null && item1.child.length > 0" class="menu-item">
                    <template slot="title">
                        <i :class="item1.icon"></i>
                        <span slot="title">{{item1.menuName}}</span>
                    </template>
                    <el-menu-item v-for="(item2,index2) in item1.child" :index="item2.path" v-bind:key="index2" class="menu-item">
                        {{item2.menuName}}
                    </el-menu-item>
                </el-submenu>
                <el-menu-item v-else :index="item1.path" v-bind:key="index1" class="menu-item">
                    <i :class="item1.icon"></i>
                    <span slot="title">{{item1.menuName}}</span>
                </el-menu-item>
            </template>
        </el-menu>
    </div>
</template>

<script>
    export default {
        name: "Menu",
        data() {
            return {
                menuList: [{
                    menuName: "仪表盘",
                    path: "/dashboard",
                    icon: "el-icon-data-line",
                    enable: true,
                    child: null
                }, {
                    menuName: "文章管理",
                    path: "/article",
                    icon: "el-icon-tickets",
                    enable: true,
                    child: [
                        {
                            menuName: "查看文章列表",
                            path: "",
                            icon: "el-icon-reading",
                            enable: true
                        }, {
                            menuName: "添加文章",
                            path: "/addArticle",
                            icon: "el-icon-edit-outline",
                            enable: true
                        }
                    ]
                }, {
                    menuName: "用户管理",
                    path: "",
                    icon: "el-icon-user",
                    enable: true,
                    child: []
                }]
            }
        },
        created() {
        },
        methods: {}
    }
</script>

<style scoped>
    .menu-container {
        height: 100vh;
        background-color: white;
        border-right: 1px solid #ededed;
    }

    .menu-vertical:not(.el-menu--collapse) {
        width: 200px;
    }

    .menu-vertical {
        border-right: 0;
        display: flex;
        flex-flow: column;
    }

    .menu-item {
        text-align: left;
    }
</style>
