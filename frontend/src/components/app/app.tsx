import { Navbar, Nav } from "react-bootstrap";
import { Routes, Route, Link,useNavigate } from "react-router-dom";
import {RequierAuth} from "../../features/auth/requier/requier-auth";
import IndexRouter from "../../routes/index-router";
import Home from "../home";
import Layout from "../layout";
import SignInForm from "../signin/sign-in-form";
import SignUpForm from "../signup/sign-up-form";
import { Auth,useAuth } from "../../features/auth/context/auth";
import CreateForm from "../create/create-album";
import AddRemoveInputField from "../input/add-remove-input-field";
import "bootstrap/dist/css/bootstrap.min.css";

//TODO:index router
const App: React.FC = () => {
  const navigate = useNavigate();
  const { setAuth } = useAuth();

  const logout = () =>{
    const auth: Auth = {
      token: "",
      role:""
    }
    setAuth(auth)
    navigate("/auth");
  }

  return (
    <>
      <Navbar collapseOnSelect expand="lg" bg="blue" variant="white">
      <Nav className="me-auto">
        <Nav.Link className="home" as={Link} to="/home">
          VGM
        </Nav.Link>
        <Nav.Link className="news" as={Link} to="/price">
          News
        </Nav.Link>
        </Nav>

        <Navbar.Toggle aria-controls="responsive-navbar-nav" />
        <Navbar.Collapse id="responsive-navbar-nav">
          <Nav className="me-auto"></Nav>
          <Nav className="links">
            <Nav.Link className="navBarLink" as={Link} to="/auth">
              Sign In
            </Nav.Link>
            <Nav.Link className="navBarLink" as={Link} to="/reg">
              Sign up
            </Nav.Link>
            <Nav.Link className="navBarLink" onClick={logout}>Log out</Nav.Link>
          </Nav>
        </Navbar.Collapse>
      </Navbar>
      <div className="App">
      <Routes>
          <Route path="/" element={<Layout />}>
            <Route path="auth" element={<SignInForm />} />
            <Route path="reg" element={<SignUpForm />} />
            <Route path="create" element={<CreateForm />} />
            <Route element = {<RequierAuth/>}>
                <Route path="home" element={<Home />} />              
            </Route>
          </Route>
        </Routes>
      </div>
    </>
  );
};

export default App;
