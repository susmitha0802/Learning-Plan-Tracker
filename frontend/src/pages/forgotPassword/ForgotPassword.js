import { useState } from "react";
import { Formik, Form as FormikForm, Field, ErrorMessage } from "formik";
import * as Yup from "yup";
import { Link } from "react-router-dom";
import { Alert, Button, Card, Form} from "react-bootstrap";
import { useAuth } from "../../contexts/AuthContext";
import "../../common/styles/Auth.css";

const initialValues = {
  email: ""
}

const validationSchema = Yup.object({
  email: Yup.string()
    .email("Invalid Email Format!")
    .required("Required!")
});

export const ForgotPassword = () => {
  const { resetPassword } = useAuth();
  const [error, setError] = useState("");
  const [message, setMessage] = useState("");
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (values) => {
    try {
      setMessage("");
      setError("");
      setLoading(true);
      await resetPassword(values.email);
      setMessage("Check your inbox for further instructions");
    } catch {
      setError("Failed to reset password");
    }

    setLoading(false);
  }

  return (
    <div className="m-lg-5 p-lg-5 d-flex align-items-center flex-column">
      <Card className="mx-lg-5 my-lg-3 p-5">
        <Card.Body>
          <h2 className="mb-4">Reset Password</h2>
          {error && <Alert variant="danger">{error}</Alert>}
          {message && <Alert variant="success">{message}</Alert>}
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
              <Button disabled={loading} className="w-100" type="submit">Reset Password</Button>
            </FormikForm>
          </Formik>
          <div className="mt-3">
            <Link to="/login">Login</Link>
          </div>
        </Card.Body>
      </Card>
      <div className="card-caption">
        Need an account? <Link to="/signup">Sign Up</Link>
      </div>
    </div>
  )
}
