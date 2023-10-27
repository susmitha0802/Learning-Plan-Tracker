import { useAuth } from "../../contexts/AuthContext";
import { useNavigate } from "react-router-dom";
import { Navbar, Container, Nav } from "react-bootstrap";

export const NavBar = () => {

  const { currentUser, logout } = useAuth();
  const navigate = useNavigate();

  const handleLogout = async () => {

    try {
      await logout();
      navigate("/login");
    } catch {
      console.log("Failed to log out");
    }
  }

  return (
    <Navbar collapseOnSelect expand="lg" className="bg-body-secondary p-3">
      <Container>
        <Navbar.Brand href="/courses">
          <span>
            <b>L</b>earning
            <b> P</b>lan
            <b> T</b>racker
          </span>
        </Navbar.Brand>
        <Navbar.Toggle aria-controls="responsive-navbar-nav" />
        <Navbar.Collapse id="responsive-navbar-nav">
          <Nav className="ms-auto">
            <Nav.Link className="mx-lg-5 px-lg-5" href="/courses">Courses</Nav.Link>
            <Nav.Link className="mx-lg-5 px-lg-5" href="/profile">Profile</Nav.Link>
            <Nav.Link className="mx-lg-5 px-lg-5" href="/mentor">Mentor</Nav.Link>
            {
              currentUser && <Nav.Link className="mx-lg-5 px-lg-5" onClick={handleLogout}>Logout</Nav.Link>
            }
          </Nav>
        </Navbar.Collapse>
      </Container>
    </Navbar>
  );
}

