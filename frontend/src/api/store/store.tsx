import { create } from 'zustand'
import { persist } from 'zustand/middleware'
// ...
type State = {
  accessToken: string
  refreshToken: string
  role:string
}

type Action = {
  setAccessToken: (token: string) => void
  setRefreshToken: (token: string) => void
  setRole: (role: string) => void
  getAccessToken:() => string
  getRefreshToken:() => string
  removeAccessToken:() => void
  removeRefreshToken:() => void
}

export const useAuthStore = create(persist<State & Action>(
  (set,get) => ({
    accessToken: "",
    refreshToken: "",
    role: "",
    setAccessToken: (accessToken: string) => set((state) => ({
      accessToken
    })),
    setRefreshToken: (refreshToken: string) => set((state) => ({
      refreshToken
    })),
    setRole: (role: string) => set((state) => ({
      role
    })),
    getAccessToken: () => get().accessToken,
    getRefreshToken: () => get().refreshToken,
    removeAccessToken:() => get().accessToken = "",
    removeRefreshToken:() => get().refreshToken = "",
  }), {
    name: 'auth'
  }
))