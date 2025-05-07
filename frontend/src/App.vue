
<template>
  <div id="app">
    <h1>Vue Frontend</h1>
    <p>{{ message }}</p>
  </div>
</template>

<script>
export default {
  data() {
    return {
      message: "Loading..."
    };
  },

mounted() {
    fetch("http://localhost:8080") // это должно быть имя сервиса (backend) и порт (8080)
        .then(res => res.text())
        .then(data => {
            this.message = data;
        })
        .catch(err => {
            this.message = "Error: " + err;
        });
}
};
</script>

<style>
body {
  font-family: sans-serif;
  padding: 2rem;
}
</style>
