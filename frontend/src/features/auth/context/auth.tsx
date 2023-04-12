import { createContext, Dispatch, ReactNode, SetStateAction, useContext, useMemo, useState } from "react";
import { AuthContextType } from "../types/auth-context-type";

type Props = {
  children: ReactNode
}

export type Auth = {
  token:string
  role:string
}

type ContextType = {
  auth: Auth
  setAuth: Dispatch<SetStateAction<Auth>>
}

const AuthContext = createContext({} as ContextType)

export function AuthProvider({ children }: Props): JSX.Element {
  const base: Auth = {
    token: "",
    role:""
  };
  const [auth, setAuth] = useState<Auth>(base)
  const value = useMemo(() => ({
    auth, setAuth
  }), [auth]);
  return (
    <AuthContext.Provider value={value}>
      {children}
    </AuthContext.Provider>
  )
}

export function useAuth(): ContextType {
  return useContext(AuthContext)
}