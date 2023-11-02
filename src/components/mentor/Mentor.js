import { useQuery } from 'react-query';
import axios from 'axios';
import { Row } from 'react-bootstrap';
import "../../App.css";
import { MenteeCard } from '../menteeCard/MenteeCard';
import { useAuth } from '../../contexts/AuthContext';

export const Mentor = () => {

  const { data: assigned, isLoading, isError } = useQuery(['assigned'], async () => {
    const response = await axios.get(`http://localhost:4000/assigned`);
    return response.data;
  });

  const { currentUser } = useAuth();

  if (isLoading) {
    return <h3>Loading...</h3>
  }

  if (isError) {
    return <h3>Loading...</h3>
  }

  const email = currentUser.email;

  const assignedMentees = assigned?.filter(assign => assign.mentor === email);

  return (
    <div className="body">
      <h1 className="p-5">My Mentees</h1>
      <Row xs={1} lg={3}>
        {
          assignedMentees?.map(assign => {
            return <MenteeCard key={assign.id}
              mentee={assign.mentee}
              courseId={assign.course}
            />
          })
        }
      </Row>
    </div>
  );
}
