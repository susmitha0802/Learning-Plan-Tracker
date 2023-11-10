import { useQuery } from 'react-query';
import axios from 'axios';
import { useParams } from "react-router-dom";
import { Accordion, Image } from 'react-bootstrap';
import { useEffect } from 'react';
import { useUser } from '../../contexts/UserContext';
import { DownloadExercise } from '../downloadExercise/DownloadExercise';
import { getTotalExercises } from '../../utils/total';
import { Progress } from '../progress/Progress';
import "../../App.css";

export const SubmittedExercises = () => {
    const params = useParams();
    const mentee = params.mentee;
    const courseId = JSON.parse(params.courseId);

    const { data: course, isLoading: isCoursesLoading, isError: isCoursesError } = useQuery(["course", courseId], async () => {
        const response = await axios.get(`http://localhost:4000/courses/${courseId}`);
        return response.data;
    });

    const { getsubmittedByUserEmail } = useUser();

    useEffect(() => {

        document.title = `${course?.name} Course`;

        const faviconLink = document.querySelector("link[rel~='icon']");
        faviconLink.href = course?.logo;

    }, [course?.name, course?.logo]);

    if (isCoursesLoading) {
        return <h3>Loading...</h3>
    }

    if (isCoursesError) {
        return <h3>Error</h3>
    }

    const submitted = [];

    getsubmittedByUserEmail(mentee).forEach(submit => {
        if (submit.courseId === courseId) {
            submitted.push(submit)
        }
    })

    var totalExercises = 0;
    if (course?.topics) {
        totalExercises = getTotalExercises(0, course?.topics);
    }

    return (
        <div className="body course-header">
            <div className="d-flex align-items-center justify-content-center">
                <Image className="w-25 mx-5 px-5" src={course.logo} />
                <div className="w-50 mx-5 px-5">
                    <h1 className="my-4">Learn {course.name} </h1>
                    <p className="mb-4 size">{course.caption}</p>
                    <p className="mb-2 size">Course Progress</p>
                    <Progress id={courseId} total={totalExercises} email={mentee} />
                    <hr />
                    <p className="mb-4 size">Contact your mentee - {mentee}</p>
                </div>
            </div>
            <div className="syllabus">
                <h1 className="pb-3">Submitted Exercises</h1>
                {submitted.length === 0 ? <h3>No exercises are submitted yet</h3> :
                    submitted.map(submit => {
                        const topicId = submit.topicId
                        const exerciseId = submit.exerciseId
                        const targetKey = `exercise-${courseId}-${topicId}-${exerciseId}`;
                        const question = course.topics[topicId - 1].exercises.filter(exercise => exercise.id === exerciseId)[0].question
                        return <Accordion key={exerciseId}>
                            <Accordion.Item eventKey={exerciseId}>
                                <Accordion.Header><h5>Exercise - {exerciseId}</h5></Accordion.Header>
                                <Accordion.Body className='mx-5 my-3'>
                                    <p className='size'>{question.split("\n").map((i, key) => {
                                        return <div key={key}>{i}</div>;
                                    })}</p>
                                    <DownloadExercise name={submit.fileName} url={submit[targetKey]} />
                                </Accordion.Body>
                            </Accordion.Item>
                        </Accordion>
                    })
                }
            </div>
        </div >
    )
}
