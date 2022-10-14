<template>
  <div>
    <!-- <title-bar></title-bar>  -->
    <menu-bar />
    <side-bar />
    <div id="main">
      <router-view></router-view>
    </div>
    <footer-bar />
  </div>
</template>

<script lang="ts" setup>
import { themeVar, appBgAlpha } from './composables/app'
import '@arco-design/web-vue/es/Message/style/css.js'
</script>

<style>
#main {
  height: calc(100vh - 56px);
  overflow: auto;
  background-color: rgba(v-bind('themeVar.bg'), v-bind('appBgAlpha'));
  color-scheme: v-bind('themeVar.scheme');
  /* color: var(--color-text-1); */
}
</style>
