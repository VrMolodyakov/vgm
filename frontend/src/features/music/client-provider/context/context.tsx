// AxiosInstanceProvider.tsx
import React, { createContext, useContext } from 'react';
import { AxiosInstance } from 'axios';

interface AxiosInstanceContextValue {
  axiosInstance: AxiosInstance;
}

const AxiosInstanceContext = createContext<AxiosInstanceContextValue | undefined>(undefined);

interface AxiosInstanceProviderProps {
  axiosInstance: AxiosInstance;
  children: React.ReactNode; // Add the 'children' property
}

export const AxiosInstanceProvider: React.FC<AxiosInstanceProviderProps> = ({ axiosInstance, children }) => {
  return (
    <AxiosInstanceContext.Provider value={{ axiosInstance }}>
      {children}
    </AxiosInstanceContext.Provider>
  );
};

export function useMusicClient(): AxiosInstance {
  const context = useContext(AxiosInstanceContext);

  if (!context) {
    throw new Error('useAxiosInstance must be used within an AxiosInstanceProvider');
  }

  return context.axiosInstance;
}