<template>
  <div>
    <!-- <title-bar></title-bar>  -->
    <div id="main">
      <router-view></router-view>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { themeVar, appBgAlpha } from './composables/app'
</script>

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
  background-color: rgba(v-bind('themeVar.bg'), v-bind('appBgAlpha'));
  color-scheme: v-bind('themeVar.scheme');
  /* color: var(--color-text-1); */
}

body {
  overscroll-behavior: none;
}
</style>
