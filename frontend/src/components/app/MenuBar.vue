<template>
  <div class="menu-bar">
    <div class="w-full flex justify-start">
      <a-button-group size="mini">
        <a-button> <icon-list /> </a-button>
        <a-button> cut </a-button>
        <a-button> find </a-button>
      </a-button-group>
    </div>
    <div class="w-full flex justify-end">
      <a-button-group size="mini">
        <a-button> copy </a-button>
        <a-button> cut </a-button>
        <a-button> find </a-button>
      </a-button-group>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { IconList } from '@arco-design/web-vue/es/icon'
</script>

<style lang="less" scoped>
.menu-demos {
  box-sizing: border-box;
  width: 100%;
  padding: 40px;
  background-color: var(--color-neutral-2);
}

.menu-bar {
  height: 28px;
  background-color: var(--color-menu-light-bg);
  display: flex;
  align-items: center;
}
</style>
