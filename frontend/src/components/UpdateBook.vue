<template>
    <div>
      <h1>Update Book</h1>
      <form @submit.prevent="updateBook">
        <input v-model="book.title" placeholder="Title" required />
        <input v-model="book.author" placeholder="Author" required />
        <input v-model="book.genre" placeholder="Genre" />
        <textarea v-model="book.description" placeholder="Description"></textarea>
        <button type="submit">Update Book</button>
      </form>
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
        },
      };
    },
    created() {
      this.fetchBook();
    },
    methods: {
      fetchBook() {
        const id = this.$route.params.id;
        fetch(`http://localhost:8080/books/${id}`)
          .then(response => response.json())
          .then(data => {
            this.book = data;
          });
      },
      updateBook() {
        const id = this.$route.params.id;
        fetch(`http://localhost:8080/books/${id}`, {
          method: 'PUT',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(this.book),
        }).then(() => {
          this.$router.push('/');
        });
      },
    },
  };
  </script>
  