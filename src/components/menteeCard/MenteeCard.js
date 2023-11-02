import { useQuery } from 'react-query';
import axios from 'axios';
import { Button, Card, Col } from 'react-bootstrap';
import { Link } from "react-router-dom";
import { FaEye } from "react-icons/fa";
// import "./CourseCard.css";

export const MenteeCard = ({ id, mentee, courseId }) => {
    const { data: course, isLoading, isError } = useQuery(['courses', courseId], async () => {
        const response = await axios.get(`http://localhost:4000/courses/${courseId}`);
        return response.data;
    });

    if (isLoading) {
        return <h3>Loading...</h3>
    }

    if (isError) {
        return <h3>Loading...</h3>
    }
    return (
        <>
            <Col key={id} className="d-flex">
                <Card>
                    {/* <Card.Header as="h5">Featured</Card.Header> */}
                    <Card.Body>
                        <Card.Title>{mentee}</Card.Title>
                        <Card.Text>{course.name}</Card.Text>
                        <Card.Text>Progress - {course.name}</Card.Text>

                        <Button variant="primary">View submitted exercises</Button>
                    </Card.Body>
                </Card>
            </Col>
        </>
    )
}