<template>
  <div>
    <input v-model="newTodo" @keyup.enter="addTodo" placeholder="Add new tod" />
    <button @click="addTodo">Add Todo</button>

    <ul>
      <li v-for="todo in todos" :key="todo.id">
        {{ todo.title }}
        <button @click="deleteTodo(todo.id)">Delete</button>
        <button @click="toggleCompleted(todo)">Toggle Completed</button>
      </li>
    </ul>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      todos: [],
      newTodo: "",
    };
  },
  created() {
    this.fetchTodos();
  },
  methods: {
    fetchTodos() {
      axios
        .get(`${process.env.VUE_APP_API_DOMAIN}/todos`)
        .then((response) => {
          this.todos = response.data;
        })
        .catch((error) => {
          console.error(error);
        });
    },
    addTodo() {
      if (this.newTodo.trim()) {
        axios
          .post(`${process.env.VUE_APP_API_DOMAIN}/todos`, {
            title: this.newTodo,
            completed: false,
          })
          .then((response) => {
            this.todos.push(response.data);
            this.newTodo = "";
          })
          .catch((error) => {
            console.error(error);
          });
      }
    },
    deleteTodo(id) {
      axios
        .delete(`${process.env.VUE_APP_API_DOMAIN}/todos/${id}`)
        .then(() => {
          this.todos = this.todos.filter((todo) => todo.id !== id);
        })
        .catch((error) => {
          console.error(error);
        });
    },
    toggleCompleted(todo) {
      axios
        .put(`${process.env.VUE_APP_API_DOMAIN}/todos/${todo.id}`, {
          ...todo,
          completed: !todo.completed,
        })
        .then((response) => {
          const index = this.todos.findIndex((t) => t.id === todo.id);
          this.todos[index] = response.data;
        })
        .catch((error) => {
          console.error(error);
        });
    },
  },
};
</script>

<style>
/* 添加一些基本的样式 */
</style>
