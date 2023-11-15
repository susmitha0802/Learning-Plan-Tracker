import { Container, Row, Col } from 'react-bootstrap';
import { UsersList } from './components/usersList/UsersList';
import { AssignButton } from './components/assignButton/AssignButton';
import { Assigned } from './components/assigned/Assigned';

export const Admin = () => {
    return (
        <Container className='body' fluid>
            <Row>
                <Col sm={3} className='border border-top-0 border-bottom-0 border-dark'>
                    <UsersList />
                </Col>
                <Col sm={9}>
                    <AssignButton />
                    <Assigned />
                </Col>
            </Row>
        </Container>
    )
}
