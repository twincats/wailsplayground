<template>
  <div class="flex flex-col h-full">
    <div>
      <div id="tool" class="text-left">
        <div class="w-150">
          <div>Quality:</div>
          <a-slider :default-value="50" />
        </div>
        <div class="w-full">
          <div>Resize :</div>
          <a-slider :default-value="50" />
        </div>
        <div style="flex: 0 0 132px; padding-right: 0">
          <div>Delete :</div>
          <a-space direction="vertical" size="large">
            <a-radio-group size="mini">
              <a-radio value="A">Yes</a-radio>
              <a-radio value="B">No</a-radio>
            </a-radio-group>
          </a-space>
        </div>
        <div style="flex: 0 0 160px; padding-right: 0">
          <div>Convert Engine :</div>
          <a-space direction="vertical" size="large">
            <a-radio-group>
              <a-radio value="A">libwebp</a-radio>
              <a-radio value="B">Ext</a-radio>
            </a-radio-group>
          </a-space>
        </div>
        <div style="flex: 0 0 188px; padding-right: 0">
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
    <div class="h-10 mb-2" style="background-color: var(--color-bg-4)">
      <a-input
        v-model="data.title"
        size="small"
        :input-attrs="{ class: 'text-center' }"
        readonly
      />
    </div>
    <div
      id="mtable"
      class="h-full overflow-auto p-3 mb-1 text-left"
      style="background-color: var(--color-bg-4)"
    >
      <a-table
        size="small"
        v-model:selected-keys="selectedKeys"
        :show-header="false"
        :pagination="false"
        :columns="columnManga"
        :data="tableManga"
        @row-click="(c)=>selectManga(c as TableManga)"
      />
    </div>
    <div>
      <div id="status_convert"></div>
      <div id="mprogress" class="mt-1">
        <a-progress
          status="success"
          :show-text="false"
          :stroke-width="20"
          :percent="0.2"
        />
      </div>
    </div>
    <div class="text-xs flex justify-between px-3 py-2 h-80px">
      <div class="status">
        <div style="min-width: 128px">
          Size Before : <strong class="text-red-500">1500 MB</strong>
        </div>
        <div style="min-width: 117px">
          Size After : <strong class="text-blue-500">1100 MB</strong>
        </div>
        <div style="min-width: 85px">
          Percent : <strong class="text-green-500">50%</strong>
        </div>
      </div>
      <div class="w-full"></div>
      <div class="w-full"></div>
      <div class="bmenu">
        <a-space>
          <a-button>Reset</a-button>
          <a-button>Cover Fix</a-button>
          <a-button>Save Log</a-button>
          <a-button type="primary">Convert</a-button>
        </a-space>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { GetFileManga } from '@wails/go/file/FileReader'

interface TableManga {
  title: string
}
const tableManga = reactive<TableManga[]>([])
const columnManga = [{ title: 'Manga', dataIndex: 'title' }]
GetFileManga().then(resp => {
  resp.forEach(item => {
    tableManga.push({ title: item })
  })
})

const selectedKeys = ref([])
const data = reactive({
  title: '',
})

const selectManga = (c: TableManga) => {
  data.title = c.title
}
</script>

<style lang="less" scoped>
#tool {
  border-bottom: 1px solid white;
  display: flex;
  justify-content: space-between;
  > div {
    padding: 0.5rem;
    font-size: 0.75rem;
    display: flex;
    flex-direction: column;
    &:not(:last-child) {
      border-right: 1px solid white;
    }
    > div {
      height: 100%;
    }
  }
}
#status_convert {
  // class="h-25 p-3 overflow-y-scroll text-left"
  height: 6rem;
  overflow-y: scroll;
  text-align: left;
  padding: 0.75rem;
  background-color: var(--color-bg-4);
}
.status {
  width: 100%;
  display: flex;
  justify-content: space-between;
  div {
    text-align: left;
    align-self: center;
    padding-right: 0.75rem;
  }
}

.bmenu {
  display: flex;
}

.arco-radio,
button,
.arco-checkbox {
  font-size: 0.75rem !important;
}
</style>

<style>
#mprogress .arco-progress-line {
  border-radius: 0 !important;
}
#mprogress .arco-progress-line-bar {
  border-radius: 0 !important;
}
#mtable .arco-table-td {
  font-size: 0.75rem !important;
}
</style>
