<template>
  {{ data }}
</template>
<script lang="ts">
import { defineComponent, ref } from "vue"

export default defineComponent({
  setup() {
    const data = ref("")
    fetch("/graphql")
      .then(it => it.text())
      .then(it => {
        data.value = it
      })
    return {
      data,
    }
  },
})
</script>