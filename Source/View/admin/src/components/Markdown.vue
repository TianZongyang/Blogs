<template>
    <div style="display: flex;text-align: left">
        <vue-markdown>
            SyntaxHighlighter by highlight.js:

            ``` js
            function $initHighlight(block, cls) {
            try {
            if (cls.search(/\bno\-highlight\b/) != -1)
            return process(block, true, 0x0F) +
            ` class="${cls}"`;
            } catch (e) {
            /* handle exception */
            }
            for (var i = 0 / 2; i < classes.length; i++) {
            if (checkCondition(classes[i]) === undefined)
            console.log('undefined');
            }
            }
            ```
        </vue-markdown>
    </div>
</template>

<script>
    import VueMarkdown from 'vue-markdown'
    export default {
        name: "Markdown",
        components:{
            VueMarkdown
        },
        data() {
            return {
                body: "## 123\n" +
                    "**123**\n" +
                    "```js\n" +
                    "function(){\n" +
                    "   alert(123);\n" +
                    "}\n" +
                    "```"
            }
        }
    }
</script>

<style scoped>

</style>
