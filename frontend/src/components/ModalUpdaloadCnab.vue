<template>
  <div class="modal">
    <div class="modal__close_modal" @click="$emit('close')"></div>
    <div class="modal__content">
      <h3>Upload CNAB file</h3>
      <span>send you CNAB file and save on database</span>
      <input ref="inputRef" type="file" @change="sendFile" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { useCounterStore } from '@/stores/transactions';
import MaterialSymbolsCancelOutlineVue from './icons/MaterialSymbolsCancelOutline.vue'
import { defineEmits, onMounted, ref } from 'vue'
import { HTTPClientApiService } from '@/services/api/http';

const apiService = new HTTPClientApiService()
const transactionStore = useCounterStore()
const inputRef = ref<HTMLInputElement>()
const emit = defineEmits(['close'])
const sendFile = async () => {
    const formData = new FormData()
    formData.append('transactions', inputRef.value?.files[0])
    if (await apiService.UploadCNAB(formData)) {
        await transactionStore.LoadStoreNames()
        emit('close')
    }
}

</script>

<style scoped lang="sass">
.modal
    width: 100%
    height: 100%
    display: flex
    align-items: center
    justify-content: center

    &__content
        width: 80%
        height: fit-content
        border: 1px solid #ccc
        background: #fff
        display: flex
        flex-direction: column
        justify-content: center
        align-items: center
        padding: 20px

    &__close_modal
        width: 10px
        height: 10px
        padding: 20px
        display: flex
        align-items: center
        justify-content: center
        border-radius: 50%
        position: absolute
        top: 20px
        right: 20px
        cursor: pointer
        background-image: url("../src/assets/MaterialSymbolsCancelOutline.svg")
        background-position: center
        background-repeat: no-repeat
        color: #ccc
        filter: invert(63%)
        transition: filter 300ms

        &:hover
            filter: invert(0%)


.modal
    z-index: 1
    position: fixed
    top: 0
    left: 0
    backdrop-filter: blur(5px)
</style>
