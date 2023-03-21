import { useLocation, Navigate, Outlet } from "react-router-dom";
import { useAuth } from "../context/auth";


const RequierAuth = () => {
    const {auth} = useAuth();
    const location = useLocation();
    console.log("require : ",auth)
    return (
        auth !== ""?<Outlet/> : <Navigate to = "/auth" state = {{from:location}} replace />
    );

}
export default RequierAuth;