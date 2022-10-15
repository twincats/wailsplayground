<template>
  <div>
    <h1>test</h1>
    {{ result }}
  </div>
</template>

<script lang="ts" setup>
import { GetMangaHome } from '@wails/go/manga/Manga'
const page = ref(1)
const limit = ref(14)
const result = ref()
GetMangaHome(page.value, limit.value).then(val => {
  result.value = val
})
</script>
