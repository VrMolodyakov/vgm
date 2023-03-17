import { ReactElement } from "react";
import { Route, Routes } from "react-router-dom";
import Home from "../components/home";
import Layout from "../components/layout";
import SignInForm from "../components/signin/sign-in-form";
import SignUpForm from "../components/signup/sign-up-form";
import RequierAuth from "../features/auth/requier/requier-auth";

const IndexRouter:React.FC = ():ReactElement => {
    return (
        <Routes>
          <Route path="/" element={<Layout />}>
            <Route path="auth" element={<SignUpForm />} />
            <Route path="reg" element={<SignInForm />} />
            <Route element = {<RequierAuth/>}>
                <Route path="home" element={<Home />} />              
            </Route>
          </Route>
        </Routes>
    )
}

export default IndexRouter;