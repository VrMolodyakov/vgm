import { useLocation, Navigate, Outlet } from "react-router-dom";
import { useAuth } from "../context/auth";

export const RequierAuth = () => {
    const {auth} = useAuth();
    const location = useLocation();
    console.log("require : ",auth)
    return (
        auth.token !== ""?<Outlet/> : <Navigate to = "/auth" state = {{from:location}} replace />
    );
}

export const RequierAdminAuth = () => {
    const {auth} = useAuth();
    const location = useLocation();
    console.log("require : ",auth)
    return (
        auth.token !== "" && auth.role !== "ADMIN"?<Outlet/> : <Navigate to = "/auth" state = {{from:location}} replace />
    );

}
