import { useState } from 'react';
import { useQuery, useMutation } from 'react-query';
import axios from 'axios';
import { Formik, Form as FormikForm, Field, ErrorMessage } from "formik";
import * as Yup from "yup";
import { Button, Form, Row, Col, Alert, Card } from "react-bootstrap";
import "./AssignCourse.css"

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
    // .array()
    //     .min(1, 'Select at least one option')
    //     .required("Required!"),
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

    const { data: assigned, isLoading: isAssignedLoading, isError: isAssignedError, refetch } = useQuery(["assigned"], async () => {
        const response = await axios.get(`http://localhost:4000/assigned`);
        return response.data;
    });

    const mutation = useMutation((values) => axios.post(`http://localhost:4000/assigned`, values));

    const [message, setMessage] = useState("");
    const [error, setError] = useState("");
    const [loading, setLoading] = useState(false);

    if (isUserDetailsLoading) {
        return <h3>Loading...</h3>
    }

    if (isUserDetailsError) {
        return <h3>Error</h3>
    }

    if (isCoursesLoading) {
        return <h3>Loading...</h3>
    }

    if (isCoursesError) {
        return <h3>Error</h3>
    }

    if (isAssignedLoading) {
        return <h3>Loading...</h3>
    }

    if (isAssignedError) {
        return <h3>Error</h3>
    }

    const mentors = userDetails?.filter(user => user.role === "mentor");

    var flag = true;
    const handleSubmit = async (values) => {
        assigned?.forEach(assign => {
            if (assign.mentor === values.mentor &&
                assign.mentee === values.mentee &&
                assign.course === values.course) {
                setError("Already assigned!")
                flag = false
            }
        });
        if (flag) {
            console.log(values)
            setLoading(true);
            mutation.mutate(values);

            setLoading(false);
            setMessage("Assigned Successfully!")
        }
        refetch()
    }

    const handleClick = () => {
        setMessage("")
        setError("")
        refetch()
    }

    return (
        <div className="p-lg-5 d-flex flex-column body">
            <Card className="m-lg-5 p-5">
                <Card.Body className="p-5 ">
                    {message ? <div className='m-5 p-5'>
                        <Alert className="text-center" variant="success">{message}</Alert>
                        <Button
                            disabled={loading}
                            onClick={handleClick}
                            type="submit"
                            className='w-100'
                        >
                            Assign Again
                        </Button>
                    </div> :
                        <div>
                            {error && <Alert variant="danger">{error}</Alert>}
                            <h2 className='text-center mb-5'>
                                Assign Mentorship and Course
                            </h2>
                            <Formik
                                initialValues={initialValues}
                                validationSchema={validationSchema}
                                onSubmit={handleSubmit}
                            >
                                <FormikForm>
                                    <Row className='d-flex space-between'>
                                        <Col className='py-3'>
                                            <Form.Group>
                                                <Form.Label
                                                    htmlFor="mentor"
                                                    className='size'
                                                >
                                                    Mentor
                                                </Form.Label>
                                                <div className="mb-3">
                                                    <Field
                                                        className="w-100"
                                                        as="select"
                                                        name="mentor"
                                                        id="mentor"
                                                        autoComplete="true"
                                                    >
                                                        <option>--- Select mentor ---</option>
                                                        {mentors.map(mentor => {
                                                            return <option
                                                                key={mentor.id}
                                                                value={mentor.email}
                                                            >
                                                                {mentor.name}
                                                            </option>
                                                        })}
                                                    </Field>
                                                    <ErrorMessage
                                                        component="div"
                                                        className="message error"
                                                        name="mentor"
                                                    />
                                                </div>
                                            </Form.Group>
                                        </Col>
                                        <Col className='px-5 py-3'>
                                            <Form.Group>
                                                <Form.Label
                                                    htmlFor="mentee"
                                                    className="size"
                                                >
                                                    Mentee
                                                </Form.Label>
                                                <div className="mb-3">
                                                    <Field
                                                        className="w-100"
                                                        as="select"
                                                        name="mentee"
                                                        id="mentee"
                                                        autoComplete="true"

                                                    >
                                                        <option>--- Select mentee ---</option>
                                                        {userDetails.map(user => {
                                                            return user.role !== "admin" && <option
                                                                key={user.id}
                                                                value={user.email}
                                                            >
                                                                {user.name}
                                                            </option>
                                                        })}
                                                    </Field>
                                                    <ErrorMessage
                                                        component="div"
                                                        className="message error"
                                                        name="mentee"
                                                    />
                                                </div>
                                            </Form.Group>
                                        </Col>
                                        <Col className='py-3'>
                                            <Form.Group>
                                                <Form.Label
                                                    htmlFor="course"
                                                    className="size"
                                                >
                                                    Course
                                                </Form.Label>
                                                <div className="mb-3">
                                                    <Field
                                                        className="w-100"
                                                        as="select"
                                                        name="course"
                                                        id="course"
                                                        autoComplete="true"
                                                    >
                                                        <option>--- Select course ---</option>
                                                        {courses.map(course => {
                                                            return <option key={course.id} value={course.id}>{course.name}</option>
                                                        })}

                                                    </Field>
                                                    <ErrorMessage
                                                        component="div"
                                                        className="message error"
                                                        name="course"
                                                    />
                                                </div>
                                            </Form.Group>
                                        </Col>
                                    </Row>
                                    <Button
                                        disabled={loading}
                                        className='mt-3 w-100 font'
                                        type="submit"
                                    >
                                        Assign
                                    </Button>
                                </FormikForm>
                            </Formik>
                        </div>
                    }
                </Card.Body>
            </Card>
        </div>
    )
}
