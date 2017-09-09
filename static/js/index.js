import Vue from 'vue';
import Hello from "./hello.vue";

new Vue({
    el: "#test",
    template: '<Hello/>',
    components: { Hello }
})
