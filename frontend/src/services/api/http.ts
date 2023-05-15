import type { Store } from "./interfaces"

const SERVER_URL = import.meta.env.SERVER_BASE_URL

export class HTTPClientApiService {
  private server_url = SERVER_URL

  constructor() {
    const envUrl = import.meta.env.VITE_SERVER_BASE_URL

    if (!envUrl) {
      throw 'define server base url on envirn'
    }
    this.server_url = envUrl
  }

  async UploadCNAB(form: FormData): Promise<boolean> {
    const response = await fetch(this.server_url + 'new', {
      method: 'POST',
    //   headers: {
    //     'Content-Type': 'multipart/form-data'
    //   },
      body: form
    })

    if (response.status === 200) {
      return true
    }

    return false
  }

  async GetStoreNames(): Promise<string[]> {
    const result = await fetch(this.server_url + 'markets')

    if (result.status === 200) {
      return result.json()
    }

    return []
  }

  async GetInfo(markeName: string): Promise<Store | null> {
    const result = await fetch(this.server_url + "detail/" + markeName)
    if (result.status === 200) {
        return result.json()
    }

    return null
  }
}
