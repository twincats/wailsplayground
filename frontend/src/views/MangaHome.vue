<template>
  <div class="p-3">
    <div class="grid grid-cols-5 lg:grid-cols-10">
      <div v-for="(manga, i) in mangaHome?.manga" :key="i" class="mb-3">
        <a-card hoverable :style="{ height: '100%' }">
          <template #cover>
            <div class="h-210px overflow-hidden">
              <img
                :style="{ width: '100%', transform: 'translateY(-20px)' }"
                alt="dessert"
                :src="`file/${MangaTitleURL(manga.title)}/cover.webp`"
              />
            </div>
          </template>
          <a-card-meta :title="manga.title">
            <template #description> Chapter {{ manga.chapter }} </template>
          </a-card-meta>
        </a-card>
      </div>
    </div>
    <div class="mt-3">
      <a-pagination
        class="justify-center"
        :total="mangaHome?.pagination?.total ? mangaHome?.pagination?.total : 0"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { GetMangaHome } from '@wails/go/manga/Manga'
import type { models } from '@wails/go/models'
import { MangaTitleURL } from '@/composables/Helper'

const mangaHome = ref<models.MangaHome | null>(null)
const nav = reactive({
  page: 1,
  limit: 10,
})

//load manga
GetMangaHome(nav.page, nav.limit).then(res => {
  mangaHome.value = res
})
</script>

<style scoped></style>
