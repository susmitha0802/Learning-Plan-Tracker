import { useQuery } from 'react-query';
import axios from 'axios';
import { useParams } from "react-router-dom";
import { Image } from 'react-bootstrap';
import { Topic } from '../topic/Topic';
import { useEffect } from 'react';
import { getTotalExercises } from '../../utils/total';
import { CompletedProvider } from '../../contexts/CompletedContext';
import { Progress } from '../progress/Progress';
import "./Course.css";
import "../../App.css";


export const Course = () => {
    const params = useParams();
    const courseId = params.courseId;

    const { data: course, isLoading, isError } = useQuery(["course", courseId], async () => {
        const response = await axios.get(`http://localhost:4000/courses/${courseId}`);
        return response.data;
    });

    useEffect(() => {

        document.title = `${course?.name} Course`;

        const faviconLink = document.querySelector("link[rel~='icon']");
        faviconLink.href = course?.logo;

    }, [course?.name, course?.logo]);

    if (isLoading) {
        return <h3>Loading...</h3>
    }

    if (isError) {
        return <h3>Loading...</h3>
    }

    var totalExercises = 0;
    if (course?.topics) {
        totalExercises = getTotalExercises(0, course?.topics);
    }

    return (
        <CompletedProvider>
            <div className="body course-header">
                <div className="d-flex align-items-center justify-content-center">
                    <Image className="w-25 mx-5 px-5" src={course.logo} />
                    <div className="w-50 mx-5 px-5">
                        <h1 className="my-4">Learn {course.name} </h1>
                        <p className="mb-4 size">{course.caption}</p>
                        <p className="mb-4 size">Course Progress</p>
                        <Progress name={course.name} total={totalExercises} />
                    </div>
                </div>
                <div className="syllabus">
                    <h1 className="pb-3">Syllabus</h1>
                    {
                        course.topics.map(topic => {
                            return <Topic
                                key={topic.id}
                                id={topic.id}
                                name={topic.name}
                                resource={topic.resource}
                                exercises={topic.exercises}
                                total={totalExercises}
                            />
                        })
                    }
                </div>
            </div>
        </CompletedProvider>
    )
}
