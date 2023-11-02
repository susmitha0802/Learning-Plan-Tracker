import { useQuery } from 'react-query';
import axios from 'axios';
import { useAuth } from "../../contexts/AuthContext";
import { useNavigate } from "react-router-dom";
import { Navbar, Container, Nav } from "react-bootstrap";

export const NavBar = () => {
  const { data: userDetails, isLoading, isError } = useQuery(["userDetails"], async () => {
    const response = await axios.get(`http://localhost:4000/userDetails`);
    return response.data;
  });
  const { currentUser, logout } = useAuth();
  const navigate = useNavigate();
  if (isLoading) {
    return <h3>Loading...</h3>
  }

  if (isError) {
    return <h3>Loading...</h3>
  }
  const handleLogout = async () => {

    try {
      await logout();
      navigate("/login");
    } catch {
      console.log("Failed to log out");
    }
  }
  const user = currentUser && userDetails?.filter(user => user.email === currentUser.email)
  const role = currentUser && user[0]?.role;
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
            {
              role === "mentor" && <Nav.Link className="mx-lg-5 px-lg-5" href="/mentor">Mentor</Nav.Link>
            }
            {
              currentUser && <Nav.Link className="mx-lg-5 px-lg-5" onClick={handleLogout}>Logout</Nav.Link>

            }
            {
              currentUser && <Nav.Link className="mx-lg-5 px-lg-5" onClick={() => localStorage.clear()}>Clear</Nav.Link>
            }
          </Nav>
        </Navbar.Collapse>
      </Container>
    </Navbar>
  );
}

