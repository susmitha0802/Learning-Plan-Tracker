import React, { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { Alert, Button, Card } from "react-bootstrap";
import { useAuth } from "../../contexts/AuthContext";
import { FaUserCircle } from "react-icons/fa";
import "../../App.css";

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
    <div className="m-lg-5 p-lg-5 d-flex align-items-center flex-column body">
      <Card className="mx-lg-5 mt-lg-5 mb-lg-3 p-5">
        <Card.Body>

          <h1 className="mb-4 d-flex align-items-center">
            <FaUserCircle />
            <span className="px-3">Profile</span>
          </h1>

          {error && <Alert variant="danger">{error}</Alert>}

          <div className="size">
            {currentUser && <p><strong>Name : </strong><i>{currentUser.displayName}</i></p>}
            {currentUser && <p><strong>Email : </strong><i>{currentUser.email}</i></p>}
            <Link to="/update-profile"> Update Profile </Link>
          </div>

        </Card.Body>
      </Card>
      <div className="w-100 text-center mt-2">
        <Button onClick={handleLogout}>Log Out</Button>
      </div>
    </div>
  )
}
