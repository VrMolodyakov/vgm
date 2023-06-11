import { Navbar, Nav } from "react-bootstrap";
import { Routes, Route, Link,useNavigate } from "react-router-dom";
import {RequierAuth} from "../../features/auth/requier/requier-auth";
import Home from "../home";
import Layout from "../layout";
import SignInForm from "../../features/signin/sign-in-form";
import SignUpForm from "../../features/signup/sign-up-form";
import { Auth,useAuth } from "../../features/auth/context/auth";
import "bootstrap/dist/css/bootstrap.min.css";
import { News } from "../../features/music/news/news";
import { PersonList } from "../../features/music/persons/list/list";
import CreateAlbumForm from "../../features/music/album/create-album";
import "./app.css"
import CreatePersonForm from "../../features/music/persons/create/create-person";

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
    <div className="app">
      <Navbar collapseOnSelect expand="lg" bg="blue" variant="white">
      <Nav className="me-auto">
        <Nav.Link className="home" as={Link} to="/home">
          VGM
        </Nav.Link>
        <Nav.Link className="news" as={Link} to="/news">
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
            <Route path="create-album" element={<CreateAlbumForm />} />
            <Route path="create-person" element={<CreatePersonForm />} />
            <Route path="persons" element={<PersonList />} />
            <Route element = {<RequierAuth/>}>
                <Route path="news" element={<News />} />  
                <Route path="home" element={<Home />} />                          
            </Route>
          </Route>
        </Routes>
      </div>
    </div>
  );
};

export default App;
