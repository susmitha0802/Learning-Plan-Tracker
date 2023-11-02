import { useQuery } from 'react-query';
import axios from 'axios';
import { Row } from 'react-bootstrap';
import { CourseCard } from '../courseCard/CourseCard';
import "../../App.css";
import { useAuth } from '../../contexts/AuthContext';

export const Courses = () => {
  const { data: assigned, isLoading: isAssignedLoading, isError: isAssignedError } = useQuery(["assigned"], async () => {
    const response = await axios.get(`http://localhost:4000/assigned`);
    return response.data;
  });

  const { currentUser } = useAuth();

  if (isAssignedLoading) {
    return <h3>Loading...</h3>
  }

  if (isAssignedError) {
    return <h3>Loading...</h3>
  }

  const email = currentUser.email;

  const assignedCourses = assigned?.filter(assign => assign.mentee === email);

  return (
    <div className="body">
      <h1 className="p-5">My Courses</h1>
      <Row xs={1} lg={3}>
        {
          assignedCourses.map(assignedCourse => {
            return <CourseCard key={assignedCourse.id}
              courseId={assignedCourse.course}
              mentor={assignedCourse.mentor}
            />
          })
        }
      </Row>
    </div>
  );
}
