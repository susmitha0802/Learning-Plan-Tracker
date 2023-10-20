import { Form } from 'react-bootstrap';

export const Exercises = ({ id, question }) => {
    return (
        <Form>
            <Form.Check
                type="checkbox"
                label={`${id}. ${question}`}
            />
        </Form>
    )
}
