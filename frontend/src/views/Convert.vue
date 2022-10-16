<template>
  <div class="flex flex-col h-full">
    <div>
      <div id="tool" class="text-left">
        <div class="w-full">
          <div>Quality:</div>
          <a-slider :default-value="50" />
        </div>
        <div class="w-full">
          <div>Resize :</div>
          <a-slider :default-value="50" />
        </div>
        <div class="w-130">
          <div>Delete :</div>
          <a-space direction="vertical" size="large">
            <a-radio-group>
              <a-radio value="A">Yes</a-radio>
              <a-radio value="B">No</a-radio>
            </a-radio-group>
          </a-space>
        </div>
        <div class="w-180">
          <div>Convert Engine :</div>
          <a-space direction="vertical" size="large">
            <a-radio-group>
              <a-radio value="A">libwebp</a-radio>
              <a-radio value="B">Ext</a-radio>
            </a-radio-group>
          </a-space>
        </div>
        <div class="w-130">
          <div>Format Image :</div>
          <a-space direction="vertical" size="large">
            <a-checkbox-group :default-value="['1']">
              <a-checkbox value="1">*.jpg, *png</a-checkbox>
              <a-checkbox value="2">*webp</a-checkbox>
            </a-checkbox-group>
          </a-space>
        </div>
      </div>
    </div>
    <div class="h-10 bg-blue-900">title</div>
    <div class="h-full bg-gray-900 overflow-auto p-3 text-left">
      <a-list>
        <a-list-item v-for="(item, i) in fileManga" :key="i">
          {{ item }}
        </a-list-item>
      </a-list>
    </div>
    <div class="bg-gray-700">status</div>
  </div>
</template>

<script setup lang="ts">
import { GetFileManga } from '@wails/go/file/FileReader'

const fileManga = ref<string[]>([])
GetFileManga().then(resp => {
  fileManga.value = resp
})
</script>

<style lang="less" scoped>
#tool {
  border-bottom: 1px solid white;
  display: flex;
  justify-content: space-between;
  > div {
    padding: 0.5rem;
    // background-color: rgb(31, 41, 55);
    display: flex;
    flex-direction: column;
    &:not(:last-child) {
      border-right: 1px solid white;
    }
  }
}
</style>
