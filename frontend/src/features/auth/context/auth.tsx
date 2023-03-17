import { createContext,useState } from "react";
import { AuthContextType } from "../types/auth-context-type";

const AuthContext = createContext<AuthContextType | null>( null );

interface Props {
    children?: React.ReactNode;
}
  
export const AuthProvider: React.FC<Props> = ({children}) => {
    const [auth, setAuth] = useState<string>();

    const saveAuth = (auth:string) =>{
        setAuth(auth)
    }

    return (
        <AuthContext.Provider value={{ auth, saveAuth }}>
            {children}
        </AuthContext.Provider>
    )
}

export default AuthContext;