import { Container, Row, Col, Button } from 'react-bootstrap';
import { UsersList } from '../usersList/UsersList';
import { AssignCourse } from '../assignCourse/AssignCourse';
import { Assigned } from '../assigned/Assigned';
import { Link } from 'react-router-dom';
import { AssignButton } from '../assignButton/AssignButton';

export const Admin = () => {
    return (
        <Container className='body' fluid>
            <Row>
                <Col sm={3} className='border border-top-0 border-bottom-0 border-dark'>
                    <UsersList />
                    {/* <Assigned /> */}
                </Col>
                <Col sm={9}>
                    <AssignButton />
                    <Assigned />
                    {/* <AssignCourse /> */}
                </Col>
            </Row>
        </Container>

        // {/* <Container className='body' fluid>
        // <Row>
        //     <Col sm={2} className='border border-right border-light'>
        //         <UsersList />
        //         <Assigned />
        //     </Col>
        //     <Col sm={10}>
        //         <AssignCourse />
        //     </Col>
        // </Row>
        // </Container> */}
    )
}
