import React, { useState, useEffect } from 'react';
import { Checkbox, Upload, Button, message } from 'antd';
import { UploadOutlined } from '@ant-design/icons';
import { useUser } from '../../contexts/UserContext';
import { useAuth } from '../../contexts/AuthContext';
import { CourseContext } from '../course/Course';
import { useContext } from 'react';


export const Exercises = ({ id, question, checked, onChange, topicId }) => {
    const { currentUser } = useAuth();
    const [completed, setCompleted] = useState(false);
    const { setsubmittedInLocalStorage, getSubmittedFromLocalStorage } = useUser();
    const courseId = useContext(CourseContext);


    const handleCheckboxChange = (e) => {
        setCompleted(e.target.checked);
    };

    const customRequest = async ({ onSuccess, onError, file }) => {
        try {
            const reader = new FileReader();
            reader.onloadend = () => {
                const data = reader.result;
                setsubmittedInLocalStorage(currentUser.email, courseId, topicId, id, data)
                setCompleted(true);
                onSuccess();
                message.success(`${file.name} uploaded successfully.`);
            };
            reader.readAsDataURL(file);
        }
        catch (error) {
            message.error(`${file.name} file upload failed.`);
            onError(error);
        }
    };

    useEffect(() => {
        const storedData = getSubmittedFromLocalStorage(currentUser.email, courseId, topicId, id);
        if (storedData) {
            setCompleted(true);
        }
    }, [currentUser.email, courseId, topicId, id, getSubmittedFromLocalStorage]);

    return (
        <>
            <Checkbox checked={completed} onChange={handleCheckboxChange}>
                Exercise {id}
            </Checkbox>
            {question.split("\n").map((i, key) => {
                return <div key={key}>{i}</div>;
            })}
            <Upload customRequest={customRequest} listType="picture" >
                <Button disabled={completed} icon={<UploadOutlined />}>Upload File</Button>
            </Upload>
        </>
    );
}
