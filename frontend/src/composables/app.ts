import { useColorMode } from '@vueuse/core'
import { computed, reactive } from 'vue'

export const mode = useColorMode({
  selector: 'body',
  attribute: 'arco-theme',
  emitAuto: true,
})

export const { next: toggleTheme } = useCycleList(['dark', 'light'], {
  initialValue: mode,
})

const darkThemeVar = reactive({
  bg: [24, 24, 24],
  text: [255, 255, 255],
  scheme: 'dark',
})
const lightThemeVar = reactive({
  bg: [250, 250, 250],
  text: [5, 5, 5],
  scheme: 'light',
})

const tempAppBgAlpha = ref(1)
export const appBgAlpha = ref(0.85)
export const toggleTransparant = () => {
  if (appBgAlpha.value == 1) {
    appBgAlpha.value = tempAppBgAlpha.value
  } else {
    tempAppBgAlpha.value = appBgAlpha.value
    appBgAlpha.value = 1
  }
}

export const themeVar = computed(() => {
  if (mode.value == 'dark') {
    return darkThemeVar
  } else {
    return lightThemeVar
  }
})
