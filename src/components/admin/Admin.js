import { Container, Row, Col } from 'react-bootstrap';
import { UsersList } from '../usersList/UsersList';
import { AssignCourse } from '../assignCourse/AssignCourse';
import { Assigned } from '../assigned/Assigned';

export const Admin = () => {
    return (
        <Container className='m-0' fluid>
            <Row>
                <Col sm={3} className='bg-light'>
                    <UsersList />
                </Col>
                <Col sm={5} >
                    <AssignCourse />
                </Col>
                <Col sm={4} style={{ backgroundColor: 'lightgrey' }}>
                    <Assigned />
                </Col>
            </Row>
        </Container>
    )
}
