import { createContext,useState } from "react";
import { AuthContextType } from "../types/auth-context-type";

const AuthContext = createContext<AuthContextType | undefined>( undefined );

interface Props {
    children?: React.ReactNode;
}
  
export const AuthProvider: React.FC<Props> = ({children}) => {
    const [auth, setAuth] = useState<string | null>(null);


    // const saveAuth = (auth:string) =>{
    //     setAuth(auth)
    // }

    return (
        <AuthContext.Provider value={{ auth, setAuth }}>
            {children}
        </AuthContext.Provider>
    )
}

export default AuthContext;