export interface Store {
  market_name: string
  owner: string
  balance: number
  operations: Operation[]
}

export interface Operation {
  id: number
  type: string
  date: string
  value: number
  cpf: string
  card: string
  time: string
  owner: string
  market: string
}
