<template>
    <div>
      <h1>Reset Book List</h1>
      <button @click="resetBooks">Reset to Original State</button>
    </div>
  </template>
  
  <script>
  export default {
    methods: {
      resetBooks() {
        fetch('http://localhost:8080/books/reset', {
          method: 'POST',
        }).then(() => {
          alert('Book list reset to original state');
        });
      },
    },
  };
  </script>
  