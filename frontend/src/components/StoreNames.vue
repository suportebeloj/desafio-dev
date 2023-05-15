<template>
  <div>
    <ul v-if="transactionStore.getStoreNames.length > 0">
      <RouterLink
        :to="'/store/' + store"
        v-for="store in transactionStore.getStoreNames"
        :key="store"
      >
        <li class="">{{ store }}</li>
      </RouterLink>
    </ul>
    <div v-else>
      <p>there are no transactions in the bank, upload a new CNAB file</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useCounterStore } from '@/stores/transactions'
import { RouterLink } from 'vue-router'
const transactionStore = useCounterStore()
</script>

<style scoped lang="sass">
ul
    width: 100%
    list-style: none
    padding-block: 10px
    font-family: 'Lucida Sans', 'Lucida Sans Regular', 'Lucida Grande', 'Lucida Sans Unicode', Geneva, Verdana, sans-serif

    a
        text-decoration: none
        color: #555

        &:hover
            color: #000

    li
        padding: 10px
        cursor: pointer

        &:hover
            background: #ccc
</style>
