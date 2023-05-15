import { Store } from './../services/api/interfaces'
import { HTTPClientApiService } from './../services/api/http'
import { defineStore } from 'pinia'

export const useCounterStore = defineStore('transactions', {
  state: () => {
    return {
      storeNames: [] as string[],
      storeInfo: {} as Store
    }
  },

  getters: {
    getStoreNames: (state) => state.storeNames,
    getStoreInfo: (state) => state.storeInfo
  },

  actions: {
    async LoadStoreNames() {
      const service = new HTTPClientApiService()
      const result = await service.GetStoreNames()

      if (Array.isArray(result) && result.every((elemento) => typeof elemento === 'string')) {
        this.storeNames = result
      }
    },
    async GetStoreInfo(storeName: string) {
      const service = new HTTPClientApiService()

      const result = await service.GetInfo(storeName)
      if (result !== null) {
        this.storeInfo = result
      }
    }
  }
})
