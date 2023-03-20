import { useContext } from "react";
import { useLocation, Navigate, Outlet } from "react-router-dom";
import AuthContext from "../context/auth";
import { AuthContextType } from "../types/auth-context-type";


const RequierAuth = () => {
    const { auth } = useContext(AuthContext) as AuthContextType;
    const location = useLocation();
    console.log("require : ",auth)
    return (
        auth != undefined && auth !== ""?<Outlet/> : <Navigate to = "/auth" state = {{from:location}} replace />
    );

}
export default RequierAuth;