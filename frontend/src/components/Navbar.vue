<template>
    <nav>
        <ul>
            <li><router-link to="/">List Books</router-link></li>
            <li><router-link to="/create">Create Book</router-link></li>
            <li><router-link to="/reset">Reset Book</router-link></li>
            <!-- Add more links here as needed -->
        </ul>
    </nav>
</template>


<script>
export default {
    /* eslint-disable */
    name: 'Navbar',
};
</script>
  
<style scoped>
nav {
    background-color: #333;
    padding: 1rem;
}

nav ul {
    list-style-type: none;
    padding: 0;
    margin: 0;
    display: flex;
    gap: 1rem;
}

nav a {
    color: white;
    text-decoration: none;
}

nav a:hover {
    text-decoration: underline;
}
</style>
  