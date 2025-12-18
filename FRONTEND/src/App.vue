<script setup>
import { RouterView } from 'vue-router'
</script>

<template>
  <RouterView />
</template>

<style>
body, html {
  margin: 0;
  padding: 0;
  width: 100%;
  height: 100%;
  overflow: hidden;
}

#app {
  width: 100vw;
  height: 100vh;
  display: block;
  background-color: #ffffff;
  color: #000000;
  font-family: sans-serif;
}

* {
  box-sizing: border-box;
}
</style>
