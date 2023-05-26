import {createSlice, PayloadAction } from "@reduxjs/toolkit"
import { RootState } from "../../../api/store/store"

export interface AuthState {
    accessToken:string
    refreshToken:string
}

const initialState: AuthState = {
    accessToken: "",
    refreshToken:""
  }

const authSlice = createSlice({
    name:"auth",
    initialState,
    reducers: {
        setCredentials: (state, action: PayloadAction<AuthState>) => {
          const {accessToken,refreshToken} = action.payload
          state.accessToken = accessToken
          state.refreshToken = refreshToken
        },
        logout: (state) => {
            state.accessToken = ""
            state.refreshToken = ""
          }
      }
})

export const {setCredentials,logout} = authSlice.actions
export default authSlice.reducer
export const currentAccessToken = (state:RootState) => state.auth.accessToken
export const currentRefreshToken = (state:RootState) => state.auth.refreshToken
