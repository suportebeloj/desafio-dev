<template>
  <div class="transaction_area">
    <div class="container">
      <div class="container__store_info">
        <div class="info">
          <div>{{ info?.market_name }}</div>
          <div>{{ info?.owner }}</div>
        </div>
        <div class="balance">
          <span>{{ info?.balance }}</span>
        </div>
      </div>
    </div>
    <div class="transaction_area__content">
      <table>
        <thead>
          <tr>
            <th>id</th>
            <th>type</th>
            <th>date</th>
            <th>value</th>
            <th>cpf</th>
            <th>card</th>
            <th>time</th>
            <th>owner</th>
            <th>market</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="transaction in info?.operations" :key="transaction.id">
            <td v-for="data in transaction" :key="data">
              {{ data }}
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    <router-link to="/">
        <div class="back">
            <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 24 24"><path fill="currentColor" d="M10 22L0 12L10 2l1.775 1.775L3.55 12l8.225 8.225L10 22Z"/></svg>
        </div>
    </router-link>
  </div>
</template>

<script lang="ts" setup>
import { useCounterStore } from '@/stores/transactions'
import { log } from 'console'
import { onBeforeMount, ref } from 'vue'
import { useRoute } from 'vue-router'
import { Store } from '../services/api/interfaces'

const route = useRoute()
const transactionStore = useCounterStore()
const info = ref<Store>()
onBeforeMount(async () => {
  await transactionStore.GetStoreInfo(route.params.name as string)
  info.value = transactionStore.getStoreInfo
  console.log(transactionStore.getStoreInfo)
})
</script>

<style scoped lang="sass">
.container
    width: 100%
    background: #e9e9e9
    display: flex
    flex-direction: column


    &__store_info
        display: flex
        justify-content: space-between
        align-items: center
        padding: 10px
        border-bottom: 1px solid #ccc
        font-family: 'Lucida Sans', 'Lucida Sans Regular', 'Lucida Grande', 'Lucida Sans Unicode', Geneva, Verdana, sans-serif

        @media screen and (max-width: 640px)
            flex-direction: column
            gap: 20px

        div
            flex: 1
            display: flex
            flex-direction: column
            align-items: center
            justify-content: center
            gap: 10px

        .info
            align-items: baseline
            @media screen and (max-width: 640px)
                align-items: center

        .balance::before
            content: "Balance"
            font-weight: bold


.transaction_area
    width: 100%
    height: 100%
    background: #f1f1f1
    position: relative

    &__content
        display: flex
        align-items: center
        justify-content: center
        overflow: auto
        padding: 10px

        @media screen and (max-width: 640px) 
            margin-inline: 20px

    .back
        position: absolute
        bottom: 1.3em
        right: 1.3em
        border: 1px solid #ccc
        width: 3em
        height: 3em
        background: #777
        color: #ccc
        border-radius: 50%
        display: flex
        justify-content: center
        align-items: center
        cursor: pointer

        &:hover
            background: #444
            color: #ccc


table
  width: 100%
  max-width: 100%
  overflow-x: auto

  
  

  th, td
    padding: 12px
    text-align: left
    white-space: nowrap

    &:first-child
      padding-left: 0

    &:last-child
      padding-right: 0



  th
    background-color: #f2f2f2
    font-weight: bold


  tr:nth-child(even)
    background-color: #f2f2f2


.transaction_area
    max-width: 100vw
    max-height: calc( 100vh - 50px )
</style>
