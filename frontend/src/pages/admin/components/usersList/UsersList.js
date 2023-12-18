import { useQuery } from 'react-query';
import axios from 'axios';
import { Tab, Tabs } from 'react-bootstrap';
import { UsersTable } from '../usersTable/UsersTable';

export const UsersList = () => {
    const { data: userDetails, isLoading, isError } = useQuery(["userDetails"], async () => {
        const response = await axios.get(`http://localhost:4000/userDetails`);
        return response.data;
    });

    if (isLoading) {
        return <h3>Loading...</h3>
    }

    if (isError) {
        return <h3>Error</h3>
    }
    const mentors = userDetails?.filter(user => user.role === "mentor");
    const mentees = userDetails?.filter(user => user.role === "mentee");

    return (
        <div className='mx-4 mt-5 p-2'>
            <h4 className='pb-3 '>
                List of mentors and mentees
            </h4>
            <Tabs defaultActiveKey="mentors">
                <Tab eventKey="mentors" title="Mentors">
                    <UsersTable users={mentors} />
                </Tab>
                <Tab eventKey="mentees" title="Mentees">
                    <UsersTable users={mentees} />
                </Tab>

            </Tabs>
        </div>
    )
}
