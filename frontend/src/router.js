import { createRouter, createWebHistory } from "vue-router"; // Use correct imports for Vue 3
import ListBooks from "./components/ListBooks.vue";
import CreateBook from "./components/CreateBook.vue";
import UpdateBook from "./components/UpdateBook.vue";
import ResetBooks from "./components/ResetBooks.vue";
import EditBook from "./components/EditBooks.vue";

const routes = [
  { path: "/", component: ListBooks },
  { path: "/create", component: CreateBook },
  { path: "/update/:id", component: UpdateBook },
  { path: "/reset", component: ResetBooks },
  {
    path: '/edit/:id',
    component: EditBook,
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
