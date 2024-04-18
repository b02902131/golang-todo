<template>
  <img alt="Vue logo" src="./assets/logo.png" />
  <TodoList />
</template>

<script>
import TodoList from "./components/TodoList.vue";

export default {
  name: "App",
  components: {
    TodoList,
  },
};
</script>

<style>
/* 你可以在這裡添加一些全局樣式 */
</style>
