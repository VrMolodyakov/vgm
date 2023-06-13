import { Outlet } from "react-router-dom"
import "./layout.css"

const Layout = () => {
    return (
        <main className="mainApp">
            <Outlet />
        </main>
    )
}

export default Layout