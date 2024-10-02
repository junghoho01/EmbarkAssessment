<template>
  <div>
    <h1>Book List</h1>
    <table>
      <thead>
        <tr>
          <th>Title</th>
          <th>Author</th>
          <th>Genre</th>
          <th>Description</th>
          <th>Published</th>
          <th>Publisher</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="book in books" :key="book.id">
          <td>{{ book.title }}</td>
          <td>{{ book.author }}</td>
          <td>{{ book.genre }}</td>
          <td>{{ book.description }}</td>
          <td>{{ book.published }}</td>
          <td>{{ book.publisher }}</td>
          <td>
            <button @click="editBook(book)">Edit</button>
            <button @click="deleteBook(book.id)">Delete</button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  data() {
    return {
      books: [],
    };
  },
  mounted() {
    this.fetchBooks();
  },
  methods: {
    fetchBooks() {
      fetch('http://localhost:8080/books')
        .then(response => response.json())
        .then(data => {
          // Sort books by ID in descending order
          this.books = data.sort((a, b) => b.id - a.id);
        })
        .catch(error => {
          console.error('Error fetching books:', error);
        });
    },
    editBook(book) {
  // Redirect to the edit page with the book ID
  this.$router.push({ path: `/edit/${book.id}` });
},
deleteBook(bookId) {
  if (confirm("Are you sure you want to delete this book?")) {
    fetch(`http://localhost:8080/books/${bookId}`, {
      method: 'DELETE',
    })
      .then(response => {
        if (response.ok) {
          // Refresh the book list after deletion
          this.fetchBooks();
        } else {
          alert('Failed to delete book.');
        }
      })
      .catch(error => {
        console.error('Error deleting book:', error);
      });
  }
},
  },
};
</script>

<style scoped>
table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 20px;
}

th,
td {
  padding: 12px;
  border: 1px solid #ddd;
  text-align: left;
}

th {
  background-color: #f2f2f2;
}

tr:hover {
  background-color: #f1f1f1;
  /* Highlight row on hover */
}

button {
  margin-right: 5px;
  /* Space between buttons */
}</style>
