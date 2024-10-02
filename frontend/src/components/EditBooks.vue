<template>
    <div>
      <h1>Edit Book</h1>
      <form @submit.prevent="updateBook">
        <input v-model="book.title" placeholder="Title" required />
        <input v-model="book.author" placeholder="Author" required />
        <input v-model="book.genre" placeholder="Genre" />
        <textarea v-model="book.description" placeholder="Description"></textarea>
        <input v-model="book.published" placeholder="Published" />
        <input v-model="book.publisher" placeholder="Publisher" />
        <button type="submit">Update Book</button>
      </form>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        book: {
          id: null, // Initialize ID as null
          title: '',
          author: '',
          genre: '',
          description: '',
          published: '',
          publisher: '',
        },
      };
    },
    mounted() {
      this.fetchBook(); // Call fetchBook when the component is mounted
    },
    methods: {
      fetchBook() {
        const bookId = this.$route.params.id; // Get the book ID from the route
        if (bookId) {
          fetch(`http://localhost:8080/books/${bookId}`)
            .then(response => response.json())
            .then(data => {
              this.book = data; // Populate the book details
            })
            .catch(error => {
              console.error('Error fetching book:', error);
            });
        } else {
          console.error('No book ID provided.');
        }
      },
      updateBook() {
        // Ensure book ID is set before submitting
        if (!this.book.id) {
          console.error('Book ID is missing.');
          return; // Stop if ID is null
        }
  
        fetch(`http://localhost:8080/books/${this.book.id}`, {
          method: 'PUT',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(this.book),
        })
          .then(response => {
            if (response.ok) {
              this.$router.push('/'); // Redirect back to the book list after updating
            } else {
              console.error('Error updating book:', response.statusText);
            }
          })
          .catch(error => {
            console.error('Error updating book:', error);
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
  