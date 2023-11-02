import { useQuery } from 'react-query';
import axios from 'axios';
import { Table } from "react-bootstrap"

export const Assigned = () => {
    const { data: assigned, isLoading, isError } = useQuery(["assigned"], async () => {
        const response = await axios.get(`http://localhost:4000/assigned`);
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
            <h4>Already Assigned</h4>
            {assigned.length !== 0 ? <Table bordered hover>
                <thead>
                    <tr>
                        <th>Mentor</th>
                        <th>Mentee</th>
                        <th>Course</th>
                    </tr>
                </thead>
                <tbody>
                    {assigned.map(assign => {
                        return <tr key={assign.id}>
                            <td>{assign.mentor}</td>
                            <td>{assign.mentee}</td>
                            <td>{assign.course}</td>
                        </tr>
                    })}
                </tbody>
            </Table> : <h4>Not assigned yet</h4>}
        </>

    )
}
