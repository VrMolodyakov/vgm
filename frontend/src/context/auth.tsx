import { createContext, Dispatch, ReactNode, useState } from "react";

type AuthContextType = {
    auth: string | null
    setAuth:Dispatch<string>
}
const [auth, setAuth] = useState({ localStorage.getItem("access_token") });

const AuthContext = createContext<AuthContextType>({
    auth:auth,
    setAuth:setAuth
});



export const AuthProvider = (children  : ReactNode) => {
    

    return (
        <AuthContext.Provider value={{ auth, setAuth }}>
            {children}
        </AuthContext.Provider>
    )
}

export default AuthContext;