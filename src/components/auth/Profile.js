import React, { useState } from "react";
import { useAuth } from "../../contexts/AuthContext";
import { Link, useNavigate } from "react-router-dom";
import { Alert, Button, Card} from "react-bootstrap";

export const Profile = () => {
  const [error, setError] = useState("");
  const { currentUser, logout } = useAuth();
  const navigate = useNavigate();

  const handleLogout = async () => {
    setError("");

    try {
      await logout();
      navigate("/login");
    } catch {
      setError("Failed to log out");
    }
  }

  return (
    <div className="m-lg-5 p-lg-5 d-flex align-items-center flex-column">
      <Card className="mx-lg-5 my-lg-3 p-5">
        <Card.Body>
          <h2 className="mb-4">Profile</h2>
          {error && <Alert variant="danger">{error}</Alert>}
          <strong>Email:</strong> {currentUser && <p>{currentUser.email}</p>}
          <Link to="/update-profile">
            Update Profile
          </Link>
        </Card.Body>
      </Card>
      <div className="w-100 text-center mt-2">
        <Button onClick={handleLogout}>Log Out</Button>
      </div>
    </div>
  )
}
