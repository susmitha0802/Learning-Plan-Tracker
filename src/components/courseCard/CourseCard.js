import { useQuery } from 'react-query';
import axios from 'axios';
import { Card, Col } from 'react-bootstrap';
import { Link } from "react-router-dom";
import { FaEye } from "react-icons/fa";
import "../courseCard/CourseCard.css";

export const CourseCard = ({ courseId }) => {
  const { data: course, isLoading, isError } = useQuery(['courses', courseId], async () => {
    const response = await axios.get(`http://localhost:4000/courses/${courseId}`);
    return response.data;
  });

  if (isLoading) {
    return <h3>Loading...</h3>
  }

  if (isError) {
    return <h3>Error</h3>
  }

  return (
    <>
      <Col key={courseId} className="d-flex">
        <Card className="my-3 w-75 card-container">
          <Link className="text-decoration-none" to={`/courses/${courseId}`}>
            <Card.Img variant="top" src={course.logo} />
            <Card.Body className='bg-body-secondary'>
              <Card.Title className="text-center mt-2 mb-4 text-black" ><h2>{course.name}</h2></Card.Title>
              <hr />
              <Card.Text className="d-flex justify-content-around">
                <span className="font text-black">Estimated time - {course.time} days</span>
                <span className="text-decoration-none d-flex align-items-center font">
                  <FaEye /> View Syllabus
                </span>
              </Card.Text>
            </Card.Body>
          </Link>
        </Card>
      </Col>
    </>
  )
}