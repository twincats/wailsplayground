<template>
  <div>
    <title-bar></title-bar>
    <div id="main">
      <router-view></router-view>
    </div>
  </div>
</template>

<script lang="ts" setup></script>

<style>
#logo {
  display: block;
  width: 50%;
  height: 50%;
  margin: auto;
  padding: 10% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}

#main {
  height: calc(100vh - 30px);
  overflow: auto;
  background-color: rgba(35, 35, 36, 0.95);
  color-scheme: dark;
}

body {
  overscroll-behavior: none;
}
</style>
