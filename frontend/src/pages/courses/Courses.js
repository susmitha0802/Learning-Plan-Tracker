import { useQuery } from 'react-query';
import axios from 'axios';
import { Card, Row } from 'react-bootstrap';
import { useAuth } from '../../contexts/AuthContext';
import { CourseCard } from './components/courseCard/CourseCard';

export const Courses = () => {
  const { data: assigned, isLoading, isError } = useQuery(["assigned"], async () => {
    const response = await axios.get(`http://localhost:4000/assigned`);
    return response.data;
  });

  const { currentUser } = useAuth();

  if (isLoading) {
    return <h3>Loading...</h3>
  }

  if (isError) {
    return <h3>Error</h3>
  }

  const email = currentUser.email;

  const assignedCourses = assigned?.filter(assign => assign.mentee === email);

  return (
    assignedCourses.length === 0 ?
      <div className='body m-5 p-5'>
        <Card className='m-5 p-5 text-center text-black bg-body-secondary' >
          <Card.Body className='m-5 p-5  '>
            <h1>No courses are assigned yet</h1>
          </Card.Body>
        </Card>
      </div> :
      <div className="body">
        <h1 className="mx-5 p-5">My Courses</h1>
        <Row className="mx-5 px-4" xs={1} lg={3}>
          {
            assignedCourses.map(assignedCourse => {
              return <CourseCard key={assignedCourse.course}
                courseId={assignedCourse.course}
              />
            })
          }
        </Row>
      </div>
  );
}
