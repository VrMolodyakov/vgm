import { create } from 'zustand'
import { persist } from 'zustand/middleware'
// ...
type State = {
  accessToken: string
}

type Action = {
  setToken: (token: string) => void
  getToken:() => string
}

export const useAuthStore = create(persist<State & Action>(
  (set,get) => ({
    accessToken: "",
    setToken: (accessToken: string) => set((state) => ({
      accessToken
    })),
    getToken: () => get().accessToken
  }), {
    name: 'auth'
  }
))