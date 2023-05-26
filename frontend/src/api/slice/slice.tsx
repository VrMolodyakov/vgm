import { BaseQueryApi, FetchArgs, createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react'
import { RootState } from '../store/store'
import { AuthState, logout, setCredentials } from '../../features/auth/slice/auth-slice'

const baseQuery = fetchBaseQuery({
    baseUrl: 'http://localhost:8080',
    credentials: 'include',
    prepareHeaders: (headers, { getState }) => {
        const token =(getState() as RootState).auth.accessToken
        if (token) {
            headers.set("authorization", `Bearer ${token}`)
        }
        return headers
    }
})

const baseQueryWithReauth = async (args:string | FetchArgs, api:BaseQueryApi, extraOptions:{}) => {
    let result = await baseQuery(args, api, extraOptions)

    if (result.error && result.error.status === 401) {
        console.log('sending refresh token')
        const refreshResult = await baseQuery('/auth/refresh', api, extraOptions)
        console.log(refreshResult)
        if (refreshResult?.data) {
            const refreshData = refreshResult.data as AuthState;
            api.dispatch(setCredentials(refreshData))
            result = await baseQuery(args, api, extraOptions)
        } else {
            api.dispatch(logout())
        }
    }

    return result
}

export const apiSlice = createApi({
    baseQuery: baseQueryWithReauth,
    endpoints: builder => ({})
})