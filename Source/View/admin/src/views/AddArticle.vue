<template>
    <div style="text-align: left;height: 100%">
        <el-row  style="height: 100%" :gutter="20">
            <el-col :span="12" style="height: 100%">
                <textarea style="height: 100%;resize: none;border-radius: 5px;border: 1px solid #cccccc;width: 100%"></textarea>
            </el-col>
            <el-col :span="12" style="height: 100%">
                <div style="border: 1px solid #cccccc;border-radius: 5px;height: 100%;overflow-y: scroll; padding: 0 20px;">
                    <vue-markdown :source="body"></vue-markdown>
                </div>
            </el-col>
        </el-row>
    </div>
</template>

<script>
    import VueMarkdown from 'vue-markdown'
    export default {
        name: "AddArticle",
        components:{
            VueMarkdown
        },
        data(){
            return{
                body: "## 123\n" +
                    "**123**\n" +
                    "```js\n" +
                    "function(){\n" +
                    "   alert(223);\n" +
                    "}\n" +
                    "```" +
                    "## 123\n" +
                    "**123**\n" +
                    "```js\n" +
                    "function(){\n" +
                    "   alert(223);\n" +
                    "}\n" +
                    "```" +
                    "## 123\n" +
                    "**123**\n" +
                    "```js\n" +
                    "function(){\n" +
                    "   alert(223);\n" +
                    "}\n" +
                    "```" +
                    "## 123\n" +
                    "**123**\n" +
                    "```js\n" +
                    "function(){\n" +
                    "   alert(223);\n" +
                    "}\n" +
                    "```" +
                    "## 123\n" +
                    "**123**\n" +
                    "```js\n" +
                    "function(){\n" +
                    "   alert(223);\n" +
                    "}\n" +
                    "```" +
                    "## 123\n" +
                    "**123**\n" +
                    "```js\n" +
                    "function(){\n" +
                    "   alert(223);\n" +
                    "}\n" +
                    "```" +
                    "## 123\n" +
                    "**123**\n" +
                    "```js\n" +
                    "function(){\n" +
                    "   alert(223);\n" +
                    "}\n" +
                    "```" +
                    "## 123\n" +
                    "**123**\n" +
                    "```js\n" +
                    "function(){\n" +
                    "   alert(223);\n" +
                    "}\n" +
                    "```" +
                    "## 123\n" +
                    "**123**\n" +
                    "```js\n" +
                    "function(){\n" +
                    "   alert(223);\n" +
                    "}\n" +
                    "```" +
                    "## 123\n" +
                    "**123**\n" +
                    "```js\n" +
                    "function(){\n" +
                    "   alert(223);\n" +
                    "}\n" +
                    "```" +
                    "## 123\n" +
                    "**123**\n" +
                    "```js\n" +
                    "function(){\n" +
                    "   alert(223);\n" +
                    "}\n" +
                    "```" +
                    "## 123\n" +
                    "**123**\n" +
                    "```js\n" +
                    "function(){\n" +
                    "   alert(223);\n" +
                    "}\n" +
                    "```" +
                    "## 123\n" +
                    "**123**\n" +
                    "```js\n" +
                    "function(){\n" +
                    "   alert(223);\n" +
                    "}\n" +
                    "```" +
                    "## 123\n" +
                    "**123**\n" +
                    "```js\n" +
                    "function(){\n" +
                    "   alert(223);\n" +
                    "}\n" +
                    "```"
            }
        }
    }
</script>

<style scoped>
</style>
