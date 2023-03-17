import { Navbar, Nav, Container } from "react-bootstrap";
import IndexRouter from "../../routes/index-router";

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
        <IndexRouter/>
      </div>
    </>
  );
};

export default App;
