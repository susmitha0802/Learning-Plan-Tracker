import { Card, Col } from 'react-bootstrap';
import { Link } from "react-router-dom";
import { FaEye } from "react-icons/fa";
import "./CourseCard.css";

export const CourseCard = ({ id, logo, name, time }) => {
  return (
    <>
      <Col key={id} className="d-flex">
        <Card className="mx-5 my-3 card-container">
          <Link className="text-decoration-none" to={`/courses/${id}`}>
            <Card.Img variant="top" src={logo} />
            <Card.Body className="bg-dark-subtle" >
              <Card.Title className="text-center mt-2 mb-4" style={{ color: 'black' }}>{name}</Card.Title>
              <hr />
              <Card.Text className="d-flex justify-content-around">
                <span style={{ color: 'black' }}>Estimated time - {time} days</span>
                <span className="text-decoration-none d-flex align-items-center">
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