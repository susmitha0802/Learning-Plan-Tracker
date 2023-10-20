import { useQuery } from 'react-query';
import axios from 'axios';
import { Row } from 'react-bootstrap';
import { CourseCard } from '../courseCard/CourseCard';

export const Courses = () => {

  const { data: courses, isLoading, isError } = useQuery(['courses'], async () => {
    const response = await axios.get(`http://localhost:4000/courses`);
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
      <h1 className="m-5">My Courses</h1>
      <Row xs={1} lg={3}>
        {
          courses?.map(course => {
            return <CourseCard key={course.id}
              id={course.id}
              logo={course.logo}
              name={course.name}
              time={course.time}
            />
          })
        }
      </Row>
    </>
  );
}
