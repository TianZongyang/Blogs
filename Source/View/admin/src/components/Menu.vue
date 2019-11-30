<template>
    <div class="menu-container">

    </div>
</template>

<script>
    export default {
        name: "Menu",
        data(){
            return{
                menuList:[{
                    menuName:"仪表盘",
                    path:"",
                    enable:true
                }]
            }
        },
        methods:{

        }
    }
</script>

<style scoped>
    .menu-container{
        height: 100vh;
        background-color: #2c3e50;
    }
</style>
