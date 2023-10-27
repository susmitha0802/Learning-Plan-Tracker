import { Accordion } from 'react-bootstrap';
import { Link } from 'react-router-dom';
import { Exercises } from '../exercises/Exercises';
import { BsBookHalf } from "react-icons/bs";
import { RiTodoFill } from "react-icons/ri";
import "../../App.css";

export const Topic = ({ id, name, resource, exercises, total }) => {

    return (
        <Accordion key={id}>
            <Accordion.Item eventKey={id}>
                <Accordion.Header><h5>{name}</h5></Accordion.Header>
                <Accordion.Body>
                    <div className='mx-3 d-flex align-items-center'>
                        <BsBookHalf className="size" />
                        <Link className="text-decoration-none px-3 size" to={resource} target="_blank">
                            <span>Lesson</span>
                        </Link>
                    </div>
                    {exercises && <div className="p-3">
                        <div className='d-flex align-items-center'>
                            <RiTodoFill className="size" />
                            <span className='px-3 size'>Exercises</span>
                        </div>
                        {exercises.map(exercise => {
                            return <Exercises
                                key={exercise.id}
                                id={exercise.id}
                                question={exercise.question}
                                topicId={id}
                            />
                        })}
                    </div>
                    }
                </Accordion.Body>
            </Accordion.Item>
        </Accordion>
    )
}


