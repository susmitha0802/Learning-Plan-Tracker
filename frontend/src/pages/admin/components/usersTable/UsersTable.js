import { Table } from "react-bootstrap"

export const UsersTable = ({ users }) => {
    return (
        <Table bordered hover>
            <thead>
                <tr>
                    <th>Name</th>
                </tr>
            </thead>
            <tbody>
                {users.map(user => {
                    return <tr key={user.id}>
                        <td>{user.name}</td>
                    </tr>
                })}
            </tbody>
        </Table>
    )
}
