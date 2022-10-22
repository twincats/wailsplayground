<template>
  <div>
    <a-typography-title :heading="2"> HomePage </a-typography-title>
    <div>
      <br /><br />
      <!-- <img id="logo" alt="Wails logo" src="./assets/images/logo-universal.png" /> -->
      <a-button type="primary" @click="runtime.WindowReloadApp()"
        >Reload</a-button
      >
      <br /><br />
      <a-space>
        <a-button type="primary" status="success" @click="toggleTransparant()"
          >Toggle Transparant</a-button
        >
        <a-button type="primary" status="warning" @click="toggleTheme()"
          >Toggle Theme</a-button
        >
      </a-space>
      <br /><br />
      <a-slider
        v-model="appBgAlpha"
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

      <div class="flex justify-center">
        <div class="divider-demo">
          <router-link to="/about">go about</router-link
          ><a-divider irection="vertical" />
          <router-link to="/test">go test</router-link>
        </div>
      </div>
    </div>
    <div class="p-3 text-left">
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
import * as runtime from '@wails/runtime'
import { GetMangas } from '@wails/go/manga/Manga'
import { models } from '@wails/go/models'
import { toggleTransparant, appBgAlpha, toggleTheme } from '@/composables/app'

const output = ref<models.Manga[] | []>()
const finds = () => {
  GetMangas().then(res => {
    output.value = res as models.Manga[]
    // alert('output is out')
  })
}
finds()
</script>

<style>
.divider-demo {
  box-sizing: border-box;
  width: 560px;
  padding: 24px;
  border: 30px solid rgb(var(--gray-2));
}
</style>
