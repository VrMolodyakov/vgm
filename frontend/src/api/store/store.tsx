import { configureStore } from '@reduxjs/toolkit'
import { apiSlice } from '../slice/slice'
import authReducer from '../../features/auth/slice/auth-slice'
// ...

export const store = configureStore({
  reducer: {
      [apiSlice.reducerPath]: apiSlice.reducer,
      auth: authReducer
  },
  middleware: getDefaultMiddleware =>
      getDefaultMiddleware().concat(apiSlice.middleware),
  devTools: true
})

export type RootState = ReturnType<typeof store.getState>