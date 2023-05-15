<template>
  <div class="container">
    <div class="container__painel">
      <button class="button" @click="showModal = true">upload CNAB</button>
    </div>
    <div class="containter__content">
      <store-names />
    </div>
    <Teleport to="body">
      <div v-if="showModal">
        <modal-updaload-cnab @close="showModal = false" />
      </div>
    </Teleport>
  </div>
</template>

<script setup lang="ts">
import { Teleport, ref, onMounted } from 'vue'
import ModalUpdaloadCnab from '@/components/ModalUpdaloadCnab.vue'
import StoreNames from '@/components/StoreNames.vue';
import { useCounterStore } from '@/stores/transactions'

const showModal = ref<boolean>(false)

const transactionStore = useCounterStore()

onMounted(async () => {
  await transactionStore.LoadStoreNames()
})
</script>

<style scoped lang="sass">
.container
    width: 100%
    height: 100%
    background: #e9e9e9
    padding-inline: 3.5em
    padding-top: 20px

    &__painel
        width: 100%
        height: fit-content
        padding-block: 0.3em
        border: 1px solid #ccc
        display: flex
        justify-content: center
        align-items: center
        border-radius: 3px


.button
    padding: 0.5em

.container
    max-width: 100vw
    max-height: calc( 100vh - 50px )
</style>
