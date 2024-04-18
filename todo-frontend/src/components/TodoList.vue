<template>
  <div>
    <input v-model="newTodo" @keyup.enter="addTodo" placeholder="Add new tod" />
    <button @click="addTodo">Add Todo</button>

    <ul>
      <li v-for="todo in todos" :key="todo.id">
        <label class="switch">
          <input
            type="checkbox"
            v-model="todo.completed"
            @click="toggleCompleted(todo)"
          />
          <span class="slider round"></span>
        </label>
        {{ todo.title }}
        <button @click="deleteTodo(todo.id)">Delete</button>
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
      console.log("Toggle completed for:", todo);
      axios
        .put(`${process.env.VUE_APP_API_DOMAIN}/todos/${todo.id}`, {
          ...todo,
          completed: !todo.completed,
        })
        .then((response) => {
          const index = this.todos.findIndex((t) => t.id === todo.id);
          this.todos[index] = response.data;
          console.log("Updated todo:", response.data);
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
/* 样式化开关的外观 */
.switch {
  position: relative;
  display: inline-block;
  width: 60px;
  height: 34px;
}

/* 隐藏默认的 HTML checkbox */
.switch input {
  opacity: 0.01; /* 将这个值从 0 改为 0.01 可以暂时查看输入框的位置和大小 */
  position: absolute;
  width: 100%;
  height: 100%;
  margin: 0;
  cursor: pointer;
}

/* Slider */
.slider {
  position: absolute;
  cursor: pointer;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: #ccc;
  transition: 0.4s;
  border-radius: 34px;
}

.slider:before {
  position: absolute;
  content: "";
  height: 26px;
  width: 26px;
  left: 4px;
  bottom: 4px;
  background-color: white;
  transition: 0.4s;
  border-radius: 50%;
}

input:checked + .slider {
  background-color: #2196f3;
}

input:checked + .slider:before {
  transform: translateX(26px);
}

/* 添加额外的样式来增强开关的可点击区域和触感 */
.slider:active:before {
  width: 28px;
}
</style>
