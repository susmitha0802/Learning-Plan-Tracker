import { useState } from "react";
import { Formik, Form as FormikForm, Field, ErrorMessage } from "formik";
import * as Yup from "yup";
import { useAuth } from "../../contexts/AuthContext";
import { Link, useNavigate } from "react-router-dom";
import { Alert, Button, Card, Form } from "react-bootstrap";
import "./Auth.css";

const initialValues = {
    email: "",
    password: "",
    confirmPassword: ""
}

const validationSchema = Yup.object({
    email: Yup.string()
        .email("Invalid Email Format!")
        .required("Required!"),
    password: Yup.string()
        .min(6, 'Password must be at least 6 characters')
        .required('Password is required!'),
    confirmPassword: Yup.string()
        .oneOf([Yup.ref('password'), null], 'Passwords must match')
        .required('Confirm Password is required!')
});

export const Signup = () => {
    const { signup } = useAuth();
    const [error, setError] = useState("");
    const [loading, setLoading] = useState(false);
    const navigate = useNavigate();

    const handleSubmit = async (values) => {
        if (values.password !== values.confirmPassword) {
            return setError("Passwords do not match");
        }

        try {
            setError("");
            setLoading(true);
            await signup(values.email, values.password);
            navigate("/courses");
        } catch {
            setError("Failed to create an account");
        }

        setLoading(false);
    }

    return (
        <div className="m-lg-5 p-lg-5 d-flex align-items-center flex-column">
            <Card className="mx-lg-5 my-lg-3 p-5">
                <Card.Body>
                    <h2 className="mb-4">Create Account</h2>
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
                            <Button disabled={loading} className="w-100" type="submit">Sign Up</Button>
                        </FormikForm>
                    </Formik>
                </Card.Body>
            </Card>
            <div className="w-100 text-center mt-2">
                Already have an account? <Link to="/login">Login</Link>
            </div>
        </div>
    )
}
