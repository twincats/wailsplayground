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
        v-model:current="nav.page"
        v-model:page-size="nav.limit"
        :auto-adjust="false"
        :total="mangaHome?.pagination?.total ? mangaHome?.pagination?.total : 0"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { GetMangaHome } from '@wails/go/manga/Manga'
import type { models } from '@wails/go/models'
import { MangaTitleURL, GetBreakPoints, SEQUENCE3 } from '@/composables/Helper'

/* INTERFACE */
interface Nav {
  page: number
  limit: number
}

/* INITIAL REACTIVE VARIABLE */
const mangaHome = ref<models.MangaHome | null>(null)
const nav = reactive<Nav>({ page: 1, limit: 10 })
const { breakpoints } = GetBreakPoints()
const lg = breakpoints.greater('lg')

/* INITIAL PRELOAD FUNCTION */
//load manga
const loadManga = (n: Nav) => {
  console.log(nav, 'berore fetching')
  GetMangaHome(n.page, n.limit).then(res => {
    mangaHome.value = res
  })
}

// ON_CREATED function loading Manga Data
loadManga(nav)

// Watching View and Pagination page change
watch([lg, () => nav.page], ([l, p], [_, op]) => {
  if (l) {
    nav.limit = 30
    if (op == p) {
      nav.page = SEQUENCE3(p)
    }
    loadManga(nav)
  } else {
    nav.limit = 10
    if (op == p) {
      nav.page = (p - 1) * 3 + 1
    }
    loadManga(nav)
  }
})
</script>

<style scoped></style>
