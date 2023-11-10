import { useQuery } from 'react-query';
import axios from 'axios';
import { Card, Col } from 'react-bootstrap';
import { Link } from "react-router-dom";
import { FaEye } from "react-icons/fa";

export const MenteeCard = ({ id, mentee, courseId }) => {
    const { data: course, isLoading: isCoursesLoading, isError: isCoursesError } = useQuery(['courses', courseId], async () => {
        const response = await axios.get(`http://localhost:4000/courses/${courseId}`);
        return response.data;
    });

    const { data: userDetails, isLoading: isUserDetailsLoading, isError: isUserDetailsError } = useQuery(["userDetails"], async () => {
        const response = await axios.get(`http://localhost:4000/userDetails`);
        return response.data;
    });

    if (isCoursesLoading) {
        return <h3>Loading...</h3>
    }

    if (isCoursesError) {
        return <h3>Error</h3>
    }

    if (isUserDetailsLoading) {
        return <h3>Loading...</h3>
    }

    if (isUserDetailsError) {
        return <h3>Error</h3>
    }

    const menteeName = userDetails?.filter(user => user.email === mentee)[0].name

    return (
        <>
            <Col key={id} className="d-flex">
                <Card className="my-3 card-container w-75">
                    <Link className="text-decoration-none" to={`/mentor/${mentee}/${courseId}`}>
                        <Card.Img variant="top" src={course.logo} />
                        <Card.Body className='bg-body-secondary' >
                            <Card.Title className="text-center mt-2 mb-4 text-black"><h2>{menteeName} - {course.name}</h2></Card.Title>
                            <hr />
                            <Card.Text className="text-decoration-none d-flex align-items-center size justify-content-center">
                                <FaEye /> View submitted exercises
                            </Card.Text>
                        </Card.Body>
                    </Link>
                </Card>
            </Col>
        </>
    )
}