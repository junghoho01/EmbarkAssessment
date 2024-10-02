<template>
  <div class="create-book">
    <h1>Create Book</h1>
    <form @submit.prevent="createBook">
      <input v-model="book.title" placeholder="Title" required />
      <input v-model="book.author" placeholder="Author" required />
      <input v-model="book.genre" placeholder="Genre" />
      <textarea v-model="book.description" placeholder="Description"></textarea>
      <input type="date" v-model="book.published" placeholder="Published Date" required />
      <input v-model="book.publisher" placeholder="Publisher" />
      <button type="submit">Create Book</button>
    </form>
    <router-link to="/">Back to Book List</router-link>
  </div>
</template>
  
<script>
export default {
  data() {
    return {
      book: {
        title: '',
        author: '',
        genre: '',
        description: '',
        published: '', // New field
        publisher: '', // New field
      },
    };
  },
  methods: {
    createBook() {
      fetch('http://localhost:8080/books', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(this.book),
      })
        .then(() => {
          this.$router.push('/'); // Redirect to book list after adding
        })
        .catch((error) => {
          console.error('Error adding book:', error);
        });
    },
  },
};
</script>

<style scoped>
/* Add some simple styling */
input, textarea {
  display: block;
  margin-bottom: 10px;
  width: 100%;
}

button {
  padding: 10px;
}

h1 {
  margin-bottom: 20px;
}
</style>