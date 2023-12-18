import { useState } from "react";
import { Formik, Form as FormikForm, Field, ErrorMessage } from "formik";
import * as Yup from "yup";
import { useMutation } from "react-query";
import axios from "axios";
import { Link, useNavigate } from "react-router-dom";
import { Alert, Button, Card, Form, Row, Col } from "react-bootstrap";
import { useAuth } from "../../contexts/AuthContext";
import { useUser } from "../../contexts/UserContext";
import "../../common/styles/Auth.css";

const initialValues = {
  name: "",
  email: "",
  password: "",
  confirmPassword: "",
  role: ""
}

const validationSchema = Yup.object({
  name: Yup.string()
    .required("Required!"),
  email: Yup.string()
    .email("Invalid Email Format!")
    .required("Required!"),
  password: Yup.string()
    .min(6, 'Password must be at least 6 characters')
    .required('Password is required!'),
  confirmPassword: Yup.string()
    .oneOf([Yup.ref('password'), null], 'Passwords must match')
    .required('Confirm Password is required!'),
  role: Yup.string()
    .required("Required!"),
});

export const Signup = () => {
  const { signup, displayName } = useAuth();
  const { addUserToLocalStorage } = useUser();
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);
  const mutation = useMutation((values) => axios.post(`http://localhost:4000/userDetails`, values));
  const navigate = useNavigate();

  const handleSubmit = async (values) => {
    if (values.password !== values.confirmPassword) {
      return setError("Passwords do not match");
    }

    try {
      setError("");
      setLoading(true);
      await signup(values.email, values.password);
      displayName(values.name);
      addUserToLocalStorage(values.email);
      const details = {
        name: values.name,
        email: values.email,
        role: values.role
      }
      mutation.mutate(details);
      navigate("/login");
    }
    catch {
      setError("Failed to create an account");
    }

    setLoading(false);
  }

  return (
    <div className="p-lg-5 d-flex align-items-center flex-column body">
      <Card className="m-lg-5 p-5">
        <Card.Body>
          <h2 className="mb-4 text-center">Create Account</h2>
          {error && <Alert variant="danger">{error}</Alert>}
          <Formik
            initialValues={initialValues}
            validationSchema={validationSchema}
            onSubmit={handleSubmit}
          >
            <FormikForm>
              <Row>
                <Col>
                  <Form.Group>
                    <Form.Label htmlFor="name">Name</Form.Label>
                    <div className="mb-3">
                      <Field
                        className="w-100"
                        name="name"
                        id="name"
                        placeholder="Enter your name"
                        autoComplete="true"
                      />
                      <ErrorMessage component="div" className="message error" name="name" />
                    </div>
                  </Form.Group>
                </Col>
                <Col>
                  <Form.Group>
                    <Form.Label htmlFor="email">Email</Form.Label>
                    <div className="mb-3">
                      <Field
                        className="w-100"
                        type="email"
                        name="email"
                        id="email"
                        placeholder="Enter your email"
                        autoComplete="true"
                      />
                      <ErrorMessage component="div" className="message error" name="email" />
                    </div>
                  </Form.Group>
                </Col>
              </Row>
              <Row>
                <Col>
                  <Form.Group>
                    <Form.Label htmlFor="password">Password</Form.Label>
                    <div className="mb-3">
                      <Field
                        className="w-100"
                        type="password"
                        name="password"
                        id="password"
                        placeholder="Enter your password"
                        autoComplete="true"
                      />
                      <ErrorMessage component="div" className="message error" name="password" />
                    </div>
                  </Form.Group>
                </Col>
                <Col>
                  <Form.Group>
                    <Form.Label htmlFor="confirmPassword">Confirm Password</Form.Label>
                    <div className="mb-3">
                      <Field
                        className="w-100"
                        type="password"
                        name="confirmPassword"
                        id="confirmPassword"
                        placeholder="Confirm your password"
                        autoComplete="true"
                      />
                      <ErrorMessage component="div" className="message error" name="confirmPassword" />
                    </div>
                  </Form.Group>
                </Col>
              </Row>
              <Row>
                <Col>
                  <Form.Group>
                    <Form.Label htmlFor="role">Role</Form.Label>
                    <div className="mb-3">
                      <Field
                        className="w-100"
                        as="select"
                        name="role"
                        id="role"
                      >
                        <option>--- Select your role ---</option>
                        <option value="admin">Admin</option>
                        <option value="mentor">Mentor</option>
                        <option value="mentee">Mentee</option>
                      </Field>
                      <ErrorMessage component="div" className="message error" name="role" />
                    </div>
                  </Form.Group>
                </Col>
              </Row>
              <Button disabled={loading} className="w-100" type="submit">Sign Up</Button>
            </FormikForm>
          </Formik>
        </Card.Body>
      </Card>
      <div className="w-100 text-center">
        Already have an account? <Link to="/login">Login</Link>
      </div>
    </div>
  )
}
