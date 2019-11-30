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
                    icon:"",
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
        background-color: white;
    }
</style>
