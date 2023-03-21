import { useContext } from "react";
import AuthContext from "../context/auth";



const useAuth = () => {
    const ctx =  useContext(AuthContext);
    if (ctx === undefined){
        throw new Error('AuthContext must be inside a AuthProvider');
    }
    return ctx
}

export default useAuth;