import { useLocation, Navigate, Outlet } from "react-router-dom";
import { useAuth } from "../context/auth";
import { useAuthStore } from "../../../api/store/store";

export const RequierAuth = () => {
    let getToken = useAuthStore(state => state.getAccessToken)
    const location = useLocation();
    console.log("require : ",getToken())
    return (
        getToken() !== ""?<Outlet/> : <Navigate to = "/auth" state = {{from:location}} replace />
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
