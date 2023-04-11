import { createContext, Dispatch, ReactNode, SetStateAction, useContext, useMemo, useState } from "react";
import { AuthContextType } from "../types/auth-context-type";

type Props = {
  children: ReactNode
}
type ContextType = {
  auth: string
  setAuth: Dispatch<SetStateAction<string>>
}

const AuthContext = createContext({} as ContextType)

export function AuthProvider({ children }: Props): JSX.Element {
  const [auth, setAuth] = useState<string>("")
  const value = useMemo(() => ({
    auth, setAuth
  }), [auth]);

  console.log("inside - ", auth)
  return (
    <AuthContext.Provider value={value}>
      {children}
    </AuthContext.Provider>
  )
}

export function useAuth(): ContextType {
  return useContext(AuthContext)
}