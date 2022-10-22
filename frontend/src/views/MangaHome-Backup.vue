<template>
  <div class="p-3">
    <div class="grid grid-cols-5 lg:grid-cols-10">
      <div v-for="(manga, i) in mangaHomeFilter" :key="i" class="mb-3">
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
        v-model:page-size="limitMangaFilter"
        @change="paginationChange"
        :total="mangaHome?.pagination?.total ? mangaHome?.pagination?.total : 0"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { GetMangaHome } from '@wails/go/manga/Manga'
import type { models } from '@wails/go/models'
import { MangaTitleURL, GetBreakPoints } from '@/composables/Helper'

const mangaHome = ref<models.MangaHome | null>(null)
const nav = reactive({
  page: 1,
  limit: 30,
})

interface Nav {
  page: number
  limit: number
}

const sliceFilter = reactive({
  start: 0,
  end: 10,
})

// watch(
//   () => sliceFilter.start,
//   n => {
//     if (n == 10) {
//       nav.page = nav.page + 1
//     } else if (n == 20) {
//       nav.page = nav.page + 2
//     } else {
//       nav.page = n + 1
//     }
//   }
// )

const loadingM = ref(false)

//load manga
const loadManga = (n: Nav) => {
  loadingM.value = true
  GetMangaHome(n.page, n.limit).then(res => {
    mangaHome.value = res
    loadingM.value = false
  })
}
loadManga(nav)

//test
const { breakpoints } = GetBreakPoints()
const lg = breakpoints.greater('lg')

const mangaHomeFilter = computed(() => {
  if (mangaHome.value) {
    if (lg.value) {
      return mangaHome.value.manga
    } else {
      return mangaHome.value.manga.slice(sliceFilter.start, sliceFilter.end)
    }
  } else {
    return []
  }
})
const limitMangaFilter = computed(() => {
  if (lg.value) {
    return 30
  } else {
    return 10
  }
})
const paginationChange = (p: number) => {
  console.log(p)

  if (lg.value) {
    nav.page = p
    loadManga(nav)
  } else {
    if (p % 3 == 1) {
      //change only every 3 page (30)
      nav.page = Math.floor(p / 3) + 1

      console.log('modulo')
      loadManga(nav)
      if (loadingM.value == false) {
        sliceFilter.start = 0
        sliceFilter.end = 10
      }
    } else {
      if (p % 3 == 2) {
        sliceFilter.start = 10
        sliceFilter.end = 20
      } else {
        sliceFilter.start = 20
        sliceFilter.end = 30
      }
      console.log('empty', p)
    }
  }
}
</script>

<style scoped></style>
