<template>
  <div id="main">
    <title-bar></title-bar>
    <br /><br />
    <!-- <img id="logo" alt="Wails logo" src="./assets/images/logo-universal.png" /> -->
    <a-button type="primary" @click="runtime.WindowReloadApp()"
      >Reload</a-button
    >
    <br /><br />
    <a-button type="primary" status="success" @click="trans"
      >Toggle Transparant</a-button
    >
    <br /><br />
    <a-slider
      v-model="bgColorMain[3]"
      :step="0.01"
      :min="0"
      :max="1"
      :style="{ width: '200px' }"
    />
    <br /><br />
    <a-button type="primary" status="success" @click="finds"
      >Get Manga</a-button
    >
    <br /><br />
    <div class="p-3 text-left overflow-auto h-120">
      <a-list>
        <template #header> Manga </template>
        <template #empty><div class="py-5">Tidak ada data</div> </template>
        <a-list-item v-for="(item, i) in output" :key="i">{{
          item.title
        }}</a-list-item>
      </a-list>
    </div>
  </div>
</template>

<script lang="ts" setup>
import * as runtime from '../wailsjs/runtime'
import { GetManga } from '../wailsjs/go/main/App'
import type { main } from '../wailsjs/go/models'

const bgColorMain = ref([35, 35, 36, 0.95])
const bgAlphaMain = ref(0.95)

const trans = () => {
  if (bgColorMain.value[3] == 1) {
    bgColorMain.value[3] = bgAlphaMain.value
  } else {
    bgAlphaMain.value = bgColorMain.value[3]
    bgColorMain.value[3] = 1
  }
}

const output = ref<main.Manga[] | []>()
const finds = () => {
  GetManga().then(res => {
    output.value = res as main.Manga[]
    // alert('output is out')
  })
}
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
  min-height: 100%;
  background-color: rgba(v-bind(bgColorMain));
}

body {
  overscroll-behavior: none;
}
</style>
