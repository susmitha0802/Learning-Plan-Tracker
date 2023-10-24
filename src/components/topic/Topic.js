import { Accordion } from 'react-bootstrap';
import { BsBookHalf } from "react-icons/bs";
import { Link } from 'react-router-dom';
import { Exercises } from '../exercises/Exercises';

export const Topic = ({ id, name, resource, exercises }) => {
    return (
        <Accordion key={id}>
            <Accordion.Item eventKey={id}>
                <Accordion.Header>{name}</Accordion.Header>
                <Accordion.Body>
                    <div>
                        <BsBookHalf /> <Link to={resource}>Lesson</Link>
                    </div>
                    {exercises && <div>
                        <b>Exercises</b>
                        {exercises.map(exercise => {
                            return <Exercises
                                key={exercise.id}
                                id={exercise.id}
                                question={exercise.question}
                                status={exercise.status}
                            />
                        })}
                    </div>
                    }
                </Accordion.Body>
            </Accordion.Item>
        </Accordion>
    )
}
