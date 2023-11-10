import { useState } from "react";
import { Formik, Form as FormikForm, Field } from "formik";
import { Link, useNavigate } from "react-router-dom";
import { Alert, Button, Card, Form} from "react-bootstrap";
import { useAuth } from "../../contexts/AuthContext";
import "./Auth.css";

export const UpdateProfile = () => {
  const { currentUser, passwordUpdate, emailUpdate } = useAuth();
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  const initialValues = {
    email: currentUser.email,
    password: "",
    confirmPassword: ""
  }

  const handleSubmit = (values) => {
    if (values.password !== values.confirmPassword) {
      return setError("Passwords do not match");
    }

    const promises = [];
    setLoading(true);
    setError("");

    if (values.email !== currentUser.email) {
      promises.push(emailUpdate(values.email));
    }
    if (values.password) {
      promises.push(passwordUpdate(values.password));
    }

    Promise.all(promises)
      .then(() => {
        navigate("/profile");
      })
      .catch(() => {
        setError("Failed to update account");
      })
      .finally(() => {
        setLoading(false);
      })
  }

  return (
    <div className="m-lg-5 p-lg-5 d-flex align-items-center flex-column">
      <Card className="mx-lg-5 my-lg-3 p-5">
        <Card.Body>
          <h2 className="mb-4">Update Profile</h2>
          {error && <Alert variant="danger">{error}</Alert>}
          <Formik 
            initialValues={initialValues}
            onSubmit={handleSubmit}
          >
            <FormikForm>
              <Form.Group>
                <Form.Label htmlFor="email">Email</Form.Label>
                <div className="mb-3">
                  <Field 
                    className="w-100"
                    type="email" 
                    name="email" 
                    id="email" 
                    required 
                  />
                </div>
              </Form.Group>
              <Form.Group>
                <Form.Label htmlFor="password">Password</Form.Label>
                <div className="mb-3">
                  <Field 
                    className="w-100" 
                    type="password" 
                    name="password" 
                    id="password" 
                    placeholder="Leave blank to keep the same" 
                  />
                </div>
              </Form.Group>
              <Form.Group>
                <Form.Label htmlFor="confirmPassword">Confirm Password</Form.Label>
                <div className="mb-3">
                  <Field 
                    className="w-100" 
                    type="password" 
                    name="confirmPassword" 
                    id="confirmPassword" 
                    placeholder="Leave blank to keep the same" 
                  />
                </div>
              </Form.Group>
              <Button disabled={loading} className="w-100" type="submit">Update</Button>
            </FormikForm>
          </Formik>
        </Card.Body>
      </Card>
      <div className="card-caption">
        <Link to="/profile">Cancel</Link>
      </div>
    </div>
  )
}
