import { Table } from "react-bootstrap"

export const UsersTable = ({ users }) => {
    return (
        <Table bordered hover>
            <thead>
                <tr>
                    <th>Name</th>
                    <th>Email</th>
                </tr>
            </thead>
            <tbody>
                {users.map(user => {
                    return <tr key={user.id}>
                        <td>{user.name}</td>
                        <td>{user.email}</td>
                    </tr>
                })}
            </tbody>
        </Table>
    )
}
