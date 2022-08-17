<template>
  <div>
    <a-typography-title :heading="2"> About </a-typography-title>
    <router-link to="/">go home</router-link>
    <br /><br />
    <div style="color: var(--color-text-1)">{{ mode }}</div>

    <a-button @click="mode = 'light'">Light</a-button>
    <a-button @click="mode = 'dark'">Dark</a-button>
    <div style="color: var(--color-text-1)">
      {{ themeVar }}
    </div>
  </div>
</template>

<script lang="ts" setup>
import { mode, themeVar } from '@/composables/app'
// const mode = ref('')
</script>
