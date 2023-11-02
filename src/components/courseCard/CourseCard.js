import { useQuery } from 'react-query';
import axios from 'axios';
import { Card, Col } from 'react-bootstrap';
import { Link } from "react-router-dom";
import { FaEye } from "react-icons/fa";
import "./CourseCard.css";

export const CourseCard = ({ courseId, mentor }) => {
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
  console.log(courseId)
  console.log(`${courseId} ${course?.name}`)
  return (
    <>
      <Col key={courseId} className="d-flex">
        <Card className="mx-5 my-3 card-container">
          <Link className="text-decoration-none" to={`/courses/${courseId}`}>
            <Card.Img variant="top" src={course.logo} />
            <Card.Body className="bg-dark-subtle" >
              <Card.Title className="text-center mt-2 mb-4" style={{ color: 'black' }}><h2>{course.name}</h2></Card.Title>
              <hr />
              <Card.Text className="d-flex justify-content-around">
                <span className="size" style={{ color: 'black' }}>Estimated time - {course.time} days</span>
                <span className="text-decoration-none d-flex align-items-center size">
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