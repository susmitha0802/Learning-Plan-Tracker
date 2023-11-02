import { useQuery } from 'react-query';
import axios from 'axios';
import { Formik, Form as FormikForm, Field, ErrorMessage } from "formik";
import * as Yup from "yup";
import { useMutation } from "react-query";
import { Alert, Button, Card, Form, Row, Col, Container } from "react-bootstrap";
import { useState } from 'react';

const initialValues = {
    mentor: "",
    mentee: "",
    course: "",
}

const validationSchema = Yup.object({
    mentor: Yup.string()
        .required("Required!"),
    mentee: Yup.string()
        .required("Required!"),
    course: Yup.string()
        .required("Required!"),
});
export const AssignCourse = () => {
    const { data: userDetails, isLoading: isUserDetailsLoading, isError: isUserDetailsError } = useQuery(["userDetails"], async () => {
        const response = await axios.get(`http://localhost:4000/userDetails`);
        return response.data;
    });

    const { data: courses, isLoading: isCoursesLoading, isError: isCoursesError } = useQuery(['courses'], async () => {
        const response = await axios.get(`http://localhost:4000/courses`);
        return response.data;
    });
    const mutation = useMutation((values) => axios.post(`http://localhost:4000/assigned`, values));

    const [loading, setLoading] = useState(false);

    if (isUserDetailsLoading) {
        return <h3>Loading...</h3>
    }

    if (isUserDetailsError) {
        return <h3>Loading...</h3>
    }
    if (isCoursesLoading) {
        return <h3>Loading...</h3>
    }

    if (isCoursesError) {
        return <h3>Loading...</h3>
    }
    const mentors = userDetails?.filter(user => user.role === "mentor");

    const handleSubmit = async (values) => {
        setLoading(true);
        await mutation.mutate(values);

        setLoading(false);
        return <h1>Assigned</h1>
    }

    return (
        <>
            <h4>Assign Here</h4>
            <Formik
                initialValues={initialValues}
                validationSchema={validationSchema}
                onSubmit={handleSubmit}
            >
                <FormikForm>
                    <Row>
                        <Col>
                            <Form.Group>
                                <Form.Label htmlFor="mentor">Mentor</Form.Label>
                                <div className="mb-3">
                                    <Field
                                        className="w-100"
                                        as="select"
                                        name="mentor"
                                        id="mentor"
                                        placeholder="Enter mentor email"
                                        autoComplete="true"
                                    >
                                        <option>--- Select mentor ---</option>
                                        {mentors.map(mentor => {
                                            return <option key={mentor.id} value={mentor.email}>{mentor.name}</option>
                                        })}
                                    </Field>
                                    <ErrorMessage component="div" className="message error" name="mentor" />
                                </div>
                            </Form.Group>
                        </Col>
                        <Col>
                            <Form.Group>
                                <Form.Label htmlFor="mentee">Mentee</Form.Label>
                                <div className="mb-3">
                                    <Field
                                        className="w-100"
                                        as="select"
                                        name="mentee"
                                        id="mentee"
                                        placeholder="Enter mentee email"
                                        autoComplete="true"
                                    >
                                        <option>--- Select mentee ---</option>
                                        {userDetails.map(user => {
                                            return user.role !== "admin" && <option key={user.id} value={user.email}>{user.name}</option>
                                        })}
                                    </Field>
                                    <ErrorMessage component="div" className="message error" name="mentee" />
                                </div>
                            </Form.Group>
                        </Col>
                        <Col>
                            <Form.Group>
                                <Form.Label htmlFor="course">Course</Form.Label>
                                <div className="mb-3">
                                    <Field
                                        className="w-100"
                                        as="select"
                                        name="course"
                                        id="course"
                                        placeholder="Enter your course"
                                        autoComplete="true"
                                    >
                                        <option>--- Select course ---</option>
                                        {courses.map(course => {
                                            return <option key={course.id} value={course.id}>{course.name}</option>
                                        })}

                                    </Field>
                                    <ErrorMessage component="div" className="message error" name="course" />
                                </div>
                            </Form.Group>
                        </Col>
                    </Row>
                    <Button disabled={loading} className="w-100" type="submit">Assign</Button>
                </FormikForm>
            </Formik>
        </>
    )
}
