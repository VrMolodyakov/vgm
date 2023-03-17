import { useContext } from "react";
import { useLocation, Navigate, Outlet } from "react-router-dom";
import AuthContext from "../context/auth";
import { AuthContextType } from "../types/auth-context-type";


const RequierAuth = () => {
    const { auth } = useContext(AuthContext) as AuthContextType;
    const location = useLocation();
    return (
        auth !== undefined?<Outlet/> : <Navigate to = "/auth" state = {{from:location}} replace />
    );

}
export default RequierAuth;