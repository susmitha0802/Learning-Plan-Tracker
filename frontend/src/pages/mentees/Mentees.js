import { useQuery } from 'react-query';
import axios from 'axios';
import { Card, Row } from 'react-bootstrap';
import { useAuth } from '../../contexts/AuthContext';
import { MenteeCard } from './components/menteeCard/MenteeCard';
import "../../App.css";

export const Mentees = () => {

  const { data: assigned, isLoading, isError } = useQuery(['assigned'], async () => {
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

  const assignedMentees = assigned?.filter(assign => assign.mentor === email);

  return (
    assignedMentees.length === 0 ?
      <div className='body p-5'>
        <Card className='m-5 p-5 d-flex justify-content-center align-items-center text-black bg-body-secondary'>
          <Card.Body className='m-5 p-5'>
            <h1>No mentees are assigned yet</h1>
          </Card.Body>
        </Card>
      </div> :
      <div className="body">
        <h1 className="mx-5 p-5">My Mentees</h1>
        <Row className="mx-5 px-4" xs={1} lg={3}>
          {
            assignedMentees?.map(assign => {
              return <MenteeCard key={assign.course}
                mentee={assign.mentee}
                courseId={assign.course}
              />
            })
          }
        </Row>
      </div>
  );
}
