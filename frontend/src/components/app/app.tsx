import { Navbar, Nav } from "react-bootstrap";
import { Routes, Route } from "react-router-dom";
import RequierAuth from "../../features/auth/requier/requier-auth";
import IndexRouter from "../../routes/index-router";
import Home from "../home";
import Layout from "../layout";
import SignInForm from "../signin/sign-in-form";
import SignUpForm from "../signup/sign-up-form";

const App: React.FC = () => {
  return (
    <>
      <Navbar collapseOnSelect expand="lg" bg="blue" variant="white">
        <Navbar.Brand className="Home" href="/">
          Home
        </Navbar.Brand>
        <Navbar.Brand className="Charts" href="/price">
          Charts
        </Navbar.Brand>

        <Navbar.Toggle aria-controls="responsive-navbar-nav" />
        <Navbar.Collapse id="responsive-navbar-nav">
          <Nav className="me-auto"></Nav>
          <Nav className="links">
            <Nav.Link className="navBarLink" href="/auth">
              Sign In
            </Nav.Link>
            <Nav.Link className="navBarLink" href="/reg">
              Sign up
            </Nav.Link>
            <Nav.Link className="navBarLink">Log out</Nav.Link>
          </Nav>
        </Navbar.Collapse>
      </Navbar>
      <div className="App">
      <Routes>
          <Route path="/" element={<Layout />}>
            <Route path="auth" element={<SignInForm />} />
            <Route path="reg" element={<SignUpForm />} />
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
