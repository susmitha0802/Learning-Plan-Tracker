import { useQuery } from 'react-query';
import axios from 'axios';
import { Card, Table } from "react-bootstrap"

export const Assigned = () => {
    const { data: assigned, isLoading: isAssignedLoading, isError: isAssignedError } = useQuery(["assigned"], async () => {
        const response = await axios.get(`http://localhost:4000/assigned`);
        return response.data;
    });

    const { data: userDetails, isLoading: isUserDetailsLoading, isError: isUserDetailsError } = useQuery(["userDetails"], async () => {
        const response = await axios.get(`http://localhost:4000/userDetails`);
        return response.data;
    });

    const { data: courses, isLoading: isCoursesLoading, isError: isCoursesError } = useQuery(['courses'], async () => {
        const response = await axios.get(`http://localhost:4000/courses`);
        return response.data;
    });

    if (isAssignedLoading) {
        return <h3>Loading...</h3>
    }

    if (isAssignedError) {
        return <h3>Error</h3>
    }

    if (isUserDetailsLoading) {
        return <h3>Loading...</h3>
    }

    if (isUserDetailsError) {
        return <h3>Error</h3>
    }

    if (isCoursesLoading) {
        return <h3>Loading...</h3>
    }

    if (isCoursesError) {
        return <h3>Error</h3>
    }

    return (
        assigned.length === 0 ? <div className='body p-5'>
            <Card className='p-5 d-flex justify-content-center align-items-center text-black bg-body-secondary' >
                <Card.Body className='m-5 p-5 d-flex justify-content-center align-items-center flex-column '>
                    <h1>Assignments will be displayed here</h1>
                </Card.Body>
            </Card>
        </div> :
            <div className='mx-5 px-5'>
                <h1 className='mx-5 mb-2 pb-3 text-center'>Current Assignments</h1>
                <div className='m-5'>
                    <Table bordered hover className='size' >
                        <thead>
                            <tr>
                                <th>Mentor</th>
                                <th>Mentee</th>
                                <th>Course</th>
                            </tr>
                        </thead>
                        <tbody>
                            {assigned.map(assign => {
                                const mentor = userDetails.filter(user => user.email === assign.mentor)
                                const mentee = userDetails.filter(user => user.email === assign.mentee)
                                const course = courses.filter(course => course.id === JSON.parse(assign.course))
                                return <tr key={assign.id}>
                                    <td>{mentor[0].name}</td>
                                    <td>{mentee[0].name}</td>
                                    <td>{course[0].name}</td>
                                </tr>
                            })}
                        </tbody>
                    </Table>
                </div>
            </div >
    )
}
