import { useEffect, useState } from "react";
import { Formik, Form as FormikForm, Field, ErrorMessage } from "formik";
import * as Yup from "yup";
import { Link, useNavigate } from "react-router-dom";
import { Alert, Button, Card, Form } from "react-bootstrap";
import { useQuery } from "react-query";
import axios from "axios";
import { useAuth } from "../../contexts/AuthContext";
import "../../common/styles/Auth.css";

const initialValues = {
  email: "",
  password: "",
}

const validationSchema = Yup.object({
  email: Yup.string()
    .email("Invalid Email Format!")
    .required("Required!"),
  password: Yup.string()
    .min(6, 'Password must be at least 6 characters')
    .required('Password is required!')
});

export const Login = () => {
  const { login, currentUser } = useAuth();
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  const { data: userDetails, isLoading, isError } = useQuery(["userDetails"], async () => {
    const response = await axios.get(`http://localhost:4000/userDetails`);
    return response.data;
  });

  useEffect(() => {
    const user = currentUser && userDetails?.filter(user => user.email === currentUser.email)
    const role = user && user[0] && user[0].role
    if (role === "admin") {
      navigate("/admin");
    }
    else if (role === "mentor") {
      navigate("/mentees");
    }
    else if (role === "mentee") {
      navigate("/courses");
    }
  }, [currentUser, userDetails, navigate])

  if (isLoading) {
    return <h3>Loading...</h3>
  }

  if (isError) {
    return <h3>Error</h3>
  }

  const handleSubmit = async (values) => {
    try {
      setError("");
      setLoading(true);
      await login(values.email, values.password);
    } catch {
      setError("Failed to log in");
    }

    setLoading(false);

  }


  return (
    <div className="p-lg-5 d-flex align-items-center flex-column body">
      <Card className="m-lg-5 p-5">
        <Card.Body>
          <h2 className="mb-4 text-center">Login</h2>
          {error && <Alert variant="danger">{error}</Alert>}
          <Formik
            initialValues={initialValues}
            validationSchema={validationSchema}
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
                    placeholder="Enter your email"
                    autoComplete="true"
                  />
                  <ErrorMessage component="div" className="message error" name="email" />
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
                    placeholder="Enter your password"
                    autoComplete="true"
                  />
                  <ErrorMessage component="div" className="message error" name="password" />
                </div>
              </Form.Group>
              <Button disabled={loading} className="w-100" type="submit">Login</Button>
            </FormikForm>
          </Formik>
          <div className="mt-3">
            <Link to="/forgot-password">Forgot Password?</Link>
          </div>
        </Card.Body>
      </Card>
      <div className="w-100 text-center">
        Need an account? <Link to="/signup"> Sign Up</Link>
      </div>
    </div>
  )
}
