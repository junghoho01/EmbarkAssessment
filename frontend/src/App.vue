<template>
  <div id="app">
    <Navbar /> <!-- Include the Navbar component -->
    <router-view></router-view> <!-- The main content area for routing -->
  </div>
</template>

<script>
import Navbar from './components/Navbar.vue'; // Import the Navbar component

export default {
  name: 'App',
  components: {
    Navbar, // Register the Navbar component
  },
};
</script>

<style>
/* Add any global styles here */
</style>
