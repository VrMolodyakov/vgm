// App.tsx (или любой другой компонент, который является корневым в вашем приложении)
import React, { ReactNode } from 'react';
import { AxiosInstanceProvider } from './context/context';
import { createMusicClient } from '../../../api/axios/axiosInstance';
import config from '../../../config/config';
import { useAuthStore } from '../../../api/store/store';

type Props = {
    children: ReactNode
  }

export function ClientProvider({ children }: Props): JSX.Element {
    let [getToken, removeAccessToken, removeRefreshToken, setAccessToken] = useAuthStore(state => [
        state.getRefreshToken,
        state.removeAccessToken,
        state.removeRefreshToken,
        state.setAccessToken
    ])

    let client = createMusicClient(
        config.MusicServerUrl,
        config.RefreshTokenUrl,
        getToken,
        removeAccessToken,
        removeRefreshToken,
        setAccessToken
    )
    return (
        <AxiosInstanceProvider axiosInstance={client}>
           {children}
        </AxiosInstanceProvider>
    )
}

