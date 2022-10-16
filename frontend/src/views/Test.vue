<template>
  <div>
    <h1>test</h1>

    <div>
      <div class="border-b mb-5">
        {{ result }}
      </div>
      <div v-if="result">
        {{ result.manga[0].title }} <br />
        D:\DATA\Manga\{{ mangaUrl(result.manga[0].title) }}\cover.webp <br />
        <img :src="`/file/${mangaUrl(result.manga[0].title)}/cover.webp`" />
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { GetMangaHome } from '@wails/go/manga/Manga'
import type { models } from '@wails/go/models'

const page = ref(1)
const limit = ref(14)
const result = ref<models.MangaHome | null>(null)
GetMangaHome(page.value, limit.value).then(val => {
  result.value = val
})

const mangaUrl = (url: string) => {
  const fixed = url.replaceAll(/[^A-Za-z0-9_\-[\]()' ~.,!@&]|\.+$/gm, '')
  return fixed.trim()
}
</script>
